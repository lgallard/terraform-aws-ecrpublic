# Changelog

## [0.13.1](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.13.0...0.13.1) (2025-12-28)


### Bug Fixes

* simplify CI by skipping terraform_docs in favor of local pre-commit ([5e1080a](https://github.com/lgallard/terraform-aws-ecrpublic/commit/5e1080a3db9ba9df5a72170f28a53125e54d23bd))

## [0.13.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.12.0...0.13.0) (2025-12-28)


### Features

* Add enhanced outputs (tags_all, gallery_url, catalog_data) ([#63](https://github.com/lgallard/terraform-aws-ecrpublic/issues/63)) ([bb06dec](https://github.com/lgallard/terraform-aws-ecrpublic/commit/bb06dec16af77d10b1ee0411178f3cbf11447aca))
* Add multiple repositories support via module-level for_each pattern ([#61](https://github.com/lgallard/terraform-aws-ecrpublic/issues/61)) ([c6e0a4e](https://github.com/lgallard/terraform-aws-ecrpublic/commit/c6e0a4e6e04b4931db8dead504ce3052e849036e))


### Bug Fixes

* add debugging and update cache key in pre-commit workflow ([282a490](https://github.com/lgallard/terraform-aws-ecrpublic/commit/282a4900cbd9f919211fa0df63e336f863b6fd55))
* clean up duplicate documentation markers in README.md ([b3667f0](https://github.com/lgallard/terraform-aws-ecrpublic/commit/b3667f0ea3cbb9d8feac6fbd2031ec7338417230))
* revert to standard pre-commit terraform-docs markers ([2c44ef7](https://github.com/lgallard/terraform-aws-ecrpublic/commit/2c44ef7f8263cb626930bb17185cfb58667a1563))
* standardize terraform-docs configuration and markers ([d2e2614](https://github.com/lgallard/terraform-aws-ecrpublic/commit/d2e26145b5ba35e82fc0b83c5f0114e9ad78d3a3))
* update example READMEs with terraform-docs ([1220381](https://github.com/lgallard/terraform-aws-ecrpublic/commit/1220381c3e51e730006fd224b0fbea85987f0463))
* update terraform-docs to v0.20.0 in CI workflow ([a0a7d64](https://github.com/lgallard/terraform-aws-ecrpublic/commit/a0a7d64c0285969c83244b23b8905325187102e6))

## [0.12.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.11.0...0.12.0) (2025-12-27)


### Features

* Add aws_ecrpublic_repository_policy resource support ([#56](https://github.com/lgallard/terraform-aws-ecrpublic/issues/56)) ([43b893b](https://github.com/lgallard/terraform-aws-ecrpublic/commit/43b893b4188448bd347904a84d562beb695cf84b))


### Bug Fixes

* correct pre-commit formatting issues ([#59](https://github.com/lgallard/terraform-aws-ecrpublic/issues/59)) ([0c0f0b5](https://github.com/lgallard/terraform-aws-ecrpublic/commit/0c0f0b5ef1202133c08df9f7ce6e75194b20ec85))

## [0.11.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.10.0...0.11.0) (2025-08-17)


### Features

* improve security and policy best practices ([#47](https://github.com/lgallard/terraform-aws-ecrpublic/issues/47)) ([6eb8844](https://github.com/lgallard/terraform-aws-ecrpublic/commit/6eb884479eebfcaa2db706aba5196fd3d5eb90ef))


### Bug Fixes

* update Go version to 1.23 and configure Renovate to prevent pre-release versions ([#49](https://github.com/lgallard/terraform-aws-ecrpublic/issues/49)) ([f13071d](https://github.com/lgallard/terraform-aws-ecrpublic/commit/f13071da5ef5c52a8a5a63a3d2408e00c644ba9b))

## [0.10.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.9.0...0.10.0) (2025-08-10)


### Features

* add comprehensive input validation for variables ([#31](https://github.com/lgallard/terraform-aws-ecrpublic/issues/31)) ([36005b6](https://github.com/lgallard/terraform-aws-ecrpublic/commit/36005b6777e401bd466dd89d07f95fd8c98a7c88))

## [0.9.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.8.0...0.9.0) (2025-08-10)


### Features

* add pre-commit workflow for automated code quality checks ([4b7f5c1](https://github.com/lgallard/terraform-aws-ecrpublic/commit/4b7f5c13803848ced32a580668c2e4dedbb96125))


### Bug Fixes

* add validation rules and fix security workflow issues ([#37](https://github.com/lgallard/terraform-aws-ecrpublic/issues/37)) ([e9f1b4c](https://github.com/lgallard/terraform-aws-ecrpublic/commit/e9f1b4c97b5883f5bf5028a38d57e37c84155b4b))

## [0.8.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.7.0...0.8.0) (2025-08-09)


### Features

* add Claude dispatch workflow for repository events ([#30](https://github.com/lgallard/terraform-aws-ecrpublic/issues/30)) ([df53765](https://github.com/lgallard/terraform-aws-ecrpublic/commit/df53765d16b4d6129bd4863bd8498db1ac4ba134))
* add MCP server support for enhanced documentation access ([#32](https://github.com/lgallard/terraform-aws-ecrpublic/issues/32)) ([df1188d](https://github.com/lgallard/terraform-aws-ecrpublic/commit/df1188ddd7ccf2cbafae08ec798945a6f476e779))

## [0.7.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.6.0...0.7.0) (2025-07-30)


### Features

* replicate security-hardened Claude Code Review workflow with PR focus ([#27](https://github.com/lgallard/terraform-aws-ecrpublic/issues/27)) ([c53d77e](https://github.com/lgallard/terraform-aws-ecrpublic/commit/c53d77e0d57953258df50b7ec5645f5996a6cdb6))

## [0.6.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.5.0...0.6.0) (2025-07-28)


### Features

* standardize GitHub Actions workflows for ECR Public module ([#24](https://github.com/lgallard/terraform-aws-ecrpublic/issues/24)) ([5e600b3](https://github.com/lgallard/terraform-aws-ecrpublic/commit/5e600b34967ae835c3a41f0c152a011b1604fee3))

## [0.5.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.4.0...0.5.0) (2025-07-28)


### Features

* add Renovate for dependency management ([#18](https://github.com/lgallard/terraform-aws-ecrpublic/issues/18)) ([d3dc2a4](https://github.com/lgallard/terraform-aws-ecrpublic/commit/d3dc2a4a74e625b79bd323e938e6691526883ee6))


### Bug Fixes

* standardize Claude Code Review workflow formatting ([24094ea](https://github.com/lgallard/terraform-aws-ecrpublic/commit/24094eaf06c14283cdf1f4d93778a687792e1f4e))

## [0.4.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.3.0...0.4.0) (2025-07-23)


### Features

* standardize release-please configuration ([#16](https://github.com/lgallard/terraform-aws-ecrpublic/issues/16)) ([04d67d9](https://github.com/lgallard/terraform-aws-ecrpublic/commit/04d67d9876fd4bc0310ea4e63834617241d9f396))

## [0.3.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.2.0...0.3.0) (2025-06-21)


### Features

* Set up automated testing and CI/CD integration ([#8](https://github.com/lgallard/terraform-aws-ecrpublic/issues/8)) ([50c5f90](https://github.com/lgallard/terraform-aws-ecrpublic/commit/50c5f90ce1ab7d4f345668c76de5381e60743362))

## [0.2.0](https://github.com/lgallard/terraform-aws-ecrpublic/compare/0.1.0...0.2.0) (2025-06-21)


### Features

* Implement release-please for automated versioning without "v" prefix ([#10](https://github.com/lgallard/terraform-aws-ecrpublic/issues/10)) ([caa0f86](https://github.com/lgallard/terraform-aws-ecrpublic/commit/caa0f868e0e3522dd1c979fd29c92e3db17c135f))

## 0.1.0 (April 9, 2021)

FEATURES:

  * Module implementation
