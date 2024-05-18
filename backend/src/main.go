package main

import (
	"log"
	"net"

	"github.com/coop-cmpers/what2do-backend/src/endpoints"

	helloWorldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"

	"google.golang.org/grpc"
)

func main() {
	port := "12345"

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s, err: %v", port, err)
	}

	s := grpc.NewServer()
	helloWorldpb.RegisterHelloWorldServiceServer(s, &endpoints.HelloWorldServer{})
	what2dopb.RegisterWhat2DoServiceServer(s, &endpoints.What2doServer{})

	log.Printf("gRPC server listening at %v", listen.Addr())

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve, err: %v", err)
	}
}
