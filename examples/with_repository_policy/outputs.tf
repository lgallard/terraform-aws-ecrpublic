output "repository_uri" {
  description = "The URI of the repository"
  value       = module.public-ecr.repository_uri
}

output "repository_name" {
  description = "The name of the repository"
  value       = module.public-ecr.repository_name
}

output "registry_id" {
  description = "The registry ID where the repository was created"
  value       = module.public-ecr.registry_id
}

output "repository_arn" {
  description = "The ARN of the repository"
  value       = module.public-ecr.repository_arn
}

output "repository_policy" {
  description = "The repository policy JSON"
  value       = module.public-ecr.repository_policy
}
