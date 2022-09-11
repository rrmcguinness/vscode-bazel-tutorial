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


go_proto_library(
  name="model_proto_lib",
  importpath = "example/model",
  proto = "//api:model",
)

go_proto_library(
  name = "service_proto_lib",
  compilers = ["@io_bazel_rules_go//proto:go_grpc"],
  importpath = "example/service",
  proto = "//api:service",
  deps = [
    ":model_proto_lib"
  ]
)

go_library(
  name = "go_model_lib",
  embed = [":model_proto_lib"],
  importpath = "example/model",
  visibility = ["//visibility:public"],
)

go_library(
  name = "go_service_lib",
  embed = [":service_proto_lib"],
  importpath = "example/service",
  visibility = ["//visibility:public"],
  deps = [
    ":go_model_lib"
  ]
)
