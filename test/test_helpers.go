package test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// generateTestTags creates common tags for resource cleanup and tracking
// This helps identify and clean up orphaned resources from interrupted test runs
func generateTestTags(testName, uniqueID string) map[string]string {
	return map[string]string{
		"Purpose":     "terratest",
		"TestRun":     uniqueID,
		"TestName":    testName,
		"CreatedAt":   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		"CreatedBy":   "terraform-aws-ecrpublic-tests",
		"Environment": "test",
		"ManagedBy":   "terratest",
	}
}

// generateMinimalCatalogData creates minimal valid catalog data for basic functionality tests
// This reduces test execution time and AWS API usage for tests that don't need comprehensive data
func generateMinimalCatalogData(repositoryName string) map[string]interface{} {
	// Note: repositoryName validation is handled by the calling function via ensureSafeTestExecution
	return map[string]interface{}{
		"description": "Test container",
		"about_text":  "# Test\nBasic test container.",
		"usage_text":  fmt.Sprintf("# Usage\n```bash\ndocker pull public.ecr.aws/registry/%s:latest\n```", repositoryName),
		"architectures": []string{"x86-64"},
		"operating_systems": []string{"Linux"},
	}
}

// generateMinimalVariableCatalogData creates minimal catalog data using individual variables
// This is the variable-based equivalent of generateMinimalCatalogData
func generateMinimalVariableCatalogData(repositoryName string) map[string]interface{} {
	// Note: repositoryName validation is handled by the calling function via ensureSafeTestExecution
	return map[string]interface{}{
		"catalog_data_description": "Test container",
		"catalog_data_about_text":  "# Test\nBasic test container.",
		"catalog_data_usage_text":  fmt.Sprintf("# Usage\n```bash\ndocker pull public.ecr.aws/registry/%s:latest\n```", repositoryName),
		"catalog_data_architectures": []string{"x86-64"},
		"catalog_data_operating_systems": []string{"Linux"},
	}
}

// loadTestData loads test data from external file and replaces placeholders
// This reduces memory usage by avoiding large embedded strings in test files
func loadTestData(filename, repositoryName string) (string, error) {
	filePath := filepath.Join("testdata", filename)

	// Check file size before loading to prevent memory issues
	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	// Limit test data files to 1MB to prevent excessive memory usage
	const maxFileSize = 1024 * 1024 // 1MB
	if info.Size() > maxFileSize {
		return "", fmt.Errorf("test data file %s is too large (%d bytes, max %d bytes)", filename, info.Size(), maxFileSize)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Replace placeholder with actual repository name (validation is handled by calling function)
	return strings.ReplaceAll(string(content), "{{REPOSITORY_NAME}}", repositoryName), nil
}

// validateRepositoryNameFormat validates repository name format for security
// This helper prevents string injection and ensures ECR Public naming compliance
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

	// Check for potentially dangerous patterns
	if strings.Contains(repositoryName, "..") {
		t.Fatal("Repository name cannot contain '..' patterns")
	}
	if strings.HasPrefix(repositoryName, "-") || strings.HasSuffix(repositoryName, "-") {
		t.Fatal("Repository name cannot start or end with hyphens")
	}
}

// sanitizeRepositoryName safely formats repository name for string interpolation
// This prevents injection attacks by validating and escaping the repository name
func sanitizeRepositoryName(t *testing.T, repositoryName string) string {
	// First validate the repository name format
	validateRepositoryNameFormat(t, repositoryName)

	// Return the validated name (since it passed validation, it's safe to use)
	return repositoryName
}

// checkECRPublicQuota checks current ECR Public repository usage against AWS limits
// This helps prevent quota exhaustion during parallel test execution
func checkECRPublicQuota(t *testing.T) {
	// Skip quota check if running in CI or if AWS_SKIP_QUOTA_CHECK is set
	if os.Getenv("CI") != "" || os.Getenv("AWS_SKIP_QUOTA_CHECK") != "" {
		t.Log("Skipping quota check (CI environment or AWS_SKIP_QUOTA_CHECK set)")
		return
	}

	t.Log("Checking ECR Public repository quota...")

	// Get current repository count
	cmd := exec.Command("aws", "ecr-public", "describe-repositories", "--region", "us-east-1", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		t.Logf("Warning: Could not check ECR Public quota: %v", err)
		return
	}

	var response struct {
		Repositories []interface{} `json:"repositories"`
	}

	if err := json.Unmarshal(output, &response); err != nil {
		t.Logf("Warning: Could not parse ECR Public response: %v", err)
		return
	}

	currentCount := len(response.Repositories)

	// ECR Public has a default limit of 10,000 repositories per region
	// We'll warn at 70% (7,000) and fail at 85% (8,500) to be more conservative during parallel testing
	const (
		warningThreshold = 7000
		errorThreshold   = 8500
	)

	t.Logf("Current ECR Public repositories: %d", currentCount)

	if currentCount >= errorThreshold {
		t.Fatalf("ECR Public repository quota nearly exhausted (%d repositories, limit ~10,000). Please clean up existing repositories before running tests. Use cleanup script: ./cleanup-orphaned-resources.sh", currentCount)
	}

	if currentCount >= warningThreshold {
		t.Logf("Warning: High ECR Public repository usage (%d repositories). Consider running cleanup script.", currentCount)
	}
}

// ensureSafeTestExecution performs pre-flight checks for safe test execution
// This includes quota checking and resource validation
func ensureSafeTestExecution(t *testing.T, repositoryName string) {
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Check AWS quota to prevent exhaustion
	checkECRPublicQuota(t)
}

// setupTestCleanup creates a comprehensive cleanup function for test resources
// This handles terraform destroy failures with enhanced error recovery
func setupTestCleanup(t *testing.T, terraformOptions *terraform.Options, repositoryName string) func() {
	return func() {
		t.Logf("Starting cleanup for repository: %s", repositoryName)

		// Comprehensive error recovery with multiple retry attempts
		destroyAttempted := false
		terraformSuccess := false

		// Attempt terraform destroy with error recovery
		if terraformOptions != nil {
			t.Logf("Attempting terraform destroy...")
			destroyOutput, err := terraform.DestroyE(t, terraformOptions)
			destroyAttempted = true

			if err != nil {
				t.Logf("Warning: Terraform destroy failed: %v", err)
				if destroyOutput != "" {
					t.Logf("Terraform destroy output: %s", destroyOutput)
				}
				t.Logf("Repository that may need manual cleanup: %s", repositoryName)
			} else {
				t.Logf("Terraform destroy succeeded for repository: %s", repositoryName)
				terraformSuccess = true
			}
		}

		// If terraform destroy failed, attempt direct AWS cleanup
		if destroyAttempted && !terraformSuccess {
			t.Logf("Attempting direct AWS cleanup as fallback for repository: %s", repositoryName)
			if cleanupErr := attemptDirectAWSCleanup(t, repositoryName); cleanupErr != nil {
				t.Logf("Direct AWS cleanup also failed: %v", cleanupErr)

				// Final attempt: wait and retry direct cleanup (sometimes resources are in transition)
				t.Logf("Waiting 10 seconds before final cleanup attempt...")
				time.Sleep(10 * time.Second)

				if retryErr := attemptDirectAWSCleanup(t, repositoryName); retryErr != nil {
					t.Logf("Final cleanup attempt also failed: %v", retryErr)

					// Provide comprehensive manual cleanup instructions
					t.Logf("=== MANUAL CLEANUP REQUIRED ===")
					t.Logf("Repository: %s", repositoryName)
					t.Logf("Region: us-east-1")
					t.Logf("AWS CLI Command:")
					t.Logf("  aws ecr-public delete-repository --region us-east-1 --repository-name %s --force", repositoryName)
					t.Logf("Alternative cleanup methods:")
					t.Logf("  1. Use cleanup script: ./test/cleanup-orphaned-resources.sh")
					t.Logf("  2. AWS Console: https://console.aws.amazon.com/ecr/repositories?region=us-east-1")
					t.Logf("  3. Check if repository exists: aws ecr-public describe-repositories --repository-names %s --region us-east-1", repositoryName)
					t.Logf("==============================")
				} else {
					t.Logf("Final cleanup attempt succeeded for repository: %s", repositoryName)
				}
			} else {
				t.Logf("Direct AWS cleanup succeeded for repository: %s", repositoryName)
			}
		}

		if !destroyAttempted {
			t.Logf("Warning: No cleanup attempted due to invalid terraform options")
		}
	}
}

// attemptDirectAWSCleanup tries to clean up ECR repository directly via AWS CLI
// This is used as a fallback when terraform destroy fails
func attemptDirectAWSCleanup(t *testing.T, repositoryName string) error {
	// Skip direct cleanup in CI environments to avoid permission issues
	if os.Getenv("CI") != "" {
		t.Log("Skipping direct AWS cleanup in CI environment")
		return fmt.Errorf("skipped in CI environment")
	}

	t.Logf("Attempting direct AWS cleanup for: %s", repositoryName)

	// First, check if repository exists
	checkCmd := exec.Command("aws", "ecr-public", "describe-repositories",
		"--repository-names", repositoryName,
		"--region", "us-east-1",
		"--output", "json")

	if err := checkCmd.Run(); err != nil {
		// Repository doesn't exist, cleanup not needed
		t.Logf("Repository %s does not exist, no cleanup needed", repositoryName)
		return nil
	}

	// Repository exists, attempt to delete it
	deleteCmd := exec.Command("aws", "ecr-public", "delete-repository",
		"--repository-name", repositoryName,
		"--region", "us-east-1",
		"--force")

	if err := deleteCmd.Run(); err != nil {
		return fmt.Errorf("failed to delete repository %s: %v", repositoryName, err)
	}

	t.Logf("Successfully deleted repository via AWS CLI: %s", repositoryName)
	return nil
}
