package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/coop-cmpers/what2do-backend/src/endpoints"
	"github.com/coop-cmpers/what2do-backend/src/helpers"
	"github.com/coop-cmpers/what2do-backend/src/store"

	helloWorldpb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	what2dopb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"

	"google.golang.org/grpc"
)

var db *sql.DB

func main() {
	port := "12345"

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s - err: %v", port, err)
		return
	}

	ctx := context.Background()
	ctx = helpers.AddEnvToCtx(ctx)

	db, err = store.Connect(ctx)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	helloWorldpb.RegisterHelloWorldServiceServer(s, &endpoints.HelloWorldServer{})
	what2dopb.RegisterWhat2DoServiceServer(s, &endpoints.What2doServer{})

	log.Printf("gRPC server listening at %v", listen.Addr())

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve - err: %v", err)
		return
	}
}
