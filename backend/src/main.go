package main

import (
	"context"
	"log"
	"net"

	"github.com/coop-cmpers/what2do-backend/src/endpoints"
	"github.com/coop-cmpers/what2do-backend/src/helpers"
	"github.com/coop-cmpers/what2do-backend/src/store"

	helloworldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"

	_ "github.com/lib/pq" // postgres driver
	"google.golang.org/grpc"
)

func main() {
	// Inject environment variables into the context
	ctx := context.Background()
	ctx = helpers.AddEnvToCtx(ctx)

	// Initialise connection to the database
	store := &store.Store{}
	_, err := store.Connect(ctx)
	if err != nil {
		return
	}

	// Initialise the gRPC server
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			helpers.DBStreamServerInterceptor(store),
		),
		grpc.ChainUnaryInterceptor(
			helpers.DBUnaryServerInterceptor(store),
		),
	)

	// Register the gRPC services
	helloworldpb.RegisterHelloWorldServiceServer(s, &endpoints.HelloWorldServer{})
	what2dopb.RegisterWhat2DoServiceServer(s, &endpoints.What2doServer{})

	// Listen and serve on port
	env, err := helpers.GetEnvFromCtx(ctx)
	if err != nil {
		return
	}

	port := env["BACKEND_PORT"]
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s - err: %v", port, err)
		return
	}

	log.Printf("gRPC server listening at %v", listen.Addr())

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve - err: %v", err)
		return
	}
}
