package service
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
  sync "sync"
  context "context"
  model "example/model"
  uuid "github.com/google/uuid"
  emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
  //codes "google.golang.org/grpc/codes"
	//status "google.golang.org/grpc/status"
	//protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	//protoimpl "google.golang.org/protobuf/runtime/protoimpl"
  //reflect "reflect"
)

type storeServer struct {
  UnimplementedStoresServer
  savedStores []*model.Store
  mu sync.Mutex
}

func NewStoreServer() *storeServer {
  return &storeServer{savedStores: make([]*model.Store,0)}
}

func (server *storeServer) List(value *emptypb.Empty, stream Stores_ListServer) (err error) {
  for _, s := range server.savedStores {
    err = stream.Send(s)
    if err != nil {
      break
    }
  }
  return err
}

func (server *storeServer) Create(ctx context.Context, store *model.Store) (*model.Store, error) {
  if len(store.Id) == 0 {
    id, _ := uuid.NewRandom()
    store.Id = id.String()
  }
  server.savedStores = append(server.savedStores, store)
  return store, nil
}

func (server *storeServer) Update(ctx context.Context, store *model.Store) (*model.Store, error) {
  server.Delete(ctx, &model.IdRequest{Id: store.Id})
  return server.Create(ctx, store)
}

func (server *storeServer) Delete(ctx context.Context, request *model.IdRequest) (*model.SimpleResponse, error) {
  new_array := make([]*model.Store, 0)
  for _, s := range server.savedStores {
    if s.Id != request.Id {
      new_array = append(new_array, s)
    }
  }

  return &model.SimpleResponse{ Ts: timestamppb.Now(), ResponseCode: 200, ErrorMessage: ""}, nil
}
