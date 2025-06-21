module "public-ecr" {

  source = "../.."

  repository_name = var.repository_name

  catalog_data_about_text        = var.catalog_data_about_text
  catalog_data_architectures     = var.catalog_data_architectures
  catalog_data_description       = var.catalog_data_description
  catalog_data_logo_image_blob   = var.catalog_data_logo_image_blob != null ? var.catalog_data_logo_image_blob : filebase64("image.png")
  catalog_data_operating_systems = var.catalog_data_operating_systems
  catalog_data_usage_text        = var.catalog_data_usage_text

}
