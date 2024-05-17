package main

import (
	"context"
	"log"
	"net"

	"github.com/coop-cmpers/what2do-backend/src/endpoints"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hi, My name is Slim Shady!"}, nil
}

func main() {
	port := "12345"

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s, err: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	what2dopb.RegisterWhat2DoServiceServer(s, &endpoints.What2doServer{})

	log.Printf("gRPC server listening at %v", listen.Addr())

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve, err: %v", err)
	}
}
