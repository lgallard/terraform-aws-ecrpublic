# Using objects example

This example creates one ECR Public repository and configures gallery catalog data through the `catalog_data` object.

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
provider "aws" {
  region = "us-east-1"
}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data = {
    about_text        = "# Public repo\nPut your description here using Markdown format"
    architectures     = ["x86-64"]
    description       = "Description"
    operating_systems = ["Linux"]
    usage_text        = "# Usage\nHow to use your image goes here. Use Markdown format."
  }
}
```

## Local example notes

The live Terraform example in this directory keeps `source = "../.."` so CI validates the checked-out module. It also includes an optional `image.png` logo file and safely reads it only when present.

<!-- BEGIN_TF_DOCS -->


## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.2, < 2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0, < 7.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_public-ecr"></a> [public-ecr](#module\_public-ecr) | ../.. | n/a |

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.

<!-- END_TF_DOCS -->
