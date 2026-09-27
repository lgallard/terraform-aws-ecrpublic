# Read image metadata from an existing ECR Public repository after images have
# been pushed. Keep this separate from repository creation so empty repositories
# do not make the module creation path depend on image discovery timing.
data "aws_ecrpublic_images" "selected" {
  repository_name = var.repository_name
  registry_id     = var.registry_id

  dynamic "image_ids" {
    for_each = var.image_ids
    iterator = image_id

    content {
      image_tag    = image_id.value.image_tag
      image_digest = image_id.value.image_digest
    }
  }
}
