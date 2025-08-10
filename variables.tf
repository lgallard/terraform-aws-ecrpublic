# General vars
variable "repository_name" {
  description = "Name of the repository."
  type        = string
}

# catalog_data
variable "catalog_data" {
  description = "Catalog data configuration for the repository."
  type        = any
  default     = {}
}

variable "catalog_data_about_text" {
  description = "A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format."
  type        = string
  default     = null
}

variable "catalog_data_architectures" {
  description = "The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: 'ARM', 'ARM 64', 'x86', 'x86-64'."
  type        = list(string)
  default     = []

  validation {
    condition = alltrue([
      for arch in var.catalog_data_architectures :
      contains(["ARM", "ARM 64", "x86", "x86-64"], arch)
    ])
    error_message = "Architectures must be one of: ARM, ARM 64, x86, x86-64."
  }
}

variable "catalog_data_description" {
  description = "A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery."
  type        = string
  default     = null
}

variable "catalog_data_logo_image_blob" {
  description = "The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute."
  type        = string
  default     = null
}

variable "catalog_data_operating_systems" {
  description = "The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`."
  type        = list(string)
  default     = null

  validation {
    condition = var.catalog_data_operating_systems == null ? true : alltrue([
      for os in var.catalog_data_operating_systems :
      contains(["Linux", "Windows"], os)
    ])
    error_message = "Operating systems must be one of: Linux, Windows."
  }
}

variable "catalog_data_usage_text" {
  description = "Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format."
  type        = string
  default     = null
}

# Timeouts
variable "timeouts" {
  description = "Timeouts map."
  type        = map(any)
  default     = {}
}

variable "timeouts_delete" {
  description = "How long to wait for a repository to be deleted."
  type        = string
  default     = null
}

# Tags
variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
  default     = {}
}
