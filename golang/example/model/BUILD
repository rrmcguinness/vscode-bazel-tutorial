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

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")


# Model Objects

# Create the go proto library for the model.
go_proto_library(
  name="model_proto_lib",
  importpath = "example/model",
  proto = "//api:model",
  visibility = ["//golang:__subpackages__"],
)

# Wrap the model libraries into a go library and extend the model behaviors
go_library(
  name = "model",
  srcs = ["utils.go"],
  embed = [":model_proto_lib"],
  importpath = "example/model",
  visibility = ["//golang:__subpackages__"],
  deps = [
    "@com_github_google_uuid//:uuid"
  ]
)

go_test(
  name = "test",
  srcs = ["utils_test.go"],
  embed = [":model"],
  deps = [
    "@com_github_stretchr_testify//assert"
  ]
)
