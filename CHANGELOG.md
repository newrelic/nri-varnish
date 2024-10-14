# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

Unreleased section should follow [Release Toolkit](https://github.com/newrelic/release-toolkit#render-markdown-and-update-markdown)

## Unreleased

## v2.6.0 - 2024-10-14

### dependency
- Upgrade go to 1.23.2

### 🚀 Enhancements
- Upgrade integrations SDK so the interval is variable and allows intervals up to 5 minutes

## v2.5.10 - 2024-09-09

### ⛓️ Dependencies
- Updated golang version to v1.23.1

## v2.5.9 - 2024-07-08

### ⛓️ Dependencies
- Updated golang version to v1.22.5

## v2.5.8 - 2024-05-13

### ⛓️ Dependencies
- Updated golang version to v1.22.3

## v2.5.7 - 2024-04-15

### ⛓️ Dependencies
- Updated golang version to v1.22.2

## v2.5.6 - 2024-03-11

### 🐞 Bug fixes
- Updated golang to version v1.21.7 to fix a vulnerability

## v2.5.5 - 2024-02-26

### ⛓️ Dependencies
- Updated github.com/newrelic/infra-integrations-sdk to v3.8.2+incompatible

## v2.5.4 - 2023-11-06

### ⛓️ Dependencies
- Updated golang version to 1.21

## 2.5.3  (2022-07-04)

### Changed
- Bump dependencies
### Added
Added support for more distributions:
- RHEL(EL) 9
- Ubuntu 22.04

## 2.5.2 (2022-05-11)
### Changed
- Use Go 1.18
- Bump SDK to version v3.7.3 that introduce the store file entries TTL.

## 2.5.1 (2022-01-10)
### Fixed
- Fix negative values (#42)

## 2.5.0 (2021-11-25)
### Added
- Add monitoring to Varnish Plus attributes

## 2.4.0 (2021-09-16)
### Added
- Add support for V1 varnishstat output present on Varnish version > v6.4
- Add varnish_name config to allow specify an instnace name for varnishstat (-n)

## 2.3.0 (2021-08-27)
### Added

Moved default config.sample to [V4](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-newer-configuration-format/), added a dependency for infra-agent version 1.20.0

Please notice that old [V3](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-standard-configuration-format/) configuration format is deprecated, but still supported.

## 2.2.1 (2021-06-10)
### Changed
- Add ARM support

## 2.2.0 (2021-05-10)
### Changed
- Update Go to v1.16.
- Migrate to Go Modules
- Update Infrastracture SDK to v3.6.7.
- Update other dependecies.

## 2.1.1 (2021-03-24)
### Changed
- Added arm packages and binaries
## 2.1.0 (2019-11-18)
### Changed
- Renamed the integration executable from nr-varnish to nri-varnish in order to be consistent with the package naming. **Important Note:** if you have any security module rules (eg. SELinux), alerts or automation that depends on the name of this binary, these will have to be updated.
## 2.0.2 - 2019-10-23
### Added
- Windows MSI resources

## 2.0.1 - 2019-09-16
### Fixed
- Segfault on non-namespaced metrics

## 2.0.0 - 2019-04-22
### Changed
- Prefixed entity namespaces for uniqueness
- Added instanceName as an identity attribute

## 1.0.0 - 2019-02-05
### Changed
- Bumped version to 1.0.0

## 0.1.0 - 2018-11-02
### Added
- Initial version: Includes Metrics and Inventory data
