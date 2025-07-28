#!/bin/bash
set -euo pipefail

# Find the most recent open PR for the given branch
BRANCH_NAME="${1:-}"
REPOSITORY="${2:-}"

if [[ -z "$BRANCH_NAME" || -z "$REPOSITORY" ]]; then
    echo "Usage: $0 <branch_name> <repository>"
    exit 1
fi

# Find the most recent open PR for this branch
PR_NUMBER=$(gh pr list --repo "$REPOSITORY" --state open --head "$BRANCH_NAME" --json number --jq '.[0].number // empty')

if [[ -z "$PR_NUMBER" ]]; then
    echo "No open PR found for branch $BRANCH_NAME"
    exit 1
fi

echo "pr_number=$PR_NUMBER" >> "$GITHUB_OUTPUT"