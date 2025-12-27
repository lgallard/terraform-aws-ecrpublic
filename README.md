![Terraform](https://lgallardo.com/images/terraform.jpg)
# terraform-aws-ecrpublic
Terraform module to create a Public [AWS ECR](https://aws.amazon.com/ecr) to share images in the [ECR Public Gallery](https://gallery.ecr.aws).

## Usage
You can use this module to create a public ECR registry using objects definition, or using the variables approach:

Check the [examples](examples/) for the **using objects**, **using variables**, and **multiple repositories** snippets.

### Using Objects example
This example creates an public ECR registry:

```
module "public-ecr" {

  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data = {
    about_text        = "# Public repo\nPut your description here using Markdown format"
    architectures     = ["ARM", "x86-64"]
    description       = "Description"
    logo_image_blob   = filebase64("image.png")
    operating_systems = ["Linux"]
    usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"
  }
}
```

### Using variables
This example creates an public ECR registry using variables

```
module "public-ecr" {

  source = "lgallard/ecrpublic/aws"

  repository_name = "lgallard-public-repo"

  catalog_data_about_text        = "# Public repo\nPut your description here using Markdown format"
  catalog_data_architectures     = ["ARM", "x86-64"]
  catalog_data_description       = "Description"
  catalog_data_logo_image_blob   = filebase64("image.png")
  catalog_data_operating_systems = ["Linux"]
  catalog_data_usage_text        = "# Usage\n How to use you image goes here. Use Markdown format"

}

```

### Multiple Repositories

You can create multiple ECR Public repositories using Terraform's native `for_each` capability at the module level:

```hcl
locals {
  repositories = {
    "frontend" = {
      description       = "Frontend application container"
      architectures     = ["x86-64", "ARM 64"]
      operating_systems = ["Linux"]
    }
    "backend" = {
      description       = "Backend API container"
      architectures     = ["x86-64"]
      operating_systems = ["Linux"]
    }
    "worker" = {
      description       = "Background worker container"
      architectures     = ["x86-64"]
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
  catalog_data_about_text = "# ${title(each.key)} Application\n\nContainer image for the ${each.key} component."

  tags = {
    Environment = "production"
    Service     = each.key
    ManagedBy   = "terraform"
  }
}

# Access individual repository outputs
output "frontend_repository_uri" {
  value = module.public-ecr["frontend"].repository_uri
}
```

For comprehensive examples and patterns, see [examples/multiple_repositories](examples/multiple_repositories/).

## Authorization Tokens for Image Push

ECR Public repositories require authentication tokens for programmatic operations like pushing container images. The `aws_ecrpublic_authorization_token` data source provides the necessary credentials for CI/CD integration.

### Important Notes

- **Regional Constraint**: Authorization tokens must be retrieved from `us-east-1` region (same as ECR Public repositories)
- **Token Expiration**: Authorization tokens expire after 12 hours and should be refreshed for long-running processes
- **CI/CD Integration**: Essential for automated image pushing in GitHub Actions, Jenkins, and other CI/CD systems

### Using Authorization Tokens

```hcl
# Authorization token for ECR Public authentication
data "aws_ecrpublic_authorization_token" "token" {}

module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-application"

  catalog_data = {
    description       = "My application container"
    architectures     = ["x86-64"]
    operating_systems = ["Linux"]
  }
}

# Outputs for CI/CD integration
output "repository_uri" {
  value = module.public-ecr.repository_uri
}

output "authorization_token" {
  value     = data.aws_ecrpublic_authorization_token.token.authorization_token
  sensitive = true
}
```

### Docker Authentication Commands

```bash
# Method 1: Using AWS CLI (Recommended)
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

# Method 2: Using Terraform output
echo "$AUTHORIZATION_TOKEN" | docker login --username AWS --password-stdin public.ecr.aws

# Push your image
docker tag my-app:latest <repository_uri>:latest
docker push <repository_uri>:latest
```

### GitHub Actions Integration

```yaml
name: Push to ECR Public
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: us-east-1

      - name: Login to ECR Public
        run: |
          aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

      - name: Build and push
        run: |
          docker build -t ${{ secrets.ECR_REPOSITORY_URI }}:latest .
          docker push ${{ secrets.ECR_REPOSITORY_URI }}:latest
```

For a complete example with authorization token integration, see [examples/with_auth_token](examples/with_auth_token/).

## Testing

This module includes comprehensive automated testing using [Terratest](https://github.com/gruntwork-io/terratest) with specialized ECR Public testing patterns following CLAUDE.md guidelines.

### Prerequisites

- **Terraform** >= 1.0
- **Go** >= 1.21
- **AWS credentials** with ECR Public permissions
- **AWS Region**: Tests must run in `us-east-1` (ECR Public constraint)

### Quick Start

```bash
# Run static tests only (no AWS resources)
make test

# Run basic integration tests (creates AWS resources)
make test-integration

# Run all comprehensive tests
make test-all
```

### Available Test Suites

#### Static Tests (No AWS Resources)
```bash
make test-static          # Format and validation checks
make fmt-check           # Terraform formatting validation
make validate            # Configuration validation
```

#### Integration Tests (Creates AWS Resources)
```bash
make test-integration    # Basic ECR Public repository tests
make test-catalog        # Catalog data validation tests
make test-gallery        # Public gallery compliance tests
make test-timeouts       # Timeout configuration tests
make test-validation     # Variable validation tests
```

#### Comprehensive Testing
```bash
make check              # All static checks
make check-all          # All checks including integration tests
make test-all           # Complete integration test suite
```

### Test Coverage

#### 1. Static Tests (`terraform_basic_test.go`)
- **Format validation**: Terraform file formatting
- **Configuration validation**: Syntax and structure validation
- **Example validation**: Both `using_objects` and `using_variables` examples
- **Fast execution**: Completes in seconds, no AWS resources

#### 2. Integration Tests (`terraform_integration_test.go`)
- **Basic repository creation**: Core ECR Public functionality
- **Variable-based catalog data**: Individual variable approach testing
- **Object-based catalog data**: Object configuration approach testing
- **Timeout configuration**: Custom timeout settings validation
- **Complete configuration**: Comprehensive feature testing

#### 3. Catalog Data Tests (`terraform_catalog_data_test.go`)
- **Content validation**: ECR Public Gallery guidelines compliance
- **Minimal configuration**: Tests minimal valid catalog data
- **Comprehensive configuration**: Tests all catalog data fields
- **Multi-architecture support**: Architecture and OS tagging validation
- **Markdown formatting**: Rich markdown content validation

#### 4. Public Gallery Tests (`terraform_public_gallery_test.go`)
- **Gallery optimization**: Discoverability and searchability features
- **Content guidelines**: ECR Public Gallery content standards
- **Regional constraints**: us-east-1 region requirement validation
- **Public accessibility**: Global public access validation

### ECR Public Specific Testing

This module follows ECR Public-specific testing patterns:

#### Regional Constraints
All tests enforce ECR Public's `us-east-1` regional constraint:
```go
awsRegion := "us-east-1"
EnvVars: map[string]string{
    "AWS_DEFAULT_REGION": awsRegion,
}
```

#### Catalog Data Validation
Tests validate ECR Public Gallery content guidelines:
- **Description length**: Maximum 256 characters
- **Architecture validation**: `x86-64`, `ARM`, `ARM 64`
- **Operating system validation**: `Linux`, `Windows`
- **Markdown formatting**: Proper markdown in about and usage text
- **Public content standards**: Appropriate content for public repositories

#### Repository Naming
Tests use unique repository names to avoid conflicts:
```go
uniqueID := strings.ToLower(random.UniqueId())
repositoryName := fmt.Sprintf("terratest-prefix-%s", uniqueID)
```

### Test Validation Patterns

The test suite includes specialized validation helpers:

#### Repository Validation
- `validateECRPublicRepository()`: Basic repository creation validation
- `validateBasicOutputs()`: Module outputs validation
- `validateCatalogDataOutputs()`: Catalog data specific validation
- `validateAllOutputs()`: Comprehensive output validation

#### Gallery Validation
- `validatePublicGalleryOptimization()`: Gallery discoverability checks
- `validatePublicGallerySearchability()`: Search optimization validation
- `validatePublicGalleryContentGuidelines()`: Content compliance validation
- `validateRegionalConstraints()`: Regional constraint validation

### Example Test Execution

#### Basic Repository Test
```bash
cd test
go test -v -run TestTerraformECRPublicBasic -timeout 30m
```

#### Catalog Data Validation
```bash
cd test
go test -v -run TestECRPublicCatalogDataValidation -timeout 45m
```

#### Complete Test Suite
```bash
cd test
go test -v -timeout 30m -parallel 2
```

### AWS Resources and Costs

⚠️ **Important**: Integration tests create real AWS resources:

- **ECR Public Repositories**: Created and configured with catalog data
- **Region Requirement**: All resources created in `us-east-1` only
- **Cost**: ECR Public repositories are free to create and maintain
- **Cleanup**: Automatic cleanup with `terraform destroy` after each test

### Required AWS Permissions

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr-public:CreateRepository",
                "ecr-public:DeleteRepository",
                "ecr-public:DescribeRepositories",
                "ecr-public:GetRepositoryCatalogData",
                "ecr-public:PutRepositoryCatalogData"
            ],
            "Resource": "*"
        }
    ]
}
```

### Continuous Integration

GitHub Actions workflow provides automated testing:

- **Static Tests**: Run on all pull requests (no AWS resources)
- **Integration Tests**: Run on main branch or manual dispatch
- **Parallel Execution**: Optimized for faster feedback
- **Comprehensive Coverage**: All test suites included

### Troubleshooting

#### Common Issues

**Region Errors**
```
Error: ECR Public repositories can only be created in us-east-1
Solution: Set AWS_DEFAULT_REGION=us-east-1
```

**Permission Errors**
```
Error: Access denied for ECR Public operations
Solution: Verify AWS credentials have ECR Public permissions
```

**Repository Conflicts**
```
Error: Repository already exists
Solution: Tests use unique IDs, manual cleanup may be needed
```

For detailed testing instructions and troubleshooting, see [test/README.md](test/README.md).

## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| aws | n/a |

## Modules

No Modules.

## Resources

| Name |
|------|
| [aws_ecrpublic_repository](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository) |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| catalog\_data | Catalog data configuration for the repository. | `any` | `{}` | no |
| catalog\_data\_about\_text | A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format. | `string` | `null` | no |
| catalog\_data\_architectures | The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`. | `list(string)` | `[]` | no |
| catalog\_data\_description | A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery. | `string` | `null` | no |
| catalog\_data\_logo\_image\_blob | The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute. | `string` | `null` | no |
| catalog\_data\_operating\_systems | The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. 'ARM', 'ARM 64', 'x86', 'x86-64'. | `list(string)` | `null` | no |
| catalog\_data\_usage\_text | Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format. | `string` | `null` | no |
| repository\_name | Name of the repository. | `string` | n/a | yes |
| timeouts | Timeouts map. | `map` | `{}` | no |
| timeouts\_delete | How long to wait for a repository to be deleted. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| arn | Full ARN of the repository |
| id | The repository name. |
| registry\_id | The registry ID where the repository was created. |
| repository\_uri | The URI of the repository. |
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
README.md updated successfully
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

<!-- BEGIN_TF_DOCS -->


## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_ecrpublic_repository.repo](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository) | resource |
| [aws_ecrpublic_repository_policy.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecrpublic_repository_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data"></a> [catalog\_data](#input\_catalog\_data) | Catalog data configuration for the repository. | `any` | `{}` | no |
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format. | `string` | `null` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: 'ARM', 'ARM 64', 'x86', 'x86-64'. | `list(string)` | `[]` | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery. | `string` | `null` | no |
| <a name="input_catalog_data_logo_image_blob"></a> [catalog\_data\_logo\_image\_blob](#input\_catalog\_data\_logo\_image\_blob) | The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute. | `string` | `null` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters: `Linux`, `Windows`. | `list(string)` | `[]` | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format. | `string` | `null` | no |
| <a name="input_create_repository_policy"></a> [create\_repository\_policy](#input\_create\_repository\_policy) | Whether to create a repository policy for controlling push access | `bool` | `false` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository. | `string` | n/a | yes |
| <a name="input_repository_policy"></a> [repository\_policy](#input\_repository\_policy) | The JSON policy document for the repository | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resource. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | Timeouts map. | `map(any)` | `{}` | no |
| <a name="input_timeouts_delete"></a> [timeouts\_delete](#input\_timeouts\_delete) | How long to wait for a repository to be deleted. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | Full ARN of the repository |
| <a name="output_catalog_data"></a> [catalog\_data](#output\_catalog\_data) | The catalog data configuration for the repository |
| <a name="output_gallery_url"></a> [gallery\_url](#output\_gallery\_url) | The URL to the repository in ECR Public Gallery |
| <a name="output_id"></a> [id](#output\_id) | The registry ID where the repository was created. |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created. |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | The repository name. |
| <a name="output_repository_policy"></a> [repository\_policy](#output\_repository\_policy) | The repository policy JSON |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository. |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags |

<!-- END_TF_DOCS -->
