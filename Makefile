.PHONY: fmt-check validate test test-static test-integration test-catalog test-gallery test-all check

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

# Test Integration - run basic integration tests (requires AWS credentials)
test-integration:
	@echo "Running basic integration tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestTerraformECRPublicBasic|TestTerraformECRPublicWithVariableCatalogData|TestTerraformECRPublicWithObjectCatalogData" -timeout 15m -parallel 1

# Test Catalog Data - run catalog data validation tests (requires AWS credentials)
test-catalog:
	@echo "Running catalog data validation tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestECRPublicCatalogData|TestECRPublicMinimalCatalogData|TestECRPublicComprehensiveCatalogData" -timeout 15m -parallel 1

# Test Gallery - run public gallery compliance tests (requires AWS credentials)
test-gallery:
	@echo "Running public gallery compliance tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestECRPublicGallery" -timeout 15m -parallel 1

# Test All Integration - run all integration tests (requires AWS credentials)
test-all:
	@echo "Running comprehensive integration test suite (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -timeout 15m -parallel 1

# Test Timeouts - run timeout configuration tests (requires AWS credentials)
test-timeouts:
	@echo "Running timeout configuration tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestTerraformECRPublicWithTimeouts|TestTerraformECRPublicCompleteConfiguration" -timeout 15m -parallel 1

# Test Validation - run variable validation tests (requires AWS credentials)
test-validation:
	@echo "Running variable validation tests (requires AWS credentials)..."
	@echo "⚠️  This will create real AWS resources and may incur costs!"
	@cd test && go mod tidy && go test -v -run "TestTerraformECRPublicVariableValidation|TestECRPublicMultiArchitectureSupport" -timeout 15m -parallel 1

# Check - run all static checks (recommended for development)
check: fmt-check validate test-static
	@echo "All static checks passed!"

# Check All - run all checks including integration tests
check-all: fmt-check validate test-static test-integration test-catalog test-gallery
	@echo "All checks including comprehensive integration tests passed!"

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

# Help - display available targets
help:
	@echo "Available targets:"
	@echo "  fmt-check      - Check Terraform code formatting"
	@echo "  validate       - Validate Terraform configuration"
	@echo "  fmt            - Format Terraform code"
	@echo "  test           - Run static tests only (no AWS resources)"
	@echo "  test-static    - Run static analysis tests"
	@echo "  test-integration - Run basic integration tests (creates AWS resources)"
	@echo "  test-catalog   - Run catalog data validation tests"
	@echo "  test-gallery   - Run public gallery compliance tests"
	@echo "  test-timeouts  - Run timeout configuration tests"
	@echo "  test-validation - Run variable validation tests"
	@echo "  test-all       - Run all integration tests"
	@echo "  check          - Run all static checks"
	@echo "  check-all      - Run all checks including integration tests"
	@echo "  clean          - Clean up test artifacts"
	@echo "  help           - Display this help message"