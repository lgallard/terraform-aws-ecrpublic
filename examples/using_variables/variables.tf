variable "repository_name" {
  description = "Name of the repository"
  type        = string
  default     = "lgallard-public-repo"
}

variable "catalog_data_about_text" {
  description = "A detailed description of the contents of the repository"
  type        = string
  default     = "# Public repo\nPut your description here using Markdown format"
}

variable "catalog_data_architectures" {
  description = "The system architecture that the images in the repository are compatible with"
  type        = list(string)
  default     = ["Linux"]
}

variable "catalog_data_description" {
  description = "A short description of the contents of the repository"
  type        = string
  default     = "Description"
}

variable "catalog_data_logo_image_blob" {
  description = "The base64-encoded repository logo payload"
  type        = string
  default     = null
}

variable "catalog_data_operating_systems" {
  description = "The operating systems that the images in the repository are compatible with"
  type        = list(string)
  default     = ["ARM"]
}

variable "catalog_data_usage_text" {
  description = "Detailed information on how to use the contents of the repository"
  type        = string
  default     = "# Usage\n How to use you image goes here. Use Markdown format"
}