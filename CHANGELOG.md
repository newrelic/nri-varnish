# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## 2.2.0 (2021-05-10)
## Changed
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
