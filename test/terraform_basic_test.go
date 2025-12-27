package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformValidate(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	// Initialize and validate the Terraform configuration
	terraform.Init(t, terraformOptions)
	terraform.Validate(t, terraformOptions)
}

func TestTerraformFormat(t *testing.T) {
	t.Parallel()

	// Check if all Terraform files are properly formatted
	_, err := terraform.RunTerraformCommandE(t, &terraform.Options{
		TerraformDir: "../",
	}, "fmt", "-check", "-recursive")

	assert.NoError(t, err, "Terraform files should be properly formatted")
}

func TestExamplesValidate(t *testing.T) {
	t.Parallel()

	examples := []string{
		"../examples/using_objects",
		"../examples/using_variables",
		"../examples/with_repository_policy",
	}

	for _, example := range examples {
		t.Run(filepath.Base(example), func(t *testing.T) {
			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: example,
			})

			terraform.Init(t, terraformOptions)
			terraform.Validate(t, terraformOptions)
		})
	}
}

func TestExamplesFormat(t *testing.T) {
	t.Parallel()

	examples := []string{
		"../examples/using_objects",
		"../examples/using_variables",
		"../examples/with_repository_policy",
	}

	for _, example := range examples {
		t.Run(filepath.Base(example), func(t *testing.T) {
			_, err := terraform.RunTerraformCommandE(t, &terraform.Options{
				TerraformDir: example,
			}, "fmt", "-check", "-recursive")

			assert.NoError(t, err, "Example %s should be properly formatted", example)
		})
	}
}
