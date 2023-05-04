# Hyperbuild


A tool for building a high number of software images in parallel.
## Badges
[![Open Source Saturday](https://img.shields.io/badge/%E2%9D%A4%EF%B8%8F-open%20source%20saturday-F64060.svg?style=for-the-badge)](https://www.meetup.com/it-IT/Open-Source-Saturday-Milano/)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mecha-ci/hyperbuild?style=for-the-badge)](https://tip.golang.org/doc/go1.20)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/mecha-ci/hyperbuild?style=for-the-badge)](https://github.com/mecha-ci/hyperbuild/releases/latest)
[![GitHub](https://img.shields.io/github/license/mecha-ci/hyperbuild?style=for-the-badge)](/LICENSE.md)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mecha-ci/hyperbuild?style=for-the-badge)](https://github.com/mecha-ci/hyperbuild)
[![GitHub repo file count (file type)](https://img.shields.io/github/directory-file-count/mecha-ci/hyperbuild?style=for-the-badge)](https://github.com/mecha-ci/hyperbuild)
[![GitHub all releases](https://img.shields.io/github/downloads/mecha-ci/hyperbuild/total?style=for-the-badge)](https://github.com/mecha-ci/hyperbuild)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/mecha-ci/hyperbuild?style=for-the-badge)](https://github.com/mecha-ci/hyperbuild/commits)
## Dependencies

Most of the tooling needed to work with hyperbuild is listed in the .tool-versions file: you can use [asdf](https://asdf-vm.com) to install them. Any dependency not supported by asdf is listed in the Brewfile, to use as a secondary installation method.
All project actions are also runnable via Docker: the provided Makefile supports both "local" and "docker" versions for each action, to allow developers to use their preferred toolchain, e.g. `make lint-yaml` and `make lint-yaml-docker`.
