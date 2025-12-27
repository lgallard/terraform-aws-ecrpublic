# Outputs for multiple repositories example

# Map of repository names to their URIs
output "repository_uris" {
  description = "Map of repository names to their URIs"
  value = {
    for k, v in module.public-ecr : k => v.repository_uri
  }
}

# Map of repository names to their ARNs
output "repository_arns" {
  description = "Map of repository names to their ARNs"
  value = {
    for k, v in module.public-ecr : k => v.arn
  }
}

# Map of repository names to their registry IDs
output "registry_ids" {
  description = "Map of repository names to their registry IDs"
  value = {
    for k, v in module.public-ecr : k => v.registry_id
  }
}

# Complete repository information
output "repositories" {
  description = "Complete repository information for all created repositories"
  value = {
    for name, repo in module.public-ecr : name => {
      name        = repo.repository_name
      uri         = repo.repository_uri
      arn         = repo.arn
      registry_id = repo.registry_id
    }
  }
}

# Object-based repositories (if enabled)
output "object_repositories" {
  description = "Repository information for object-based configuration example"
  value = var.enable_object_example ? {
    for name, repo in module.public-ecr-object : name => {
      name        = repo.repository_name
      uri         = repo.repository_uri
      arn         = repo.arn
      registry_id = repo.registry_id
    }
  } : {}
}

# Specific repository outputs for common use cases
output "frontend_repository_uri" {
  description = "URI of the frontend repository"
  value       = module.public-ecr["frontend"].repository_uri
}

output "api_server_repository_uri" {
  description = "URI of the API server repository"
  value       = module.public-ecr["api-server"].repository_uri
}

output "worker_repository_uri" {
  description = "URI of the worker repository"
  value       = module.public-ecr["worker"].repository_uri
}

# Docker pull commands for easy copy-paste
output "docker_pull_commands" {
  description = "Docker pull commands for all repositories"
  value = {
    for name, repo in module.public-ecr : name =>
    "docker pull ${repo.repository_uri}:latest"
  }
}

# Summary information
output "repository_summary" {
  description = "Summary of created repositories"
  value = {
    total_repositories = length(module.public-ecr)
    repository_names   = keys(module.public-ecr)
    registry_alias     = var.registry_alias
    environment        = var.environment
    project_name       = var.project_name
  }
}
