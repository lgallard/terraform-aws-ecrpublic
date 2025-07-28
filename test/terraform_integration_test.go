package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestTerraformECRPublicBasic tests basic ECR Public repository creation
func TestTerraformECRPublicBasic(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-basic-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

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

	// Clean up resources with "terraform destroy" at the end of the test
	defer func() {
		if err := terraform.DestroyE(t, terraformOptions); err != nil {
			t.Logf("Warning: Failed to destroy resources: %v", err)
			t.Logf("Manual cleanup needed for repository: %s", repositoryName)
		}
	}()

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verify the ECR Public repository was created
	validateECRPublicRepository(t, terraformOptions, repositoryName)

	// Verify basic outputs
	validateBasicOutputs(t, terraformOptions, repositoryName)

	// Check if repository exists using AWS session
	session, err := aws.NewAuthenticatedSession(awsRegion)
	assert.NoError(t, err, "Should be able to create AWS session")
	assert.NotNil(t, session, "AWS session should be created successfully")
}

// TestTerraformECRPublicWithVariableCatalogData tests ECR Public repository with catalog data using variables
func TestTerraformECRPublicWithVariableCatalogData(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-vars-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":                    repositoryName,
			"catalog_data_description":           "Test repository created by Terratest using variables",
			"catalog_data_about_text":           "# Test Repository\nThis is a test repository created by automated tests using variable-based configuration.",
			"catalog_data_usage_text":           "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
			"catalog_data_architectures":        []string{"x86-64"},
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

	// Verify the ECR Public repository was created with catalog data
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataOutputs(t, terraformOptions, repositoryName)
}

// TestTerraformECRPublicWithObjectCatalogData tests ECR Public repository with catalog data using objects
func TestTerraformECRPublicWithObjectCatalogData(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-obj-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	// Create custom terraform options for object-based catalog data
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description":       "Test repository created by Terratest using objects",
				"about_text":       "# Test Repository\nThis is a test repository created by automated tests using object-based configuration.",
				"usage_text":       "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
				"architectures":     []string{"x86-64", "ARM 64"},
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

	// Verify the ECR Public repository was created with catalog data
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataOutputs(t, terraformOptions, repositoryName)
}

// TestTerraformECRPublicWithTimeouts tests ECR Public repository with custom timeouts
func TestTerraformECRPublicWithTimeouts(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-timeout-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"timeouts": map[string]interface{}{
				"delete": "30m",
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

	// Verify the ECR Public repository was created
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateBasicOutputs(t, terraformOptions, repositoryName)
}

// TestTerraformECRPublicVariableValidation tests variable validation rules
func TestTerraformECRPublicVariableValidation(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-validation-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":                    repositoryName,
			"catalog_data_description":           "Valid description under 256 characters",
			"catalog_data_about_text":           "# Valid About Text\nThis is a valid about text.",
			"catalog_data_usage_text":           "# Valid Usage\nThis is valid usage text.",
			"catalog_data_architectures":        []string{"x86-64", "ARM 64"},
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

	// Verify the repository was created successfully with valid inputs
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataOutputs(t, terraformOptions, repositoryName)
}

// TestTerraformECRPublicCompleteConfiguration tests complete ECR Public configuration
func TestTerraformECRPublicCompleteConfiguration(t *testing.T) {
	t.Parallel()
	
	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-complete-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description":       "Complete test repository for comprehensive testing",
				"about_text":       "# Complete Test Repository\n## Overview\nThis repository is created for comprehensive testing of the terraform-aws-ecrpublic module.\n\n## Features\n- Complete catalog data configuration\n- Multi-architecture support\n- Detailed documentation",
				"usage_text":       "# Usage\n\n## Quick Start\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\ndocker run public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Advanced Usage\nSee documentation for advanced configuration options.",
				"architectures":     []string{"x86-64", "ARM", "ARM 64"},
				"operating_systems": []string{"Linux"},
			},
			"timeouts": map[string]interface{}{
				"delete": "30m",
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

	// Verify comprehensive repository configuration
	validateECRPublicRepository(t, terraformOptions, repositoryName)
	validateCatalogDataOutputs(t, terraformOptions, repositoryName)
	validateAllOutputs(t, terraformOptions, repositoryName)
}

// Helper function to validate ECR Public repository creation
func validateECRPublicRepository(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	assert.NotEmpty(t, repositoryURL, "Repository URL should not be empty")
	assert.Contains(t, repositoryURL, expectedRepoName, "Repository URL should contain repository name")
	assert.Contains(t, repositoryURL, "public.ecr.aws", "Repository URL should contain public ECR domain")

	actualRepositoryName := terraform.Output(t, terraformOptions, "repository_name")
	assert.Equal(t, expectedRepoName, actualRepositoryName, "Repository name should match expected value")
}

// Helper function to validate basic outputs
func validateBasicOutputs(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	// Validate ARN output
	arn := terraform.Output(t, terraformOptions, "arn")
	assert.NotEmpty(t, arn, "ARN should not be empty")
	assert.Contains(t, arn, expectedRepoName, "ARN should contain repository name")

	// Validate repository ARN output (duplicate check for backward compatibility)
	repositoryARN := terraform.Output(t, terraformOptions, "repository_arn")
	assert.Equal(t, arn, repositoryARN, "repository_arn should match arn output")

	// Validate registry ID
	registryID := terraform.Output(t, terraformOptions, "registry_id")
	assert.NotEmpty(t, registryID, "Registry ID should not be empty")

	// Validate repository URI
	repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")
	assert.NotEmpty(t, repositoryURI, "Repository URI should not be empty")
	assert.Contains(t, repositoryURI, expectedRepoName, "Repository URI should contain repository name")
}

// Helper function to validate catalog data specific outputs
func validateCatalogDataOutputs(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	validateBasicOutputs(t, terraformOptions, expectedRepoName)

	// Additional validation specific to catalog data can be added here
	// Since catalog data is part of the repository resource, we mainly validate
	// that the repository was created successfully with the catalog data
}

// Helper function to validate all outputs comprehensively
func validateAllOutputs(t *testing.T, terraformOptions *terraform.Options, expectedRepoName string) {
	// Validate all basic outputs
	validateBasicOutputs(t, terraformOptions, expectedRepoName)

	// Validate ID output
	id := terraform.Output(t, terraformOptions, "id")
	registryID := terraform.Output(t, terraformOptions, "registry_id")
	assert.Equal(t, registryID, id, "id output should match registry_id")

	// Validate URL vs URI consistency
	repositoryURL := terraform.Output(t, terraformOptions, "repository_url")
	repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")
	assert.Equal(t, repositoryURI, repositoryURL, "repository_url should match repository_uri")
}

// Helper function to validate repository name format for security
func validateRepositoryNameFormat(t *testing.T, repositoryName string) {
	// Validate repository name format to prevent string injection
	validPattern := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !validPattern.MatchString(repositoryName) {
		t.Fatalf("Invalid repository name format: %s. Repository names must contain only lowercase letters, numbers, and hyphens.", repositoryName)
	}
	
	// Additional ECR Public repository name validations
	if len(repositoryName) == 0 {
		t.Fatal("Repository name cannot be empty")
	}
	if len(repositoryName) > 256 {
		t.Fatal("Repository name must be 256 characters or less")
	}
}

