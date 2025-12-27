# Issue: Consider multiple repository support

**Title:** `feat: Evaluate multiple repository creation support`

**Labels:** `enhancement`, `priority: medium`, `discussion`

---

## Summary
Evaluate and potentially implement support for creating multiple ECR Public repositories with shared configurations in a single module call.

## Background
Currently, the module creates a single repository per module call. Some users may want to create multiple related repositories with shared configurations (e.g., a microservices application with multiple container images).

## Options to Consider

### Option 1: Native for_each Support (Recommended for evaluation)

Add a `repositories` variable that accepts a map:

```hcl
variable "repositories" {
  description = "Map of repository configurations to create"
  type = map(object({
    catalog_data = optional(object({
      about_text        = optional(string)
      architectures     = optional(list(string))
      description       = optional(string)
      logo_image_blob   = optional(string)
      operating_systems = optional(list(string))
      usage_text        = optional(string)
    }))
    tags = optional(map(string))
  }))
  default = {}
}

resource "aws_ecrpublic_repository" "repos" {
  for_each        = var.repositories
  repository_name = each.key

  dynamic "catalog_data" {
    for_each = each.value.catalog_data != null ? [each.value.catalog_data] : []
    content {
      about_text        = catalog_data.value.about_text
      architectures     = catalog_data.value.architectures
      description       = catalog_data.value.description
      logo_image_blob   = catalog_data.value.logo_image_blob
      operating_systems = catalog_data.value.operating_systems
      usage_text        = catalog_data.value.usage_text
    }
  }

  tags = merge(var.tags, try(each.value.tags, {}))
}
```

### Option 2: Document Module-Level for_each (No code changes)

Keep the current single-repository module and document how to use it with `for_each`:

```hcl
locals {
  repositories = {
    "app-frontend" = {
      description = "Frontend application"
    }
    "app-backend" = {
      description = "Backend API"
    }
    "app-worker" = {
      description = "Background worker"
    }
  }
}

module "public-ecr" {
  source   = "lgallard/ecrpublic/aws"
  for_each = local.repositories

  repository_name          = each.key
  catalog_data_description = each.value.description
  catalog_data_architectures     = ["x86-64"]
  catalog_data_operating_systems = ["Linux"]
}
```

### Option 3: Separate Submodule

Create a `modules/multiple/` submodule specifically for multiple repository creation.

## Recommendation

**Option 2** (documentation) is the simplest and maintains backward compatibility. It leverages Terraform's native module `for_each` capability without adding complexity to the module.

**Option 1** could be considered if there's significant community demand or specific use cases that Option 2 doesn't address well.

## Priority
**Medium** - Useful but workarounds exist

## Discussion Points
- Is there community demand for this feature?
- Would Option 2 (documentation) be sufficient?
- Are there specific use cases that require native support?

## Tasks
- [ ] Gather community feedback on need
- [ ] Document module-level for_each pattern in README
- [ ] Decide on implementation approach
- [ ] Implement if native support is chosen
