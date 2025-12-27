# Issue: Add example with repository policy configuration

**Title:** `docs: Add example with repository policy configuration`

**Labels:** `documentation`, `priority: high`

---

## Summary
Create a new example directory demonstrating how to use the repository policy feature with common CI/CD access patterns.

## Background
Once repository policy support is added (see issue #001), users will need clear examples showing:
- How to grant push access to CI/CD roles
- Cross-account access patterns
- Common policy patterns for ECR Public

## Proposed Implementation

### New Example Directory
Create `examples/with_repository_policy/` containing:

**main.tf:**
```hcl
provider "aws" {
  region = "us-east-1"
}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-app"

  create_repository_policy = true
  repository_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowPushFromCICD"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:aws:iam::123456789012:role/GitHubActionsRole"
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
    description       = "Public container image with CI/CD push access"
    about_text        = "# My Public App\n\nThis image is automatically built and published via CI/CD."
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }

  tags = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}
```

**outputs.tf:**
```hcl
output "repository_uri" {
  description = "The URI of the repository"
  value       = module.public-ecr.repository_uri
}

output "repository_policy" {
  description = "The repository policy"
  value       = module.public-ecr.repository_policy
}
```

**README.md:**
Document the example with:
- Prerequisites
- Common policy patterns
- CI/CD integration guidance

## Priority
**High** - Examples are essential for adoption and understanding

## Dependencies
- Requires: Issue #001 (Repository Policy Support)

## Tasks
- [ ] Create `examples/with_repository_policy/` directory
- [ ] Add `main.tf` with policy configuration
- [ ] Add `outputs.tf`
- [ ] Add `provider.tf`
- [ ] Add `README.md` with documentation
- [ ] Add test for the new example
