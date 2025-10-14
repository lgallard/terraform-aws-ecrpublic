---
name: ECR Public Deprecation Notice
about: Handle deprecation of ECR Public features in the AWS provider
title: 'chore: Handle deprecation of [FEATURE_NAME]'
labels: 'deprecation,breaking-change,auto-discovered'
assignees: ''
---

## Deprecation Discovery Details

**AWS Provider Version:** <!-- e.g., 5.82.0 -->
**Resource/Data Source:** <!-- e.g., aws_ecrpublic_repository -->
**Discovery Date:** <!-- Auto-filled by workflow -->
**Deprecation Timeline:** <!-- If known from provider documentation -->

## Deprecated Feature Information

<!-- Provide details about what has been deprecated in the AWS ECR Public provider -->

### Deprecation Type
- [ ] Entire resource deprecated
- [ ] Specific argument/attribute deprecated
- [ ] Configuration pattern deprecated
- [ ] Catalog data field deprecated
- [ ] Authentication method deprecated
- [ ] Other: ___________

### Technical Details

**Deprecated Item:** 
**Replacement/Alternative:** 
**Provider Documentation:** 
**Deprecation Warning Message:** 

### Impact Assessment

#### Current Module Usage
- [ ] Feature is currently implemented in the module
- [ ] Feature is used in examples
- [ ] Feature is documented in README
- [ ] Feature has associated tests

#### Breaking Change Analysis
- [ ] Immediate breaking change (removes functionality)
- [ ] Gradual deprecation (warnings only)
- [ ] Alternative available in same provider version
- [ ] Migration path exists

### Migration Strategy

#### Phase 1: Preparation
- [ ] Add deprecation warnings to variable descriptions
- [ ] Update documentation with migration guidance
- [ ] Add validation warnings for deprecated usage
- [ ] Create migration examples

#### Phase 2: Implementation
- [ ] Implement alternative/replacement feature
- [ ] Update all examples to use new approach
- [ ] Add comprehensive tests for new implementation
- [ ] Maintain backward compatibility with warnings

#### Phase 3: Cleanup (Future Release)
- [ ] Remove deprecated feature support
- [ ] Update variable validation to reject deprecated options
- [ ] Clean up documentation and examples
- [ ] Coordinate breaking change release

### ECR Public Specific Considerations

#### Public Gallery Impact
- [ ] Check if deprecation affects catalog data visibility
- [ ] Verify repository accessibility patterns remain intact
- [ ] Update public gallery optimization strategies
- [ ] Review content validation rules

#### Region Constraints
- [ ] Confirm us-east-1 region requirements unchanged
- [ ] Verify provider configuration patterns
- [ ] Update authentication examples if needed

#### Security Implications
- [ ] Review if deprecated feature creates security concerns
- [ ] Update input validation for replacement features
- [ ] Maintain content security for public repositories

### Implementation Timeline

#### Immediate Actions (Current Release)
- [ ] Document deprecation in CHANGELOG.md
- [ ] Add deprecation warnings to affected variables
- [ ] Update README with migration guidance
- [ ] Add tests for deprecated feature warnings

#### Next Minor Release
- [ ] Implement replacement feature if available
- [ ] Update examples to demonstrate new approach
- [ ] Provide comprehensive migration documentation
- [ ] Add validation for smooth transition

#### Future Major Release
- [ ] Remove deprecated feature support completely
- [ ] Update module to only support replacement approach
- [ ] Coordinate with semantic versioning strategy
- [ ] Communicate breaking changes clearly

### Validation Checklist

#### Documentation Updates
- [ ] CHANGELOG.md entry added
- [ ] README.md migration section added
- [ ] Variable descriptions include deprecation warnings
- [ ] Examples updated or marked as deprecated

#### Backward Compatibility
- [ ] Current users not immediately broken
- [ ] Clear migration path provided
- [ ] Warnings guide users to alternatives
- [ ] Comprehensive testing covers transition

#### Communication Strategy
- [ ] Release notes clearly explain deprecation
- [ ] Migration guide is comprehensive
- [ ] Examples demonstrate best practices
- [ ] Deprecation timeline is clear

### Testing Requirements

#### Deprecation Warnings
- [ ] Test that deprecation warnings are shown appropriately
- [ ] Verify deprecated features still function during transition
- [ ] Test validation prevents problematic usage
- [ ] Add tests for replacement features

#### Migration Testing
- [ ] Test migration from deprecated to new approach
- [ ] Verify no functional regression during transition
- [ ] Test both configuration patterns work during overlap
- [ ] Validate ECR Public Gallery compatibility

### Implementation Notes

<!-- Add any additional context about the deprecation, migration challenges, or timeline considerations -->

### Definition of Done

- [ ] Deprecation impact fully assessed
- [ ] Migration strategy documented and implemented
- [ ] Backward compatibility maintained during transition
- [ ] Users have clear path forward
- [ ] Tests validate both deprecated and new approaches
- [ ] Documentation guides users through migration
- [ ] Timeline for complete deprecation established

---

**Auto-discovered by:** ECR Public Feature Discovery Bot
**Migration Guide:** <!-- Link to detailed migration documentation when available -->
**Related Documentation:** [AWS ECR Public Documentation](https://docs.aws.amazon.com/AmazonECR/latest/public-userguide/)