// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: api/proto/wso2/discovery/api/api.proto

package api

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

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServiceClient interface {
	CreateAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error)
	UpdateAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error)
	DeleteAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) CreateAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/wso2.discovery.api.APIService/createAPI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) UpdateAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/wso2.discovery.api.APIService/updateAPI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) DeleteAPI(ctx context.Context, in *API, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/wso2.discovery.api.APIService/deleteAPI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
// All implementations must embed UnimplementedAPIServiceServer
// for forward compatibility
type APIServiceServer interface {
	CreateAPI(context.Context, *API) (*Response, error)
	UpdateAPI(context.Context, *API) (*Response, error)
	DeleteAPI(context.Context, *API) (*Response, error)
	mustEmbedUnimplementedAPIServiceServer()
}

// UnimplementedAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (UnimplementedAPIServiceServer) CreateAPI(context.Context, *API) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPI not implemented")
}
func (UnimplementedAPIServiceServer) UpdateAPI(context.Context, *API) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPI not implemented")
}
func (UnimplementedAPIServiceServer) DeleteAPI(context.Context, *API) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPI not implemented")
}
func (UnimplementedAPIServiceServer) mustEmbedUnimplementedAPIServiceServer() {}

// UnsafeAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServiceServer will
// result in compilation errors.
type UnsafeAPIServiceServer interface {
	mustEmbedUnimplementedAPIServiceServer()
}

func RegisterAPIServiceServer(s grpc.ServiceRegistrar, srv APIServiceServer) {
	s.RegisterService(&APIService_ServiceDesc, srv)
}

func _APIService_CreateAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(API)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).CreateAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wso2.discovery.api.APIService/createAPI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).CreateAPI(ctx, req.(*API))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_UpdateAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(API)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).UpdateAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wso2.discovery.api.APIService/updateAPI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).UpdateAPI(ctx, req.(*API))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_DeleteAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(API)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).DeleteAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wso2.discovery.api.APIService/deleteAPI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).DeleteAPI(ctx, req.(*API))
	}
	return interceptor(ctx, in, info, handler)
}

// APIService_ServiceDesc is the grpc.ServiceDesc for APIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wso2.discovery.api.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAPI",
			Handler:    _APIService_CreateAPI_Handler,
		},
		{
			MethodName: "updateAPI",
			Handler:    _APIService_UpdateAPI_Handler,
		},
		{
			MethodName: "deleteAPI",
			Handler:    _APIService_DeleteAPI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/wso2/discovery/api/api.proto",
}
