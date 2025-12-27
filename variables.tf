# General vars
variable "repository_name" {
  description = "Name of the repository."
  type        = string

  validation {
    condition     = can(regex("^[a-z0-9](?:[a-z0-9._-]*[a-z0-9])?$", var.repository_name))
    error_message = "Repository name must start and end with an alphanumeric character and can only contain lowercase letters, numbers, hyphens, underscores, and periods."
  }

  validation {
    condition     = length(var.repository_name) >= 2 && length(var.repository_name) <= 256
    error_message = "Repository name must be between 2 and 256 characters long."
  }

  validation {
    condition     = !can(regex("(?i)^(admin|aws|amazon|ecr|public|private|root|system|test|null|undefined)$", var.repository_name))
    error_message = "Repository name must not use reserved or potentially confusing names for security and clarity."
  }
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

  validation {
    condition     = var.catalog_data_about_text == null || length(var.catalog_data_about_text) > 0
    error_message = "About text cannot be empty when provided."
  }

  validation {
    condition     = var.catalog_data_about_text == null || length(var.catalog_data_about_text) <= 16384
    error_message = "About text must be 16384 characters or less for ECR Public Gallery."
  }

  validation {
    condition = var.catalog_data_about_text == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", var.catalog_data_about_text))
    error_message = "About text must not contain potentially malicious scripts or executable content for security."
  }
}

variable "catalog_data_architectures" {
  description = "The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: 'ARM', 'ARM 64', 'x86', 'x86-64'."
  type        = list(string)
  default     = []

  validation {
    condition = length(var.catalog_data_architectures) == 0 || alltrue([
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

  validation {
    condition     = var.catalog_data_description == null || length(var.catalog_data_description) <= 256
    error_message = "Description must be 256 characters or less for ECR Public Gallery visibility."
  }

  validation {
    condition = var.catalog_data_description == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", var.catalog_data_description))
    error_message = "Description must not contain potentially malicious scripts or executable content for security."
  }
}

variable "catalog_data_logo_image_blob" {
  description = "The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute."
  type        = string
  default     = null

  validation {
    condition     = var.catalog_data_logo_image_blob == null || can(base64decode(var.catalog_data_logo_image_blob))
    error_message = "Logo image blob must be valid base64-encoded data."
  }

  validation {
    condition     = var.catalog_data_logo_image_blob == null || length(var.catalog_data_logo_image_blob) <= 2097152
    error_message = "Logo image must be under 2MB when base64-encoded to prevent resource exhaustion."
  }
}

variable "catalog_data_operating_systems" {
  description = "The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`."
  type        = list(string)
  default     = []

  validation {
    condition = length(var.catalog_data_operating_systems) == 0 || alltrue([
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

  validation {
    condition     = var.catalog_data_usage_text == null || length(var.catalog_data_usage_text) > 0
    error_message = "Usage text cannot be empty when provided."
  }

  validation {
    condition     = var.catalog_data_usage_text == null || length(var.catalog_data_usage_text) <= 10240
    error_message = "Usage text must be 10240 characters or less for ECR Public Gallery."
  }

  validation {
    condition = var.catalog_data_usage_text == null || !can(regex("(?i)(<script\\b|javascript:|vbscript:|data:[^,]*script|\\bon\\w+\\s*=|&#x?[0-9a-f]*;)", var.catalog_data_usage_text))
    error_message = "Usage text must not contain potentially malicious scripts or executable content for security."
  }
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

  validation {
    condition     = var.timeouts_delete == null || can(regex("^[0-9]+(ns|us|µs|ms|s|m|h)$", var.timeouts_delete))
    error_message = "Timeout must be a valid duration string (e.g., '30s', '5m', '1h')."
  }
}

# Repository policy
variable "create_repository_policy" {
  description = "Whether to create a repository policy for controlling push access"
  type        = bool
  default     = false
}

variable "repository_policy" {
  description = "The JSON policy document for the repository"
  type        = string
  default     = null

  validation {
    condition     = var.repository_policy == null || can(jsondecode(var.repository_policy))
    error_message = "Repository policy must be valid JSON when provided."
  }

  validation {
    condition     = var.repository_policy == null || (
      can(jsondecode(var.repository_policy)) &&
      can(lookup(jsondecode(var.repository_policy), "Version", null)) &&
      can(lookup(jsondecode(var.repository_policy), "Statement", null))
    )
    error_message = "Repository policy must be a valid IAM policy document with Version and Statement fields."
  }

  validation {
    condition     = var.repository_policy == null || length(var.repository_policy) <= 10240
    error_message = "Repository policy must be 10240 characters or less."
  }
}

# Tags
variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
  default     = {}
}
