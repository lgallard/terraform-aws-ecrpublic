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
