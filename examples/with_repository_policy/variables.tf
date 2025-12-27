variable "repository_name" {
  description = "Name of the ECR Public repository"
  type        = string
  default     = "my-public-app"
}

variable "create_repository_policy" {
  description = "Whether to create a repository policy for controlling push access"
  type        = bool
  default     = true
}

variable "repository_policy" {
  description = "The JSON policy document for the repository. This example shows CI/CD push access."
  type        = string
  default     = null
}

variable "catalog_data" {
  description = "Catalog data configuration for the repository"
  type = object({
    description       = optional(string)
    about_text        = optional(string)
    usage_text        = optional(string)
    architectures     = optional(list(string))
    operating_systems = optional(list(string))
    logo_image_blob   = optional(string)
  })
  default = {
    description       = "Public container image with CI/CD push access"
    about_text        = "# My Public App\n\nThis image is automatically built and published via CI/CD."
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }
}

variable "tags" {
  description = "A map of tags to assign to the repository"
  type        = map(string)
  default = {
    Environment = "production"
    ManagedBy   = "terraform"
  }
}
