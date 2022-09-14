---
title: "Setting up the Workspace"
weight: 2
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

## Introduction

In your home directory, create a folder called "projects". This only matters due to some limitations in the workspace configuration file. Vscode extensions for Java and Go WILL NOT support variable expressions for paths.

{{< hint type=note icon=gdoc_github title=Limitations >}}
This project is currently setup in "/opt/projects". Why /opt? Because writing a turorial
to create a '/data/projects' directory on MacOS seems to be a hit-and-miss since / is a read-only
mount, and /etc/synthetic.conf has other issues.
{{< /hint >}}

## Steps

```shell
# Create a projects directory, this example creates it on the /opt file system
# This IS NOT recommended, but is done so for the sake of a conistant development environment.
# Please feel free to put this any path you have access, but you need to remember for a couple of changes.
sudo mkdir /opt/projects

cd /opt/projects

# Clone the repository
git clone https://github.com/rrmcguinness/vscode-bazel-tutorial

# Run Bazel once before starting
bazel build //...

# Open Visual Studio Code

# If you changed the download location:
code .

# If you didn't change the download location:
code vscode-bazel-tutorial.code-workspace
```

### Edit the workspace (IF YOU CHANGED THE DOWNLOAD LOCATION)

Once Visual Studio Code has opened, edit the `vscode-bazel-tutorial.code-workspace` file.

Change all occurances of "/opt/projects" to your download directoy. DO NOT attempt to use variables or relative links, they ARE NOT honored by the extensions.

Lastly, in the `java.configuration.runtimes`, if you ARE NOT using Apple OS X, then please change `_macos` with `_linux` or `_windows` accordingly.

> Again, this is because a relative path, and command variable expansion IS NOT supported by the workpace file. Otherwise ${command:os_type} would work.

Once that change has been made, you may either "open workspace from file", or close and reopen from the terminal with the last command above.

### Start the Java Workspace

Navigate to java > src > service > CustomerService.java and open it.

> When you open this file, it will start the Java Language Server, this will in-turn read the workspace file to determine where to find Java, and which dependencies to add to the classpath. \*\* More about that below.
