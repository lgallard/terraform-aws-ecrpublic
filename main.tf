resource "aws_ecrpublic_repository" "repo" {

  repository_name = var.repository_name

  # Security lifecycle checks
  lifecycle {
    precondition {
      condition     = local.secure_content_check
      error_message = "All catalog data fields must pass security validation checks to prevent malicious content."
    }
  }

  # Catalog data configuration
  dynamic "catalog_data" {
    for_each = local.catalog_data
    content {
      about_text        = lookup(catalog_data.value, "about_text", null)
      architectures     = lookup(catalog_data.value, "architectures", null)
      description       = lookup(catalog_data.value, "description", null)
      logo_image_blob   = lookup(catalog_data.value, "logo_image_blob", null)
      operating_systems = lookup(catalog_data.value, "operating_systems", null)
      usage_text        = lookup(catalog_data.value, "usage_text", null)
    }
  }

  # Timeouts
  dynamic "timeouts" {
    for_each = local.timeouts
    content {
      delete = lookup(timeouts.value, "delete", null)
    }
  }

  # Tags for resource management and automated cleanup
  tags = var.tags
}

# Repository policy for controlling push access
resource "aws_ecrpublic_repository_policy" "this" {
  count           = var.create_repository_policy ? 1 : 0
  repository_name = aws_ecrpublic_repository.repo.repository_name
  policy          = var.repository_policy
}

locals {
  # Helper values to determine if catalog_data should be created
  _catalog_data_about_text        = lookup(var.catalog_data, "about_text", var.catalog_data_about_text)
  _catalog_data_architectures     = lookup(var.catalog_data, "architectures", length(var.catalog_data_architectures) > 0 ? var.catalog_data_architectures : null)
  _catalog_data_description       = lookup(var.catalog_data, "description", var.catalog_data_description)
  _catalog_data_logo_image_blob   = lookup(var.catalog_data, "logo_image_blob", var.catalog_data_logo_image_blob)
  _catalog_data_operating_systems = lookup(var.catalog_data, "operating_systems", var.catalog_data_operating_systems)
  _catalog_data_usage_text        = lookup(var.catalog_data, "usage_text", var.catalog_data_usage_text)

  # Security validation for all text-based catalog data fields
  secure_content_check = alltrue([
    for field in [local._catalog_data_description, local._catalog_data_about_text, local._catalog_data_usage_text] :
    field == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", field))
  ])

  # Only create catalog_data block if at least one field has a value
  _has_catalog_data = (
    local._catalog_data_about_text != null ||
    local._catalog_data_architectures != null ||
    local._catalog_data_description != null ||
    local._catalog_data_logo_image_blob != null ||
    local._catalog_data_operating_systems != null ||
    local._catalog_data_usage_text != null
  )

  # catalog_data
  catalog_data = local._has_catalog_data ? [
    {
      about_text        = local._catalog_data_about_text
      architectures     = local._catalog_data_architectures
      description       = local._catalog_data_description
      logo_image_blob   = local._catalog_data_logo_image_blob
      operating_systems = local._catalog_data_operating_systems
      usage_text        = local._catalog_data_usage_text
    }
  ] : []

  # Timeouts
  # If no timeouts block is provided, build one using the default values
  timeouts = (var.timeouts_delete != null || length(var.timeouts) > 0) ? [{
    delete = coalesce(lookup(var.timeouts, "delete", null), var.timeouts_delete)
  }] : []
}
