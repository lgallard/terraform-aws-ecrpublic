# Testing terraform-aws-ecrpublic

This directory contains comprehensive automated tests for the terraform-aws-ecrpublic module. The tests use [Terratest](https://github.com/gruntwork-io/terratest), a Go library that provides utilities for testing Terraform code, following ECR Public-specific testing patterns and validation approaches.

## Prerequisites

1. **Go** (version 1.21 or later)
2. **Terraform** (version 1.0 or later)
3. **AWS credentials** configured with ECR Public permissions
4. **AWS Region**: Tests must run in `us-east-1` (ECR Public constraint)

### Required AWS Permissions

Your AWS credentials need the following permissions for ECR Public:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr-public:CreateRepository",
                "ecr-public:DeleteRepository",
                "ecr-public:DescribeRepositories",
                "ecr-public:GetRepositoryCatalogData",
                "ecr-public:PutRepositoryCatalogData"
            ],
            "Resource": "*"
        }
    ]
}
```

## Test Structure

The test suite is organized into focused test files, each covering specific aspects of the terraform-aws-ecrpublic module:

### 1. Basic Tests (`terraform_basic_test.go`)
- **Static validation tests**: Format checking and Terraform validation
- **Example validation**: Ensures all examples are valid and properly formatted
- **No AWS resources created**

### 2. Integration Tests (`terraform_integration_test.go`)
- **Basic repository creation**: Tests minimal ECR Public repository setup
- **Variable-based catalog data**: Tests catalog data using individual variables
- **Object-based catalog data**: Tests catalog data using object configuration
- **Timeout configuration**: Tests custom timeout settings
- **Variable validation**: Tests input validation rules
- **Complete configuration**: Tests comprehensive setup with all features

### 3. Catalog Data Tests (`terraform_catalog_data_test.go`)
- **Catalog data validation**: Tests ECR Public Gallery content guidelines
- **Minimal configuration**: Tests minimal valid catalog data
- **Comprehensive configuration**: Tests all catalog data fields
- **Multi-architecture support**: Tests architecture and OS tagging
- **Markdown formatting**: Tests proper markdown in catalog data fields

### 4. Public Gallery Tests (`terraform_public_gallery_test.go`)
- **Gallery optimization**: Tests repository configuration for maximum discoverability
- **Searchability**: Tests content optimized for ECR Public Gallery search
- **Content guidelines**: Tests compliance with public gallery content standards
- **Regional constraints**: Tests us-east-1 region requirement

## Running Tests

### Quick Start

```bash
# Run static tests only (no AWS resources)
make test

# Run basic integration tests (creates AWS resources)
make test-integration

# Run all tests including comprehensive suites
make test-all
```

### Individual Test Suites

```bash
# Static analysis only
make test-static

# Catalog data validation tests
make test-catalog

# Public gallery compliance tests
make test-gallery

# Timeout configuration tests
make test-timeouts

# Variable validation tests
make test-validation
```

### Direct Go Test Execution

```bash
cd test

# Run specific test
go test -v -run TestTerraformECRPublicBasic

# Run all tests in a file
go test -v -run TestECRPublicCatalogData

# Run with timeout and parallel execution
go test -v -timeout 30m -parallel 2
```

## Test Categories

### Static Tests (No AWS Resources)
- **Format validation**: Ensures all Terraform files are properly formatted
- **Configuration validation**: Validates Terraform syntax and structure
- **Example validation**: Checks that all examples are valid
- **Fast execution**: Completes in seconds

### Integration Tests (Creates AWS Resources)
- **Basic repository creation**: Tests core functionality
- **Catalog data configuration**: Tests both object and variable approaches
- **Output validation**: Verifies all module outputs
- **Resource cleanup**: Automatic cleanup with `terraform destroy`

### Catalog Data Tests (Advanced Validation)
- **Content guidelines compliance**: Ensures catalog data follows ECR Public standards
- **Multi-architecture support**: Tests proper architecture tagging
- **Markdown validation**: Tests rich markdown content in catalog fields
- **Comprehensive documentation**: Tests detailed about and usage text

### Public Gallery Tests (ECR Public Specific)
- **Gallery optimization**: Tests discoverability features
- **Search optimization**: Tests content for maximum searchability
- **Regional constraints**: Validates us-east-1 requirement
- **Public accessibility**: Tests public repository features

## Test Configuration

### ECR Public Specific Requirements

All tests are configured for ECR Public constraints:

- **Region**: All tests use `us-east-1` (ECR Public requirement)
- **Repository naming**: Unique names to avoid conflicts
- **Catalog data**: Tests follow ECR Public Gallery content guidelines
- **Public accessibility**: Tests validate public repository features

### Test Data Patterns

```go
// Unique repository naming
uniqueID := strings.ToLower(random.UniqueId())
repositoryName := fmt.Sprintf("terratest-prefix-%s", uniqueID)

// Region constraint
awsRegion := "us-east-1"

// Catalog data validation
terraformOptions := &terraform.Options{
    TerraformDir: "../examples/using_variables",
    Vars: map[string]interface{}{
        "repository_name": repositoryName,
        "catalog_data_description": "Valid ECR Public description",
        "catalog_data_architectures": []string{"x86-64", "ARM 64"},
        "catalog_data_operating_systems": []string{"Linux"},
    },
    EnvVars: map[string]string{
        "AWS_DEFAULT_REGION": awsRegion,
    },
}
```

## Validation Helpers

The test suite includes specialized validation functions:

### Repository Validation
- `validateECRPublicRepository()`: Basic ECR Public repository validation
- `validateBasicOutputs()`: Tests all module outputs
- `validateCatalogDataOutputs()`: Catalog data specific validation
- `validateAllOutputs()`: Comprehensive output validation

### Gallery Validation
- `validatePublicGalleryOptimization()`: Gallery discoverability checks
- `validatePublicGallerySearchability()`: Search optimization validation
- `validatePublicGalleryContentGuidelines()`: Content compliance checks
- `validateRegionalConstraints()`: us-east-1 constraint validation

## Test Examples

### Basic Repository Test
```go
func TestTerraformECRPublicBasic(t *testing.T) {
    t.Parallel()
    
    uniqueID := strings.ToLower(random.UniqueId())
    repositoryName := fmt.Sprintf("terratest-basic-%s", uniqueID)
    
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../",
        Vars: map[string]interface{}{
            "repository_name": repositoryName,
        },
        EnvVars: map[string]string{
            "AWS_DEFAULT_REGION": "us-east-1",
        },
    })
    
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    
    validateECRPublicRepository(t, terraformOptions, repositoryName)
}
```

### Catalog Data Test
```go
func TestECRPublicCatalogDataValidation(t *testing.T) {
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../examples/using_variables",
        Vars: map[string]interface{}{
            "repository_name": repositoryName,
            "catalog_data_description": "Production-ready container with security hardening",
            "catalog_data_about_text": "# Container\n## Features\n- Security hardened\n- Performance optimized",
            "catalog_data_usage_text": "# Usage\n```bash\ndocker pull public.ecr.aws/registry/repo:latest\n```",
            "catalog_data_architectures": []string{"x86-64", "ARM 64"},
            "catalog_data_operating_systems": []string{"Linux"},
        },
    })
    
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    
    validateCatalogDataCompliance(t, terraformOptions)
}
```

## AWS Resources and Costs

### Resource Creation
- **ECR Public repositories**: Tests create real ECR Public repositories
- **Catalog data**: Tests populate repositories with comprehensive metadata
- **Automatic cleanup**: All resources are destroyed after test completion

### Cost Considerations
- **ECR Public repositories**: Free to create and maintain
- **Catalog data**: No additional charges for metadata
- **Region constraint**: Tests run only in us-east-1 to comply with ECR Public requirements

### Resource Cleanup

All tests include automatic cleanup:

```go
defer terraform.Destroy(t, terraformOptions)
```

If tests fail unexpectedly, manual cleanup may be required:

```bash
# List ECR Public repositories
aws ecr-public describe-repositories --region us-east-1

# Delete specific repository
aws ecr-public delete-repository --repository-name <repository-name> --region us-east-1
```

## CI/CD Integration

### GitHub Actions
Tests integrate with GitHub Actions workflows:

```yaml
- name: Run Terratest
  run: |
    cd test
    go test -v -timeout 30m -parallel 2
  env:
    AWS_DEFAULT_REGION: us-east-1
```

### Makefile Integration
Use Makefile targets for consistent execution:

```bash
# Development workflow
make fmt-check validate test-static

# Pre-release validation
make check-all

# Specific test suites
make test-catalog test-gallery
```

## Troubleshooting

### Common Issues

**Region Errors**
```
Error: ECR Public repositories can only be created in us-east-1
Solution: Ensure AWS_DEFAULT_REGION=us-east-1
```

**Repository Name Conflicts**
```
Error: Repository already exists
Solution: Tests use unique IDs, but manual cleanup may be needed
```

**Permissions Errors**
```
Error: Access denied for ECR Public operations
Solution: Verify AWS credentials have ECR Public permissions
```

### Debug Mode

Enable debug logging:

```bash
# Verbose test output
go test -v -run TestName

# Terraform debug
TF_LOG=DEBUG go test -v -run TestName
```

### Manual Validation

Verify test results manually:

```bash
# Check repository in ECR Public
aws ecr-public describe-repositories --region us-east-1

# Get catalog data
aws ecr-public get-repository-catalog-data --repository-name <name> --region us-east-1

# Test public accessibility
curl -s https://gallery.ecr.aws/
```

## Contributing

### Adding New Tests

1. **Follow naming patterns**: Use descriptive test function names
2. **Include cleanup**: Always use `defer terraform.Destroy()`
3. **Use unique names**: Generate unique repository names
4. **Validate thoroughly**: Use helper functions for consistent validation
5. **Document purpose**: Add clear comments explaining test objectives

### Test Guidelines

- **Parallel execution**: Use `t.Parallel()` for independent tests
- **Timeout management**: Set appropriate timeouts for AWS operations
- **Error handling**: Provide clear error messages with context
- **Resource isolation**: Ensure tests don't interfere with each other

### Example Test Template

```go
func TestECRPublicNewFeature(t *testing.T) {
    t.Parallel()
    
    // Generate unique repository name
    uniqueID := strings.ToLower(random.UniqueId())
    repositoryName := fmt.Sprintf("terratest-feature-%s", uniqueID)
    
    // Configure terraform options
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../",
        Vars: map[string]interface{}{
            "repository_name": repositoryName,
            // Add test-specific variables
        },
        EnvVars: map[string]string{
            "AWS_DEFAULT_REGION": "us-east-1",
        },
    })
    
    // Ensure cleanup
    defer terraform.Destroy(t, terraformOptions)
    
    // Execute terraform
    terraform.InitAndApply(t, terraformOptions)
    
    // Validate results
    validateECRPublicRepository(t, terraformOptions, repositoryName)
    // Add feature-specific validations
}
```

This testing framework provides comprehensive coverage of ECR Public functionality while following best practices for Terraform module testing and ECR Public Gallery compliance.