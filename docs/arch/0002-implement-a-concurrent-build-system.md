# 2. Implement a concurrent build system

Date: 2023-04-01

## Status

Accepted

## Context

Companies often needs to build images using a variety of tools and for a variety of platforms.
It sometimes happens those images begin to be high in number, depending on each other, and in need to be built often, perhaps even several times a day.

## Decision

We are going to build a modular, concurrent images build tool that is able to handle graph of dependencies and build them at the highest level of concurrency.
We already have several executors in mind, such as Docker, Buildah and Packer, as well as a shell one to use as an 'escape hatch'.

## Consequences

We do not intend to build an actual CI system for the moment, but we do have in mind the possibility to evolve hyperbuild in that direction in the future: the initial focus is building graphs of (dependent) images with the highest possible concurrency.
