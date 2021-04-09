resource "aws_ecrpublic_repository" "repo" {

  repository_name = var.repository_name

  # Image scanning configuration
  dynamic "catalog_data" {
    for_each = local.catalog_data
    content {
      about_text        = lookup(catalog_data.value, "about_text")
      architectures     = lookup(catalog_data.value, "architectures")
      description       = lookup(catalog_data.value, "description")
      logo_image_blob   = lookup(catalog_data.value, "logo_image_blob")
      operating_systems = lookup(catalog_data.value, "operating_systems")
      usage_text        = lookup(catalog_data.value, "usage_text")
    }
  }

  # Timeouts
  dynamic "timeouts" {
    for_each = local.timeouts
    content {
      delete = lookup(timeouts.value, "delete")
    }
  }
}

locals {
  # catalog_data
  catalog_data = [
    {
      about_text        = lookup(var.catalog_data, "about_text", null) == null ? var.catalog_data_about_text : lookup(var.catalog_data, "about_text", null)
      architectures     = lookup(var.catalog_data, "architectures", []) == null ? var.catalog_data_architectures : lookup(var.catalog_data, "architectures", [])
      description       = lookup(var.catalog_data, "description", null) == null ? var.catalog_data_description : lookup(var.catalog_data, "description", null)
      logo_image_blob   = lookup(var.catalog_data, "logo_image_blob", null) == null ? var.catalog_data_logo_image_blob : lookup(var.catalog_data, "logo_image_blob", null)
      operating_systems = lookup(var.catalog_data, "operating_systems", []) == null ? var.catalog_data_operating_systems : lookup(var.catalog_data, "operating_systems", [])
      usage_text        = lookup(var.catalog_data, "usage_text", null) == null ? var.catalog_data_usage_text : lookup(var.catalog_data, "usage_text", null)
    }
  ]

  # Timeouts
  # If no timeouts block is provided, build one using the default values
  timeouts = var.timeouts_delete == null && length(var.timeouts) == 0 ? [] : [
    {
      delete = lookup(var.timeouts, "delete", null) == null ? var.timeouts_delete : lookup(var.timeouts, "delete")
    }
  ]
}
