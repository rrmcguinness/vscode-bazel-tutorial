# Visual Studio Code - Bazel Workspace

## TL;DR;

This repository contains an example project that uses Protocol Buffers (protobuf) to create a domain model and services. There are two example implementations, one in Go, the other in Java. The functionality is stubbed out, but shows how to build a multiple language project using Bazel. Lastly, it ties it all together with a predefined workspace and task to make it repeatable.

![Overview](/docs/content/en/arch-view.png)

### What this IS NOT

This repository IS NOT a general purpose tutorial for how to use Bazel, Visual Studio Code, Java, or Go. It is intended for those who want their IDE to work with a mono-repo / multi-project build system. That being said, I've attempted to add a README in each module to explain how it was built.

## Prerequisites

- Visual Studio Code (vscode) - [Download and Install](https://code.visualstudio.com/)
- Bazel - [Install Instructions](https://bazel.build/start)
- Knowledge of your Terminal / Command Prompt

> Once Bazel is installed, make sure you can execute `bazel help` from the terminal.

## Verified

- :white_check_mark: Linux (Ubuntu / Linux Mint)
- :white_check_mark: Mac OS X
- :white_check_mark: Windows 10 with Windows Subsystem for Linux (WSL)
- :white_check_mark: Chrome OS (Version 105+) with Developer Linux Tools
- :x: Windows 10 Native - Not working due to path requirements.
- :white_check_mark: Java Language Server and Auto Complete
- :large_orange_diamond: - Go Language Server and Auto Complete

## Known Issues

- :bangbang: Many of the language tools require fixed, non-relative, non-expanding paths to toolsets. This is painful when your build tool can construct the tool-chain.
    - [Java Extension Open Bug](https://github.com/redhat-developer/vscode-java/issues/1357)
- :heavy_exclamation_mark: Gopls Server and Protobufs not honoring the $GOBIN variable, nor have I found a setting to add it to the server.

## Setup

- [Document Pages](docs/content/en/setup.md)

## Build

- [How to Build](docs/content/en/build/_index.md)
- [Understanding the Build Process](docs/content/en/build/details.md)
