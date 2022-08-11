// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protocol/protocol.proto

package protocol

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

// ProductCatalogClient is the client API for ProductCatalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogClient interface {
	FetchProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
}

type productCatalogClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogClient(cc grpc.ClientConnInterface) ProductCatalogClient {
	return &productCatalogClient{cc}
}

func (c *productCatalogClient) FetchProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/server.ProductCatalog/FetchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServer is the server API for ProductCatalog service.
// All implementations must embed UnimplementedProductCatalogServer
// for forward compatibility
type ProductCatalogServer interface {
	FetchProducts(context.Context, *ProductsRequest) (*ProductsResponse, error)
	mustEmbedUnimplementedProductCatalogServer()
}

// UnimplementedProductCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServer struct {
}

func (UnimplementedProductCatalogServer) FetchProducts(context.Context, *ProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchProducts not implemented")
}
func (UnimplementedProductCatalogServer) mustEmbedUnimplementedProductCatalogServer() {}

// UnsafeProductCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServer will
// result in compilation errors.
type UnsafeProductCatalogServer interface {
	mustEmbedUnimplementedProductCatalogServer()
}

func RegisterProductCatalogServer(s grpc.ServiceRegistrar, srv ProductCatalogServer) {
	s.RegisterService(&ProductCatalog_ServiceDesc, srv)
}

func _ProductCatalog_FetchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServer).FetchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ProductCatalog/FetchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServer).FetchProducts(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalog_ServiceDesc is the grpc.ServiceDesc for ProductCatalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.ProductCatalog",
	HandlerType: (*ProductCatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchProducts",
			Handler:    _ProductCatalog_FetchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/protocol.proto",
}
