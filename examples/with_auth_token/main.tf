provider "aws" {
  region = "us-east-1"
}

# Authorization token for ECR Public
# Note: Must be used in us-east-1 region
data "aws_ecrpublic_authorization_token" "token" {}

module "public-ecr" {
  source = "../.."

  repository_name = var.repository_name

  catalog_data = {
    description       = var.catalog_data_description
    about_text        = var.catalog_data_about_text
    usage_text        = var.catalog_data_usage_text
    architectures     = var.catalog_data_architectures
    operating_systems = var.catalog_data_operating_systems
  }
}
