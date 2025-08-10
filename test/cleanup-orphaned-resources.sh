#!/bin/bash

# Cleanup script for orphaned ECR Public repositories from interrupted tests
# This script identifies and removes ECR Public repositories that were created by tests
# but not properly cleaned up due to test interruptions or failures.

set -e

# Configuration
DEFAULT_REGION="us-east-1"  # ECR Public is only available in us-east-1
DEFAULT_MAX_AGE_HOURS=6     # Clean up resources older than 6 hours by default
DEFAULT_DRY_RUN="true"      # Default to dry run mode for safety

# Parse command line arguments
REGION="${1:-$DEFAULT_REGION}"
MAX_AGE_HOURS="${2:-$DEFAULT_MAX_AGE_HOURS}"
DRY_RUN="${3:-$DEFAULT_DRY_RUN}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== ECR Public Repository Cleanup Script ===${NC}"
echo -e "${BLUE}Region: ${REGION}${NC}"
echo -e "${BLUE}Max age: ${MAX_AGE_HOURS} hours${NC}"
echo -e "${BLUE}Dry run: ${DRY_RUN}${NC}"
echo ""

# Check if required tools are available
if ! command -v aws &> /dev/null; then
    echo -e "${RED}Error: AWS CLI is not installed or not in PATH${NC}"
    exit 1
fi

if ! command -v jq &> /dev/null; then
    echo -e "${RED}Error: jq is not installed or not in PATH${NC}"
    echo "Please install jq: https://stedolan.github.io/jq/download/"
    exit 1
fi

# Check AWS credentials
if ! aws sts get-caller-identity &> /dev/null; then
    echo -e "${RED}Error: AWS credentials not configured or invalid${NC}"
    exit 1
fi

echo -e "${GREEN}✓ AWS CLI, jq, and credentials verified${NC}"

# Function to check if a repository is older than max age
is_repository_old() {
    local repo_name="$1"
    local created_at

    # Get repository creation date
    created_at=$(aws ecr-public describe-repositories \
        --region "$REGION" \
        --repository-names "$repo_name" \
        --query 'repositories[0].createdAt' \
        --output text 2>/dev/null || echo "")

    if [[ -z "$created_at" ]]; then
        return 1
    fi

    # Convert to epoch time
    local created_epoch
    created_epoch=$(date -d "$created_at" +%s 2>/dev/null || date -j -f "%Y-%m-%dT%H:%M:%S" "${created_at%.*}" +%s 2>/dev/null || echo "0")

    local current_epoch
    current_epoch=$(date +%s)

    local age_seconds=$((current_epoch - created_epoch))
    local max_age_seconds=$((MAX_AGE_HOURS * 3600))

    if [[ $age_seconds -gt $max_age_seconds ]]; then
        return 0  # Repository is old
    else
        return 1  # Repository is not old enough
    fi
}

# Function to get repository tags
get_repository_tags() {
    local repo_arn="$1"
    aws ecr-public list-tags-for-resource \
        --region "$REGION" \
        --resource-arn "$repo_arn" \
        --query 'tags' \
        --output json 2>/dev/null || echo "[]"
}

# Function to check if repository has terratest tags
has_terratest_tags() {
    local tags_json="$1"

    # Check if tags contain Purpose=terratest
    echo "$tags_json" | jq -r '.[] | select(.key=="Purpose") | .value' | grep -q "terratest"
}

# Function to delete repository
delete_repository() {
    local repo_name="$1"

    echo -e "${YELLOW}Deleting repository: $repo_name${NC}"

    if [[ "$DRY_RUN" == "true" ]]; then
        echo -e "${BLUE}[DRY RUN] Would delete repository: $repo_name${NC}"
        return 0
    fi

    # Delete the repository
    if aws ecr-public delete-repository \
        --region "$REGION" \
        --repository-name "$repo_name" \
        --force &> /dev/null; then
        echo -e "${GREEN}✓ Successfully deleted repository: $repo_name${NC}"
        return 0
    else
        echo -e "${RED}✗ Failed to delete repository: $repo_name${NC}"
        return 1
    fi
}

# Main cleanup logic
echo -e "${BLUE}Scanning for ECR Public repositories...${NC}"

# Get all repositories
repositories=$(aws ecr-public describe-repositories \
    --region "$REGION" \
    --query 'repositories[].repositoryName' \
    --output text 2>/dev/null || echo "")

if [[ -z "$repositories" ]]; then
    echo -e "${GREEN}No repositories found in region $REGION${NC}"
    exit 0
fi

echo -e "${BLUE}Found repositories: $repositories${NC}"
echo ""

total_repos=0
terratest_repos=0
old_repos=0
deleted_repos=0
failed_deletions=0

# Process each repository
for repo_name in $repositories; do
    total_repos=$((total_repos + 1))

    echo -e "${BLUE}Analyzing repository: $repo_name${NC}"

    # Get repository ARN for tagging
    repo_arn=$(aws ecr-public describe-repositories \
        --region "$REGION" \
        --repository-names "$repo_name" \
        --query 'repositories[0].repositoryArn' \
        --output text 2>/dev/null || echo "")

    if [[ -z "$repo_arn" ]]; then
        echo -e "${YELLOW}  Warning: Could not get ARN for repository $repo_name${NC}"
        continue
    fi

    # Get tags
    tags_json=$(get_repository_tags "$repo_arn")

    # Check if it's a terratest repository
    if has_terratest_tags "$tags_json"; then
        terratest_repos=$((terratest_repos + 1))
        echo -e "${GREEN}  ✓ Repository has terratest tags${NC}"

        # Check if repository is old enough to delete
        if is_repository_old "$repo_name"; then
            old_repos=$((old_repos + 1))
            echo -e "${YELLOW}  ✓ Repository is older than $MAX_AGE_HOURS hours${NC}"

            # Extract additional tag information for logging
            test_run=$(echo "$tags_json" | jq -r '.[] | select(.key=="TestRun") | .value' 2>/dev/null || echo "unknown")
            created_at=$(echo "$tags_json" | jq -r '.[] | select(.key=="CreatedAt") | .value' 2>/dev/null || echo "unknown")

            echo -e "${BLUE}  TestRun: $test_run, CreatedAt: $created_at${NC}"

            # Delete the repository
            if delete_repository "$repo_name"; then
                deleted_repos=$((deleted_repos + 1))
            else
                failed_deletions=$((failed_deletions + 1))
            fi
        else
            echo -e "${BLUE}  Repository is not old enough (< $MAX_AGE_HOURS hours)${NC}"
        fi
    else
        echo -e "${BLUE}  Repository does not have terratest tags, skipping${NC}"
    fi

    echo ""
done

# Summary
echo -e "${BLUE}=== Cleanup Summary ===${NC}"
echo -e "${BLUE}Total repositories scanned: $total_repos${NC}"
echo -e "${GREEN}Terratest repositories found: $terratest_repos${NC}"
echo -e "${YELLOW}Old repositories eligible for cleanup: $old_repos${NC}"

if [[ "$DRY_RUN" == "true" ]]; then
    echo -e "${BLUE}Repositories that would be deleted: $old_repos${NC}"
    echo -e "${YELLOW}Run with DRY_RUN=false to actually delete repositories${NC}"
else
    echo -e "${GREEN}Successfully deleted repositories: $deleted_repos${NC}"
    if [[ $failed_deletions -gt 0 ]]; then
        echo -e "${RED}Failed deletions: $failed_deletions${NC}"
    fi
fi

echo ""
echo -e "${BLUE}Usage examples:${NC}"
echo -e "${BLUE}  ./cleanup-orphaned-resources.sh                    # Dry run with defaults${NC}"
echo -e "${BLUE}  ./cleanup-orphaned-resources.sh us-east-1 12 false # Clean resources older than 12h${NC}"
echo -e "${BLUE}  ./cleanup-orphaned-resources.sh us-east-1 1 false  # Clean resources older than 1h${NC}"

exit 0
