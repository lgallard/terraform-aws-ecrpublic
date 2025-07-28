package test

import (
	"os"
	"path/filepath"
	"strings"
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

// loadTestData loads test data from external file and replaces placeholders
// This reduces memory usage by avoiding large embedded strings in test files
func loadTestData(filename, repositoryName string) (string, error) {
	filePath := filepath.Join("testdata", filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	
	// Replace placeholder with actual repository name
	return strings.ReplaceAll(string(content), "{{REPOSITORY_NAME}}", repositoryName), nil
}