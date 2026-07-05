# Provider configuration for multiple repositories example

terraform {
  required_version = ">= 1.2, < 2.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0, < 7.0"
    }
  }
}

# ECR Public repositories must be created in us-east-1
provider "aws" {
  region = "us-east-1"
}
