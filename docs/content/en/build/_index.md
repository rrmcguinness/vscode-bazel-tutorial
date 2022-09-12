---
title: "The Build"
weight: 3
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

## Visual Studio Code

Using the command pallet, you may execute the build tasks specified in the 'tasks' section of the build file. These are IDE extension to the [Command Line Arguments](#command_line).

- [Apple OSX](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-macos.pdf) - cmd + opt + p
- [Windows](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf) - ctrl + shift + p
- [Linux](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-linux.pdf) - ctrl + shift + p

1. Open the command pallet
2. Type: Tasks: Run Tasks (This will autocomplete)
3. Select Desired Target
4. Don't Monitor Output

Available Tasks:

## General

- Build All
- Test All

### API

- Build API

### Documents

- Build Docs
- Run Document Site Locally

### Go

- Build Go Projects
- Run Go Protobuf Server
- Run Go Protobuf Client

### Java

- Build Java Projects
- Run Java Protobuf Server
- Run Java Protobuf Client

<a name='command_line'></a>

## Command Line

### Well know Targets

```shell

# Build All
bazel build //...

# Test All
bazel test //...
```

### API (Protocol Buffer Targets)

```shell
# Build APIs
bazel build //api
```

### Documentation Site

```shell
# Build Docs
bazel build //docs

# Serve Docs Locally
bazel run //docs:serve
```

### Go Project: Server & Client

```shell
# Build Go Projects
bazel build //golang/...

# Run Go Protobuf Server
bazel run //golang/example/server

# Run Go Protobuf Client
bazel run //golang/example/client
```

### Java Project: Server & Client

```shell
# Build Java Projects
bazel build //java/...

# Run Java Protobuf Server
bazel run //java/example/server

# Run Java Protobuf Client
bazel run //java/example/client
```
