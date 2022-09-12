package model
// Copyright 2022 Ryan McGuinness
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
  "github.com/google/uuid"
)

// Add some helper methods for stores.
func NewStore(name string, zipcode string) *Store {
  id, _ := uuid.NewRandom()

  return &Store{
    Id : id.String(),
    Name: name,
    Zipcode: zipcode,
    Meta: make(map[string]string),
  }
}

func (store *Store) addMeta(key string, value string) {
  store.Meta[key] = value
}
