module "public-ecr" {
  source = "../.."

  repository_name = var.repository_name

  # Repository policy configuration
  create_repository_policy = var.create_repository_policy
  repository_policy        = var.repository_policy

  # Catalog data
  catalog_data = var.catalog_data

  # Tags
  tags = var.tags
}
