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

output "authorization_token" {
  description = "The authorization token (base64 encoded)"
  value       = data.aws_ecrpublic_authorization_token.token.authorization_token
  sensitive   = true
}

output "token_expires_at" {
  description = "Token expiration timestamp"
  value       = data.aws_ecrpublic_authorization_token.token.expires_at
}

output "docker_login_command" {
  description = "Docker login command for ECR Public (use with authorization_token output)"
  value       = "echo '<authorization_token>' | docker login --username AWS --password-stdin public.ecr.aws"
  sensitive   = false
}

output "aws_cli_login_command" {
  description = "AWS CLI command to login to ECR Public"
  value       = "aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws"
  sensitive   = false
}
