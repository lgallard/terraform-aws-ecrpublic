# Ephemeral authorization token for ECR Public.
# Token values are available during the Terraform operation but are not persisted
# in Terraform state or plan files.
ephemeral "aws_ecrpublic_authorization_token" "token" {
  count = var.run_docker_login ? 1 : 0
}

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

# Optional example of consuming the ephemeral token during apply without storing
# it in state. Disabled by default so validation and example applies do not run
# Docker login unless explicitly requested.
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
