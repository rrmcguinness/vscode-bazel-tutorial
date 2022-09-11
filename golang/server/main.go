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
package main

import (
  "flag"
  "fmt"
  "log"
  service "example/service"
  "net"

  "google.golang.org/grpc"
)

var(
  tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
  certFile   = flag.String("cert_file", "", "The TLS cert file")
  keyFile    = flag.String("key_file", "", "The TLS key file")
  jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
  port       = flag.Int("port", 50051, "The server port")
)
func main() {
  flag.Parse()

  fmt.Println("Starting Server...")

  host_info := fmt.Sprintf("localhost:%d", *port)

  lis, err := net.Listen("tcp", host_info)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  var opts []grpc.ServerOption

  //if *tls {
	//	if *certFile == "" {
	//		*certFile = data.Path("x509/server_cert.pem")
	//	}
	//	if *keyFile == "" {
	//		*keyFile = data.Path("x509/server_key.pem")
	//	}
	//	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//	if err != nil {
	//		log.Fatalf("Failed to generate credentials %v", err)
	//	}
	//	opts = []grpc.ServerOption{grpc.Creds(creds)}
	//}

  grpcServer := grpc.NewServer(opts...)
  fmt.Println("Created new Server")

	service.RegisterStoresServer(grpcServer, service.NewStoreServer())
  fmt.Println("Registered Stores Service")
  fmt.Printf("Running on: %s \n\n", host_info)
	grpcServer.Serve(lis)
}

