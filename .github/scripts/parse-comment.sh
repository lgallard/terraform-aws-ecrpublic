#!/bin/bash
set -euo pipefail

# Parse Claude code review comment commands
COMMENT_BODY="${1:-}"
EVENT_NAME="${2:-}"

# Default values
echo "mode=review" >> "$GITHUB_OUTPUT"
echo "focus=code-quality,security,performance" >> "$GITHUB_OUTPUT"
echo "verbose=false" >> "$GITHUB_OUTPUT"
echo "include_tests=true" >> "$GITHUB_OUTPUT"

# Parse comment content for commands
if [[ "$EVENT_NAME" == "issue_comment" || "$EVENT_NAME" == "pull_request_review_comment" ]]; then
    COMMENT="$COMMENT_BODY"

    # Extract command and parameters
    if echo "$COMMENT" | grep -qi "codebot hunt"; then
        echo "mode=hunt" >> "$GITHUB_OUTPUT"
        echo "focus=bugs,security,performance" >> "$GITHUB_OUTPUT"
        echo "verbose=false" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "codebot analyze"; then
        echo "mode=analyze" >> "$GITHUB_OUTPUT"
        echo "focus=architecture,patterns,complexity" >> "$GITHUB_OUTPUT"
        echo "verbose=true" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "codebot security"; then
        echo "mode=security" >> "$GITHUB_OUTPUT"
        echo "focus=security,vulnerabilities,compliance" >> "$GITHUB_OUTPUT"
        echo "verbose=true" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "codebot performance"; then
        echo "mode=performance" >> "$GITHUB_OUTPUT"
        echo "focus=performance,optimization,bottlenecks" >> "$GITHUB_OUTPUT"
        echo "verbose=true" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "codebot review"; then
        echo "mode=review" >> "$GITHUB_OUTPUT"
        echo "focus=code-quality,security,performance" >> "$GITHUB_OUTPUT"
        echo "verbose=true" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "codebot"; then
        # Default to hunt mode for simple "codebot" command
        echo "mode=hunt" >> "$GITHUB_OUTPUT"
        echo "focus=bugs,security,performance" >> "$GITHUB_OUTPUT"
        echo "verbose=false" >> "$GITHUB_OUTPUT"
    fi

    # Check for verbose flag
    if echo "$COMMENT" | grep -qi "verbose\|detailed"; then
        echo "verbose=true" >> "$GITHUB_OUTPUT"
    fi

    # Check for specific focus areas
    if echo "$COMMENT" | grep -qi "security"; then
        echo "focus=security,vulnerabilities,compliance" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "performance"; then
        echo "focus=performance,optimization,bottlenecks" >> "$GITHUB_OUTPUT"
    elif echo "$COMMENT" | grep -qi "tests"; then
        echo "focus=test-coverage,test-quality" >> "$GITHUB_OUTPUT"
    fi
fi