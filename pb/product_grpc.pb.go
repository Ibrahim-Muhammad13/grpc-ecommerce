// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.1
// source: proto/product.proto

package pb

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	SearchProduct(ctx context.Context, in *SearchProductRequest, opts ...grpc.CallOption) (*SearchProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SearchProduct(ctx context.Context, in *SearchProductRequest, opts ...grpc.CallOption) (*SearchProductResponse, error) {
	out := new(SearchProductResponse)
	err := c.cc.Invoke(ctx, "/ProductService/SearchProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	SearchProduct(context.Context, *SearchProductRequest) (*SearchProductResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductServiceServer) SearchProduct(context.Context, *SearchProductRequest) (*SearchProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SearchProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SearchProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SearchProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SearchProduct(ctx, req.(*SearchProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "SearchProduct",
			Handler:    _ProductService_SearchProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/product.proto",
}