package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
)

type HelloWorldServer struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *HelloWorldServer) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hi, My name is Slim Shady!"}, nil
}