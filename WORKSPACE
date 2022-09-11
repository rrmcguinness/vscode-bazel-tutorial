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
workspace(name = "vscode_bazel_tutorial")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Load the Google API definitions for use with Protobufs
load("//:conf/repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = False,
    csharp = False,
    gapic = False,
    go = True,
    grpc = True,
    java = True,
    nodejs = False,
    php = False,
    python = False,
    ruby = False,
)

###############################################################################
# GO Tool Chain
###############################################################################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "16e9fca53ed6bd4ff4ad76facc9b7b651a89db1689a2877d6fd7b82aa824e366",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.34.0/rules_go-v0.34.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "501deb3d5695ab658e82f6f6f549ba681ea3ca2a5fb7911154b5aa45596183fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.26.0/bazel-gazelle-v0.26.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.13")

gazelle_dependencies()

###############################################################################
# Protobuf Tool Chain
###############################################################################

http_archive(
    name = "com_google_protobuf",
    sha256 = "7ba0cb2ecfd9e5d44a6fa9ce05f254b7e5cd70ec89fafba0b07448f3e258310c",
    strip_prefix = "protobuf-21.5",
    urls = [
        "https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protobuf-all-21.5.tar.gz",
    ],
)

http_archive(
    name = "rules_proto",
    sha256 = "e017528fd1c91c5a33f15493e3a398181a9e821a804eb7ff5acdd1d2d6c2b18d",
    strip_prefix = "rules_proto-4.0.0-3.20.0",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0-3.20.0.tar.gz",
    ],
)

http_archive(
    name = "rules_proto_grpc",
    sha256 = "bbe4db93499f5c9414926e46f9e35016999a4e9f6e3522482d3760dc61011070",
    strip_prefix = "rules_proto_grpc-4.2.0",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.2.0.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("//:conf/go_deps.bzl", "go_dependencies")

# gazelle:repository_macro conf/go_deps.bzl%go_dependencies
go_dependencies()

protobuf_deps()

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "e3bb0dc8b0274ea1aca75f1f8c0c835adbe589708ea89bf698069d0790701ea3",
    strip_prefix = "buildtools-5.1.0",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/5.1.0.tar.gz",
    ],
)

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")

rules_proto_grpc_go_repos()

load("@rules_proto_grpc//doc:repositories.bzl", rules_proto_grpc_doc_repos = "doc_repos")

rules_proto_grpc_doc_repos()

###############################################################################
# Java Tool Chain
###############################################################################

# Variables
RULES_JVM_EXTERNAL_TAG = "4.2"

RULES_JVM_EXTERNAL_SHA = (
    "cd1a77b7b02e8e008439ca76fd34f5b07aecb8c752961f9640dea15e9e5ba1ca"
)

JUNIT_JUPITER_VERSION = "5.8.2"

JUNIT_PLATFORM_VERSION = "1.8.2"

PROJECT_MAVEN_DEPENDENCIES = [
    "org.apache.commons:commons-lang3:3.12.0",
    "io.grpc:grpc-core:1.48.1",
    "io.grpc:grpc-netty-shaded:1.48.1",
    "io.grpc:grpc-protobuf:1.48.1",
    "io.grpc:grpc-stub:1.48.1",
    "org.apache.tomcat:annotations-api:6.0.53",
    "org.apache.logging.log4j:log4j-api:2.18.0",
    "org.apache.logging.log4j:log4j-core:2.18.0",
    "com.google.protobuf:protoc:3.21.5",
    "io.netty:netty-all:4.1.79.Final",
]

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" %
          RULES_JVM_EXTERNAL_TAG,
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

load(
    "@rules_proto_grpc//java:repositories.bzl",
    rules_proto_grpc_java_repos = "java_repos",
)

rules_proto_grpc_java_repos()

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = PROJECT_MAVEN_DEPENDENCIES,
    generate_compat_repositories = True,
    repositories = [
        "https://repo.maven.apache.org/maven2/",
    ],
)

load("@maven//:compat.bzl", "compat_repositories")

compat_repositories()

###############################################################################
# JUnit 5 Tool Chain
###############################################################################
load(
    "//:conf/junit.bzl",
    "junit_jupiter_java_repositories",
    "junit_platform_java_repositories",
)

junit_jupiter_java_repositories(
    version = JUNIT_JUPITER_VERSION,
)

junit_platform_java_repositories(
    version = JUNIT_PLATFORM_VERSION,
)

###############################################################################
# Hugo Tool Chain
###############################################################################
# Update these to latest
RULES_HUGO_COMMIT = "02234789fa9f2112807c1642eacb9f9728fc179d"

RULES_HUGO_SHA256 = "4ce20c981ad50ac0c956e85ef991e59b204778bde59d81e40be05450259ae969"

# Use a modified version of the popular rules_hugo, this modified version
# allows modern Hugo themes to be used.
http_archive(
    name = "build_stack_rules_hugo",
    sha256 = RULES_HUGO_SHA256,
    strip_prefix = "rules_hugo-%s" % RULES_HUGO_COMMIT,
    url = "https://github.com/rrmcguinness/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
)

load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_repository")

#
# Load hugo binary itself
#
# Optionally, load a specific version of Hugo, with the 'version' argument
hugo_repository(
    name = "hugo",
    extended = True,
    version = "0.101.0",
)

# Create a readable archive from a GitHub Hugo Theme that DOES NOT support the theme layout.
http_archive(
    name = "theme_geekdoc",
    build_file_content = """
filegroup(
    name = "files",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
    """,
    sha256 = "7fdd57f7d4450325a778629021c0fff5531dc8475de6c4ec70ab07e9484d400e",
    url = "https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.34.2/hugo-geekdoc.tar.gz",
)
