# Issue: Add enhanced module outputs

**Title:** `feat: Add enhanced outputs (tags_all, gallery_url, catalog_data)`

**Labels:** `enhancement`, `priority: low`

---

## Summary
Add additional outputs to provide more useful information about the created ECR Public repository.

## Background
The current module provides basic outputs. Additional outputs would improve usability:
- `tags_all`: Shows merged tags including provider default_tags
- `gallery_url`: Direct link to the repository in ECR Public Gallery
- `catalog_data`: Composite output of all catalog data fields

## Proposed Implementation

### New Outputs

Add to `outputs.tf`:

```hcl
output "tags_all" {
  description = "A map of tags assigned to the resource, including those inherited from the provider default_tags"
  value       = aws_ecrpublic_repository.repo.tags_all
}

output "gallery_url" {
  description = "The URL to the repository in ECR Public Gallery"
  value       = "https://gallery.ecr.aws/${aws_ecrpublic_repository.repo.registry_id}/${aws_ecrpublic_repository.repo.repository_name}"
}

output "catalog_data" {
  description = "The catalog data configuration for the repository"
  value = length(aws_ecrpublic_repository.repo.catalog_data) > 0 ? {
    about_text        = try(aws_ecrpublic_repository.repo.catalog_data[0].about_text, null)
    architectures     = try(aws_ecrpublic_repository.repo.catalog_data[0].architectures, null)
    description       = try(aws_ecrpublic_repository.repo.catalog_data[0].description, null)
    operating_systems = try(aws_ecrpublic_repository.repo.catalog_data[0].operating_systems, null)
    usage_text        = try(aws_ecrpublic_repository.repo.catalog_data[0].usage_text, null)
  } : null
}
```

## Priority
**Low** - Nice to have but not blocking any functionality

## Benefits
- `tags_all`: Useful for debugging and compliance when using provider default_tags
- `gallery_url`: Convenience for users to quickly access their repository in the Gallery
- `catalog_data`: Allows downstream modules/resources to reference catalog configuration

## Tasks
- [ ] Add `tags_all` output
- [ ] Add `gallery_url` output
- [ ] Add `catalog_data` composite output
- [ ] Update README documentation
- [ ] Update example outputs to demonstrate new outputs
