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

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
  name = "client_lib",
  visibility = ["//golang/example:__subpackages__"],
  srcs = ["client.go"],
  importpath = "example/client",
  deps = [
    "//golang/example/model",
    "//golang/example/service",
    "@org_golang_google_grpc//:grpc",
    "@org_golang_google_protobuf//types/known/emptypb",
  ],
)

go_binary(
  name = "client",
  srcs = [
    "main.go"
  ],
  deps = [
    ":client_lib",
    "//golang/example/model",
    "//golang/example/service",
    "@org_golang_google_grpc//:grpc",
    "@org_golang_google_grpc//credentials",
  ]
)
