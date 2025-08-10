# TFLint configuration for terraform-aws-ecrpublic
# https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/config.md

config {
  # Enable all rules by default
  disabled_by_default = false

  # Plugin cache directory
  plugin_dir = "~/.tflint.d/plugins"
}

# AWS plugin for additional AWS-specific rules
plugin "aws" {
  enabled = true
  version = "0.24.1"
  source  = "github.com/terraform-linters/tflint-ruleset-aws"
}

# Enable specific rules for deprecated syntax
rule "terraform_deprecated_lookup" {
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
