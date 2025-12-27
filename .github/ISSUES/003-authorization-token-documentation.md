# Issue: Document aws_ecrpublic_authorization_token data source usage

**Title:** `docs: Add authorization token usage example and documentation`

**Labels:** `documentation`, `priority: medium`

---

## Summary
Add documentation and an example showing how to use the `aws_ecrpublic_authorization_token` data source alongside this module for programmatic image pushing.

## Background
The `aws_ecrpublic_authorization_token` data source provides authentication tokens needed to push images to ECR Public repositories. While not part of this module directly, it's commonly used alongside it and should be documented.

## Proposed Implementation

### New Example Directory
Create `examples/with_auth_token/` containing:

**main.tf:**
```hcl
provider "aws" {
  region = "us-east-1"
}

# Authorization token for ECR Public
# Note: Must be used in us-east-1 region
data "aws_ecrpublic_authorization_token" "token" {}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-app"

  catalog_data = {
    description       = "My application container"
    architectures     = ["x86-64"]
    operating_systems = ["Linux"]
  }
}
```

**outputs.tf:**
```hcl
output "repository_uri" {
  description = "The URI of the repository"
  value       = module.public-ecr.repository_uri
}

output "docker_login_command" {
  description = "Docker login command for ECR Public"
  value       = "echo $TOKEN | docker login --username AWS --password-stdin public.ecr.aws"
  sensitive   = false
}

output "authorization_token" {
  description = "The authorization token (base64 encoded)"
  value       = data.aws_ecrpublic_authorization_token.token.authorization_token
  sensitive   = true
}

output "token_expiration" {
  description = "Token expiration timestamp"
  value       = data.aws_ecrpublic_authorization_token.token.expires_at
}
```

### README Updates
Add a section to the main README explaining:
- How to get authorization tokens
- Docker login process
- Token expiration considerations
- Regional constraints (us-east-1 requirement)

## Priority
**Medium** - Useful for CI/CD integration but not core functionality

## References
- [aws_ecrpublic_authorization_token documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_authorization_token)
- Known issue: Data source must use us-east-1 region

## Tasks
- [ ] Create `examples/with_auth_token/` directory
- [ ] Add example Terraform configuration
- [ ] Add README.md for the example
- [ ] Update main README with auth token section
- [ ] Document regional constraints and workarounds
