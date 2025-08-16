# Terraform AWS ECR Public Module - Development Guidelines

## Overview
This document outlines Terraform-specific development guidelines for the terraform-aws-ecrpublic module, focusing on best practices for AWS ECR Public Gallery repository management infrastructure as code. This module specializes in creating and managing public container repositories that are accessible via the ECR Public Gallery, designed for open-source projects and public container distribution.

## Module Structure & Organization

### File Organization
- **main.tf** - ECR Public repository resource definitions and locals (48 lines)
- **variables.tf** - Input variable definitions for catalog data and timeouts (62 lines)
- **outputs.tf** - Output value definitions for repository attributes
- **versions.tf** - Provider version constraints
- **examples/** - Example configurations for different use cases
  - **using_objects/** - Object-based catalog data configuration
  - **using_variables/** - Variable-based catalog data configuration
- **test/** - Go-based Terratest integration tests

### Code Organization Principles
- Focus on single ECR Public repository resource management
- Use descriptive locals for catalog data configuration
- Support both object-based and variable-based configuration patterns
- Maintain backward compatibility with existing variable names
- Implement flexible catalog data management for ECR Public Gallery

## Terraform Best Practices

### ECR Public Repository Creation Patterns
**Single ECR Public repository with flexible catalog data configuration:**

```hcl
# ECR Public repository with dynamic catalog data configuration
resource "aws_ecrpublic_repository" "repo" {
  repository_name = var.repository_name

  # Dynamic catalog data configuration for ECR Public Gallery
  dynamic "catalog_data" {
    for_each = local.catalog_data
    content {
      about_text        = lookup(catalog_data.value, "about_text", null)
      architectures     = lookup(catalog_data.value, "architectures", null)
      description       = lookup(catalog_data.value, "description", null)
      logo_image_blob   = lookup(catalog_data.value, "logo_image_blob", null)
      operating_systems = lookup(catalog_data.value, "operating_systems", null)
      usage_text        = lookup(catalog_data.value, "usage_text", null)
    }
  }

  # Optional timeouts for repository operations
  dynamic "timeouts" {
    for_each = local.timeouts
    content {
      delete = lookup(timeouts.value, "delete", null)
    }
  }
}
```

### Catalog Data Management
**Implement flexible catalog data configuration supporting both object and variable approaches:**

```hcl
# Support both object-based and variable-based catalog data configuration
locals {
  catalog_data = [
    {
      about_text        = coalesce(lookup(var.catalog_data, "about_text", null), var.catalog_data_about_text)
      architectures     = coalesce(lookup(var.catalog_data, "architectures", null), var.catalog_data_architectures)
      description       = coalesce(lookup(var.catalog_data, "description", null), var.catalog_data_description)
      logo_image_blob   = coalesce(lookup(var.catalog_data, "logo_image_blob", null), var.catalog_data_logo_image_blob)
      operating_systems = coalesce(lookup(var.catalog_data, "operating_systems", null), var.catalog_data_operating_systems)
      usage_text        = coalesce(lookup(var.catalog_data, "usage_text", null), var.catalog_data_usage_text)
    }
  ]
}
```

### Timeout Configuration
**Implement flexible timeout management for ECR Public operations:**

```hcl
# Flexible timeout configuration
variable "timeouts" {
  description = "Timeouts map for repository operations"
  type        = map(any)
  default     = {}
}

variable "timeouts_delete" {
  description = "How long to wait for a repository to be deleted"
  type        = string
  default     = null
}

locals {
  # Build timeouts configuration conditionally
  timeouts = (var.timeouts_delete != null || length(var.timeouts) > 0) ? [{
    delete = coalesce(lookup(var.timeouts, "delete", null), var.timeouts_delete)
  }] : []
}
```

### Variable Validation Patterns
**Implement robust validation for ECR Public Gallery content:**

```hcl
# Example: Enhanced variable validation for public content
variable "catalog_data_description" {
  description = "Public description visible in ECR Public Gallery"
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_description == null || length(var.catalog_data_description) <= 256
    error_message = "Description must be 256 characters or less for ECR Public Gallery visibility."
  }
}

variable "catalog_data_architectures" {
  description = "Supported architectures for container images"
  type        = list(string)
  default     = []

  validation {
    condition = alltrue([
      for arch in var.catalog_data_architectures :
      contains(["ARM", "ARM 64", "x86", "x86-64"], arch)
    ])
    error_message = "Architectures must be one of: ARM, ARM 64, x86, x86-64."
  }
}

variable "catalog_data_operating_systems" {
  description = "Supported operating systems for container images"
  type        = list(string)
  default     = []

  validation {
    condition = alltrue([
      for os in var.catalog_data_operating_systems :
      contains(["Linux", "Windows"], os)
    ])
    error_message = "Operating systems must be one of: Linux, Windows."
  }
}
```

## Testing Requirements

### Terratest Integration
**Use Go-based testing for ECR Public resources:**

```go
// Example: ECR Public testing pattern
func TestTerraformECRPublicExample(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../examples/using_variables",
        Vars: map[string]interface{}{
            "repository_name": fmt.Sprintf("test-public-repo-%s", random.UniqueId()),
            "catalog_data_description": "Test repository for Terratest",
            "catalog_data_about_text": "# Test Repository\nThis is a test repository created by Terratest",
            "catalog_data_architectures": []string{"x86-64"},
            "catalog_data_operating_systems": []string{"Linux"},
        },
    }

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Validate ECR Public repository creation
    repositoryName := terraform.Output(t, terraformOptions, "repository_name")
    repositoryURI := terraform.Output(t, terraformOptions, "repository_uri")

    assert.NotEmpty(t, repositoryName)
    assert.Contains(t, repositoryURI, "public.ecr.aws")

    // Validate catalog data was applied
    registryID := terraform.Output(t, terraformOptions, "registry_id")
    assert.NotEmpty(t, registryID)
}
```

### Test Coverage Strategy
**Comprehensive testing for ECR Public functionality:**
- **Create corresponding test files** in `test/` directory
- **Test both object-based and variable-based configuration patterns**
- **Validate catalog data content and formatting**
- **Test public repository accessibility and URI format**
- **Verify timeout configuration handling**
- **Test example configurations in `examples/` directory**
- **Validate variable validation rules**

## ECR Public Gallery Considerations

### Region Constraints
**ECR Public repositories must be created in us-east-1:**

```hcl
# ECR Public provider configuration - must use us-east-1
provider "aws" {
  alias  = "ecr_public"
  region = "us-east-1"
}

# ECR Public repositories can only be created in us-east-1
resource "aws_ecrpublic_repository" "repo" {
  provider = aws.ecr_public

  repository_name = var.repository_name
  # Repository configuration...
}
```

### Public Repository Management
**ECR Public repositories are globally accessible by design:**

```hcl
# ECR Public repositories are automatically accessible to all AWS users
# No additional IAM policies or resource policies are required for public access
resource "aws_ecrpublic_repository" "repo" {
  repository_name = var.repository_name

  # Catalog data is publicly visible in the ECR Public Gallery
  catalog_data {
    description = "This description will be publicly visible"
    about_text  = "# Public Repository\nThis content is publicly accessible"
    usage_text  = "# Usage Instructions\nPublic usage documentation"
  }
}
```

### Content Guidelines
**Ensure appropriate content for public repositories:**

```hcl
# Guidelines for public repository content
variable "catalog_data_description" {
  description = "Public description visible in ECR Public Gallery"
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_description == null || length(var.catalog_data_description) <= 256
    error_message = "Description must be 256 characters or less."
  }
}

variable "catalog_data_about_text" {
  description = "Public about text in markdown format"
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_about_text == null || length(var.catalog_data_about_text) > 0
    error_message = "About text cannot be empty when provided."
  }
}
```

## ECR Public Development Patterns

### Catalog Data Flexibility
**Support both object and variable-based configuration:**

```hcl
# Users can choose between object-based or individual variable configuration
# Object-based approach
module "public-ecr-object" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-app"

  catalog_data = {
    description = "My application"
    about_text  = "# About\nDetailed information"
    usage_text  = "# Usage\nInstructions here"
  }
}

# Variable-based approach
module "public-ecr-variables" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-app"

  catalog_data_description = "My application"
  catalog_data_about_text  = "# About\nDetailed information"
  catalog_data_usage_text  = "# Usage\nInstructions here"
}
```

### Gallery Optimization
**Optimize repository for ECR Public Gallery discoverability:**

```hcl
# Best practices for ECR Public Gallery visibility
resource "aws_ecrpublic_repository" "optimized" {
  repository_name = var.repository_name

  catalog_data {
    # Clear, searchable description
    description = "Production-ready Node.js application container"

    # Comprehensive about text with markdown formatting
    about_text = <<-EOT
      # ${var.repository_name}

      ## Description
      This container provides a production-ready Node.js application environment.

      ## Features
      - Multi-stage builds for optimized size
      - Security hardening
      - Health checks included
    EOT

    # Detailed usage instructions
    usage_text = <<-EOT
      # Usage

      ## Quick Start
      ```bash
      docker pull public.ecr.aws/your-registry/${var.repository_name}:latest
      docker run -p 3000:3000 public.ecr.aws/your-registry/${var.repository_name}:latest
      ```
    EOT

    # Proper architecture and OS tagging for filters
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }
}
```

## Development Workflow

### Pre-commit Requirements
- **Run `terraform fmt`** on all modified files
- **Execute `terraform validate`** to ensure syntax correctness
- **Test examples** in `examples/` directory
- **Validate catalog data** content and formatting
- **Update documentation** for variable or output changes
- **Check markdown formatting** in catalog data fields

### ECR Public Testing
**Test ECR Public repository creation and catalog data:**

```bash
# Example testing approach for ECR Public
terraform init

# Use single timestamp to avoid race conditions
TIMESTAMP=$(date +%s)
REPO_NAME="test-repo-$TIMESTAMP"

terraform plan -var="repository_name=$REPO_NAME" \
               -var="catalog_data_description=Test repository" \
               -var="catalog_data_about_text=# Test\nTest repository" \
               -var="catalog_data_architectures=[\"x86-64\"]" \
               -var="catalog_data_operating_systems=[\"Linux\"]"

terraform apply -auto-approve

# Verify repository in ECR Public Gallery
aws ecr-public describe-repositories --repository-names "$REPO_NAME" --region us-east-1

# Test repository accessibility (ECR Public is always in us-east-1)
aws ecr-public get-login-token --region us-east-1

# Clean up
terraform destroy -auto-approve
```

### Manual Validation
**Verify ECR Public Gallery integration:**

```bash
# Check repository in ECR Public Gallery
aws ecr-public describe-repositories --region us-east-1
aws ecr-public describe-repository-creation-template --region us-east-1

# Validate catalog data
aws ecr-public get-repository-catalog-data --repository-name <repository-name> --region us-east-1

# Test public accessibility (no authentication required for public repositories)
docker pull public.ecr.aws/<registry-alias>/<repository-name>:latest
```

### Release Management
- **Use conventional commit messages** for proper automation
- **Follow semantic versioning** principles
- **DO NOT manually update CHANGELOG.md** - use release-please
- **Test all examples** before releasing

## Common ECR Public Patterns

### 1. **Flexible Configuration**
Support both object-based and variable-based catalog data configuration

### 2. **Gallery Optimization**
Design catalog data for maximum discoverability in ECR Public Gallery

### 3. **Markdown Content**
Provide rich, well-formatted descriptions and usage instructions

### 4. **Public Access Design**
Repositories are publicly accessible by design with no additional IAM configuration

### 5. **Architecture Tagging**
Proper categorization with architecture and operating system metadata

### 6. **Timeout Management**
Configurable timeouts for repository operations

### 7. **Content Validation**
Input validation for catalog data fields and formatting

### 8. **Backward Compatibility**
Maintain compatibility while adding new catalog data features

## Example Configurations

### Simple ECR Public Repository
```hcl
module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-app"

  catalog_data_description = "My public application container"
  catalog_data_about_text  = "# My Public App\nDetailed description here"
  catalog_data_usage_text  = "# Usage\nDocker pull instructions here"
}
```

### Complete ECR Public with Full Catalog Data
```hcl
module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-complete-public-app"

  catalog_data = {
    about_text        = "# Public Application\nComprehensive description in Markdown"
    architectures     = ["ARM", "x86-64"]
    description       = "Production-ready public container image"
    # Validate file existence and size before using filebase64()
    logo_image_blob   = can(fileexists("${path.module}/logo.png")) ? filebase64("${path.module}/logo.png") : null
    operating_systems = ["Linux"]
    usage_text        = "# Usage\n\n```bash\ndocker pull public.ecr.aws/myregistry/my-complete-public-app:latest\n```"
  }

  timeouts = {
    delete = "30m"
  }
}
```

## Provider Version Management

```hcl
terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}
```

## Key Module Features

1. **ECR Public Repository Management** - Single public repository creation with us-east-1 region constraint
2. **Flexible Catalog Data Configuration** - Support for both object-based and variable-based configuration approaches
3. **Public Gallery Integration** - Optimized metadata for ECR Public Gallery visibility and discoverability
4. **Rich Markdown Content Support** - Comprehensive descriptions and usage documentation with markdown formatting
5. **Architecture & OS Tagging** - Proper categorization with validated architecture and operating system metadata
6. **Logo Image Support** - Visual branding support for verified AWS accounts (base64 encoded images)
7. **Flexible Timeout Configuration** - Configurable operation timeouts for repository management
8. **Comprehensive Validation Rules** - Input validation for catalog data fields and content formatting
9. **Dual Configuration Patterns** - Backward compatibility with both object and individual variable approaches
10. **Global Public Access** - Repositories are globally accessible without additional IAM configuration
11. **Gallery Content Optimization** - Best practices for searchable, discoverable public repositories
12. **Example-Driven Documentation** - Multiple configuration examples for different use cases
13. **Security-First Design** - Enhanced input validation, content security, and least privilege principles

## Security Best Practices

### Content Security
**ECR Public repositories are publicly accessible, requiring careful content validation:**

```hcl
# Enhanced security validations prevent malicious content
variable "catalog_data_description" {
  description = "Public description visible in ECR Public Gallery"
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_description == null || length(var.catalog_data_description) <= 256
    error_message = "Description must be 256 characters or less for ECR Public Gallery visibility."
  }

  validation {
    condition = var.catalog_data_description == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", var.catalog_data_description))
    error_message = "Description must not contain potentially malicious scripts or executable content for security."
  }
}
```

### Repository Naming Security
**Prevent reserved names and potential security issues:**

```hcl
variable "repository_name" {
  description = "Name of the repository"
  type        = string

  validation {
    condition     = can(regex("^[a-z0-9](?:[a-z0-9._-]*[a-z0-9])?$", var.repository_name))
    error_message = "Repository name must start and end with an alphanumeric character and can only contain lowercase letters, numbers, hyphens, underscores, and periods."
  }

  validation {
    condition     = !can(regex("(?i)^(admin|aws|amazon|ecr|public|private|root|system|test|null|undefined)$", var.repository_name))
    error_message = "Repository name must not use reserved or potentially confusing names for security and clarity."
  }
}
```

### Encryption Considerations
**ECR Public Repository Encryption Limitations:**

- **No Encryption Configuration**: Unlike private ECR repositories, ECR Public repositories do not support custom encryption configuration
- **AWS Managed Encryption**: ECR Public uses AWS-managed encryption at rest by default
- **Data in Transit**: All API communications use HTTPS/TLS encryption
- **Content Validation**: Focus on input validation since encryption configuration is not available

### IAM and Access Control
**ECR Public repositories follow different security principles:**

```hcl
# ECR Public repositories are inherently public - no IAM policies needed for public access
# Focus on AWS account-level controls for repository management

# Example: Limit repository creation to specific roles (implemented outside this module)
data "aws_iam_policy_document" "ecr_public_admin" {
  statement {
    effect = "Allow"
    actions = [
      "ecr-public:CreateRepository",
      "ecr-public:DeleteRepository",
      "ecr-public:DescribeRepositories",
      "ecr-public:PutRepositoryCatalogData"
    ]
    resources = ["*"]
    
    condition {
      string_equals = {
        "aws:RequestedRegion" = "us-east-1"
      }
    }
  }
}
```

### Logo Image Security
**Secure handling of base64-encoded logo images:**

```hcl
variable "catalog_data_logo_image_blob" {
  description = "Base64-encoded repository logo payload (Only visible for verified accounts)"
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_logo_image_blob == null || can(base64decode(var.catalog_data_logo_image_blob))
    error_message = "Logo image blob must be valid base64-encoded data."
  }

  validation {
    condition = var.catalog_data_logo_image_blob == null || length(var.catalog_data_logo_image_blob) <= 2097152  # 2MB in base64 ~= 1.5MB raw
    error_message = "Logo image must be under 2MB when base64-encoded to prevent resource exhaustion."
  }
}
```

### Security Validation Patterns
**Comprehensive input validation for public content:**

```hcl
# Content security validation function (example pattern)
locals {
  # Validate all text inputs for potential security issues
  secure_content_check = alltrue([
    for field in [local._catalog_data_description, local._catalog_data_about_text, local._catalog_data_usage_text] :
    field == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", field))
  ])
}

# Assert security validation in main resource
resource "aws_ecrpublic_repository" "repo" {
  repository_name = var.repository_name

  # Fail fast if content security validation fails
  lifecycle {
    precondition {
      condition     = local.secure_content_check
      error_message = "All catalog data fields must pass security validation checks."
    }
  }

  # Repository configuration...
}
```

### Security Monitoring and Compliance
**Recommended practices for ECR Public security:**

1. **Tag-Based Resource Management**
   ```hcl
   tags = merge(var.tags, {
     "Security:Public"     = "true"
     "Security:Encryption" = "aws-managed"
     "Compliance:SOC2"     = "compliant"
   })
   ```

2. **Content Review Workflow**
   - Implement peer review for all catalog data changes
   - Use automated content scanning for security patterns
   - Regular audit of public repository content

3. **Access Logging**
   - Enable CloudTrail for ECR Public API calls
   - Monitor repository creation and modification events
   - Alert on unauthorized repository changes

### Threat Model Considerations
**Security considerations specific to public repositories:**

1. **Content Exposure**: All catalog data is publicly visible
2. **Name Squatting**: Prevent malicious actors from claiming similar names
3. **Brand Protection**: Logo images and descriptions represent your organization
4. **Information Disclosure**: Avoid exposing sensitive information in public content

### Security Testing Requirements
**Enhanced testing for security validations:**

```go
// Example security test pattern
func TestSecurityValidations(t *testing.T) {
    maliciousInputs := []string{
        "<script>alert('xss')</script>",
        "javascript:alert('xss')",
        "vbscript:msgbox('xss')",
        "data:text/html,<script>alert('xss')</script>",
        "onclick=\"alert('xss')\"",
    }

    for _, input := range maliciousInputs {
        terraformOptions := &terraform.Options{
            TerraformDir: "../examples/using_variables",
            Vars: map[string]interface{}{
                "repository_name": "test-repo",
                "catalog_data_description": input,
            },
        }

        _, err := terraform.InitAndApplyE(t, terraformOptions)
        assert.Error(t, err, "Should reject malicious input: %s", input)
    }
}
```

## MCP Server Configuration

### Available MCP Servers
This project is configured to use the following Model Context Protocol (MCP) servers for enhanced documentation access:

#### Terraform MCP Server
**Purpose**: Access up-to-date Terraform and AWS provider documentation
**Package**: `@modelcontextprotocol/server-terraform`

**Local Configuration** (`.mcp.json`):
```json
{
  "mcpServers": {
    "terraform": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-terraform@latest"]
    }
  }
}
```

**Usage Examples**:
- `Look up aws_ecrpublic_repository resource documentation`
- `Find the latest ECR Public catalog data examples`
- `Search for AWS ECR Public Terraform modules`
- `Get documentation for ECR Public Gallery integration`

#### Context7 MCP Server
**Purpose**: Access general library and framework documentation
**Package**: `@upstash/context7-mcp`

**Local Configuration** (`.mcp.json`):
```json
{
  "mcpServers": {
    "context7": {
      "command": "npx",
      "args": ["-y", "@upstash/context7-mcp@latest"]
    }
  }
}
```

**Usage Examples**:
- `Look up Go testing patterns for Terratest`
- `Find AWS CLI ECR Public commands documentation`
- `Get current container registry best practices`
- `Search for GitHub Actions workflow patterns`

### GitHub Actions Integration
The MCP servers are automatically available in GitHub Actions through the claude.yml workflow configuration. Claude can access the same documentation in PRs and issues as available locally.

### Usage Tips
1. **Be Specific**: When requesting documentation, specify the exact resource or concept
2. **Version Awareness**: Both servers provide current, version-specific documentation
3. **Combine Sources**: Use Terraform MCP for ECR Public-specific docs, Context7 for general development patterns
4. **Local vs CI**: Same MCP servers work in both local development and GitHub Actions

### Example Workflows

**ECR Public Resource Development**:
```
@claude I need to add support for ECR Public registry scanning configuration. Can you look up the latest aws_ecrpublic_repository documentation and show me how to implement enhanced security features?
```

**Testing Pattern Research**:
```
@claude Look up current Terratest patterns for testing ECR Public repositories and help me add comprehensive tests for the catalog data functionality.
```

**Gallery Optimization**:
```
@claude Research the latest ECR Public Gallery best practices and help me implement better discoverability and searchability for public repositories.
```

**Catalog Data Enhancement**:
```
@claude Look up the latest ECR Public catalog data patterns and help me improve the markdown content and metadata organization for better Gallery presentation.
```

*Note: This module focuses exclusively on AWS ECR Public Gallery best practices and patterns for public container distribution and open-source project hosting.*
