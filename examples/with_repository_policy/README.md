# ECR Public Repository with Policy Example

This example demonstrates how to create an ECR Public repository with a repository policy to control push access for CI/CD systems and cross-account scenarios.

## Overview

Repository policies for ECR Public allow you to control who can push images to your public repository while maintaining public read access. This is particularly useful for:

- **CI/CD Systems**: Grant push access to specific roles used by GitHub Actions, GitLab CI, or other automation tools
- **Cross-Account Access**: Allow other AWS accounts to push images to your repository
- **Team-based Access**: Control which IAM users or roles within your organization can push updates

## Prerequisites

- Terraform >= 1.0
- AWS CLI configured with appropriate credentials
- ECR Public repositories must be created in the `us-east-1` region

## Repository Policy Patterns

### 1. CI/CD Push Access

Grant push access to a specific CI/CD role (e.g., GitHub Actions):

```hcl
module "public-ecr" {
  source = "lgallard/ecrpublic/aws"

  repository_name = "my-public-app"

  create_repository_policy = true
  repository_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowPushFromCICD"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:aws:iam::123456789012:role/GitHubActionsRole"
        }
        Action = [
          "ecr-public:BatchCheckLayerAvailability",
          "ecr-public:PutImage",
          "ecr-public:InitiateLayerUpload",
          "ecr-public:UploadLayerPart",
          "ecr-public:CompleteLayerUpload"
        ]
      }
    ]
  })

  catalog_data = {
    description       = "Public container image with CI/CD push access"
    about_text        = "# My Public App\n\nThis image is automatically built and published via CI/CD."
    architectures     = ["x86-64", "ARM 64"]
    operating_systems = ["Linux"]
  }
}
```

### 2. Cross-Account Access

Allow another AWS account to push images:

```hcl
repository_policy = jsonencode({
  Version = "2012-10-17"
  Statement = [
    {
      Sid       = "AllowCrossAccountPush"
      Effect    = "Allow"
      Principal = {
        AWS = "arn:aws:iam::987654321098:root"
      }
      Action = [
        "ecr-public:BatchCheckLayerAvailability",
        "ecr-public:PutImage",
        "ecr-public:InitiateLayerUpload",
        "ecr-public:UploadLayerPart",
        "ecr-public:CompleteLayerUpload"
      ]
      Condition = {
        StringEquals = {
          "ecr-public:username" = "allowed-user"
        }
      }
    }
  ]
})
```

### 3. Multiple Roles with Different Permissions

Grant different levels of access to multiple principals:

```hcl
repository_policy = jsonencode({
  Version = "2012-10-17"
  Statement = [
    {
      Sid       = "AllowDevelopmentPush"
      Effect    = "Allow"
      Principal = {
        AWS = [
          "arn:aws:iam::123456789012:role/DeveloperRole",
          "arn:aws:iam::123456789012:role/GitHubActionsRole"
        ]
      }
      Action = [
        "ecr-public:BatchCheckLayerAvailability",
        "ecr-public:PutImage",
        "ecr-public:InitiateLayerUpload",
        "ecr-public:UploadLayerPart",
        "ecr-public:CompleteLayerUpload"
      ]
    },
    {
      Sid       = "AllowProductionPushWithCondition"
      Effect    = "Allow"
      Principal = {
        AWS = "arn:aws:iam::123456789012:role/ProductionDeployRole"
      }
      Action = [
        "ecr-public:BatchCheckLayerAvailability",
        "ecr-public:PutImage",
        "ecr-public:InitiateLayerUpload",
        "ecr-public:UploadLayerPart",
        "ecr-public:CompleteLayerUpload"
      ]
      Condition = {
        StringLike = {
          "ecr-public:image-tag" = ["v*", "release-*"]
        }
      }
    }
  ]
})
```

## Usage

### Basic Example

1. **Configure variables** in `terraform.tfvars`:

```hcl
repository_name = "my-company-app"

create_repository_policy = true

repository_policy = jsonencode({
  Version = "2012-10-17"
  Statement = [
    {
      Sid       = "AllowPushFromCICD"
      Effect    = "Allow"
      Principal = {
        AWS = "arn:aws:iam::123456789012:role/GitHubActionsRole"
      }
      Action = [
        "ecr-public:BatchCheckLayerAvailability",
        "ecr-public:PutImage",
        "ecr-public:InitiateLayerUpload",
        "ecr-public:UploadLayerPart",
        "ecr-public:CompleteLayerUpload"
      ]
    }
  ]
})

catalog_data = {
  description       = "My company's public application"
  about_text        = "# My Company App\n\nA production-ready containerized application."
  usage_text        = "# Usage\n\n```bash\ndocker pull public.ecr.aws/my-registry/my-company-app:latest\ndocker run -p 8080:8080 public.ecr.aws/my-registry/my-company-app:latest\n```"
  architectures     = ["x86-64", "ARM 64"]
  operating_systems = ["Linux"]
}

tags = {
  Environment = "production"
  Team        = "platform"
  ManagedBy   = "terraform"
}
```

2. **Deploy the infrastructure**:

```bash
terraform init
terraform plan
terraform apply
```

### With GitHub Actions Integration

For GitHub Actions integration, your workflow needs the appropriate IAM role. Here's an example setup:

**GitHub Actions Workflow:**
```yaml
name: Build and Push to ECR Public

on:
  push:
    branches: [main]

permissions:
  id-token: write
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: arn:aws:iam::123456789012:role/GitHubActionsRole
          aws-region: us-east-1

      - name: Login to Amazon ECR Public
        id: login-ecr-public
        uses: aws-actions/amazon-ecr-login@v2
        with:
          registry-type: public

      - name: Build and push image
        env:
          REGISTRY: ${{ steps.login-ecr-public.outputs.registry }}
          REPOSITORY: my-public-app
          IMAGE_TAG: ${{ github.sha }}
        run: |
          docker build -t $REGISTRY/$REPOSITORY:$IMAGE_TAG .
          docker push $REGISTRY/$REPOSITORY:$IMAGE_TAG
```

## Important Notes

### ECR Public Specifics

- **Region Constraint**: ECR Public repositories must be created in `us-east-1`
- **Public Access**: All images are publicly accessible for pull; policies only control push access
- **Global Availability**: Images can be pulled from any AWS region despite being stored in `us-east-1`

### Security Considerations

- **Principle of Least Privilege**: Only grant the minimum permissions required
- **Condition-Based Access**: Use conditions to restrict access based on image tags, time, or other factors
- **Regular Review**: Periodically review and update repository policies
- **Logging**: Enable CloudTrail to monitor repository access and changes

### Policy Actions

Common ECR Public actions for repository policies:

- `ecr-public:BatchCheckLayerAvailability`: Check if layers exist
- `ecr-public:PutImage`: Push completed images
- `ecr-public:InitiateLayerUpload`: Start uploading image layers
- `ecr-public:UploadLayerPart`: Upload parts of image layers
- `ecr-public:CompleteLayerUpload`: Complete layer upload

### Troubleshooting

**Common Issues:**

1. **Permission Denied**: Ensure the IAM role/user has the necessary ECR Public permissions
2. **Policy Validation**: Repository policies must be valid JSON with proper IAM syntax
3. **Region Issues**: Remember that ECR Public operations must target `us-east-1`

**Verification:**

```bash
# Check repository policy
aws ecr-public get-repository-policy --repository-name my-public-app --region us-east-1

# List repositories
aws ecr-public describe-repositories --region us-east-1

# Test push access (after proper authentication)
docker tag my-app:latest public.ecr.aws/my-registry/my-public-app:latest
docker push public.ecr.aws/my-registry/my-public-app:latest
```

## Outputs

This example provides the following outputs:

- `repository_uri`: The full URI for pulling/pushing images
- `repository_name`: The name of the created repository
- `registry_id`: The registry ID where the repository was created
- `repository_arn`: The ARN of the repository
- `repository_policy`: The applied repository policy JSON

## Resources Created

- **aws_ecrpublic_repository**: The ECR Public repository
- **aws_ecrpublic_repository_policy**: The repository access policy (if `create_repository_policy = true`)

## Clean Up

To remove all resources created by this example:

```bash
terraform destroy
```

Note: Ensure the repository is empty before destruction, as repositories containing images cannot be deleted.