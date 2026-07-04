# ECR Public with authorization token example

This example creates an ECR Public repository and retrieves an authorization token for Docker login and push workflows.

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
provider "aws" {
  region = "us-east-1"
}

# Authorization tokens for ECR Public must be requested from us-east-1.
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

output "authorization_token" {
  value     = data.aws_ecrpublic_authorization_token.token.authorization_token
  sensitive = true
}
```

## Docker login

```bash
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
```

## Notes

- Authorization tokens expire after 12 hours.
- Treat token outputs as sensitive and avoid logging them.
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
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.53.0 |

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
| <a name="output_authorization_token"></a> [authorization\_token](#output\_authorization\_token) | The authorization token (base64 encoded) |
| <a name="output_aws_cli_login_command"></a> [aws\_cli\_login\_command](#output\_aws\_cli\_login\_command) | AWS CLI command to login to ECR Public |
| <a name="output_docker_login_command"></a> [docker\_login\_command](#output\_docker\_login\_command) | Docker login command for ECR Public (use with authorization\_token output) |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | Name of the repository |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_token_expires_at"></a> [token\_expires\_at](#output\_token\_expires\_at) | Token expiration timestamp |

<!-- END_TF_DOCS -->
