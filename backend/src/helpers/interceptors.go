package helpers

import (
	"context"

	"github.com/coop-cmpers/what2do-backend/src/constants"
	"github.com/coop-cmpers/what2do-backend/src/store"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// Injects the database connection into the context

func DBUnaryServerInterceptor(store *store.Store) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, constants.Store, store), req)
	}
}

func DBStreamServerInterceptor(store *store.Store) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = context.WithValue(stream.Context(), constants.Store, store)
		return handler(srv, wrapped)
	}
}
