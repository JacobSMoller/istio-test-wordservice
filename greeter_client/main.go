/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "helloworld/helloworld"

	"github.com/e-conomic/ctxvml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "istio.vml.visma.ai:443"
)

var (
	host     = flag.String("host", "localhost:50051", "Specify host and port")
	insecure = flag.Bool("insecure", false, "Use insecure transport")
	name     = flag.String("name", "world", "name to greet")
)

func main() {
	flag.Parse()
	opt := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if *insecure {
		opt = grpc.WithInsecure()
	}

	conn, err := grpc.Dial(
		*host,
		opt,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(16*1024*1024),
		),
		grpc.WithUnaryInterceptor(
			ctxvml.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatalln("grpc Dial failed")
	}

	// Set up a connection to the server.
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
