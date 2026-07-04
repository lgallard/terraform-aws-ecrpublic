# ECR Public repository with policy example

This example creates an ECR Public repository and attaches an optional repository policy to control push access for CI/CD systems or trusted AWS principals.

## Copy/paste usage

Use this Registry source address in consumer configurations:

```hcl
provider "aws" {
  region = "us-east-1"
}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-app"

  create_repository_policy = true
  repository_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowPushFromCICD"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:aws:iam::123456789012:role/GitHubActionsRole"
        }
        Action = [
          "ecr-public:BatchCheckLayerAvailability",
          "ecr-public:PutImage",
          "ecr-public:InitiateLayerUpload",
          "ecr-public:UploadLayerPart",
          "ecr-public:CompleteLayerUpload"
        ]
      }
    ]
  })

  catalog_data = {
    description       = "Public container image with CI/CD push access"
    about_text        = "# My Public App\n\nThis image is automatically built and published via CI/CD."
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }
}
```

## Notes

- ECR Public repositories must be managed in `us-east-1`.
- Public pull access is inherent to ECR Public; repository policies here control push-related access.
- Replace the example IAM principal ARN with a real role from your account.
- The live Terraform example in this directory keeps `source = "../.."` so CI validates the checked-out module.

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

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data"></a> [catalog\_data](#input\_catalog\_data) | Catalog data configuration for the repository | <pre>object({<br/>    description       = optional(string)<br/>    about_text        = optional(string)<br/>    usage_text        = optional(string)<br/>    architectures     = optional(list(string))<br/>    operating_systems = optional(list(string))<br/>    logo_image_blob   = optional(string)<br/>  })</pre> | <pre>{<br/>  "about_text": "# My Public App\n\nThis image is automatically built and published via CI/CD.",<br/>  "architectures": [<br/>    "x86-64",<br/>    "ARM 64"<br/>  ],<br/>  "description": "Public container image with CI/CD push access",<br/>  "operating_systems": [<br/>    "Linux"<br/>  ]<br/>}</pre> | no |
| <a name="input_create_repository_policy"></a> [create\_repository\_policy](#input\_create\_repository\_policy) | Whether to create a repository policy for controlling push access | `bool` | `true` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the ECR Public repository | `string` | `"my-public-app"` | no |
| <a name="input_repository_policy"></a> [repository\_policy](#input\_repository\_policy) | The JSON policy document for the repository. This example shows CI/CD push access. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the repository | `map(string)` | <pre>{<br/>  "Environment": "production",<br/>  "ManagedBy": "terraform"<br/>}</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | The ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | The name of the repository |
| <a name="output_repository_policy"></a> [repository\_policy](#output\_repository\_policy) | The repository policy JSON |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository |

<!-- END_TF_DOCS -->
