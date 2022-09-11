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

# Go Project

## TL;DR;

This project is a Go project that implements a service and a client from the Protocol Buffers.
Done to demonstrate how to implement both sides of a micro-service.

## Structure

- client - a client implementation
- model - a model extension
- service - the service business model implementation
- server - an example server

## Build Targets

### Build the libraries

```shell
bazel build //golang
```

### Run the server

```shell
bazel run //golang/server
```

### Run the client library

```shell
bazel run //golang/client
```

## Project Creation

### Creating the directory

#### Prerequists

This really depends on how familiar you are with Bazel, but if you're a native Go developer and have the Go toolchain available on your command line, then there is nothing new here. We simply wrap the toolchain with Bazel to make it buildable in any environment.

If you are not a native Go developer, then all you really have to do is install Go according to the go.dev tutorial.

## Creating the project

> Open your terminal / command prompt and excute the following commands from the project directory.

```shell

# Create the project directory structure
mkdir -p golang/client
mkdir golang/model
mkdir golang/server
mkdir golang/service

# Create the example module
cd golang
go mod init example

## Greate your project dependencies
## Test assertions
go get github.com/stretchr/testify@latest

## UUID
go get github.com/google/uuid@latest

## Protocol Generator for Go
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest

## GRPC Generator for Go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## Verify your mod file
tail go.mod

## You should see the dependencies there.

## Generate the dependency file
cd .. # Project root
bazel run //:gazelle-update-repos
```

## Time to code

Once complete, you can start adding business logic.
