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

load("@rules_java//java:defs.bzl", "java_library")
load("@rules_proto_grpc//java:defs.bzl", "java_proto_library", "java_grpc_library")

package(default_visibility = ["//:__subpackages__"])

java_library(
    name = "service_deps",
    visibility = ["//visibility:public"],
    exports = [
        "@maven//:io_grpc_grpc_api",
        "@maven//:io_grpc_grpc_netty_shaded",
        "@maven//:io_grpc_grpc_protobuf",
        "@maven//:io_grpc_grpc_stub",
        "@maven//:org_apache_commons_commons_lang3",
        "@maven//:org_apache_logging_log4j_log4j_api",
        "@maven//:org_apache_logging_log4j_log4j_core",
        "@maven//:org_apache_tomcat_annotations_api",
    ],
)

# Build the services, and depend on the model.
java_grpc_library(
    name = "service_proto",
    protos = [
        "//api:service",
    ],
    deps = [
        "//java/model:model_proto",
    ],
)

java_library(
    name = "service",
    srcs = glob([
        "example/service/*.java",
    ]),
    deps = [
        ":service_deps",
        "//java/model:model_proto",
        "//java/model",
        ":service_proto",
    ],
)
