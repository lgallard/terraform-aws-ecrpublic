# Multiple Repositories Example

This example demonstrates how to create multiple ECR Public repositories using the module's native `for_each` support pattern.

## Overview

The terraform-aws-ecrpublic module is designed following Terraform best practices for single-resource management. To create multiple repositories, you can leverage Terraform's native `for_each` meta-argument at the module level.

## Benefits

- **Clean and maintainable**: Uses standard Terraform patterns
- **Flexible configuration**: Each repository can have unique settings
- **Shared configuration**: Common settings can be defined once
- **Backward compatible**: No changes to the existing module interface

## Usage

### Basic Multiple Repositories

```hcl
locals {
  repositories = {
    "frontend" = {
      description = "Frontend application container"
      architectures = ["x86-64", "ARM 64"]
      operating_systems = ["Linux"]
    }
    "backend" = {
      description = "Backend API container"
      architectures = ["x86-64"]
      operating_systems = ["Linux"]
    }
    "worker" = {
      description = "Background worker container"
      architectures = ["x86-64"]
      operating_systems = ["Linux"]
    }
  }
}

module "public-ecr" {
  source   = "lgallard/ecrpublic/aws"
  for_each = local.repositories

  repository_name                = each.key
  catalog_data_description       = each.value.description
  catalog_data_architectures     = each.value.architectures
  catalog_data_operating_systems = each.value.operating_systems

  # Shared configuration
  catalog_data_about_text = "# ${title(each.key)} Application\n\nContainer image for the ${each.key} component of our microservices architecture."
  catalog_data_usage_text = "# Usage\n\n```bash\ndocker pull public.ecr.aws/${var.registry_alias}/${each.key}:latest\n```"

  tags = {
    Environment = "production"
    Service     = each.key
    ManagedBy   = "terraform"
  }
}
```

### Advanced Configuration with Object-Based Catalog Data

```hcl
locals {
  microservices = {
    "api-gateway" = {
      catalog_data = {
        description       = "API Gateway service for microservices"
        about_text        = "# API Gateway\n\nCentral API gateway handling authentication, routing, and rate limiting."
        usage_text        = "# Usage\n\n```bash\ndocker run -p 8080:8080 public.ecr.aws/myregistry/api-gateway:latest\n```"
        architectures     = ["x86-64", "ARM 64"]
        operating_systems = ["Linux"]
      }
      tags = {
        Type = "gateway"
        Tier = "frontend"
      }
    }
    "user-service" = {
      catalog_data = {
        description       = "User management microservice"
        about_text        = "# User Service\n\nHandles user authentication, authorization, and profile management."
        usage_text        = "# Usage\n\nRequires database connection and JWT configuration."
        architectures     = ["x86-64"]
        operating_systems = ["Linux"]
      }
      tags = {
        Type = "service"
        Tier = "backend"
      }
    }
    "notification-service" = {
      catalog_data = {
        description       = "Notification delivery service"
        about_text        = "# Notification Service\n\nAsynchronous notification processing with multiple delivery channels."
        usage_text        = "# Usage\n\nConfigure SMTP, SMS, and push notification providers."
        architectures     = ["x86-64"]
        operating_systems = ["Linux"]
      }
      tags = {
        Type = "service"
        Tier = "backend"
      }
    }
  }
}

module "microservices" {
  source   = "lgallard/ecrpublic/aws"
  for_each = local.microservices

  repository_name = each.key
  catalog_data    = each.value.catalog_data

  # Merge common and service-specific tags
  tags = merge(
    {
      Project     = "microservices-platform"
      Environment = "production"
      ManagedBy   = "terraform"
    },
    each.value.tags
  )
}
```

### Mixed Configuration Approaches

```hcl
locals {
  # Applications with different configuration patterns
  applications = {
    # Object-based configuration
    "web-app" = {
      type = "object"
      config = {
        description       = "Main web application"
        about_text        = "# Web Application\n\nReact-based frontend application."
        architectures     = ["x86-64", "ARM 64"]
        operating_systems = ["Linux"]
      }
    }
    # Variable-based configuration
    "api-server" = {
      type = "variables"
      description = "REST API server"
      about_text  = "# API Server\n\nNode.js REST API with Express framework."
      architectures = ["x86-64"]
      operating_systems = ["Linux"]
    }
  }
}

module "applications" {
  source   = "lgallard/ecrpublic/aws"
  for_each = local.applications

  repository_name = each.key

  # Conditional configuration based on type
  catalog_data = each.value.type == "object" ? each.value.config : {}

  # Variable-based configuration
  catalog_data_description       = each.value.type == "variables" ? each.value.description : null
  catalog_data_about_text        = each.value.type == "variables" ? each.value.about_text : null
  catalog_data_architectures     = each.value.type == "variables" ? each.value.architectures : []
  catalog_data_operating_systems = each.value.type == "variables" ? each.value.operating_systems : []

  tags = {
    Application = each.key
    ManagedBy   = "terraform"
  }
}
```

## Outputs

When using `for_each` with modules, outputs become maps keyed by the `for_each` key:

```hcl
# Output all repository URIs
output "repository_uris" {
  description = "Map of repository names to their URIs"
  value = {
    for k, v in module.public-ecr : k => v.repository_uri
  }
}

# Output specific repository information
output "frontend_repository_uri" {
  description = "URI of the frontend repository"
  value       = module.public-ecr["frontend"].repository_uri
}

# Output all repository information
output "repositories" {
  description = "Complete repository information"
  value = {
    for name, repo in module.public-ecr : name => {
      name         = repo.repository_name
      uri          = repo.repository_uri
      arn          = repo.arn
      registry_id  = repo.registry_id
    }
  }
}
```

## Variable Configuration

```hcl
variable "registry_alias" {
  description = "ECR Public registry alias"
  type        = string
  default     = "myorganization"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "project_name" {
  description = "Project name for tagging"
  type        = string
  default     = "microservices"
}
```

## Best Practices

1. **Use descriptive keys**: Repository names should be clear and meaningful
2. **Leverage locals**: Define repository configurations in locals for better organization
3. **Share common configuration**: Use merge() and conditional logic for shared settings
4. **Consistent naming**: Follow a consistent naming convention for repositories
5. **Proper tagging**: Include environment, project, and service tags for resource management
6. **Catalog data quality**: Ensure each repository has appropriate ECR Public Gallery content

## Testing

Test the configuration with:

```bash
terraform init
terraform plan
terraform apply
```

Verify repositories in the ECR Public Gallery:

```bash
aws ecr-public describe-repositories --region us-east-1
```

## Regional Considerations

Remember that ECR Public repositories must be created in `us-east-1`:

```hcl
provider "aws" {
  alias  = "ecr_public"
  region = "us-east-1"
}

module "public-ecr" {
  providers = {
    aws = aws.ecr_public
  }
  # ... rest of configuration
}
```