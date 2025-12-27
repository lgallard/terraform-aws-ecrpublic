# Multiple repositories example using module-level for_each

# Define multiple repositories with shared and unique configurations
locals {
  # Repository configurations
  repositories = {
    "frontend" = {
      description       = "Frontend application container image"
      architectures     = ["x86-64", "ARM 64"]
      operating_systems = ["Linux"]
      about_text        = "# Frontend Application\n\nReact-based frontend application for the public web interface."
      usage_text        = "# Usage\n\n```bash\ndocker pull public.ecr.aws/${var.registry_alias}/frontend:latest\ndocker run -p 3000:3000 public.ecr.aws/${var.registry_alias}/frontend:latest\n```"
      tags = {
        Tier = "frontend"
        Type = "web-application"
      }
    }
    "api-server" = {
      description       = "REST API server container image"
      architectures     = ["x86-64"]
      operating_systems = ["Linux"]
      about_text        = "# API Server\n\nNode.js REST API server with Express framework providing backend services."
      usage_text        = "# Usage\n\n```bash\ndocker pull public.ecr.aws/${var.registry_alias}/api-server:latest\ndocker run -p 8080:8080 -e DATABASE_URL=... public.ecr.aws/${var.registry_alias}/api-server:latest\n```"
      tags = {
        Tier = "backend"
        Type = "api-service"
      }
    }
    "worker" = {
      description       = "Background worker container image"
      architectures     = ["x86-64"]
      operating_systems = ["Linux"]
      about_text        = "# Background Worker\n\nAsynchronous task processor for handling background jobs and data processing."
      usage_text        = "# Usage\n\n```bash\ndocker pull public.ecr.aws/${var.registry_alias}/worker:latest\ndocker run -e REDIS_URL=... public.ecr.aws/${var.registry_alias}/worker:latest\n```"
      tags = {
        Tier = "backend"
        Type = "worker-service"
      }
    }
  }
}

# Create multiple ECR Public repositories using for_each
module "public-ecr" {
  source   = "../.."
  for_each = local.repositories

  repository_name = each.key

  # Catalog data configuration
  catalog_data_description       = each.value.description
  catalog_data_about_text        = each.value.about_text
  catalog_data_usage_text        = each.value.usage_text
  catalog_data_architectures     = each.value.architectures
  catalog_data_operating_systems = each.value.operating_systems

  # Merge common and repository-specific tags
  tags = merge(
    {
      Environment = var.environment
      Project     = var.project_name
      ManagedBy   = "terraform"
      CreatedBy   = "terraform-aws-ecrpublic-module"
    },
    each.value.tags
  )
}

# Alternative example using object-based catalog data configuration
module "public-ecr-object" {
  source   = "../.."
  for_each = var.enable_object_example ? local.object_repositories : {}

  repository_name = each.key
  catalog_data    = each.value.catalog_data

  tags = merge(
    {
      Environment   = var.environment
      Project       = var.project_name
      ManagedBy     = "terraform"
      ConfigMethod  = "object-based"
    },
    each.value.tags
  )
}

# Object-based configuration example
locals {
  object_repositories = {
    "database" = {
      catalog_data = {
        description       = "Database container with PostgreSQL"
        about_text        = "# Database Service\n\nPostgreSQL database container optimized for production workloads."
        usage_text        = "# Usage\n\nContainer includes automated backups and health monitoring."
        architectures     = ["x86-64"]
        operating_systems = ["Linux"]
      }
      tags = {
        Tier = "data"
        Type = "database"
      }
    }
    "cache" = {
      catalog_data = {
        description       = "Redis cache container"
        about_text        = "# Cache Service\n\nRedis cache container for session management and data caching."
        usage_text        = "# Usage\n\nConfigured with persistence and clustering support."
        architectures     = ["x86-64"]
        operating_systems = ["Linux"]
      }
      tags = {
        Tier = "data"
        Type = "cache"
      }
    }
  }
}