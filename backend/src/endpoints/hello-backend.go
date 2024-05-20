package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
)

func (s *HelloWorldService) HelloBackend(ctx context.Context, req *pb.HelloBackendRequest) (*pb.HelloBackendResponse, error) {
	helloMessage := "Hi, My name is Slim Shady!"

	return &pb.HelloBackendResponse{Message: helloMessage}, nil
}
