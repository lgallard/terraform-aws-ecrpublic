config {
  module = true
  force = false
  disabled_by_default = false
}

plugin "aws" {
  enabled = true
  version = "0.30.0"
  source  = "github.com/terraform-linters/tflint-ruleset-aws"
}

plugin "terraform" {
  enabled = true
  version = "0.6.0"
  source  = "github.com/terraform-linters/tflint-ruleset-terraform"
}

# ECR Public specific rules - using actual TFLint AWS rules
# Note: Custom ECR Public rules would need to be implemented as custom TFLint plugins

# General AWS and Terraform rules
rule "terraform_deprecated_interpolation" {
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = true
}

rule "terraform_comment_syntax" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_module_pinned_source" {
  enabled = true
}

rule "terraform_naming_convention" {
  enabled = true
  format  = "snake_case"
}

rule "terraform_standard_module_structure" {
  enabled = true
}

# AWS specific rules relevant to ECR Public
rule "aws_resource_missing_tags" {
  enabled = false  # ECR Public repositories don't support tags
}

rule "aws_ecr_repository_image_tag_mutability" {
  enabled = false  # ECR Public has different configuration than ECR
}

rule "aws_ecr_repository_image_scan_on_push" {
  enabled = false  # ECR Public has different scan configuration
}