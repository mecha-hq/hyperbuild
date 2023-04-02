# Hyperbuild

[![Open Source Saturday](https://img.shields.io/badge/%E2%9D%A4%EF%B8%8F-open%20source%20saturday-F64060.svg)](https://www.meetup.com/it-IT/Open-Source-Saturday-Milano/)

A tool for building a high number of software images in parallel.

## Dependencies

Most of the tooling needed to work with hyperbuild is listed in the .tool-versions file: you can use [asdf](https://asdf-vm.com) to install them. Any dependency not supported by asdf is listed in the Brewfile, to use as a secondary installation method.
All project actions are also runnable via Docker: the provided Makefile supports both "local" and "docker" versions for each action, to allow developers to use their preferred toolchain, e.g. `make lint-yaml` and `make lint-yaml-docker`.
