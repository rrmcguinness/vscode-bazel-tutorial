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

# API

## TL;DR;

The API defines the Nouns and the Verbs for a domain model. These are defined via Protocol Buffers
to ensure a language agnostic approach. In turn this will allow us to generate the language specific implementations WITHOUT the overhead of managing multiple models.

## Build Targets

- `bazel build //api` Builds all protos
- `bazel build //api:model` build the model protos
- `bazel build //api:service` builds the service protos
- `bazel build //api:docs` Builds all documentation for the protos
- `bazel build //api:model-docs` Generates a markdown file from the comments in the model.proto file.
- `bazel build //api:service-docs` Generates a markdown file from the comments in the service.proto file.

## Model (Nouns)

The model represents the 'things' and 'ideas' within a given domain. This specific domain is a simplified retail data model.

- Store - A physical location.
- Customer - A consumer of store goods.
- Item - A thing to be sold at a store.
- Line Item - An enriched Item with transaction specific information.
- Transaction - A simple state machine relating the Store, Customer, and Line items. Akin to a receipt.

## Service (Verbs)

The service definitions are the Verbs of the system, how can service consumers appropriatly interact with the nouns. These are CRUD (Create, Read, Update, Delete) and business operation services.

- Stores - CRUD operations for Store Objects.
- Customers - CRUD operations for Customer Objects.
- Items - CRUD operations for Item Objects.
- Transactions - Business and Crud operations for managing transactions.
