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

package example.service;

import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import java.util.ArrayList;
import java.util.concurrent.atomic.AtomicLong;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import example.model.Item;
import example.model.ItemSearch;
import example.model.IdRequest;
import example.model.SimpleResponse;

public class ItemService extends ItemsGrpc.ItemsImplBase {

  @Override
  public void find(ItemSearch itemSearch, StreamObserver<Item> responseObserver) {

  }

  @Override
  public void get(IdRequest request, StreamObserver<Item> responseObserver) {

  }

  @Override
  public void create(Item item, StreamObserver<Item> responseObserver) {

  }

  @Override
  public void update(Item item, StreamObserver<Item> responseObserver) {

  }

  @Override
  public void delete(IdRequest request, StreamObserver<SimpleResponse> responseObserver) {

  }

  @Override
  public void restore(IdRequest request, StreamObserver<SimpleResponse> responseObserver) {

  }
}
