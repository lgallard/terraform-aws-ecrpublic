![Terraform](https://lgallardo.com/images/terraform.jpg)
# terraform-aws-ecrpublic

Terraform module to create and manage public [Amazon ECR Public](https://aws.amazon.com/ecr/) repositories for sharing container images in the [ECR Public Gallery](https://gallery.ecr.aws/).

## Terraform Registry compatibility

This repository follows Terraform Registry module expectations:

- The GitHub repository is named `terraform-aws-ecrpublic` using the `terraform-<PROVIDER>-<NAME>` pattern.
- Root module files live at the repository root: `main.tf`, `variables.tf`, `outputs.tf`, and `versions.tf`.
- Reusable examples live under [`examples/`](examples/) and their documentation shows copy/paste usage with the public registry source address `lgallard/ecrpublic/aws`.
- Inputs and outputs are generated with `terraform-docs` in this README and in example READMEs.
- This repository uses Release Please to publish semantic version tags.

ECR Public repositories and authorization tokens are managed through `us-east-1`; configure the AWS provider for `us-east-1` when using this module.

## Complete example

The [`examples/using_variables`](examples/using_variables/) directory is the complete minimal example for Terraform Registry users. Its live Terraform files use `source = "../.."` so maintainers can inspect and exercise the checked-out module locally when needed. The copy/paste snippet below uses the public Registry source address.

```hcl
provider "aws" {
  region = "us-east-1"
}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-repo"

  catalog_data_description       = "Public container image"
  catalog_data_about_text        = "# My Public Repository\n\nContainer image published to ECR Public."
  catalog_data_usage_text        = "# Usage\n\ndocker pull public.ecr.aws/my-registry/my-public-repo:latest"
  catalog_data_architectures     = ["x86-64"]
  catalog_data_operating_systems = ["Linux"]
}
```

## Examples

| Example | Purpose |
|---|---|
| [`using_variables`](examples/using_variables/) | Complete minimal example using individual catalog-data variables. |
| [`using_objects`](examples/using_objects/) | Configure catalog data through the `catalog_data` object. |
| [`with_repository_policy`](examples/with_repository_policy/) | Attach an ECR Public repository policy for push access. |
| [`with_ephemeral_auth_token`](examples/with_ephemeral_auth_token/) | Preferred short-lived authorization token example that avoids storing credentials in state. |
| [`with_auth_token`](examples/with_auth_token/) | Legacy data source authorization token example with state-storage caveats. |
| [`multiple_repositories`](examples/multiple_repositories/) | Create multiple repositories with module-level `for_each`. |

## Basic usage

### Object-based catalog data

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

### Variable-based catalog data

```hcl
provider "aws" {
  region = "us-east-1"
}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data_about_text        = "# Public repo\nPut your description here using Markdown format"
  catalog_data_architectures     = ["x86-64"]
  catalog_data_description       = "Description"
  catalog_data_operating_systems = ["Linux"]
  catalog_data_usage_text        = "# Usage\nHow to use your image goes here. Use Markdown format."
}
```

## Authorization tokens for image push

ECR Public repositories require authentication tokens for programmatic operations like pushing container images. Authorization tokens must be retrieved from `us-east-1` and expire after 12 hours.

Prefer one of these patterns:

- **AWS CLI login** for CI/CD jobs and local pushes. This avoids putting token values in Terraform configuration, plan files, or state.
- **Ephemeral `aws_ecrpublic_authorization_token` resource** when a Terraform operation must pass short-lived registry credentials to a provider or provisioner. See [`examples/with_ephemeral_auth_token`](examples/with_ephemeral_auth_token/). Requires Terraform/OpenTofu `>= 1.10` and AWS provider `>= 6.31`.
- **Legacy `data.aws_ecrpublic_authorization_token` data source** only for older Terraform compatibility. Data source attributes, including sensitive token/password values, are stored in Terraform state. See [`examples/with_auth_token`](examples/with_auth_token/).

```bash
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
```

## Validation and maintenance

This repository keeps lightweight pre-commit checks for Terraform formatting, validation, linting, and generated docs. It intentionally does not maintain the legacy Terratest suite or broader CI/security/test workflow stack. Changes are reviewed through pre-commit, maintainer inspection, AI-assisted/codebot review, and user reports.

When validating changes locally, use the smallest direct Terraform commands that match the change. For example:

```bash
pre-commit run --all-files
terraform fmt -check -recursive
terraform init -backend=false -input=false
terraform validate
```

ECR Public resources and authorization tokens are managed through `us-east-1`. Do not run Terraform operations that create, update, or delete AWS resources unless you explicitly intend to manage live infrastructure.

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

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_ecrpublic_repository.repo](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository) | resource |
| [aws_ecrpublic_repository_policy.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data"></a> [catalog\_data](#input\_catalog\_data) | Catalog data configuration for the repository. | `any` | `{}` | no |
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format. | `string` | `null` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: 'ARM', 'ARM 64', 'x86', 'x86-64'. | `list(string)` | `[]` | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery. | `string` | `null` | no |
| <a name="input_catalog_data_logo_image_blob"></a> [catalog\_data\_logo\_image\_blob](#input\_catalog\_data\_logo\_image\_blob) | The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute. | `string` | `null` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`. | `list(string)` | `[]` | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format. | `string` | `null` | no |
| <a name="input_create_repository_policy"></a> [create\_repository\_policy](#input\_create\_repository\_policy) | Whether to create a repository policy for controlling push access | `bool` | `false` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository. | `string` | n/a | yes |
| <a name="input_repository_policy"></a> [repository\_policy](#input\_repository\_policy) | The JSON policy document for the repository | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resource. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | Timeouts map. | `map(any)` | `{}` | no |
| <a name="input_timeouts_delete"></a> [timeouts\_delete](#input\_timeouts\_delete) | How long to wait for a repository to be deleted. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | Full ARN of the repository |
| <a name="output_catalog_data"></a> [catalog\_data](#output\_catalog\_data) | The catalog data configuration for the repository |
| <a name="output_gallery_url"></a> [gallery\_url](#output\_gallery\_url) | The URL to the repository in ECR Public Gallery |
| <a name="output_id"></a> [id](#output\_id) | The registry ID where the repository was created. |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created. |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | The repository name. |
| <a name="output_repository_policy"></a> [repository\_policy](#output\_repository\_policy) | The repository policy JSON |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository. |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags |

<!-- END_TF_DOCS -->
