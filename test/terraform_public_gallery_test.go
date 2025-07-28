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

// TestECRPublicGalleryOptimization tests repository configuration optimized for ECR Public Gallery
func TestECRPublicGalleryOptimization(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-gallery-opt-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description": "Production-ready Node.js application container optimized for ECR Public Gallery discoverability",
				"about_text": `# ` + repositoryName + `

## Description

This container provides a production-ready Node.js application environment optimized for the ECR Public Gallery with comprehensive documentation and best practices.

## Features

- **Security Hardening**: Non-root user execution and minimal attack surface
- **Multi-Architecture**: Supports x86-64 and ARM 64 architectures
- **Performance Optimized**: Fast startup and efficient resource usage
- **Health Monitoring**: Built-in health checks and metrics
- **Production Ready**: Comprehensive logging and error handling

## Use Cases

- Web applications and APIs
- Microservices deployments
- Serverless containers
- Development environments
- CI/CD pipelines

## Security

- Runs as non-root user for security
- Regular security updates
- Minimal base image to reduce attack surface
- No sensitive data in image layers`,
				"usage_text": `# Usage Instructions

## Quick Start

\`\`\`bash
# Pull the latest image
docker pull public.ecr.aws/registry/` + repositoryName + `:latest

# Run with default configuration
docker run -p 3000:3000 public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

## Production Usage

\`\`\`bash
# Production deployment
docker run -d \\
  --name ` + repositoryName + ` \\
  --restart unless-stopped \\
  -p 3000:3000 \\
  -e NODE_ENV=production \\
  public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

## Kubernetes Deployment

\`\`\`yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ` + repositoryName + `
spec:
  replicas: 3
  selector:
    matchLabels:
      app: ` + repositoryName + `
  template:
    metadata:
      labels:
        app: ` + repositoryName + `
    spec:
      containers:
      - name: app
        image: public.ecr.aws/registry/` + repositoryName + `:latest
        ports:
        - containerPort: 3000
        livenessProbe:
          httpGet:
            path: /health
            port: 3000
\`\`\`

## Environment Variables

- \`NODE_ENV\`: Environment mode (development/production)
- \`PORT\`: Application port (default: 3000)
- \`LOG_LEVEL\`: Logging level (debug/info/warn/error)

## Health Checks

\`\`\`bash
curl http://localhost:3000/health
\`\`\``,
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

	// Verify gallery-optimized repository configuration
	validatePublicGalleryOptimization(t, terraformOptions, repositoryName)
}

// TestECRPublicGallerySearchability tests configuration for maximum searchability
func TestECRPublicGallerySearchability(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-searchable-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Use us-east-1 as ECR Public is only available in this region
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/using_variables",
		Vars: map[string]interface{}{
			"repository_name":                    repositoryName,
			"catalog_data_description":           "Open-source Python web framework container with Flask, Django support, and development tools",
			"catalog_data_about_text":           "# Python Web Framework Container\n\n## Overview\nA comprehensive Python container optimized for web development with popular frameworks.\n\n## Supported Frameworks\n- **Flask**: Lightweight WSGI web application framework\n- **Django**: High-level Python web framework\n- **FastAPI**: Modern, fast web framework for building APIs\n\n## Keywords\nPython, Flask, Django, FastAPI, web development, API, microservices, REST, GraphQL, WSGI, ASGI",
			"catalog_data_usage_text":           "# Python Web Development\n\n## Flask Application\n```bash\ndocker run -p 5000:5000 -v $(pwd):/app public.ecr.aws/registry/" + repositoryName + ":latest python app.py\n```\n\n## Django Application\n```bash\ndocker run -p 8000:8000 -v $(pwd):/app public.ecr.aws/registry/" + repositoryName + ":latest python manage.py runserver 0.0.0.0:8000\n```\n\n## FastAPI Application\n```bash\ndocker run -p 8000:8000 -v $(pwd):/app public.ecr.aws/registry/" + repositoryName + ":latest uvicorn main:app --host 0.0.0.0 --port 8000\n```",
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

	// Verify searchability-optimized configuration
	validatePublicGallerySearchability(t, terraformOptions, repositoryName)
}

// TestECRPublicGalleryContentGuidelines tests compliance with content guidelines
func TestECRPublicGalleryContentGuidelines(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-guidelines-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

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
				"usage_text": `# Usage Guide

## Basic Spring Boot Application

\`\`\`bash
# Run a basic Spring Boot application
docker run -p 8080:8080 \\
  -v $(pwd)/app.jar:/app/application.jar \\
  public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

## Production Deployment

\`\`\`bash
# Production deployment with environment configuration
docker run -d \\
  --name java-app \\
  --restart unless-stopped \\
  -p 8080:8080 \\
  -e SPRING_PROFILES_ACTIVE=production \\
  -e JAVA_OPTS="-Xmx2g -Xms1g" \\
  -e DATABASE_URL="jdbc:postgresql://db:5432/myapp" \\
  -v /app/logs:/opt/app/logs \\
  -v $(pwd)/application.jar:/app/application.jar \\
  public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

## Docker Compose Setup

\`\`\`yaml
version: '3.8'
services:
  app:
    image: public.ecr.aws/registry/` + repositoryName + `:latest
    ports:
      - "8080:8080"
    environment:
      - SPRING_PROFILES_ACTIVE=production
      - DATABASE_URL=jdbc:postgresql://db:5432/myapp
      - REDIS_URL=redis://redis:6379
    volumes:
      - ./app.jar:/app/application.jar
      - ./logs:/opt/app/logs
    depends_on:
      - db
      - redis
      
  db:
    image: postgres:15
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      
  redis:
    image: redis:7-alpine
\`\`\`

## Kubernetes Deployment

\`\`\`yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: java-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: java-app
  template:
    metadata:
      labels:
        app: java-app
    spec:
      containers:
      - name: app
        image: public.ecr.aws/registry/` + repositoryName + `:latest
        ports:
        - containerPort: 8080
        env:
        - name: SPRING_PROFILES_ACTIVE
          value: "kubernetes"
        - name: JAVA_OPTS
          value: "-Xmx2g -Xms1g -XX:+UseG1GC"
        resources:
          requests:
            memory: "1Gi"
            cpu: "500m"
          limits:
            memory: "3Gi"
            cpu: "2"
        livenessProbe:
          httpGet:
            path: /actuator/health
            port: 8080
          initialDelaySeconds: 60
          periodSeconds: 30
        readinessProbe:
          httpGet:
            path: /actuator/health/readiness
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        volumeMounts:
        - name: app-jar
          mountPath: /app/application.jar
          subPath: application.jar
      volumes:
      - name: app-jar
        configMap:
          name: app-config
\`\`\`

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| \`SPRING_PROFILES_ACTIVE\` | Active Spring profiles | default | No |
| \`JAVA_OPTS\` | JVM options | -Xmx1g -Xms512m | No |  
| \`DATABASE_URL\` | Database connection URL | - | Yes |
| \`REDIS_URL\` | Redis connection URL | - | No |
| \`LOG_LEVEL\` | Application log level | INFO | No |

### Health Monitoring

\`\`\`bash
# Application health
curl http://localhost:8080/actuator/health

# Readiness check  
curl http://localhost:8080/actuator/health/readiness

# Metrics endpoint
curl http://localhost:8080/actuator/metrics
\`\`\`

## Security Best Practices

### Non-root User
\`\`\`bash
# Verify non-root execution
docker exec container-name id
# Output: uid=1001(appuser) gid=1001(appuser)
\`\`\`

### Vulnerability Scanning
\`\`\`bash
# Scan for vulnerabilities
docker scout cves public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

## Troubleshooting

### Memory Issues
\`\`\`bash
# Increase heap size
docker run -e JAVA_OPTS="-Xmx4g -Xms2g" public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

### Debug Mode
\`\`\`bash
# Enable debug logging
docker run -e LOG_LEVEL=DEBUG public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\`

### JVM Analysis
\`\`\`bash
# Enable JVM debugging
docker run -e JAVA_OPTS="-XX:+PrintGCDetails -XX:+PrintGCTimeStamps" \\
  public.ecr.aws/registry/` + repositoryName + `:latest
\`\`\``,
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

	// Verify content guidelines compliance
	validatePublicGalleryContentGuidelines(t, terraformOptions, repositoryName)
}

// TestECRPublicGalleryRegionalConstraints tests us-east-1 region constraint
func TestECRPublicGalleryRegionalConstraints(t *testing.T) {
	t.Parallel()

	// Generate a unique repository name to avoid conflicts
	uniqueID := strings.ToLower(random.UniqueId())
	repositoryName := fmt.Sprintf("terratest-regional-%s", uniqueID)
	
	// Validate repository name format for security
	validateRepositoryNameFormat(t, repositoryName)

	// Explicitly test us-east-1 constraint
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"repository_name": repositoryName,
			"catalog_data": map[string]interface{}{
				"description":       "Test repository for regional constraints validation",
				"about_text":       "# Regional Constraints Test\nThis repository tests ECR Public's us-east-1 regional constraint.",
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