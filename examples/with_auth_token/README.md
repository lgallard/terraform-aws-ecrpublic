# ECR Public with legacy authorization token data source example

This example creates an ECR Public repository and retrieves an authorization token with the legacy `data.aws_ecrpublic_authorization_token` data source.

Use this example only when you need compatibility with Terraform versions that do not support ephemeral resources. Data source values are persisted in Terraform state, including sensitive attributes such as `authorization_token` and `password`, even if an output is marked `sensitive`.

For Terraform/OpenTofu versions that support ephemeral resources, prefer [`../with_ephemeral_auth_token`](../with_ephemeral_auth_token/) so short-lived registry credentials are not stored in state or plan files.

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
provider "aws" {
  region = "us-east-1"
}

# Authorization tokens for ECR Public must be requested from us-east-1.
# WARNING: Data source values are stored in Terraform state.
data "aws_ecrpublic_authorization_token" "token" {}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-application"

  catalog_data = {
    description       = "My application container"
    about_text        = "# My Application\n\nProduction-ready container for my public application."
    usage_text        = "# Usage\n\ndocker pull public.ecr.aws/your-registry/my-application:latest"
    architectures     = ["x86-64"]
    operating_systems = ["Linux"]
  }
}

output "repository_uri" {
  value = module.public-ecr.repository_uri
}

output "token_expires_at" {
  value = data.aws_ecrpublic_authorization_token.token.expires_at
}
```

## Docker login

Prefer the AWS CLI for Docker login because it avoids persisting token values in Terraform state:

```bash
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
```

## Notes

- Authorization tokens expire after 12 hours.
- Do not output `authorization_token`, `password`, or decoded credentials from this example.
- Sensitive Terraform outputs are still stored in state. Marking an output `sensitive` only hides it from normal CLI display.
- The live Terraform example in this directory keeps `source = "../.."` so CI validates the checked-out module.

<!-- BEGIN_TF_DOCS -->


## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.2, < 2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0, < 7.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.55.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_public-ecr"></a> [public-ecr](#module\_public-ecr) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_ecrpublic_authorization_token.token](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_authorization_token) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | Public about text in markdown format | `string` | `"# My Application\n\n## Description\nThis container provides a sample application for demonstrating ECR Public authentication.\n\n## Features\n- Lightweight container image\n- Production-ready configuration\n- Easy deployment\n"` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | Supported architectures for container images | `list(string)` | <pre>[<br/>  "x86-64"<br/>]</pre> | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | Public description visible in ECR Public Gallery | `string` | `"My application container"` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | Supported operating systems for container images | `list(string)` | <pre>[<br/>  "Linux"<br/>]</pre> | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Public usage instructions in markdown format | `string` | `"# Usage\n\n## Authentication\n`<pre>bash\n# Get authorization token and login to ECR Public\naws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws\n</pre>\n\n## Pull Image\n<pre>bash\ndocker pull public.ecr.aws/your-registry/my-app:latest\n</pre>\n\n## Push Image\n<pre>bash\n# Tag your image\ndocker tag my-app:latest public.ecr.aws/your-registry/my-app:latest\n\n# Push to ECR Public\ndocker push public.ecr.aws/your-registry/my-app:latest\n</pre>\n" | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository | `string` | `"my-app"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_cli_login_command"></a> [aws\_cli\_login\_command](#output\_aws\_cli\_login\_command) | AWS CLI command to login to ECR Public without storing Terraform token values |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | Name of the repository |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_token_expires_at"></a> [token\_expires\_at](#output\_token\_expires\_at) | Token expiration timestamp from the legacy data source |

<!-- END_TF_DOCS -->
