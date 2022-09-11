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

# Configuration (conf)

The files in this directory are either generated (via Gazelle), or created to be helper methods for the primary WORKSPACE or BUILD functions.

## Build Targets

- `bazel build //api/...` builds all service targets.
- `bazel build //api:model` builds the model objects.
- `bazel build //api:services` builds the service objects.
- `bazel build //api:docs` generates markdown documentation from the proto comments.

## copy_file_groups

Is a helper method for copying file groups into another module. This is mostly useful for static third_party assets.

## go_deps.bzl

Is auto genderated using the `bazel run //:gazelle-update-repos` command. This reads the `golang/go.mod` file and fetches all of the module dependencies into a Bazel friedly filegroup.

## junit.bzl

This is used to extend Bazel default test functionality with JUnit 5 test cases.

## repository_rules.bzl

Repository rules and macros which are expected to be called from WORKSPACE file of either
googleapis itself or any third_party repository which consumes googleapis as its dependency.
