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

output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags"
  value       = module.public-ecr.tags_all
}

output "gallery_url" {
  description = "The URL to the repository in ECR Public Gallery"
  value       = module.public-ecr.gallery_url
}

output "catalog_data" {
  description = "The catalog data configuration for the repository"
  value       = module.public-ecr.catalog_data
}
