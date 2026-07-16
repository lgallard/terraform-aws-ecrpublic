variable "repository_name" {
  description = "Name of the repository"
  type        = string
  default     = "my-app"
}

variable "run_docker_login" {
  description = "Whether to run the local Docker login example during apply. Disabled by default to keep the example safe for validation."
  type        = bool
  default     = false
}

variable "catalog_data_description" {
  description = "Public description visible in ECR Public Gallery"
  type        = string
  default     = "My application container"
}

variable "catalog_data_about_text" {
  description = "Public about text in markdown format"
  type        = string
  default     = <<-EOT
    # My Application

    ## Description
    This container provides a sample application for demonstrating ephemeral ECR Public authentication.

    ## Features
    - Lightweight container image
    - Production-ready configuration
    - Auth token not persisted in Terraform state
  EOT
}

variable "catalog_data_usage_text" {
  description = "Public usage instructions in markdown format"
  type        = string
  default     = <<-EOT
    # Usage

    ## Authentication
    ```bash
    # Preferred: Get a short-lived password through AWS CLI
    aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
    ```

    ## Pull Image
    ```bash
    docker pull public.ecr.aws/your-registry/my-app:latest
    ```
  EOT
}

variable "catalog_data_architectures" {
  description = "Supported architectures for container images"
  type        = list(string)
  default     = ["x86-64"]
}

variable "catalog_data_operating_systems" {
  description = "Supported operating systems for container images"
  type        = list(string)
  default     = ["Linux"]
}
