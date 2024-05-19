package main

import (
	"context"
	"log"
	"net"

	"github.com/coop-cmpers/what2do-backend/src/endpoints"
	"github.com/coop-cmpers/what2do-backend/src/helpers"
	"github.com/coop-cmpers/what2do-backend/src/store"
	"go.uber.org/zap"

	helloworldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

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
		log.Fatalf("Failed to initialise databsae connection")
	}

	// Initialise Zap logger
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialise Zap logger")
	}
	defer zapLogger.Sync()

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(helpers.ZapCodeToLevel),
	}
	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)

	// Initialise the gRPC server
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			helpers.DBStreamServerInterceptor(store),
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger, zapOpts...),
			grpc_recovery.StreamServerInterceptor(), // TODO: implement graceful panic recovery
		),
		grpc.ChainUnaryInterceptor(
			helpers.DBUnaryServerInterceptor(store),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, zapOpts...),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	// Register the gRPC services
	helloworldpb.RegisterHelloWorldServiceServer(s, &endpoints.HelloWorldServer{})
	what2dopb.RegisterWhat2DoServiceServer(s, &endpoints.What2doServer{})

	// Listen and serve on port
	env, err := helpers.GetEnvFromCtx(ctx)
	if err != nil {
		log.Fatalf("Failed to find environment variables")
	}

	port := env["BACKEND_PORT"]
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s - err: %v", port, err)
	}

	log.Printf("gRPC server listening at %v", listen.Addr())

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve - err: %v", err)
	}
}
