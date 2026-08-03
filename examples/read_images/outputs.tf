output "image_digests" {
  description = "Image digests returned by the ECR Public images data source"
  value       = [for image in data.aws_ecrpublic_images.selected.images : image.digest if image.digest != null]
}

output "image_tags" {
  description = "Distinct image tags returned by the ECR Public images data source"
  value       = distinct(flatten([for image in data.aws_ecrpublic_images.selected.images : image.tags]))
}

output "images" {
  description = "Image metadata returned by the ECR Public images data source"
  value = [
    for image in data.aws_ecrpublic_images.selected.images : {
      digest                    = image.digest
      tags                      = image.tags
      size_in_bytes             = image.size_in_bytes
      pushed_at                 = image.pushed_at
      artifact_media_type       = image.artifact_media_type
      image_manifest_media_type = image.image_manifest_media_type
      registry_id               = image.registry_id
      repository_name           = image.repository_name
    }
  ]
}
