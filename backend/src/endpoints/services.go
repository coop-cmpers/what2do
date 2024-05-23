package endpoints

import (
	"context"

	helloworldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"github.com/coop-cmpers/what2do-backend/src/clients"
	"github.com/coop-cmpers/what2do-backend/src/store"
	"go.uber.org/zap"
)

type HelloWorldService struct {
	helloworldpb.UnimplementedHelloWorldServiceServer

	store  *store.Store
	logger *zap.SugaredLogger
}

type What2doService struct {
	what2dopb.UnimplementedWhat2DoServiceServer

	store  *store.Store
	client *clients.Client
	logger *zap.SugaredLogger
}

func NewHelloWorldService(ctx context.Context, store *store.Store, logger *zap.Logger) *HelloWorldService {
	return &HelloWorldService{
		store:  store,
		logger: logger.Sugar(),
	}
}

func NewWhat2doService(ctx context.Context, store *store.Store, client *clients.Client, logger *zap.Logger) *What2doService {
	return &What2doService{
		store:  store,
		client: client,
		logger: logger.Sugar(),
	}
}
