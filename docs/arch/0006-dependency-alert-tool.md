# 6. Dependency alert tool

Date: 2023-04-02

## Status

Accepted

## Context

We need a tool that helps us update the project dependencies automatically, keeping it as secure as possible.

## Decision

We are going to use [Depedabot](https://github.com/dependabot) as the dependency maintainer: it is a GitHub built-in software that simplifies the update process, opening a PR when a new dependency update is available.

Another possible solution is the [Renovate bot](https://github.com/renovatebot/renovate) which performs the same operations and is portable in different workflows like Gitlab CI.

## Consequences

To keep things as much easier as possible, we will implement the Dependabot solution.
