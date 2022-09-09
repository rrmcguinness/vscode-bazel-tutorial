# Copyright 2022 Google LLC
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

load("@rules_java//java:defs.bzl", "java_library")
load("@rules_proto_grpc//java:defs.bzl", "java_proto_library")

java_library(
    name = "model_deps",
    visibility = ["//visibility:public"],
    exports = [
        "@maven//:org_apache_commons_commons_lang3",
        "@maven//:org_apache_logging_log4j_log4j_api",
    ],
)

# Build the model into a local library.
java_proto_library(
    name = "model_proto",
    visibility = ["//visibility:public"],
    protos = [
        "//api:model",
    ],
)

java_library(
    name = "model",
    visibility = ["//visibility:public"],
    srcs = glob([
        "example/model/*.java",
        "example/utils/*.java",
    ]),
    deps = [
        ":model_deps",
        ":model_proto",
    ]
)
