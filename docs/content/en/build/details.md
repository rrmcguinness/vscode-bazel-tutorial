---
title: "The Build in Detail"
weight: 1
icon: "pen-solid"
---

<!--
 Copyright 2022 Ryan McGuinness

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

## What happend

1. When you ran `bazel build //...`, bazel read the .bazelrc file in this directory.
2. .bazelrc instructed bazel to download all of the dependencies to the `build` directory under projects.
   - The dependency cache is placed here so we can reference our GoLang, Java, Protobuf, and Downloaded dependencies using relative paths.
3. Bazel then reads the WORKSPACE, and each recursive BUILD file to create a set of dependencies, this includes all of the binaries for Java, Go, etc.
4. You either had the right directory, or you made some changes and edited the workspace.
   - This is how the Java and Go environments are setup, so it's pretty important.
   - Again, with the workspace limitations, Java Projects MAY NOT have multiple Java dependencies, this ONLY impacts the IDEs ability to do auto complete, but DOES NOT represent the compile or runtime dependencies. This MAY cause problems.

## Understanding the Layout

- api - This is where the protocol buffer files are located.
- conf - These are configuration files used by Bazel.
- docs - a Hugo website (WIP).
- golang - The Go Project implementing the generated APIs.
- java - The Java Project implementing the generated APIS.
- third_party - A best practice directory for storing third_party assets.
- tools - Helper scripts for VSCode to utilize Go with Bazel.

### Important Files

- vscode-bazel-tutorial.code-workspace - responsible for configuring vscode.
- WORKSPACE - responsible for dependency management in Bazel.
- api/BUILD - the build file for the API model and GRPC.
- golang/BUILD - the Go build file with dependencies on the API.
- java/BUILD - the Java build file with dependencies on the API.

## Build Details

Just so you know what's happening behind the scenes when we execute `bazel build //...`. From above, we have a rough understanding, this section will provide a little more detail, starting from the WORKSPACE.

1. As bazel reads the workspace, it determines which languages and features will be supported. From top to bottom:
   - Protocol Buffers
   - The Protol Buffer Tool Chain (Protoc)
   - GRPC
   - GRPC Documentation Generator
   - Go Toolchain (version 1.19.1)
   - Gazelle (A build and dependency helper for bazel)
   - Go GRPC rules.
   - The Java Tool Chain (Java 17)
   - The internal Maven Tool Chain
   - The project Maven Dependencies
   - Hugo Site Generation Tools
   - A Hug Theme (GeekDoc)
2. Onces those tool chains have been loaded, each directory BUILD file is read, and a dependency graph is created. That dependency graph determines the build order.
   - API - Is built first because Java and Go depend on it, it uses the GRPC toolchain.
   - Java - Is built, using the Java Toolchain, maven dependencies, and the protocol definitions. The output of the generated files are in the 'build' directory.
   - Go - Is built, using the Go Tool Chain, and the dependencies from the Gazelle (go_deps.bzl) file.
3. The docs directory is built using the Hugo toolchain and theme.
