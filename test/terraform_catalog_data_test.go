package test

import (
	"fmt"
	"regexp"
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
		Vars: map[string]interface{}{
			"repository_name":                    repositoryName,
			"catalog_data_description":           "Production-ready container for web applications with security hardening and performance optimizations",
			"catalog_data_about_text":           "# Web Application Container\n\n## Overview\nThis container provides a secure, production-ready environment for web applications.\n\n## Features\n- Security hardened base image\n- Performance optimized configuration\n- Multi-architecture support\n- Comprehensive logging\n\n## Security\n- Non-root user execution\n- Minimal attack surface\n- Regular security updates",
			"catalog_data_usage_text":           "# Usage Instructions\n\n## Quick Start\n```bash\n# Pull the latest image\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n\n# Run with default configuration\ndocker run -p 8080:8080 public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Configuration\n\nEnvironment variables:\n- `PORT`: Application port (default: 8080)\n- `ENV`: Environment mode (development/production)\n\n## Health Checks\n```bash\ncurl http://localhost:8080/health\n```",
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
				"usage_text": "# Usage Guide\n\n## Quick Start\n\n### Basic Deployment\n```bash\n# Pull the latest stable release\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n\n# Run with default configuration\ndocker run -d \\\n  --name my-app \\\n  -p 8080:8080 \\\n  public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n### Production Deployment\n```bash\n# Production deployment with custom configuration\ndocker run -d \\\n  --name production-app \\\n  --restart unless-stopped \\\n  -p 8080:8080 \\\n  -e NODE_ENV=production \\\n  -e LOG_LEVEL=info \\\n  -e METRICS_ENABLED=true \\\n  -v /var/log/app:/app/logs \\\n  public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Configuration Options\n\n### Environment Variables\n\n| Variable | Description | Default | Required |\n|----------|-------------|---------|----------|\n| `NODE_ENV` | Environment mode | development | No |\n| `PORT` | Application port | 8080 | No |\n| `LOG_LEVEL` | Logging level | info | No |\n| `METRICS_ENABLED` | Enable metrics | false | No |\n| `DATABASE_URL` | Database connection | - | Yes |\n| `REDIS_URL` | Redis connection | - | No |\n\n### Volume Mounts\n\n- `/app/logs`: Application log files\n- `/app/config`: Configuration files\n- `/app/data`: Persistent data storage\n\n## Health Monitoring\n\n### Health Check Endpoints\n\n```bash\n# Application health\ncurl http://localhost:8080/health\n\n# Readiness check\ncurl http://localhost:8080/ready\n\n# Metrics endpoint\ncurl http://localhost:8080/metrics\n```\n\n### Expected Responses\n\n```json\n{\n  \"status\": \"healthy\",\n  \"timestamp\": \"2024-01-01T12:00:00Z\",\n  \"uptime\": 3600,\n  \"version\": \"1.0.0\"\n}\n```\n\n## Kubernetes Deployment\n\n### Basic Deployment\n```yaml\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: enterprise-app\nspec:\n  replicas: 3\n  selector:\n    matchLabels:\n      app: enterprise-app\n  template:\n    metadata:\n      labels:\n        app: enterprise-app\n    spec:\n      containers:\n      - name: app\n        image: public.ecr.aws/registry/" + repositoryName + ":latest\n        ports:\n        - containerPort: 8080\n        env:\n        - name: NODE_ENV\n          value: \"production\"\n        livenessProbe:\n          httpGet:\n            path: /health\n            port: 8080\n          initialDelaySeconds: 30\n          periodSeconds: 10\n        readinessProbe:\n          httpGet:\n            path: /ready\n            port: 8080\n          initialDelaySeconds: 5\n          periodSeconds: 5\n```\n\n## Security Considerations\n\n### Running as Non-root\nThis container runs as a non-root user by default:\n\n```bash\n# Verify non-root execution\ndocker run --rm public.ecr.aws/registry/" + repositoryName + ":latest id\n# Output: uid=1001(appuser) gid=1001(appuser) groups=1001(appuser)\n```\n\n### Scanning for Vulnerabilities\n```bash\n# Scan image for vulnerabilities\ndocker scout cves public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Troubleshooting\n\n### Common Issues\n\n**Port Already in Use**\n```bash\n# Use different port\ndocker run -p 8081:8080 public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n**Permission Issues**\n```bash\n# Check container user\ndocker run --rm public.ecr.aws/registry/" + repositoryName + ":latest whoami\n```\n\n### Debug Mode\n```bash\n# Run in debug mode\ndocker run -e LOG_LEVEL=debug public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Support & Contributing\n\n- **Issues**: Report issues on the project repository\n- **Documentation**: Comprehensive docs at project homepage\n- **Security**: Report security issues privately\n- **Contributing**: Pull requests welcome with proper testing",
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
			"catalog_data_description":           "Multi-architecture container supporting x86-64, ARM, and ARM 64 platforms",
			"catalog_data_about_text":           "# Multi-Architecture Container\n\nThis container is built for multiple architectures:\n- x86-64 for traditional Intel/AMD systems\n- ARM 64 for modern ARM systems including Apple Silicon\n- ARM for IoT and embedded devices",
			"catalog_data_usage_text":           "# Multi-Architecture Usage\n\nDocker will automatically pull the correct architecture:\n```bash\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\nTo pull a specific architecture:\n```bash\ndocker pull --platform linux/amd64 public.ecr.aws/registry/" + repositoryName + ":latest\ndocker pull --platform linux/arm64 public.ecr.aws/registry/" + repositoryName + ":latest\n```",
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
			"catalog_data_description":      "Container with comprehensive markdown documentation",
			"catalog_data_about_text":      "# Application Container\n\n## Features\n\n- **Security**: Hardened base image\n- **Performance**: Optimized for speed\n- **Monitoring**: Built-in observability\n\n### Code Example\n\n```javascript\nconst app = require('./app');\napp.listen(3000);\n```\n\n> **Note**: This container follows security best practices.",
			"catalog_data_usage_text":      "# Getting Started\n\n## Installation\n\n1. Pull the image\n2. Configure environment\n3. Run the container\n\n```bash\n# Step 1: Pull\ndocker pull public.ecr.aws/registry/" + repositoryName + ":latest\n\n# Step 2: Run\ndocker run -p 3000:3000 public.ecr.aws/registry/" + repositoryName + ":latest\n```\n\n## Configuration\n\n| Variable | Description |\n|----------|-------------|\n| `PORT` | Application port |\n| `ENV` | Environment mode |",
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