module "public-ecr" {

  source = "../.."

  repository_name = "lgallard-public-repo"

  catalog_data = {
    about_text        = "# Public repo\nPut your description here using Markdown format"
    architectures     = ["Linux"]
    description       = "Description"
    logo_image_blob   = filebase64("image.png")
    operating_systems = ["ARM"]
    usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"
  }
}
