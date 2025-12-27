# Variables for multiple repositories example

variable "registry_alias" {
  description = "ECR Public registry alias for repository URIs"
  type        = string
  default     = "myorganization"

  validation {
    condition     = can(regex("^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$", var.registry_alias))
    error_message = "Registry alias must contain only lowercase letters, numbers, and hyphens."
  }
}

variable "environment" {
  description = "Environment name (e.g., dev, staging, production)"
  type        = string
  default     = "production"

  validation {
    condition     = contains(["dev", "staging", "production"], var.environment)
    error_message = "Environment must be one of: dev, staging, production."
  }
}

variable "project_name" {
  description = "Project name for resource tagging and organization"
  type        = string
  default     = "microservices-platform"

  validation {
    condition     = length(var.project_name) >= 2 && length(var.project_name) <= 50
    error_message = "Project name must be between 2 and 50 characters."
  }
}

variable "enable_object_example" {
  description = "Whether to enable the object-based configuration example"
  type        = bool
  default     = false
}

variable "common_architectures" {
  description = "Default architectures for all repositories"
  type        = list(string)
  default     = ["x86-64"]

  validation {
    condition = alltrue([
      for arch in var.common_architectures :
      contains(["ARM", "ARM 64", "x86", "x86-64"], arch)
    ])
    error_message = "Architectures must be one of: ARM, ARM 64, x86, x86-64."
  }
}

variable "common_operating_systems" {
  description = "Default operating systems for all repositories"
  type        = list(string)
  default     = ["Linux"]

  validation {
    condition = alltrue([
      for os in var.common_operating_systems :
      contains(["Linux", "Windows"], os)
    ])
    error_message = "Operating systems must be one of: Linux, Windows."
  }
}

variable "additional_tags" {
  description = "Additional tags to apply to all repositories"
  type        = map(string)
  default     = {}

  validation {
    condition     = can(keys(var.additional_tags))
    error_message = "Additional tags must be a valid map of strings."
  }
}