# Terraform AWS ECR Public Module - Development Guidelines

## Overview

This repository contains the `lgallard/ecrpublic/aws` Terraform module for creating and managing Amazon ECR Public repositories and related repository policy/catalog metadata.

ECR Public is a global public gallery service with API operations handled through `us-east-1`. Provider examples should use `us-east-1` unless a specific provider behavior proves otherwise.

## Repository layout

- `main.tf` - ECR Public repository and optional repository policy resources
- `variables.tf` - module inputs and validation rules
- `outputs.tf` - module outputs
- `versions.tf` - Terraform and AWS provider constraints
- `examples/` - registry-facing examples
- `.github/workflows/pre-commit.yml` - lightweight Terraform formatting/validation/linting checks
- `.github/workflows/release-please.yml` - semantic version/tag automation
- `.github/workflows/claude*.yml` - AI/codebot automation

This repo intentionally keeps pre-commit as the lightweight automated quality gate. It no longer contains the legacy Terratest suite or broader CI/security/test workflow stack.

## Development rules

1. Keep the module focused on ECR Public repositories and ECR Public repository policies.
2. Preserve Terraform Registry compatibility:
   - root module files at the repo root
   - examples under `examples/`
   - README usage should use `source = "lgallard/ecrpublic/aws"`
3. Keep examples copy/paste friendly and configured for `us-east-1`.
4. Avoid storing short-lived authorization token secrets in Terraform state. Prefer AWS CLI login or the ephemeral token example where Terraform/OpenTofu compatibility allows it.
5. Do not run live Terraform operations that create, update, delete, or import AWS resources without explicit maintainer approval.
6. Keep Release Please and Claude/codebot workflows intact unless the task explicitly targets those workflows.

## Validation approach

The maintainer preference for this repository is lightweight pre-commit plus review instead of maintaining the old Terratest/CI/security workflow stack. Use:

- pre-commit for Terraform formatting, validation, linting, and generated docs
- direct local Terraform commands for targeted checks when useful
- AI/codebot review on pull requests
- maintainer inspection and user reports for behavior validation

Useful local commands when the touched files justify them:

```bash
pre-commit run --all-files
terraform fmt -check -recursive
terraform init -backend=false -input=false
terraform validate
```

For examples, run `terraform init -backend=false -input=false` and `terraform validate` inside the specific example directory being changed. Do not run commands that create or mutate AWS resources unless explicitly approved.

## ECR Public implementation notes

- ECR Public repository APIs must use `us-east-1`.
- Catalog data is publicly visible in the ECR Public Gallery; avoid secrets, internal-only wording, or sensitive operational details.
- `logo_image_blob` is a base64-encoded payload and can be large; keep examples small and clear.
- Repository policy support should remain optional and controlled by `create_repository_policy` plus `repository_policy`.
- Keep backward-compatible variable patterns unless a breaking change is explicitly planned and released.

## PR expectations

- Keep PRs focused and avoid mixing resource behavior changes with maintenance-only cleanup.
- Include a concise summary and targeted validation evidence.
- Reference the GitHub issue being closed.
- If codebot reports actionable findings, fix them and rerun the review loop.
