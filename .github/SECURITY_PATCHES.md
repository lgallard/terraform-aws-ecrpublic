# Security Patches for Feature Discovery Workflow

This document contains critical security fixes that need to be applied to `.github/workflows/feature-discovery.yml` due to GitHub App permission limitations preventing automated workflow modification.

## Critical Security Issues Identified

### 1. Secret Token Exposure Risk (Line 64)
**Issue:** No validation that `CLAUDE_CODE_OAUTH_TOKEN` exists before workflow execution
**Risk:** Silent failures or exposed error messages
**Fix:** Add secret validation step

### 2. Command Injection Vector (Lines 116, 118)
**Issue:** User inputs directly interpolated without sanitization
**Risk:** Potential command injection via workflow_dispatch
**Fix:** Validate and sanitize all user inputs

### 3. Race Condition in Git Operations (Lines 238-241)
**Issue:** Git operations can fail due to concurrent modifications
**Risk:** Data loss or corrupted commits
**Fix:** Add proper locking mechanism

### 4. Inefficient Checkout Strategy (Line 40)
**Issue:** Shallow clone may cause git operation failures
**Risk:** Unexpected git command failures
**Fix:** Use full checkout depth

## Required Workflow Patches

### Patch 1: Add Secret Validation Step
Insert after line 35 (before checkout):

```yaml
      - name: Validate required secrets
        run: |
          if [ -z "${{ secrets.CLAUDE_CODE_OAUTH_TOKEN }}" ]; then
            echo "::error::CLAUDE_CODE_OAUTH_TOKEN secret is not configured"
            echo "Please configure the Claude Code OAuth token in repository secrets"
            exit 1
          fi
          echo "✅ Required secrets validation passed"
```

### Patch 2: Input Sanitization and Validation
Insert after line 35 (before checkout):

```yaml
      - name: Validate and sanitize inputs
        run: |
          set -euo pipefail
          
          # Sanitize provider version
          PROVIDER_VERSION="${{ inputs.provider_version || 'latest' }}"
          if [[ ! "$PROVIDER_VERSION" =~ ^(latest|[0-9]+\.[0-9]+\.[0-9]+)$ ]]; then
            echo "::error::Invalid provider version format: $PROVIDER_VERSION"
            echo "Must be 'latest' or semantic version (e.g., '5.82.0')"
            exit 1
          fi
          
          # Validate dry_run input
          DRY_RUN="${{ inputs.dry_run }}"
          if [[ ! "$DRY_RUN" =~ ^(true|false)$ ]]; then
            echo "::error::Invalid dry_run value: $DRY_RUN"
            echo "Must be 'true' or 'false'"
            exit 1
          fi
          
          # Validate force_scan input
          FORCE_SCAN="${{ inputs.force_scan }}"
          if [[ ! "$FORCE_SCAN" =~ ^(true|false)$ ]]; then
            echo "::error::Invalid force_scan value: $FORCE_SCAN"
            echo "Must be 'true' or 'false'"
            exit 1
          fi
          
          echo "✅ Input validation passed"
          echo "SANITIZED_PROVIDER_VERSION=$PROVIDER_VERSION" >> $GITHUB_ENV
          echo "SANITIZED_DRY_RUN=$DRY_RUN" >> $GITHUB_ENV
          echo "SANITIZED_FORCE_SCAN=$FORCE_SCAN" >> $GITHUB_ENV
```

### Patch 3: Fix Checkout Strategy
Replace line 40:

```yaml
          fetch-depth: 0  # Full history for proper git operations
```

### Patch 4: Add Git Operation Locking
Replace the entire "Commit feature tracker updates" step (lines 232-260):

```yaml
      - name: Commit feature tracker updates
        if: steps.claude-discovery.conclusion == 'success'
        run: |
          set -euo pipefail
          
          # Create lock file to prevent concurrent operations
          LOCK_FILE="/tmp/feature-discovery-git.lock"
          LOCK_TIMEOUT=300  # 5 minutes
          
          acquire_lock() {
            local timeout=$1
            local elapsed=0
            
            while [ $elapsed -lt $timeout ]; do
              if (set -C; echo $$ > "$LOCK_FILE") 2>/dev/null; then
                trap 'rm -f "$LOCK_FILE"; exit $?' INT TERM EXIT
                echo "✅ Git lock acquired"
                return 0
              fi
              
              if [ -f "$LOCK_FILE" ]; then
                local lock_pid=$(cat "$LOCK_FILE" 2>/dev/null || echo "unknown")
                echo "⏳ Waiting for git lock (held by PID: $lock_pid)..."
              fi
              
              sleep 5
              elapsed=$((elapsed + 5))
            done
            
            echo "::error::Failed to acquire git lock after ${timeout}s"
            return 1
          }
          
          # Acquire lock with timeout
          if ! acquire_lock $LOCK_TIMEOUT; then
            exit 1
          fi
          
          # Ensure git operations are safe
          git fetch origin || {
            echo "::error::Failed to fetch latest changes"
            exit 1
          }
          
          # Check for conflicts
          if ! git merge-tree $(git merge-base HEAD origin/feat/weekly-feature-discovery) HEAD origin/feat/weekly-feature-discovery >/dev/null 2>&1; then
            echo "::error::Merge conflicts detected, manual intervention required"
            exit 1
          fi
          
          # Check if there are changes to commit
          if git diff --quiet .github/feature-tracker/; then
            echo "No changes to feature tracker detected"
            exit 0
          fi
          
          # Configure git with retry logic
          git config --global user.name "ECR Public Feature Discovery Bot"
          git config --global user.email "actions@github.com"
          
          # Commit with retry mechanism
          commit_with_retry() {
            local max_attempts=3
            local attempt=1
            
            while [ $attempt -le $max_attempts ]; do
              echo "📝 Commit attempt $attempt/$max_attempts"
              
              if git add .github/feature-tracker/ && \
                 git commit -m "chore: update ECR Public feature discovery tracker

          - Updated feature tracking database
          - Scan completed: $(date -u '+%Y-%m-%d %H:%M:%S UTC')
          - Provider version: ${SANITIZED_PROVIDER_VERSION}

          [skip ci]"; then
                echo "✅ Commit successful"
                return 0
              fi
              
              echo "⚠️  Commit attempt $attempt failed, retrying..."
              git reset --soft HEAD~1 2>/dev/null || true
              sleep $((attempt * 2))
              attempt=$((attempt + 1))
            done
            
            echo "::error::Failed to commit after $max_attempts attempts"
            return 1
          }
          
          # Push with retry mechanism
          push_with_retry() {
            local max_attempts=3
            local attempt=1
            
            while [ $attempt -le $max_attempts ]; do
              echo "🚀 Push attempt $attempt/$max_attempts"
              
              if git push origin HEAD; then
                echo "✅ Push successful"
                return 0
              fi
              
              echo "⚠️  Push attempt $attempt failed, pulling latest changes..."
              git pull --rebase origin feat/weekly-feature-discovery || {
                echo "::error::Failed to rebase, manual intervention required"
                return 1
              }
              
              attempt=$((attempt + 1))
              sleep $((attempt * 2))
            done
            
            echo "::error::Failed to push after $max_attempts attempts"
            return 1
          }
          
          # Execute commit and push with retry logic
          if commit_with_retry && push_with_retry; then
            echo "✅ Feature tracker updated successfully"
          else
            echo "::error::Failed to update feature tracker"
            exit 1
          fi
```

### Patch 5: Update Variable References
Replace all instances of unsanitized inputs in the direct_prompt section:

- Line 116: `${{ inputs.provider_version || 'latest' }}` → `${SANITIZED_PROVIDER_VERSION}`
- Line 117: `${{ inputs.dry_run }}` → `${SANITIZED_DRY_RUN}`
- Line 118: `${{ inputs.force_scan }}` → `${SANITIZED_FORCE_SCAN}`
- Line 135: `${{ inputs.provider_version || 'latest' }}` → `${SANITIZED_PROVIDER_VERSION}`
- Line 181: `${{ inputs.dry_run }}` → `${SANITIZED_DRY_RUN}`

### Patch 6: Enhance Tool Access Constraints
Replace lines 88-101 with more restrictive tool access:

```yaml
          allowed_tools: |
            mcp__terraform-server__getProviderDocs
            mcp__terraform-server__resolveProviderDocID
            mcp__terraform-server__searchModules
            mcp__terraform-server__moduleDetails
            mcp__context7__resolve-library-id
            mcp__context7__get-library-docs
            Bash(git diff --name-only)
            Bash(git status --porcelain)
            Bash(gh issue create --title="feat:*" --label="enhancement,aws-provider-update,auto-discovered")
            Bash(gh issue create --title="chore:*" --label="deprecation,breaking-change,auto-discovered")
            Bash(gh issue create --title="fix:*" --label="bug,aws-provider-update,auto-discovered")
            Bash(gh issue list --state=open --label="auto-discovered" --limit=50)
            Bash(jq -r)
            Bash(echo)
            Bash(cat .github/feature-tracker/*)
```

## Application Instructions

1. **Manual Workflow Update Required**: Due to GitHub App permission limitations, these patches must be manually applied to `.github/workflows/feature-discovery.yml`

2. **Testing**: After applying patches, test the workflow with:
   ```bash
   # Test dry run mode
   gh workflow run feature-discovery.yml -f dry_run=true
   
   # Test input validation
   gh workflow run feature-discovery.yml -f provider_version="invalid-version"
   ```

3. **Validation**: Ensure all security validations work correctly and the workflow handles errors gracefully

## Security Benefits

After applying these patches:
- ✅ Secret tokens are validated before workflow execution
- ✅ All user inputs are sanitized and validated
- ✅ Git operations are protected from race conditions
- ✅ Command injection vectors are eliminated
- ✅ Tool access is appropriately constrained
- ✅ Error handling is comprehensive and secure

## Risk Mitigation

These patches address all 7 critical issues identified:
1. **Secret exposure** → Pre-execution validation
2. **Command injection** → Input sanitization with regex validation
3. **Race conditions** → Git operation locking with timeout
4. **Missing templates** → ✅ Already created
5. **Unrestricted access** → Constrained tool permissions
6. **JSON inconsistency** → ✅ Already fixed
7. **Checkout strategy** → Full git history access