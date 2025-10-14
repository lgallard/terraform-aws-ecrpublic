---
name: New ECR Public Feature
about: Request implementation of a new AWS ECR Public feature discovered in the provider
title: 'feat: Add support for [FEATURE_NAME]'
labels: 'enhancement,aws-provider-update,auto-discovered'
assignees: ''
---

## Feature Discovery Details

**AWS Provider Version:** <!-- e.g., 5.82.0 -->
**Resource/Data Source:** <!-- e.g., aws_ecrpublic_repository -->
**Discovery Date:** <!-- Auto-filled by workflow -->

## New Feature Description

<!-- Provide a clear and concise description of the new ECR Public feature discovered -->

### Feature Type
- [ ] New resource (`aws_ecrpublic_*`)
- [ ] New data source (`data.aws_ecrpublic_*`)
- [ ] New argument/attribute on existing resource
- [ ] New catalog data field for ECR Public Gallery
- [ ] New repository configuration option
- [ ] New authentication/authorization feature
- [ ] New registry scanning capability
- [ ] Other: ___________

### Technical Details

**Resource/Argument Name:** 
**Documentation Link:** 
**Provider Changelog Reference:** 

### Implementation Requirements

#### Primary Changes Needed
- [ ] Update `main.tf` with new resource/argument
- [ ] Add variables in `variables.tf` with appropriate validation
- [ ] Update `outputs.tf` if new attributes are available
- [ ] Add support in locals for flexible configuration

#### Configuration Pattern
- [ ] Support object-based configuration (via `var.config_object`)
- [ ] Support individual variable configuration (via `var.config_*`)
- [ ] Maintain backward compatibility with existing variables
- [ ] Follow ECR Public Gallery best practices

#### Documentation & Examples
- [ ] Update module README with new capabilities
- [ ] Add example in `examples/using_objects/`
- [ ] Add example in `examples/using_variables/`
- [ ] Update variable descriptions and validation rules

#### Testing Requirements
- [ ] Add Terratest integration tests for new feature
- [ ] Test both configuration patterns (object vs variables)
- [ ] Validate ECR Public Gallery integration if applicable
- [ ] Test catalog data validation rules
- [ ] Add negative test cases for validation

### ECR Public Specific Considerations

#### Region Constraints
- [ ] Verify feature works in us-east-1 (ECR Public requirement)
- [ ] Update provider configuration examples if needed

#### Public Gallery Impact
- [ ] Check if feature affects ECR Public Gallery visibility
- [ ] Update catalog data configuration if applicable
- [ ] Verify public repository accessibility patterns

#### Security & Content Validation
- [ ] Add input validation for public-facing content
- [ ] Implement security checks for user-provided data
- [ ] Follow content guidelines for public repositories

### Validation Checklist

#### Code Quality
- [ ] Follow existing module patterns and conventions
- [ ] Use descriptive variable names and documentation
- [ ] Implement comprehensive input validation
- [ ] Maintain consistent coding style

#### Compatibility
- [ ] Maintain backward compatibility
- [ ] Support both configuration approaches
- [ ] Follow semantic versioning for changes
- [ ] Update CHANGELOG.md appropriately

#### ECR Public Best Practices
- [ ] Optimize for ECR Public Gallery discoverability
- [ ] Support rich markdown content in catalog data
- [ ] Implement proper architecture and OS tagging
- [ ] Follow public repository security practices

### Implementation Notes

<!-- Add any additional context, concerns, or implementation notes discovered during analysis -->

### Definition of Done

- [ ] Feature implemented with both configuration patterns
- [ ] Comprehensive tests added and passing
- [ ] Documentation updated (README, examples, variables)
- [ ] Backward compatibility maintained
- [ ] Security validation implemented
- [ ] ECR Public Gallery optimization considered
- [ ] Module released with proper versioning

---

**Auto-discovered by:** ECR Public Feature Discovery Bot
**Related Documentation:** [AWS ECR Public Documentation](https://docs.aws.amazon.com/AmazonECR/latest/public-userguide/)