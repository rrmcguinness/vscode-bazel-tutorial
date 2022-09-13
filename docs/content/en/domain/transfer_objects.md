---
title: Transfer Object
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
{{< mermaid class="text-center">}}
classDiagram


class SimpleResponse {
  +TimeStamp Ts
  +int ResponseCode
  +string ErrorMessage
}

class SearchType {
<< enumeration >>
  CONTAINS
  BEGINS_WITH
  ENDS_WITH
}

class OneOf {
  << interface >>
}

CustomerSearchPredicate ..|> OneOf

class CustomerSearchPredicate {
  +string id
  +string name
}

class CustomerSearch {
  +SearchType type
  +CustomerSearchPredicate search
}

CustomerSearch --|> SearchType
CustomerSearch --* CustomerSearchPredicate

ItemSearchPredicate ..|> OneOf
class ItemSearchPredicate {
  +string sku
  +string name
}


class ItemSearch {
  +SearchType type
  +ItemSearchPredicate search
}
ItemSearch --* ItemSearchPredicate

class IdRequest {
  +string Id
}

class TransactionRequest {
  +string TransactionId;
}


class TransactionResponse {
  +string TransactionId
  +Transaction State
}
TransactionResponse --> Transaction


class TransactionCustomerRequest {
  +string TransactionId
  +Customer Customer
}
TransactionCustomerRequest --> Customer

class TransactionLineItemRequest {
  +string TransactionId
  +LineItem LineItem
}
TransactionLineItemRequest --> LineItem
{{< /mermaid >}}

