# Issue: Add aws_ecrpublic_repository_policy resource support

**Title:** `feat: Add aws_ecrpublic_repository_policy resource support`

**Labels:** `enhancement`, `priority: high`

---

## Summary
Add support for the `aws_ecrpublic_repository_policy` resource to enable IAM-based access control for ECR Public repositories.

## Background
While ECR Public repositories are publicly readable by design, repository policies allow controlling who can **push** images. This is essential for:
- CI/CD pipeline access control
- Cross-account image push permissions
- Restricting push access to specific IAM roles/users

## Proposed Implementation

### New Variables
```hcl
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
```

### New Resource
```hcl
resource "aws_ecrpublic_repository_policy" "this" {
  count           = var.create_repository_policy ? 1 : 0
  repository_name = aws_ecrpublic_repository.repo.repository_name
  policy          = var.repository_policy
}
```

### New Outputs
```hcl
output "repository_policy" {
  description = "The repository policy JSON"
  value       = try(aws_ecrpublic_repository_policy.this[0].policy, null)
}
```

## Priority
**High** - This provides feature parity with [terraform-aws-modules/ecr](https://github.com/terraform-aws-modules/terraform-aws-ecr) and enables enterprise-ready access control.

## References
- [aws_ecrpublic_repository_policy documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository_policy)
- Feature Analysis: `FEATURE_ANALYSIS.md`

## Tasks
- [ ] Add `create_repository_policy` variable with validation
- [ ] Add `repository_policy` variable
- [ ] Add `aws_ecrpublic_repository_policy` resource with conditional creation
- [ ] Add `repository_policy` output
- [ ] Add tests for policy resource
- [ ] Update documentation
