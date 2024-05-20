package endpoints

import (
	"context"

	helloworldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
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
	logger *zap.SugaredLogger
}

func NewHelloWorldService(ctx context.Context, store *store.Store, logger *zap.Logger) (*HelloWorldService, error) {
	helloWorldService := &HelloWorldService{}

	helloWorldService.store = store
	helloWorldService.logger = logger.Sugar()

	return helloWorldService, nil
}

func NewWhat2doService(ctx context.Context, store *store.Store, logger *zap.Logger) (*What2doService, error) {
	what2doService := &What2doService{}

	what2doService.store = store
	what2doService.logger = logger.Sugar()

	return what2doService, nil
}
