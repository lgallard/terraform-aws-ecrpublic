output "repository_name" {
  description = "Name of the repository"
  value       = module.public-ecr.repository_name
}

output "repository_uri" {
  description = "The URI of the repository"
  value       = module.public-ecr.repository_uri
}

output "repository_url" {
  description = "The URL of the repository"
  value       = module.public-ecr.repository_url
}

output "repository_arn" {
  description = "Full ARN of the repository"
  value       = module.public-ecr.repository_arn
}

output "registry_id" {
  description = "The registry ID where the repository was created"
  value       = module.public-ecr.registry_id
}

output "aws_cli_login_command" {
  description = "AWS CLI command to login to ECR Public without storing Terraform token values"
  value       = "aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws"
}
