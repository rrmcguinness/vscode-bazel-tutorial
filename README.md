# Visual Studio Code + Bazel Tutorial

## Finding Generated Source Files

```shell
# Model
ls /${workspace}/bazel-out/darwin-fastbuild/bin/java/model/_javac/model_proto/libmodel_proto_tmp/example/model

# Services
ls /${workspace}/bazel-out/darwin-fastbuild/bin/java/service/_javac/service_proto/libserice_proto_tmp/example/service/*.java


```

## Classpath Entries

```shell

"${workspaceFolder}/bazel-bin/java/**/lib*.jar"
"${workspaceFolder}/bazel-bin/java/**/lib*_proto.jar"

"${workspaceFolder}/bazel-bin/java/model/libmodel_deps.jar",
"${workspaceFolder}/bazel-bin/java/model/libmodel.jar",
"${workspaceFolder}/bazel-bin/java/service/libservice_deps.jar",
"${workspaceFolder}/bazel-bin/java/service/libservice.jar",

```
