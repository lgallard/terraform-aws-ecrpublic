# AWS Resource Cleanup for Terratest

This document explains the resource cleanup mechanisms implemented to prevent orphaned AWS resources from interrupted test runs.

## Problem

When Terratest integration tests are interrupted (due to CI failures, manual cancellation, or other issues), AWS ECR Public repositories may be left behind, leading to:

- **Cost Accumulation**: Orphaned resources continue to incur charges
- **Resource Quota Issues**: May hit AWS service limits
- **Environment Pollution**: Test resources mixing with production resources

## Solution

### 1. Resource Tagging Strategy

All test resources are tagged with consistent metadata for identification and automated cleanup:

```hcl
tags = {
  "Purpose"     = "terratest"
  "TestRun"     = "unique-test-id"
  "TestName"    = "TestTerraformECRPublicBasic"
  "CreatedAt"   = "2024-01-01T12:00:00Z"
  "CreatedBy"   = "terraform-aws-ecrpublic-tests"
  "Environment" = "test"
  "ManagedBy"   = "terratest"
}
```

### 2. Automated Cleanup Script

The `cleanup-orphaned-resources.sh` script can identify and remove orphaned resources:

#### Basic Usage
```bash
# Dry run (safe, shows what would be deleted)
./test/cleanup-orphaned-resources.sh

# Clean resources older than 6 hours (default)
./test/cleanup-orphaned-resources.sh us-east-1 6 false

# Clean resources older than 1 hour  
./test/cleanup-orphaned-resources.sh us-east-1 1 false
```

#### Script Features
- **Safe by Default**: Runs in dry-run mode unless explicitly disabled
- **Age-Based Filtering**: Only removes resources older than specified threshold
- **Tag-Based Identification**: Only targets resources with `Purpose=terratest` tag
- **Comprehensive Logging**: Shows what resources are found and processed
- **Error Handling**: Graceful handling of API failures and edge cases

### 3. Test Integration

#### Automatic Tagging
All tests automatically include cleanup tags through the `generateTestTags()` helper:

```go
terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../",
    Vars: map[string]interface{}{
        "repository_name": repositoryName,
        "tags":           generateTestTags("TestTerraformECRPublicBasic", uniqueID),
    },
    EnvVars: map[string]string{
        "AWS_DEFAULT_REGION": awsRegion,
    },
})
```

#### Enhanced Error Recovery
Defer statements include comprehensive error logging for manual cleanup:

```go
defer func() {
    if err := terraform.DestroyE(t, terraformOptions); err != nil {
        t.Logf("Warning: Failed to destroy resources: %v", err)
        t.Logf("Manual cleanup needed for repository: %s", repositoryName)
        t.Logf("Use cleanup script: ./test/cleanup-orphaned-resources.sh")
    }
}()
```

## Usage in CI/CD

### GitHub Actions Integration
Add cleanup steps to your workflows:

```yaml
- name: Cleanup Orphaned Resources
  if: always()  # Run even if tests fail
  run: |
    # Clean up resources older than 2 hours
    ./test/cleanup-orphaned-resources.sh us-east-1 2 false
  env:
    AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
    AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
```

### Scheduled Cleanup
Set up periodic cleanup jobs:

```yaml
name: Cleanup Orphaned Resources
on:
  schedule:
    - cron: '0 */6 * * *'  # Every 6 hours
jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Cleanup Resources
        run: |
          chmod +x ./test/cleanup-orphaned-resources.sh
          ./test/cleanup-orphaned-resources.sh us-east-1 6 false
```

## Manual Cleanup

### Find Orphaned Resources
```bash
# List all ECR Public repositories
aws ecr-public describe-repositories --region us-east-1

# List repositories with terratest tags
aws ecr-public describe-repositories --region us-east-1 --query 'repositories[?contains(keys(@), `repositoryName`)]' | jq '.[] | select(.repositoryName | contains("terratest"))'
```

### Clean Specific Resources
```bash
# Delete specific repository
aws ecr-public delete-repository --region us-east-1 --repository-name "terratest-repo-xyz" --force

# Get repository tags for verification
aws ecr-public list-tags-for-resource --region us-east-1 --resource-arn "arn:aws:ecr-public::123456789012:repository/repo-name"
```

## Best Practices

### For Test Development
1. **Always Use Unique IDs**: Ensure repository names include unique identifiers
2. **Test Cleanup Locally**: Run cleanup script periodically during development
3. **Monitor Resource Usage**: Check AWS console for orphaned resources
4. **Use Short Timeouts**: Set reasonable test timeouts to prevent hanging resources

### For CI/CD Pipelines
1. **Post-Test Cleanup**: Always run cleanup after test suites
2. **Timeout Configuration**: Set aggressive timeouts for test execution
3. **Resource Monitoring**: Monitor AWS costs and usage patterns
4. **Alert on Failures**: Set up notifications for cleanup script failures

### For Production Usage
1. **Separate Accounts**: Use dedicated AWS accounts for testing
2. **Cost Monitoring**: Set up billing alerts for unexpected charges
3. **Resource Limits**: Implement AWS service quotas to prevent runaway usage
4. **Regular Audits**: Periodic review of test resource usage patterns

## Troubleshooting

### Cleanup Script Issues
```bash
# Enable debug output
export AWS_CLI_FILE_ENCODING=UTF-8
export AWS_DEFAULT_OUTPUT=json

# Test AWS permissions
aws sts get-caller-identity
aws ecr-public describe-repositories --region us-east-1 --max-items 1
```

### Common Problems
- **Permission Denied**: Ensure AWS credentials have ECR Public permissions
- **Region Issues**: ECR Public is only available in us-east-1
- **Tag Filtering**: Verify repository has correct `Purpose=terratest` tag
- **Age Calculation**: Check timestamp format and timezone handling

### Manual Recovery
If automated cleanup fails, use AWS CLI or Console:

1. List all repositories: `aws ecr-public describe-repositories --region us-east-1`
2. Identify test repositories by name pattern (contains "terratest")
3. Check tags: `aws ecr-public list-tags-for-resource --resource-arn <arn>`
4. Delete repository: `aws ecr-public delete-repository --repository-name <name> --force`

## Security Considerations

- **Least Privilege**: Cleanup scripts use minimal required permissions
- **Tag Validation**: Only resources with specific tags are eligible for cleanup
- **Age Verification**: Resources must be older than threshold to prevent accidental deletion
- **Dry Run Default**: Scripts default to safe dry-run mode
- **Audit Logging**: All cleanup actions are logged for audit trails

## Cost Management

The cleanup mechanisms help control costs by:
- **Preventing Resource Accumulation**: Removes orphaned resources promptly
- **Automated Scheduling**: Regular cleanup reduces manual intervention
- **Age-Based Deletion**: Removes only old, likely orphaned resources
- **Comprehensive Coverage**: Handles all ECR Public repository types created by tests

This approach ensures reliable test environments while minimizing AWS costs and resource pollution.