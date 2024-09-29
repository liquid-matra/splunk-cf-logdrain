# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## v0.1.0

Initial OSS release

## v0.2.0

- Bump github.com/philips-software/go-hsdp-api from 0.79.1 to 0.80.1 by @dependabot in #14
- Bump golang from 1.20.2-buster to 1.20.4-buster in /docker by @dependabot in #12
- Bump fluent/fluent-bit from 2.0.9-debug to 2.1.2-debug in /docker by @dependabot in #11

## v0.3.0

- Updated to suppoerted GO version `v1.23`
- Removed direct dependency `viper` and replaced with standard library
- Removed direct dependency `srslog` and replaced with standard library
- Replaced deprecated `v2` go-syslog package with `github.com/leodido/go-syslog/v4`
- Updated fluentbit version to `v3.1.9` in Dockerfile