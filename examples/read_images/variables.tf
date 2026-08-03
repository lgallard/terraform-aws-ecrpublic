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
}
