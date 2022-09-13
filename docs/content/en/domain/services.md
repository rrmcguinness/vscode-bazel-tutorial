---
title: Services
weight: 3
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

## Service Objects

{{< mermaid class="text-center" >}}
classDiagram
  class Stores {
    << interface >>
    +List(Empty) Stream~Store~
    +Create(Store) Store
    +Update(Store) Store
    +Delete(IdRequest) SimpleResponse
  }
  Stores --> Empty
  Stores --> Store
  Stores --> SimpleResponse
  Stores --> Stream

  class Customers {
    << interface >>
    +Find(CustomerSearch) Stream~Customer~
    +Create(Customer) Customer
    +Update(Customer) Customer
    +Delete(IdRequest) SimpleResponse
  }
  Customer --> CustomerSearch
  Customer --> Customer
  Customer --> IdRequest
  Customer --> SimpleResponse
  Customer --> Stream

  class Items {
    << interface >>
    +Find(ItemSearch) Stream~Item~
    +Get(IdRequest) Item
    +Create(Item) Item
    +Update(Item) Item
    +Delete(IdRequest) SimpleResponse
    +Restore(IdRequest) SimpleResponse
  }
  Items --> IteamSearch
  Items --> IdRequest
  Items --> Item
  Items --> SimpleResponse
  Items --> Streams

  class Transactions {
    << interface >>
    +Create(Store) TransactionResponse
    +AddCustomer(TransactionCustomerRequest) TransactionResponse
    +AddItem(TransactionItemRequest) TransactionResponse
    +RemoveItem(TransactionItemRequest) TransactionResponse
    +PrepareForTender(TransactionRequest) TransactionResponse
  }
  Transactions --> Store
  Transactions --> TranactionCustomerRequest
  Transactions --> TransactionItemRequest
  Transactions --> TransactionItemRequest
  Transactions --> TransactionRequest
  Transactions --> TransactionResponse
{{< /mermaid >}}
