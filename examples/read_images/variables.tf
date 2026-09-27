variable "repository_name" {
  description = "Name of the existing ECR Public repository to inspect"
  type        = string
  default     = "my-application"
}

variable "registry_id" {
  description = "AWS account ID associated with the public registry that contains the repository"
  type        = string
  default     = null
}

variable "image_ids" {
  description = "Optional image filters. Each object can specify an image_tag, an image_digest, or both. Leave empty to list repository images."
  type = list(object({
    image_tag    = optional(string)
    image_digest = optional(string)
  }))
  default = []

  validation {
    condition = alltrue([
      for image_id in var.image_ids : image_id.image_tag != null || image_id.image_digest != null
    ])
    error_message = "Each image_ids entry must specify at least one of image_tag or image_digest."
  }
}
