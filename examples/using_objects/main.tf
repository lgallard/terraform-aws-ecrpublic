module "public-ecr" {

  source = "../.."

  repository_name = "lgallard-public-repo"

  catalog_data = {
    about_text        = "# Public repo\nPut your description here using Markdown format"
    architectures     = ["x86-64"]
    description       = "Description"
    logo_image_blob   = can(fileexists("${path.module}/image.png")) ? filebase64("${path.module}/image.png") : null
    operating_systems = ["Linux"]
    usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"
  }
}
