package test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
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
	return map[string]interface{}{
		"description": "Test container",
		"about_text":  "# Test\nBasic test container.",
		"usage_text":  "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
		"architectures": []string{"x86-64"},
		"operating_systems": []string{"Linux"},
	}
}

// generateMinimalVariableCatalogData creates minimal catalog data using individual variables
// This is the variable-based equivalent of generateMinimalCatalogData
func generateMinimalVariableCatalogData(repositoryName string) map[string]interface{} {
	return map[string]interface{}{
		"catalog_data_description": "Test container",
		"catalog_data_about_text":  "# Test\nBasic test container.",
		"catalog_data_usage_text":  "# Usage\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```",
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
	
	// Replace placeholder with actual repository name
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
}