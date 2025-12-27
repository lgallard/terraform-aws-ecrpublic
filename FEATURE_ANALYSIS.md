# Feature Analysis: terraform-aws-ecrpublic

This document provides a sequential analysis of the current module implementation compared to available AWS ECR Public Terraform resources and identifies potential features for implementation.

## Current State Analysis

### Implemented Features

The module currently implements the `aws_ecrpublic_repository` resource with:

| Feature | Status | Notes |
|---------|--------|-------|
| Repository creation | Implemented | Core functionality |
| Catalog data (object-based) | Implemented | Full support for all 6 catalog fields |
| Catalog data (variable-based) | Implemented | Alternative configuration pattern |
| About text | Implemented | With security and length validation |
| Architectures | Implemented | ARM, ARM 64, x86, x86-64 validation |
| Description | Implemented | 256 char limit validation |
| Logo image blob | Implemented | Base64 validation, 2MB limit |
| Operating systems | Implemented | Linux, Windows validation |
| Usage text | Implemented | 10240 char limit validation |
| Timeouts (delete) | Implemented | Duration format validation |
| Tags | Implemented | Standard tagging support |
| Security validations | Implemented | XSS/script injection prevention |
| Repository naming validation | Implemented | Reserved names prevention |

### Available AWS Provider Resources (Not Implemented)

Based on research of the [Terraform AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs):

| Resource/Data Source | Status | Priority |
|---------------------|--------|----------|
| `aws_ecrpublic_repository_policy` | **Not implemented** | High |
| `aws_ecrpublic_authorization_token` data source | Not exposed/documented | Medium |

---

## Feature Gap Analysis

### 1. Repository Policy Support (HIGH PRIORITY)

**Resource:** `aws_ecrpublic_repository_policy`

**Current state:** Not implemented in module

**AWS Provider availability:** Available since v3.70.0

**Description:** Allows managing IAM-based access policies for ECR Public repositories. While ECR Public repositories are publicly readable, you can control who can push images using repository policies.

**Example use cases:**
- Allow specific IAM roles/users to push images
- Cross-account image push permissions
- CI/CD pipeline access control

**Proposed implementation:**
```hcl
# New variables
variable "create_repository_policy" {
  description = "Whether to create a repository policy"
  type        = bool
  default     = false
}

variable "repository_policy" {
  description = "The policy document for the repository"
  type        = string
  default     = null
}

# New resource
resource "aws_ecrpublic_repository_policy" "this" {
  count           = var.create_repository_policy ? 1 : 0
  repository_name = aws_ecrpublic_repository.repo.repository_name
  policy          = var.repository_policy
}

# New outputs
output "repository_policy" {
  description = "The repository policy JSON"
  value       = try(aws_ecrpublic_repository_policy.this[0].policy, null)
}
```

**Benefits:**
- Feature parity with [terraform-aws-modules/ecr](https://github.com/terraform-aws-modules/terraform-aws-ecr)
- Enterprise-ready access control
- CI/CD integration support

---

### 2. Authorization Token Data Source Documentation (MEDIUM PRIORITY)

**Data Source:** `aws_ecrpublic_authorization_token`

**Current state:** Not documented or exposed in module

**Description:** Provides authorization token for ECR Public API operations. Useful for pushing images programmatically.

**Proposed implementation:** Add documentation example showing how to use alongside the module:

```hcl
# In examples/with_auth_token/main.tf
data "aws_ecrpublic_authorization_token" "token" {
  # Must use us-east-1 region
}

module "public-ecr" {
  source          = "lgallard/ecrpublic/aws"
  repository_name = "my-public-repo"
  # ...
}

# For Docker login
output "docker_login_command" {
  value = "echo ${data.aws_ecrpublic_authorization_token.token.password} | docker login --username AWS --password-stdin public.ecr.aws"
  sensitive = true
}
```

---

### 3. Multiple Repository Support (MEDIUM PRIORITY)

**Current state:** Module creates single repository

**Description:** Many use cases require creating multiple related repositories with shared configurations.

**Proposed implementation:**
```hcl
variable "repositories" {
  description = "Map of repository configurations"
  type = map(object({
    catalog_data = optional(object({
      about_text        = optional(string)
      architectures     = optional(list(string))
      description       = optional(string)
      logo_image_blob   = optional(string)
      operating_systems = optional(list(string))
      usage_text        = optional(string)
    }))
    tags = optional(map(string))
  }))
  default = {}
}

resource "aws_ecrpublic_repository" "repos" {
  for_each        = var.repositories
  repository_name = each.key
  # ... dynamic catalog_data block
}
```

**Alternative:** Keep current single-repository focus and document using `for_each` at module call level.

---

### 4. Force Delete Support (LOW PRIORITY)

**Current state:** Delete timeout configurable but no force delete

**Description:** AWS ECR repositories can be force-deleted even with images present.

**Note:** The `aws_ecrpublic_repository` resource does NOT have a `force_delete` argument (unlike private ECR). This is an AWS API limitation, not a module gap.

---

### 5. Enhanced Outputs (LOW PRIORITY)

**Current state:** Basic outputs implemented

**Proposed additions:**
```hcl
output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags"
  value       = aws_ecrpublic_repository.repo.tags_all
}

output "catalog_data" {
  description = "The catalog data for the repository"
  value = try({
    about_text        = aws_ecrpublic_repository.repo.catalog_data[0].about_text
    architectures     = aws_ecrpublic_repository.repo.catalog_data[0].architectures
    description       = aws_ecrpublic_repository.repo.catalog_data[0].description
    operating_systems = aws_ecrpublic_repository.repo.catalog_data[0].operating_systems
    usage_text        = aws_ecrpublic_repository.repo.catalog_data[0].usage_text
  }, null)
}

output "gallery_url" {
  description = "The URL to the repository in ECR Public Gallery"
  value       = "https://gallery.ecr.aws/${aws_ecrpublic_repository.repo.registry_id}/${aws_ecrpublic_repository.repo.repository_name}"
}
```

---

### 6. Complete Example with Policy (HIGH PRIORITY)

**Current state:** Two examples (using_objects, using_variables)

**Proposed addition:** `examples/with_repository_policy/`

```hcl
module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-repo"

  create_repository_policy = true
  repository_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowPush"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:aws:iam::123456789012:role/CICDRole"
        }
        Action = [
          "ecr-public:BatchCheckLayerAvailability",
          "ecr-public:PutImage",
          "ecr-public:InitiateLayerUpload",
          "ecr-public:UploadLayerPart",
          "ecr-public:CompleteLayerUpload"
        ]
      }
    ]
  })

  catalog_data = {
    description       = "Public container with CI/CD push access"
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }
}
```

---

## Features NOT Applicable to ECR Public

The following features are available for private ECR but **NOT available for ECR Public**:

| Feature | Private ECR | ECR Public | Reason |
|---------|-------------|------------|--------|
| Lifecycle policies | Yes | No | AWS API limitation |
| Image scanning | Yes | No | AWS API limitation |
| Tag immutability | Yes | No | AWS API limitation |
| KMS encryption | Yes | No | AWS-managed encryption only |
| Pull-through cache | Yes | No | Not applicable to public |
| Replication rules | Yes | Limited | Auto-replication only |
| Registry policy | Yes | No | Different API |
| VPC endpoints | Yes | No | Public by design |

---

## Implementation Priority Matrix

| Feature | Priority | Effort | Impact | Recommended |
|---------|----------|--------|--------|-------------|
| Repository Policy | High | Low | High | Yes - Phase 1 |
| Example with Policy | High | Low | Medium | Yes - Phase 1 |
| Enhanced Outputs | Low | Low | Medium | Yes - Phase 1 |
| Auth Token Example | Medium | Low | Low | Yes - Phase 2 |
| Multiple Repos | Medium | Medium | Medium | Consider for Phase 2 |

---

## Recommended Implementation Plan

### Phase 1: Repository Policy Support
1. Add `aws_ecrpublic_repository_policy` resource
2. Add related variables (`create_repository_policy`, `repository_policy`)
3. Add policy-related outputs
4. Create `examples/with_repository_policy/` example
5. Update documentation and tests

### Phase 2: Enhanced Documentation
1. Add authorization token usage example
2. Add `gallery_url` and `tags_all` outputs
3. Add `catalog_data` composite output
4. Document ECR Public vs Private feature differences

### Phase 3: Consider Multiple Repository Support
1. Evaluate community demand
2. Consider separate submodule approach
3. Document module-level `for_each` pattern as alternative

---

## Sources

- [aws_ecrpublic_repository](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository)
- [aws_ecrpublic_repository_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository_policy)
- [aws_ecrpublic_authorization_token](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_authorization_token)
- [terraform-aws-modules/ecr](https://github.com/terraform-aws-modules/terraform-aws-ecr)
- [ECR Public Gallery](https://gallery.ecr.aws/)
- [Amazon ECR Public Documentation](https://docs.aws.amazon.com/AmazonECR/latest/public/what-is-ecr.html)
