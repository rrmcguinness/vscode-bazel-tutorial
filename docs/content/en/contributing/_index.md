---
title: "Contributing"
weight: 1

resources:
  - name: sdlc
    src: git-process.png
    title: "SDLC"
---

<!---
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
--->

## Understanding the Layout

- api - This is where the protocol buffer files are located.
- conf - These are configuration files used by Bazel.
- docs - a Hugo website (WIP).
- golang - The Go Project implementing the generated APIs.
- java - The Java Project implementing the generated APIS.
- third_party - A best practice directory for storing third_party assets.
- tools - Helper scripts for VSCode to utilize Go with Bazel.

### Important Files

- vscode-bazel-tutorial.code-workspace - responsible for configuring vscode.
- WORKSPACE - responsible for dependency management in Bazel.
- api/BUILD - the build file for the API model and GRPC.
- golang/BUILD - the Go build file with dependencies on the API.
- java/BUILD - the Java build file with dependencies on the API.

## SDLC

{{< img name="sdlc" size="large" lazy=false >}}

### Steps

1. Ensure you have a [GitHub](https://www.github.com) account.
1. Request access if the repository is private.
1. Fork the repository. This makes a copy to your local GitHub account.
1. Clone the newly created fork to your developer machine.
   - `git clone <repository name>`
1. Make any changes or additions to the code.
1. Write automated test cases.
1. Verify all tests are passing and all code is commented and meets the style
   guide requirements.
   _ `bazel test //...` runs all tests
   _ `bazel coverage //java/...` runs test coverage on a specific target.
1. Commit your code, and push back to your fork.
   - `git commit add .`
   - `git commit -a`
   - Add comments, if you are referencing a feature or bug, please indicate the number first.
   - `git push`
1. Verify the build runs in GitHub actions on your own branch. This is important, especially if
   you are updating BUILD or WORKSPACE files.
1. From the GitHub interface, create a pull request and ensure you comment the
   reasons and thought processes behind the change or additions. \* [Understanding Pull Requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests)
1. Once created, you will collaborate with the mantainers for QA purposes,
   unless they simply accept the changed and merge it.
1. If not, you may be asked to make some changes and create a new PR.

## Licensing

All content on this site is licened under the Apache 2.0 license. Please see the LICENSE file in the source code. If you contribute to the project, please add your name to the contributors file and link to your GitHub profile.
