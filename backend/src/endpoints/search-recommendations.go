package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
)

func (s *What2doService) SearchRecommendations(ctx context.Context, req *pb.SearchRecommendationsRequest) (*pb.SearchRecommendationsResponse, error) {
	return &pb.SearchRecommendationsResponse{
		Recommendations: []*pb.Recommendation{
			{
				Rank:        1,
				Address:     "AA",
				Description: "BB",
			},
		},
	}, nil
}
