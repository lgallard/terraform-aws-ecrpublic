# TFLint configuration for terraform-aws-ecrpublic
# Keep this lightweight for pre-commit: catch deprecated lookup usage without
# turning legacy naming/example conventions into required changes.

config {
  disabled_by_default = true
}

rule "terraform_deprecated_lookup" {
  enabled = true
}
