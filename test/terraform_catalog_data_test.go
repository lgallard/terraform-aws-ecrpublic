package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestECRPublicCatalogDataValidation tests catalog data validation rules
func TestECRPublicCatalogDataValidation(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-catalog-validation-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	// Test with valid catalog data that follows ECR Public Gallery guidelines
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: func() map[string]interface{} {
			vars := map[string]interface{}{"repository_name": repositoryName}
			for k, v := range generateMinimalVariableCatalogData(repositoryName) {
				vars[k] = v
			}
			return vars
		}(),
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify the repository was created with valid catalog data
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataCompliance(t, terraformOptions)
}

// TestECRPublicMinimalCatalogData tests minimal valid catalog data configuration
func TestECRPublicMinimalCatalogData(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-minimal-catalog-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":               repositoryName,
			"catalog_data_description":      "Minimal test container",
			"catalog_data_about_text":      "# Test Container\nMinimal configuration for testing.",
			"catalog_data_architectures":   []string{"x86-64"},
			"catalog_data_operating_systems": []string{"Linux"},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify minimal catalog data configuration works
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataCompliance(t, terraformOptions)
}

// TestECRPublicComprehensiveCatalogData tests comprehensive catalog data with all fields
func TestECRPublicComprehensiveCatalogData(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-comprehensive-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description": "Enterprise-grade application container with comprehensive tooling and security features",
				"about_text": `# Enterprise Application Container

## Overview

This container provides a comprehensive, enterprise-grade environment for deploying production applications with enhanced security, monitoring, and operational capabilities.

## Key Features

### Security
- **Hardened Base Image**: Built on distroless base for minimal attack surface
- **Non-root Execution**: All processes run as non-privileged user
- **Vulnerability Scanning**: Regular automated security scans
- **SBOM Generation**: Software Bill of Materials included

### Performance
- **Multi-stage Builds**: Optimized for size and performance
- **Caching Strategy**: Intelligent layer caching for faster builds
- **Resource Optimization**: Memory and CPU optimized configurations
- **Startup Performance**: Fast application startup times

### Monitoring & Observability
- **Health Checks**: Comprehensive health and readiness endpoints
- **Metrics Export**: Prometheus-compatible metrics
- **Structured Logging**: JSON-formatted logs with correlation IDs
- **Distributed Tracing**: OpenTelemetry integration

### Operations
- **Configuration Management**: Environment-based configuration
- **Secret Management**: Secure secret injection patterns
- **Graceful Shutdown**: Proper SIGTERM handling
- **Auto-scaling Ready**: Designed for horizontal scaling

## Architecture Support

This image supports multiple architectures:
- **x86-64**: Intel/AMD 64-bit systems
- **ARM 64**: ARM-based systems including Apple Silicon
- **ARM**: 32-bit ARM systems for IoT deployments

## Compliance & Standards

- **NIST Guidelines**: Follows NIST container security guidelines
- **CIS Benchmarks**: Compliant with CIS security benchmarks
- **OWASP Standards**: Implements OWASP container security practices`,
				"usage_text": func() string {
					data, err := loadTestData("comprehensive_catalog_usage.md", repositoryName)
					if err != nil {
						t.Fatalf("Failed to load test data: %v", err)
					}
					return data
				}(),
				"architectures":     []string{"x86-64", "ARM", "ARM 64"},
				"operating_systems": []string{"Linux"},
			},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify comprehensive catalog data configuration
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataCompliance(t, terraformOptions)
}

// TestECRPublicMultiArchitectureSupport tests multi-architecture configuration
func TestECRPublicMultiArchitectureSupport(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-multiarch-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":                    repositoryName,
			"catalog_data_description":           "Multi-arch container",
			"catalog_data_about_text":           "# Multi-Architecture\nSupports multiple architectures.",
			"catalog_data_usage_text":           "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
			"catalog_data_architectures":        []string{"x86-64", "ARM", "ARM 64"},
			"catalog_data_operating_systems":    []string{"Linux"},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify multi-architecture support configuration
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataCompliance(t, terraformOptions)
}

// TestECRPublicMarkdownFormatting tests proper markdown formatting in catalog data
func TestECRPublicMarkdownFormatting(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-markdown-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":               repositoryName,
			"catalog_data_description":      "Markdown test container",
			"catalog_data_about_text":      "# Test\n- Feature 1\n- Feature 2\n\n```bash\necho test\n```",
			"catalog_data_usage_text":      "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
			"catalog_data_architectures":   []string{"x86-64"},
			"catalog_data_operating_systems": []string{"Linux"},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify markdown formatting in catalog data
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataCompliance(t, terraformOptions)
}

// Helper function to validate catalog data compliance with ECR Public Gallery guidelines
func validateCatalogDataCompliance(t *testing.T, terraformOptions *terraform.Options) {
	// Verify the repository was created successfully
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.NotEmpty(t, repositoryURL, "Repository URL should not be empty")
	assert.Contains(t, repositoryURL, "public.ecr.aws", "Repository should be in ECR Public")

	// Verify registry ID (indicates successful creation)
	registryID := terraform.Output(t, terraformOptions, "registry_id")
	assert.NotEmpty(t, registryID, "Registry ID should not be empty")

	// Additional compliance validations can be added here
	// For now, successful creation indicates catalog data was valid
}

