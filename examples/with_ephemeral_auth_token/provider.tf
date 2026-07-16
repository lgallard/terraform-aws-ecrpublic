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
