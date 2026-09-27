# Read ECR Public image metadata example

This example reads image metadata from an existing ECR Public repository with `data.aws_ecrpublic_images`.

Use this after images have already been pushed. A newly-created or empty repository may return no images, and tying image discovery to repository creation can make plans unstable when image pushes happen outside Terraform.

## Copy/paste usage

Use this with repositories created by `lgallard/ecrpublic/aws` or any existing ECR Public repository:

```hcl
provider "aws" {
  region = "us-east-1"
}

# Requires hashicorp/aws >= 6.19.
data "aws_ecrpublic_images" "selected" {
  repository_name = "my-application"

  # Optional: restrict lookup to a known tag or digest.
  image_ids {
    image_tag = "latest"
  }
}

output "image_digests" {
  value = [for image in data.aws_ecrpublic_images.selected.images : image.digest if image.digest != null]
}

output "image_tags" {
  value = distinct(flatten([for image in data.aws_ecrpublic_images.selected.images : image.tags]))
}
```

## Notes

- ECR Public API operations should use `us-east-1`.
- The repository must already exist and should contain pushed images before this data source is useful.
- Leave `image_ids` empty to list repository images, or provide tag/digest filters when you need a specific image.
- Amazon ECR Public currently supports up to 100,000 images per repository and up to 1,000 tags per image by default.

<!-- BEGIN_TF_DOCS -->


## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3, < 2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 6.19, < 7.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.66.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_ecrpublic_images.selected](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_images) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_image_ids"></a> [image\_ids](#input\_image\_ids) | Optional image filters. Each object can specify an image\_tag, an image\_digest, or both. Leave empty to list repository images. | <pre>list(object({<br/>    image_tag    = optional(string)<br/>    image_digest = optional(string)<br/>  }))</pre> | `[]` | no |
| <a name="input_registry_id"></a> [registry\_id](#input\_registry\_id) | AWS account ID associated with the public registry that contains the repository | `string` | `null` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the existing ECR Public repository to inspect | `string` | `"my-application"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_image_digests"></a> [image\_digests](#output\_image\_digests) | Image digests returned by the ECR Public images data source |
| <a name="output_image_tags"></a> [image\_tags](#output\_image\_tags) | Deduplicated flat list of all tags across every returned image |
| <a name="output_images"></a> [images](#output\_images) | Image metadata returned by the ECR Public images data source |

<!-- END_TF_DOCS -->
