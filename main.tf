resource "aws_ecrpublic_repository" "repo" {

  repository_name = var.repository_name

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

locals {
  # catalog_data
  catalog_data = [
    {
      about_text        = coalesce(lookup(var.catalog_data, "about_text", null), var.catalog_data_about_text)
      architectures     = coalesce(lookup(var.catalog_data, "architectures", null), var.catalog_data_architectures)
      description       = coalesce(lookup(var.catalog_data, "description", null), var.catalog_data_description)
      logo_image_blob   = coalesce(lookup(var.catalog_data, "logo_image_blob", null), var.catalog_data_logo_image_blob)
      operating_systems = coalesce(lookup(var.catalog_data, "operating_systems", null), var.catalog_data_operating_systems)
      usage_text        = coalesce(lookup(var.catalog_data, "usage_text", null), var.catalog_data_usage_text)
    }
  ]

  # Timeouts
  # If no timeouts block is provided, build one using the default values
  timeouts = (var.timeouts_delete != null || length(var.timeouts) > 0) ? [{
    delete = coalesce(lookup(var.timeouts, "delete", null), var.timeouts_delete)
  }] : []
}
