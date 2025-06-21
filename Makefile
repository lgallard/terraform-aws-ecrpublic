.PHONY: fmt-check validate test check

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

# Test - run terratest suite
test:
	@echo "Running Terratest suite..."
	@cd test && go mod tidy && go test -v -parallel 4

# Check - run all checks (recommended for development)
check: fmt-check validate test
	@echo "All checks passed!"

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