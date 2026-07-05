# Multiple repositories example

This example creates multiple ECR Public repositories using Terraform's native `for_each` support at the module call site.

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
provider "aws" {
  region = "us-east-1"
}

locals {
  repositories = {
    frontend = {
      description       = "Frontend application container image"
      architectures     = ["x86-64", "ARM 64"]
      operating_systems = ["Linux"]
    }
    "api-server" = {
      description       = "REST API server container image"
      architectures     = ["x86-64"]
      operating_systems = ["Linux"]
    }
    worker = {
      description       = "Background worker container image"
      architectures     = ["x86-64"]
      operating_systems = ["Linux"]
    }
  }
}

module "public-ecr" {
  source   = "lgallard/ecrpublic/aws"
  for_each = local.repositories

  repository_name = each.key

  catalog_data_description       = each.value.description
  catalog_data_architectures     = each.value.architectures
  catalog_data_operating_systems = each.value.operating_systems
  catalog_data_about_text        = "# ${title(each.key)} Application\n\nContainer image for the ${each.key} component."
  catalog_data_usage_text        = "# Usage\n\ndocker pull public.ecr.aws/your-registry/${each.key}:latest"
}
```

## Notes

- ECR Public repositories must be managed in `us-east-1`.
- The live Terraform example in this directory keeps `source = "../.."` so CI validates the checked-out module.
- Use descriptive `for_each` keys because they become repository names.

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
| <a name="module_public-ecr-object"></a> [public-ecr-object](#module\_public-ecr-object) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_additional_tags"></a> [additional\_tags](#input\_additional\_tags) | Additional tags to apply to all repositories | `map(string)` | `{}` | no |
| <a name="input_common_architectures"></a> [common\_architectures](#input\_common\_architectures) | Default architectures for all repositories | `list(string)` | <pre>[<br/>  "x86-64"<br/>]</pre> | no |
| <a name="input_common_operating_systems"></a> [common\_operating\_systems](#input\_common\_operating\_systems) | Default operating systems for all repositories | `list(string)` | <pre>[<br/>  "Linux"<br/>]</pre> | no |
| <a name="input_enable_object_example"></a> [enable\_object\_example](#input\_enable\_object\_example) | Whether to enable the object-based configuration example | `bool` | `false` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Environment name (e.g., dev, staging, production) | `string` | `"production"` | no |
| <a name="input_project_name"></a> [project\_name](#input\_project\_name) | Project name for resource tagging and organization | `string` | `"microservices-platform"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_api_server_repository_uri"></a> [api\_server\_repository\_uri](#output\_api\_server\_repository\_uri) | URI of the API server repository |
| <a name="output_docker_pull_commands"></a> [docker\_pull\_commands](#output\_docker\_pull\_commands) | Docker pull commands for all repositories |
| <a name="output_frontend_repository_uri"></a> [frontend\_repository\_uri](#output\_frontend\_repository\_uri) | URI of the frontend repository |
| <a name="output_gallery_urls"></a> [gallery\_urls](#output\_gallery\_urls) | ECR Public Gallery URLs for all repositories |
| <a name="output_object_repositories"></a> [object\_repositories](#output\_object\_repositories) | Repository information for object-based configuration example |
| <a name="output_registry_ids"></a> [registry\_ids](#output\_registry\_ids) | Map of repository names to their registry IDs |
| <a name="output_repositories"></a> [repositories](#output\_repositories) | Complete repository information for all created repositories |
| <a name="output_repository_arns"></a> [repository\_arns](#output\_repository\_arns) | Map of repository names to their ARNs |
| <a name="output_repository_summary"></a> [repository\_summary](#output\_repository\_summary) | Summary of created repositories |
| <a name="output_repository_uris"></a> [repository\_uris](#output\_repository\_uris) | Map of repository names to their URIs |
| <a name="output_worker_repository_uri"></a> [worker\_repository\_uri](#output\_worker\_repository\_uri) | URI of the worker repository |

<!-- END_TF_DOCS -->
