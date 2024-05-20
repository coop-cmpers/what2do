package endpoints

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"github.com/coop-cmpers/what2do-backend/src/clients"
)


func (s *What2doService) SearchRecommendations(ctx context.Context, req *pb.SearchRecommendationsRequest) (*pb.SearchRecommendationsResponse, error) {
	recommendation, err := clients.FetchPlacesFromPlacesAPI(ctx, req.SearchParam, req.Location, req.EventTime)
	if err != nil {
		return nil, err
	}
	return &pb.SearchRecommendationsResponse{Recommendations: recommendation}, err
}
