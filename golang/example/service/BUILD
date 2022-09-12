# Copyright 2022 Ryan McGuinness
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

# Create the go proto library for the services
# with a dependency on the model.
go_proto_library(
  name = "service_proto_lib",
  visibility = ["//golang:__subpackages__"],
  compilers = ["@io_bazel_rules_go//proto:go_grpc"],
  importpath = "example/service",
  proto = "//api:service",
  deps = [
    "//golang/example/model:model_proto_lib"
  ]
)

## Usable libraries

# Wrap the service libaraies into a go library and expose thier import path.
# by doing so, this will give our Go code direct access by simply extending the
# 'service' package.
go_library(
  name = "service",
  srcs = [
    "common.go",
    "customers.go",
    "items.go",
    "stores.go",
    "transactions.go"
  ],
  embed = [":service_proto_lib"],
  importpath = "example/service",
  visibility = ["//visibility:public"],
  deps = [
    "@com_github_google_uuid//:uuid",
    "//golang/example/model",
  ]
)
