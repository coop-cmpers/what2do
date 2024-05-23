package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
)

func (s *What2doService) SearchRecommendations(ctx context.Context, req *pb.SearchRecommendationsRequest) (*pb.SearchRecommendationsResponse, error) {
	recommendation, err := s.client.FetchPlaces(ctx, req.SearchParam, req.Location, req.EventTime)
	if err != nil {
		return nil, err
	}
	return &pb.SearchRecommendationsResponse{Recommendations: recommendation}, err
}
