package main

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
  "flag"
  "log"
  client "example/client"
  grpc "google.golang.org/grpc"
  model "example/model"
  service "example/service"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
  flag.Parse()

  var opts []grpc.DialOption

  //opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
  opts = append(opts, grpc.WithInsecure())

  conn, err := grpc.Dial(*serverAddr, opts...)
  if err != nil {
    log.Fatalf("Failed to dial server: %v", err)
  }

  defer conn.Close()
  svc_client := service.NewStoresClient(conn)

  // Create a couple of stores
  client.CreateStore(svc_client, model.NewStore("Test Store 1", "30101"))
  client.CreateStore(svc_client, model.NewStore("Test Store 2", "30102"))

  client.ListStores(svc_client)
}
