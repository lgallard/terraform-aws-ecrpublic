# Using variables example
This example creates an public ECR registry using variables.

```
module "public-ecr" {

  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data_about_text        = "# Public repo\nPut your description here using Markdown format"
  catalog_data_architectures     = ["Linux"]
  catalog_data_description       = "Description"
  catalog_data_logo_image_blob   = filebase64("image.png")
  catalog_data_operating_systems = ["ARM"]
  catalog_data_usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"

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

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | A detailed description of the contents of the repository | `string` | `"# Public repo\nPut your description here using Markdown format"` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | The system architecture that the images in the repository are compatible with | `list(string)` | <pre>[<br/>  "x86-64"<br/>]</pre> | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | A short description of the contents of the repository | `string` | `"Description"` | no |
| <a name="input_catalog_data_logo_image_blob"></a> [catalog\_data\_logo\_image\_blob](#input\_catalog\_data\_logo\_image\_blob) | The base64-encoded repository logo payload | `string` | `null` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | The operating systems that the images in the repository are compatible with | `list(string)` | <pre>[<br/>  "Linux"<br/>]</pre> | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Detailed information on how to use the contents of the repository | `string` | `"# Usage\n How to use you image goes here. Use Markdown format"` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository | `string` | `"lgallard-public-repo"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_catalog_data"></a> [catalog\_data](#output\_catalog\_data) | The catalog data configuration for the repository |
| <a name="output_gallery_url"></a> [gallery\_url](#output\_gallery\_url) | The URL to the repository in ECR Public Gallery |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | Name of the repository |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags |

<!-- END_TF_DOCS -->
