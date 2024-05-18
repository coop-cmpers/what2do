package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
)

func (s *HelloWorldServer) HelloDB(ctx context.Context, req *pb.HelloDBRequest) (*pb.HelloDBResponse, error) {
	store, err := GetStore(ctx)
	if err != nil {
		return nil, err
	}

	testDBObject, err := store.HelloDB(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	resp := &pb.HelloDBResponse{
		TestDbObject: testDBObject,
	}

	return resp, nil
}
