# ECR Public with Authorization Token Example

This example demonstrates how to create an ECR Public repository alongside retrieving authorization tokens for programmatic image pushing.

## Overview

The `aws_ecrpublic_authorization_token` data source provides authentication tokens needed to push images to ECR Public repositories. This example shows how to use this data source with the ECR Public module for complete CI/CD integration.

## Key Features

- ECR Public repository creation with catalog data
- Authorization token retrieval for Docker authentication
- Pre-configured outputs for Docker login commands
- Regional constraint handling (us-east-1 requirement)

## Important Notes

### Regional Constraints

- **ECR Public repositories must be created in `us-east-1` region**
- **Authorization tokens must also be retrieved from `us-east-1`**
- This is an AWS service limitation for ECR Public Gallery

### Token Expiration

- Authorization tokens expire after 12 hours
- Tokens should be refreshed for long-running CI/CD processes
- Use the `token_expires_at` output to monitor expiration

## Usage

### Basic Example

```hcl
module "public-ecr-with-auth" {
  source = "lgallard/ecrpublic/aws//examples/with_auth_token"

  repository_name = "my-application"
  
  catalog_data_description = "My public application container"
  catalog_data_about_text = <<-EOT
    # My Application
    
    Production-ready container for my public application.
  EOT
}
```

### CI/CD Integration

```bash
# Retrieve outputs
REPO_URI=$(terraform output -raw repository_uri)
AUTH_TOKEN=$(terraform output -raw authorization_token)

# Docker login using authorization token
echo "$AUTH_TOKEN" | docker login --username AWS --password-stdin public.ecr.aws

# Or use AWS CLI (recommended)
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

# Tag and push image
docker tag my-app:latest $REPO_URI:latest
docker push $REPO_URI:latest
```

### GitHub Actions Example

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

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| repository_name | Name of the repository | `string` | `"my-app"` | no |
| catalog_data_description | Public description visible in ECR Public Gallery | `string` | `"My application container"` | no |
| catalog_data_about_text | Public about text in markdown format | `string` | See variables.tf | no |
| catalog_data_usage_text | Public usage instructions in markdown format | `string` | See variables.tf | no |
| catalog_data_architectures | Supported architectures for container images | `list(string)` | `["x86-64"]` | no |
| catalog_data_operating_systems | Supported operating systems for container images | `list(string)` | `["Linux"]` | no |

## Outputs

| Name | Description | Sensitive |
|------|-------------|:---------:|
| repository_name | Name of the repository | no |
| repository_uri | The URI of the repository | no |
| repository_url | The URL of the repository | no |
| repository_arn | Full ARN of the repository | no |
| registry_id | The registry ID where the repository was created | no |
| authorization_token | The authorization token (base64 encoded) | yes |
| token_expires_at | Token expiration timestamp | no |
| docker_login_command | Docker login command for ECR Public | no |
| aws_cli_login_command | AWS CLI command to login to ECR Public | no |

## Security Considerations

1. **Token Sensitivity**: Authorization tokens are marked as sensitive and should not be logged
2. **Token Expiration**: Implement token refresh logic for long-running processes
3. **Regional Compliance**: Ensure all ECR Public operations use us-east-1 region
4. **Access Control**: Limit ECR Public permissions to necessary operations only

## Common Issues

### "Repository does not exist" Error
Ensure the repository is created before attempting authentication operations.

### "No authorization token" Error
Verify that:
- AWS credentials are properly configured
- The AWS provider is configured for us-east-1 region
- IAM permissions include `ecr-public:GetAuthorizationToken`

### Token Expiration
Authorization tokens expire after 12 hours. Implement refresh logic:

```bash
# Check token expiration
TOKEN_EXPIRES=$(terraform output -raw token_expires_at)
CURRENT_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

if [[ "$CURRENT_TIME" > "$TOKEN_EXPIRES" ]]; then
  echo "Token expired, refreshing..."
  terraform refresh
fi
```

## References

- [aws_ecrpublic_authorization_token](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_authorization_token)
- [ECR Public User Guide](https://docs.aws.amazon.com/AmazonECR/latest/public/)
- [Docker CLI Reference](https://docs.docker.com/engine/reference/commandline/login/)

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 5.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_public-ecr"></a> [public-ecr](#module\_public-ecr) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_ecrpublic_authorization_token.token](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecrpublic_authorization_token) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_catalog_data_about_text"></a> [catalog\_data\_about\_text](#input\_catalog\_data\_about\_text) | Public about text in markdown format | `string` | `"# My Application\n\n## Description\nThis container provides a sample application for demonstrating ECR Public authentication.\n\n## Features\n- Lightweight container image\n- Production-ready configuration\n- Easy deployment\n"` | no |
| <a name="input_catalog_data_architectures"></a> [catalog\_data\_architectures](#input\_catalog\_data\_architectures) | Supported architectures for container images | `list(string)` | <pre>[<br>  "x86-64"<br>]</pre> | no |
| <a name="input_catalog_data_description"></a> [catalog\_data\_description](#input\_catalog\_data\_description) | Public description visible in ECR Public Gallery | `string` | `"My application container"` | no |
| <a name="input_catalog_data_operating_systems"></a> [catalog\_data\_operating\_systems](#input\_catalog\_data\_operating\_systems) | Supported operating systems for container images | `list(string)` | <pre>[<br>  "Linux"<br>]</pre> | no |
| <a name="input_catalog_data_usage_text"></a> [catalog\_data\_usage\_text](#input\_catalog\_data\_usage\_text) | Public usage instructions in markdown format | `string` | `"# Usage\n\n## Authentication\n```bash\n# Get authorization token and login to ECR Public\naws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws\n```\n\n## Pull Image\n```bash\ndocker pull public.ecr.aws/your-registry/my-app:latest\n```\n\n## Push Image\n```bash\n# Tag your image\ndocker tag my-app:latest public.ecr.aws/your-registry/my-app:latest\n\n# Push to ECR Public\ndocker push public.ecr.aws/your-registry/my-app:latest\n```\n"` | no |
| <a name="input_repository_name"></a> [repository\_name](#input\_repository\_name) | Name of the repository | `string` | `"my-app"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_authorization_token"></a> [authorization\_token](#output\_authorization\_token) | The authorization token (base64 encoded) |
| <a name="output_aws_cli_login_command"></a> [aws\_cli\_login\_command](#output\_aws\_cli\_login\_command) | AWS CLI command to login to ECR Public |
| <a name="output_docker_login_command"></a> [docker\_login\_command](#output\_docker\_login\_command) | Docker login command for ECR Public (use with authorization_token output) |
| <a name="output_registry_id"></a> [registry\_id](#output\_registry\_id) | The registry ID where the repository was created |
| <a name="output_repository_arn"></a> [repository\_arn](#output\_repository\_arn) | Full ARN of the repository |
| <a name="output_repository_name"></a> [repository\_name](#output\_repository\_name) | Name of the repository |
| <a name="output_repository_uri"></a> [repository\_uri](#output\_repository\_uri) | The URI of the repository |
| <a name="output_repository_url"></a> [repository\_url](#output\_repository\_url) | The URL of the repository |
| <a name="output_token_expires_at"></a> [token\_expires\_at](#output\_token\_expires\_at) | Token expiration timestamp |
<!-- END_TF_DOCS -->