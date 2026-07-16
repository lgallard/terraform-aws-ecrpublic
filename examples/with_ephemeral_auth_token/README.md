# ECR Public with ephemeral authorization token example

This example creates an ECR Public repository and retrieves short-lived registry credentials with the ephemeral `aws_ecrpublic_authorization_token` resource.

Ephemeral resources are available during the current Terraform operation but are not written to Terraform state or plan files. This is preferred for short-lived ECR Public credentials when your Terraform/OpenTofu version and AWS provider version support ephemeral resources.

## Compatibility

- Terraform/OpenTofu: `>= 1.10, < 2.0`
- AWS provider: `>= 6.31, < 7.0`
- Region: `us-east-1`

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
terraform {
  required_version = ">= 1.10, < 2.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 6.31, < 7.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

variable "run_docker_login" {
  description = "Whether to run Docker login during apply."
  type        = bool
  default     = false
}

# Token values are not persisted in state or plan files.
ephemeral "aws_ecrpublic_authorization_token" "token" {
  count = var.run_docker_login ? 1 : 0
}

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

resource "terraform_data" "docker_login" {
  count = var.run_docker_login ? 1 : 0

  triggers_replace = [module.public-ecr.repository_uri]

  provisioner "local-exec" {
    interpreter = ["/bin/bash", "-c"]
    command     = "printf '%s' \"$ECR_PUBLIC_PASSWORD\" | docker login --username \"$ECR_PUBLIC_USERNAME\" --password-stdin public.ecr.aws"

    environment = {
      ECR_PUBLIC_USERNAME = ephemeral.aws_ecrpublic_authorization_token.token[0].user_name
      ECR_PUBLIC_PASSWORD = ephemeral.aws_ecrpublic_authorization_token.token[0].password
    }
  }
}
```

## Docker login

For most CI/CD workflows, the AWS CLI is still the simplest way to avoid putting token values in Terraform state:

```bash
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
```

## Notes

- Authorization tokens expire after 12 hours.
- Do not output token or password values. Ephemeral values cannot be persisted as normal outputs, and outputting credentials would defeat the purpose.
- The `run_docker_login` variable defaults to `false` so local validation and example applies do not fetch an authorization token or run Docker login unexpectedly.
- If you cannot use ephemeral resources, see [`../with_auth_token`](../with_auth_token/) for the legacy data source example and its state-storage caveat.
- The live Terraform example in this directory keeps `source = "../.."` so CI validates the checked-out module.

<!-- BEGIN_TF_DOCS -->


## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.10, < 2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 6.31, < 7.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_terraform"></a> [terraform](#provider\_terraform) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_public-ecr"></a> [public-ecr](#module\_public-ecr) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [terraform_data.docker_login](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/resources/data) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | Public about text in markdown format | `string` | `"# My Application\n\n## Description\nThis container provides a sample application for demonstrating ephemeral ECR Public authentication.\n\n## Features\n- Lightweight container image\n- Production-ready configuration\n- Auth token not persisted in Terraform state\n"` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | Supported architectures for container images | `list(string)` | <pre>[<br/>  "x86-64"<br/>]</pre> | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | Public description visible in ECR Public Gallery | `string` | `"My application container"` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | Supported operating systems for container images | `list(string)` | <pre>[<br/>  "Linux"<br/>]</pre> | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Public usage instructions in markdown format | `string` | `"# Usage\n\n## Authentication\n`<pre>bash\n# Preferred: Get a short-lived password through AWS CLI\naws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws\n</pre>\n\n## Pull Image\n<pre>bash\ndocker pull public.ecr.aws/your-registry/my-app:latest\n</pre>\n" | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository | `string` | `"my-app"` | no |
| <a name="input_run_docker_login"></a> [run\_docker\_login](#input\_run\_docker\_login) | Whether to run the local Docker login example during apply. Disabled by default to keep the example safe for validation. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_cli_login_command"></a> [aws\_cli\_login\_command](#output\_aws\_cli\_login\_command) | AWS CLI command to login to ECR Public without storing Terraform token values |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | Name of the repository |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |

<!-- END_TF_DOCS -->
