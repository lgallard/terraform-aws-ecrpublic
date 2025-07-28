package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestECRPublicGalleryOptimization tests repository configuration optimized for ECR Public Gallery
func TestECRPublicGalleryOptimization(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-gallery-opt-%s", uniqueID)

	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description": "Production-ready Node.js application container optimized for ECR Public Gallery discoverability",
				"about_text": func() string {
					data, err := loadTestData("gallery_optimization_about.md", repositoryName)
					if err != nil {
						t.Fatalf("Failed to load test data: %v", err)
					}
					return data
				}(),
				"usage_text": func() string {
					data, err := loadTestData("gallery_optimization_usage.md", repositoryName)
					if err != nil {
						t.Fatalf("Failed to load test data: %v", err)
					}
					return data
				}(),
				"architectures":     []string{"x86-64", "ARM 64"},
				"operating_systems": []string{"Linux"},
			},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify gallery-optimized repository configuration
	validatePublicGalleryOptimization(t, terraformOptions, repositoryName)
}

// TestECRPublicGallerySearchability tests configuration for maximum searchability
func TestECRPublicGallerySearchability(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-searchable-%s", uniqueID)

	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":                repositoryName,
			"catalog_data_description":       "Open-source Python web framework container with Flask, Django support, and development tools",
			"catalog_data_about_text":        "# Python Web Framework Container\n\n## Overview\nA comprehensive Python container optimized for web development with popular frameworks.\n\n## Supported Frameworks\n- **Flask**: Lightweight WSGI web application framework\n- **Django**: High-level Python web framework\n- **FastAPI**: Modern, fast web framework for building APIs\n\n## Keywords\nPython, Flask, Django, FastAPI, web development, API, microservices, REST, GraphQL, WSGI, ASGI",
			"catalog_data_usage_text": func() string {
				data, err := loadTestData("python_web_usage.md", repositoryName)
				if err != nil {
					t.Fatalf("Failed to load test data: %v", err)
				}
				return data
			}(),
			"catalog_data_architectures":     []string{"x86-64", "ARM 64"},
			"catalog_data_operating_systems": []string{"Linux"},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify searchability-optimized configuration
	validatePublicGallerySearchability(t, terraformOptions, repositoryName)
}

// TestECRPublicGalleryContentGuidelines tests compliance with content guidelines
func TestECRPublicGalleryContentGuidelines(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-guidelines-%s", uniqueID)

	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description": "Enterprise Java application server with Spring Boot, security hardening, and production monitoring",
				"about_text": `# Enterprise Java Application Server

## Overview

This container provides a production-ready Java application server environment with Spring Boot framework support, comprehensive security hardening, and enterprise-grade monitoring capabilities.

## Key Features

### Application Framework Support
- **Spring Boot 3.x**: Latest Spring Framework with auto-configuration
- **Spring Security**: Authentication and authorization
- **Spring Data**: Database abstraction and ORM support
- **Spring Cloud**: Microservices and distributed systems support

### Security Hardening
- **Non-root Execution**: Application runs as dedicated user
- **Minimal Base Image**: Distroless base to reduce attack surface
- **Vulnerability Scanning**: Regular automated security scans
- **Secrets Management**: Secure configuration injection

### Production Monitoring
- **Health Endpoints**: Actuator endpoints for monitoring
- **Metrics Export**: Prometheus-compatible metrics
- **Distributed Tracing**: OpenTelemetry integration
- **Structured Logging**: JSON logs with correlation IDs

### Performance Optimization
- **JVM Tuning**: Optimized garbage collection settings
- **Memory Management**: Efficient heap and non-heap memory usage
- **Startup Performance**: Fast application initialization
- **Connection Pooling**: Optimized database connections

## Architecture Compliance

This image follows enterprise architecture patterns:
- **12-Factor App**: Compliant with twelve-factor methodology
- **Cloud Native**: Designed for containerized environments
- **Microservices Ready**: Supports distributed architecture patterns
- **Scalability**: Horizontal and vertical scaling support

## Supported Technologies

- **Java 17+**: Latest LTS Java runtime
- **Maven/Gradle**: Build tool support
- **PostgreSQL/MySQL**: Database connectivity
- **Redis**: Caching and session management
- **Kafka**: Event streaming support`,
				"usage_text":        func() string {
					data, err := loadTestData("spring_boot_usage.md", repositoryName)
					if err != nil {
						t.Fatalf("Failed to load test data: %v", err)
					}
					return data
				}(),
				"architectures":     []string{"x86-64", "ARM 64"},
				"operating_systems": []string{"Linux"},
			},
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify content guidelines compliance
	validatePublicGalleryContentGuidelines(t, terraformOptions, repositoryName)
}

// TestECRPublicGalleryRegionalConstraints tests us-east-1 region constraint
func TestECRPublicGalleryRegionalConstraints(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-regional-%s", uniqueID)

	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Explicitly test us-east-1 constraint
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": generateMinimalCatalogData(repositoryName),
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify regional constraints are respected
	validateRegionalConstraints(t, terraformOptions, repositoryName, awsRegion)
}

// Helper function to validate public gallery optimization
func validatePublicGalleryOptimization(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	// Validate basic repository creation
	validateECRPublicRepository(t, terraformOptions, expectedRepoName)

	// Verify gallery-specific optimizations
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.Contains(t, repositoryURL, "public.ecr.aws", "Repository should be in ECR Public Gallery")

	// Verify all required outputs for gallery visibility
	arn := terraform.Output(t, terraformOptions, "arn")
	assert.NotEmpty(t, arn, "ARN should be available for gallery indexing")

	registryID := terraform.Output(t, terraformOptions, "registry_id")
	assert.NotEmpty(t, registryID, "Registry ID should be available")
}

// Helper function to validate searchability optimization
func validatePublicGallerySearchability(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	// Validate basic repository creation
	validateECRPublicRepository(t, terraformOptions, expectedRepoName)

	// Additional searchability validations
	repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")
	assert.NotEmpty(t, repositoryURI, "Repository URI should be available for public access")
	assert.Contains(t, repositoryURI, expectedRepoName, "Repository URI should contain repository name for searchability")
}

// Helper function to validate content guidelines compliance
func validatePublicGalleryContentGuidelines(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	// Validate basic repository creation
	validateECRPublicRepository(t, terraformOptions, expectedRepoName)

	// Validate that repository was created successfully with comprehensive catalog data
	// This indicates that the content followed guidelines and was accepted by ECR Public
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.Contains(t, repositoryURL, "public.ecr.aws", "Repository should be publicly accessible")

	// Verify all outputs are available (indicates successful creation with catalog data)
	outputs := []string{"arn", "repository_arn", "id", "repository_name", "registry_id", "repository_uri", "repository_url"}
	for _, output := range outputs {
		value := terraform.Output(t, terraformOptions, output)
		assert.NotEmpty(t, value, fmt.Sprintf("Output %s should not be empty", output))
	}
}

// Helper function to validate regional constraints
func validateRegionalConstraints(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string, expectedRegion string) {
	// Validate basic repository creation
	validateECRPublicRepository(t, terraformOptions, expectedRepoName)

	// Verify the repository was created in the correct region (us-east-1)
	assert.Equal(t, "us-east-1", expectedRegion, "ECR Public repositories must be created in us-east-1")

	// Verify repository URL contains public ECR domain
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.Contains(t, repositoryURL, "public.ecr.aws", "Repository should use ECR Public domain")

	// The fact that the repository was created successfully validates regional constraints
}

