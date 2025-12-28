# Using objects example

```
module "public-ecr" {

  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data = {
    about_text        = "# Public repo\nPut your description here using Markdown format"
    architectures     = ["Linux"]
    description       = "Description"
    logo_image_blob   = filebase64("image.png")
    operating_systems = ["ARM"]
    usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"
  }
}
```
<!-- BEGIN_TF_DOCS -->


## Requirements

No requirements.

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
