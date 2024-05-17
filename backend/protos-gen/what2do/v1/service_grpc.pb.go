// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: what2do/v1/service.proto

package what2do_protobufs_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	What2DoService_SearchRecommendations_FullMethodName = "/what2do.v1.What2DoService/SearchRecommendations"
)

// What2DoServiceClient is the client API for What2DoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type What2DoServiceClient interface {
	SearchRecommendations(ctx context.Context, in *SearchRecommendationsRequest, opts ...grpc.CallOption) (*SearchRecommendationsResponse, error)
}

type what2DoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWhat2DoServiceClient(cc grpc.ClientConnInterface) What2DoServiceClient {
	return &what2DoServiceClient{cc}
}

func (c *what2DoServiceClient) SearchRecommendations(ctx context.Context, in *SearchRecommendationsRequest, opts ...grpc.CallOption) (*SearchRecommendationsResponse, error) {
	out := new(SearchRecommendationsResponse)
	err := c.cc.Invoke(ctx, What2DoService_SearchRecommendations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// What2DoServiceServer is the server API for What2DoService service.
// All implementations must embed UnimplementedWhat2DoServiceServer
// for forward compatibility
type What2DoServiceServer interface {
	SearchRecommendations(context.Context, *SearchRecommendationsRequest) (*SearchRecommendationsResponse, error)
	mustEmbedUnimplementedWhat2DoServiceServer()
}

// UnimplementedWhat2DoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWhat2DoServiceServer struct {
}

func (UnimplementedWhat2DoServiceServer) SearchRecommendations(context.Context, *SearchRecommendationsRequest) (*SearchRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRecommendations not implemented")
}
func (UnimplementedWhat2DoServiceServer) mustEmbedUnimplementedWhat2DoServiceServer() {}

// UnsafeWhat2DoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to What2DoServiceServer will
// result in compilation errors.
type UnsafeWhat2DoServiceServer interface {
	mustEmbedUnimplementedWhat2DoServiceServer()
}

func RegisterWhat2DoServiceServer(s grpc.ServiceRegistrar, srv What2DoServiceServer) {
	s.RegisterService(&What2DoService_ServiceDesc, srv)
}

func _What2DoService_SearchRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(What2DoServiceServer).SearchRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: What2DoService_SearchRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(What2DoServiceServer).SearchRecommendations(ctx, req.(*SearchRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// What2DoService_ServiceDesc is the grpc.ServiceDesc for What2DoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var What2DoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "what2do.v1.What2DoService",
	HandlerType: (*What2DoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchRecommendations",
			Handler:    _What2DoService_SearchRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "what2do/v1/service.proto",
}
