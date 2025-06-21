.PHONY: fmt-check validate test test-static test-integration check

# Format check - ensure all Terraform files are properly formatted
fmt-check:
	@echo "Checking Terraform formatting..."
	@terraform fmt -check -recursive

# Validate - validate Terraform configuration
validate:
	@echo "Validating main module..."
	@terraform init -backend=false
	@terraform validate
	@echo "Validating examples..."
	@cd examples/using_objects && terraform init -backend=false && terraform validate
	@cd examples/using_variables && terraform init -backend=false && terraform validate

# Test - run static terratest suite only
test: test-static

# Test Static - run static analysis tests (no AWS resources created)
test-static:
	@echo "Running static Terratest suite..."
	@cd test && go mod tidy && go test -v -run "TestTerraformValidate|TestTerraformFormat|TestExamplesValidate|TestExamplesFormat" -parallel 4

# Test Integration - run integration tests (requires AWS credentials)
test-integration:
	@echo "Running integration tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestTerraformECRPublic" -timeout 30m

# Check - run all static checks (recommended for development)
check: fmt-check validate test-static
	@echo "All static checks passed!"

# Check All - run all checks including integration tests
check-all: fmt-check validate test-static test-integration
	@echo "All checks including integration tests passed!"

# Format - format all Terraform files
fmt:
	@echo "Formatting Terraform files..."
	@terraform fmt -recursive

# Clean - clean up test artifacts
clean:
	@echo "Cleaning up..."
	@find . -name ".terraform" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name ".terraform.lock.hcl" -exec rm -f {} + 2>/dev/null || true
	@find . -name "terraform.tfstate*" -exec rm -f {} + 2>/dev/null || true