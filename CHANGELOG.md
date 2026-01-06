# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2025-01-06

### Added
- GoReleaser configuration for multi-platform builds
- GitHub Actions workflow for automated releases
- GitHub Actions workflow for CI/CD testing
- GPG signing for releases
- Automated changelog generation
- GPG setup documentation

### Changed
- Release archives now include LICENSE and README.md files
- Improved .gitignore to exclude dist/ directory

### Fixed
- Deprecated GoReleaser configuration options

## [0.1.0] - 2025-01-05

### Added
- Initial release of Aeza Terraform Provider
- Support for VPS service creation and management
- Data sources: products, services, OS, groups, types, recipes
- Resources: service, service_actions, service_prolong
- API client with Legacy and V2 API support
- Examples for basic usage, data sources, and resources
- MPL-2.0 license
- Comprehensive README with usage examples

[Unreleased]: https://github.com/scinfra-pro/terraform-provider-aeza/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/scinfra-pro/terraform-provider-aeza/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/scinfra-pro/terraform-provider-aeza/releases/tag/v0.1.0

