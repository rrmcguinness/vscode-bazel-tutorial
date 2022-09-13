---
title: Model
weight: 1
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

## Model Objects

{{< mermaid class="text-center">}}
classDiagram
  class Store{
    +string Id
    +string Name
    +string Zipcode
    +Map Meta
  }

  class Customer{
    +string Id
    +string Name
    +string Zipcode
  }

  class Item{
    +string Sku
    +string Name
    +TimeStamp EffectiveDate
    +float Price
    +bool Deleted
  }

  class LineItem{
    +int Ordinal
    +int Quantity
    +float item_total
    +float item_tax
    +float item_discounts
    +float line_total
    +Item item
    +Map~string,string~ discount_details
    +addItem(int qty, Item item)
    +removeItem(int qty, Item item)
  }
  LineItem --o Item

  class Transaction {
    +string Id
    +TimeStamp StartTime
    +TimeStamp EndTime
    +Store Store
    +Customer Customer
    +Set~LineItem~ LineItems
    +addLineItem(int qty, LineItem lineItem)
    +removeLineItem(int qty, LineItem lineItem)
  }
  Transaction --|> Store
  Transaction --|> Customer
  Transaction --> "1..*" LineItem

{{< /mermaid >}}

