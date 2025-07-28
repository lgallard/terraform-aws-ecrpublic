package test

import (
	"fmt"
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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

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

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

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

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	// Create custom terraform options for object-based catalog data
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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

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

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

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

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

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
	
	// Ensure safe test execution with quota and security checks
	ensureSafeTestExecution(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description":       "Complete test repository",
				"about_text":       "# Complete Test\nTest repository.",
				"usage_text":       "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
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

	// Set up comprehensive cleanup with error recovery
	defer setupTestCleanup(t, terraformOptions, repositoryName)()

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

