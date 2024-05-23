package clients

import (
	"context"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	*googlePlacesClient
}

type ClientConfig struct {
	Logger              *zap.Logger
	GooglePlacesAPIKey  string
	GooglePlacesBaseURL string
}

func NewClient(ctx context.Context, config *ClientConfig) *Client {
	return &Client{
		NewGooglePlacesClient(ctx, config.Logger, config.GooglePlacesAPIKey, config.GooglePlacesBaseURL),
	}
}

// Any new functions that should be exposed must be added to this interface
type GooglePlacesClient interface {
	FetchPlacesFromPlacesAPI(ctx context.Context, searchType string, location string, eventTime *timestamppb.Timestamp) ([]*pb.Recommendation, error)
}

type googlePlacesClient struct {
	logger  *zap.SugaredLogger
	apiKey  string
	baseURL string
}

func NewGooglePlacesClient(ctx context.Context, logger *zap.Logger, apiKey string, baseURL string) *googlePlacesClient {
	return &googlePlacesClient{
		logger:  logger.Sugar(),
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}
