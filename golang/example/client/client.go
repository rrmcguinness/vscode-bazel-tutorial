package client

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
  context "context"
  "io"
  "log"
  "time"
  model "example/model"
  service "example/service"
  emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ListStores(client service.StoresClient) {
  log.Printf("Listing Stores")
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  empty := &emptypb.Empty{}
  stream, err := client.List(ctx, empty)
  if err != nil {
    log.Fatalf("Client failed to get list: %v", err)
  }
  for {
    store, err := stream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatalf("Client recieved an error on read %v", err)
    }
    log.Printf("Store %s - %s :: %s", store.Id, store.Name, store.Zipcode)
  }
}

func CreateStore(client service.StoresClient, store *model.Store) {
  log.Printf("Creating Store: %s", store.Name)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  out, err := client.Create(ctx, store)
  if err != nil {
    log.Fatalf("Failed to save store: %v", err)
  }
  log.Printf("Store %s - %s :: %s", out.Id, out.Name, out.Zipcode)
}


