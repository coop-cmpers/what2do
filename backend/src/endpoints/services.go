package endpoints

import (
	"context"
	"errors"
	"log"

	helloworldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"github.com/coop-cmpers/what2do-backend/src/constants"
	"github.com/coop-cmpers/what2do-backend/src/store"
)

type HelloWorldServer struct {
	helloworldpb.UnimplementedHelloWorldServiceServer
}

type What2doServer struct {
	what2dopb.UnimplementedWhat2DoServiceServer
}

func GetStore(ctx context.Context) (*store.Store, error) {
	store, ok := ctx.Value(constants.Store).(*store.Store)

	if !ok || store == nil {
		log.Fatalf("Failed to get store from context")
		return nil, errors.New("failed to get store from context")
	}

	return store, nil
}
