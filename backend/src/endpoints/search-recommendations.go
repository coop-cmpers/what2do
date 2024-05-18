package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
)

type What2doServer struct {
	pb.UnimplementedWhat2DoServiceServer
}

func (s *What2doServer) SearchRecommendations(ctx context.Context, req *pb.SearchRecommendationsRequest) (*pb.SearchRecommendationsResponse, error) {
	return &pb.SearchRecommendationsResponse{
		Recommendation: []*pb.Recommendation{
			{
				Rank: 1, 
				Address: "AA", 
				Description: "BB",
			},
		},
	}, nil
}