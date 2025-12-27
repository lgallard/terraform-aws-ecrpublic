package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryPolicyExample(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name for this test
	repositoryName := fmt.Sprintf("test-repo-policy-%s", strings.ToLower(random.UniqueId()))

	// Example CI/CD role ARN for testing (using a realistic but non-existent ARN)
	testRoleArn := "arn:aws:iam::123456789012:role/GitHubActionsRole"

	// Policy document for CI/CD push access
	policyDocument := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Sid": "AllowPushFromCICD",
				"Effect": "Allow",
				"Principal": {
					"AWS": "%s"
				},
				"Action": [
					"ecr-public:BatchCheckLayerAvailability",
					"ecr-public:PutImage",
					"ecr-public:InitiateLayerUpload",
					"ecr-public:UploadLayerPart",
					"ecr-public:CompleteLayerUpload"
				]
			}
		]
	}`, testRoleArn)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/with_repository_policy",
		Vars: map[string]interface{}{
			"repository_name":           repositoryName,
			"create_repository_policy":  true,
			"repository_policy":         policyDocument,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")
	outputRepositoryName := terraform.Output(t, terraformOptions, "repository_name")
	registryID := terraform.Output(t, terraformOptions, "registry_id")
	repositoryArn := terraform.Output(t, terraformOptions, "repository_arn")
	outputRepositoryPolicy := terraform.Output(t, terraformOptions, "repository_policy")

	// Verify repository name
	assert.Equal(t, repositoryName, outputRepositoryName)

	// Verify repository URI format for ECR Public
	assert.Contains(t, repositoryURI, "public.ecr.aws")
	assert.Contains(t, repositoryURI, repositoryName)

	// Verify registry ID is not empty
	assert.NotEmpty(t, registryID)

	// Verify ARN format
	assert.Contains(t, repositoryArn, "arn:aws:ecr-public")
	assert.Contains(t, repositoryArn, repositoryName)

	// Verify repository policy was applied
	assert.NotEmpty(t, outputRepositoryPolicy)
	assert.Contains(t, outputRepositoryPolicy, testRoleArn)
	assert.Contains(t, outputRepositoryPolicy, "ecr-public:PutImage")
}

func TestRepositoryPolicyValidation(t *testing.T) {
	t.Parallel()

	repositoryName := fmt.Sprintf("test-repo-validation-%s", strings.ToLower(random.UniqueId()))

	// Test with invalid JSON policy
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/with_repository_policy",
		Vars: map[string]interface{}{
			"repository_name":           repositoryName,
			"create_repository_policy":  true,
			"repository_policy":         "invalid json",
		},
	})

	// This should fail due to invalid JSON in policy
	_, err := terraform.InitAndApplyE(t, terraformOptions)
	assert.Error(t, err, "Should fail with invalid JSON policy")

	// Clean up any partial resources (if any were created)
	terraform.Destroy(t, terraformOptions)
}

func TestRepositoryWithoutPolicy(t *testing.T) {
	t.Parallel()

	repositoryName := fmt.Sprintf("test-repo-no-policy-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/with_repository_policy",
		Vars: map[string]interface{}{
			"repository_name":          repositoryName,
			"create_repository_policy": false,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// This should succeed without creating a policy
	terraform.InitAndApply(t, terraformOptions)

	// Verify outputs
	repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")
	outputRepositoryName := terraform.Output(t, terraformOptions, "repository_name")
	outputRepositoryPolicy := terraform.Output(t, terraformOptions, "repository_policy")

	// Verify basic repository creation
	assert.Equal(t, repositoryName, outputRepositoryName)
	assert.Contains(t, repositoryURI, "public.ecr.aws")

	// Verify no policy was created
	assert.Empty(t, outputRepositoryPolicy)
}
