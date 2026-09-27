output "image_digests" {
  description = "Image digests returned by the ECR Public images data source"
  value       = [for image in data.aws_ecrpublic_images.selected.images : image.image_digest if image.image_digest != null]
}

output "image_tags" {
  description = "Deduplicated flat list of all tags across every returned image"
  value       = distinct(flatten([for image in data.aws_ecrpublic_images.selected.images : coalesce(image.image_tags, [])]))
}

output "images" {
  description = "Image metadata returned by the ECR Public images data source"
  value = [
    for image in data.aws_ecrpublic_images.selected.images : {
      digest                    = image.image_digest
      tags                      = coalesce(image.image_tags, [])
      size_in_bytes             = image.image_size_in_bytes
      pushed_at                 = image.image_pushed_at
      artifact_media_type       = image.artifact_media_type
      image_manifest_media_type = image.image_manifest_media_type
      registry_id               = image.registry_id
      repository_name           = image.repository_name
    }
    if image.image_digest != null
  ]
}
