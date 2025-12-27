output "arn" {
  description = "Full ARN of the repository"
  value       = aws_ecrpublic_repository.repo.arn
}

output "repository_arn" {
  description = "Full ARN of the repository"
  value       = aws_ecrpublic_repository.repo.arn
}

output "id" {
  description = "The registry ID where the repository was created."
  value       = aws_ecrpublic_repository.repo.registry_id
}

output "repository_name" {
  description = "The repository name."
  value       = aws_ecrpublic_repository.repo.repository_name
}

output "registry_id" {
  description = "The registry ID where the repository was created."
  value       = aws_ecrpublic_repository.repo.registry_id
}

output "repository_uri" {
  description = "The URI of the repository."
  value       = aws_ecrpublic_repository.repo.repository_uri
}

output "repository_url" {
  description = "The URL of the repository"
  value       = aws_ecrpublic_repository.repo.repository_uri
}

output "repository_policy" {
  description = "The repository policy JSON"
  value       = try(aws_ecrpublic_repository_policy.this[0].policy, null)
}

output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags"
  value       = aws_ecrpublic_repository.repo.tags_all
}

output "gallery_url" {
  description = "The URL to the repository in ECR Public Gallery"
  value       = format("https://gallery.ecr.aws/%s/%s", 
                      urlencode(aws_ecrpublic_repository.repo.registry_id),
                      urlencode(aws_ecrpublic_repository.repo.repository_name))
}

output "catalog_data" {
  description = "The catalog data configuration for the repository"
  value = length(aws_ecrpublic_repository.repo.catalog_data) > 0 ? {
    about_text        = try(aws_ecrpublic_repository.repo.catalog_data[0].about_text, null)
    architectures     = try(aws_ecrpublic_repository.repo.catalog_data[0].architectures, null)
    description       = try(aws_ecrpublic_repository.repo.catalog_data[0].description, null)
    logo_image_blob   = try(aws_ecrpublic_repository.repo.catalog_data[0].logo_image_blob, null)
    operating_systems = try(aws_ecrpublic_repository.repo.catalog_data[0].operating_systems, null)
    usage_text        = try(aws_ecrpublic_repository.repo.catalog_data[0].usage_text, null)
  } : null
}
