![Terraform](https://lgallardo.com/images/terraform.jpg)
# terraform-aws-ecrpublic
Terraform module to create a Public [AWS ECR](https://aws.amazon.com/ecr) to share images in the [ECR Public Gallery](https://gallery.ecr.aws).

## Usage
You can use this module to create a public ECR registry using objects definition, or using the variables approach:

Check the [examples](examples/) for the **using objects** and the **using variables* snippets.

### Using Objects example
This example creates an public ECR registry:

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

### Using variables
This example creates an public ECR registry using variables

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
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| aws | n/a |

## Modules

No Modules.

## Resources

| Name |
|------|
| [aws_ecrpublic_repository](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository) |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| catalog\_data | Catalog data configuration for the repository. | `any` | `{}` | no |
| catalog\_data\_about\_text | A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format. | `string` | `null` | no |
| catalog\_data\_architectures | The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`. | `list(string)` | `[]` | no |
| catalog\_data\_description | A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery. | `string` | `null` | no |
| catalog\_data\_logo\_image\_blob | The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute. | `string` | `null` | no |
| catalog\_data\_operating\_systems | The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. 'ARM', 'ARM 64', 'x86', 'x86-64'. | `list(string)` | `null` | no |
| catalog\_data\_usage\_text | Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format. | `string` | `null` | no |
| repository\_name | Name of the repository. | `string` | n/a | yes |
| timeouts | Timeouts map. | `map` | `{}` | no |
| timeouts\_delete | How long to wait for a repository to be deleted. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| arn | Full ARN of the repository |
| id | The repository name. |
| registry\_id | The registry ID where the repository was created. |
| repository\_uri | The URI of the repository. |
