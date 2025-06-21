output "arn" {
  description = "Full ARN of the repository"
  value       = aws_ecrpublic_repository.repo.arn
}

output "repository_arn" {
  description = "Full ARN of the repository"
  value       = aws_ecrpublic_repository.repo.arn
}

output "id" {
  description = "The repository name."
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
