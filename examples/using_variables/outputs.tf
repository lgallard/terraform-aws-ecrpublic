output "repository_name" {
  description = "Name of the repository"
  value       = module.public-ecr.repository_name
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