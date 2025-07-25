// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProductService_CreateProduct_FullMethodName     = "/product.ProductService/CreateProduct"
	ProductService_FindById_FullMethodName          = "/product.ProductService/FindById"
	ProductService_FindMany_FullMethodName          = "/product.ProductService/FindMany"
	ProductService_ListStream_FullMethodName        = "/product.ProductService/ListStream"
	ProductService_List_FullMethodName              = "/product.ProductService/List"
	ProductService_UpdateProduct_FullMethodName     = "/product.ProductService/UpdateProduct"
	ProductService_ReserveInventory_FullMethodName  = "/product.ProductService/ReserveInventory"
	ProductService_ReleaseInventory_FullMethodName  = "/product.ProductService/ReleaseInventory"
	ProductService_ActivateProduct_FullMethodName   = "/product.ProductService/ActivateProduct"
	ProductService_DeleteProduct_FullMethodName     = "/product.ProductService/DeleteProduct"
	ProductService_UpdateStock_FullMethodName       = "/product.ProductService/UpdateStock"
	ProductService_CreateCategory_FullMethodName    = "/product.ProductService/CreateCategory"
	ProductService_FindAllCategories_FullMethodName = "/product.ProductService/FindAllCategories"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Product
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error)
	FindMany(ctx context.Context, in *FindManyRequest, opts ...grpc.CallOption) (*FindManyResponse, error)
	ListStream(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProductData], error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReserveInventory(ctx context.Context, in *ReserveInventoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReleaseInventory(ctx context.Context, in *ReleaseInventoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ActivateProduct(ctx context.Context, in *ActivateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Category
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindAllCategories(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindAllCategoriesResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindByIdResponse)
	err := c.cc.Invoke(ctx, ProductService_FindById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindMany(ctx context.Context, in *FindManyRequest, opts ...grpc.CallOption) (*FindManyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindManyResponse)
	err := c.cc.Invoke(ctx, ProductService_FindMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListStream(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProductData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], ProductService_ListStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListRequest, ProductData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_ListStreamClient = grpc.ServerStreamingClient[ProductData]

func (c *productServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ProductService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ReserveInventory(ctx context.Context, in *ReserveInventoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_ReserveInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ReleaseInventory(ctx context.Context, in *ReleaseInventoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_ReleaseInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ActivateProduct(ctx context.Context, in *ActivateProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_ActivateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_UpdateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProductService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindAllCategories(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindAllCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindAllCategoriesResponse)
	err := c.cc.Invoke(ctx, ProductService_FindAllCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility.
type ProductServiceServer interface {
	// Product
	CreateProduct(context.Context, *CreateProductRequest) (*emptypb.Empty, error)
	FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error)
	FindMany(context.Context, *FindManyRequest) (*FindManyResponse, error)
	ListStream(*ListRequest, grpc.ServerStreamingServer[ProductData]) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*emptypb.Empty, error)
	ReserveInventory(context.Context, *ReserveInventoryRequest) (*emptypb.Empty, error)
	ReleaseInventory(context.Context, *ReleaseInventoryRequest) (*emptypb.Empty, error)
	ActivateProduct(context.Context, *ActivateProductRequest) (*emptypb.Empty, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*emptypb.Empty, error)
	// Category
	CreateCategory(context.Context, *CreateCategoryRequest) (*emptypb.Empty, error)
	FindAllCategories(context.Context, *emptypb.Empty) (*FindAllCategoriesResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedProductServiceServer) FindMany(context.Context, *FindManyRequest) (*FindManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMany not implemented")
}
func (UnimplementedProductServiceServer) ListStream(*ListRequest, grpc.ServerStreamingServer[ProductData]) error {
	return status.Errorf(codes.Unimplemented, "method ListStream not implemented")
}
func (UnimplementedProductServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) ReserveInventory(context.Context, *ReserveInventoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveInventory not implemented")
}
func (UnimplementedProductServiceServer) ReleaseInventory(context.Context, *ReleaseInventoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseInventory not implemented")
}
func (UnimplementedProductServiceServer) ActivateProduct(context.Context, *ActivateProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateStock(context.Context, *UpdateStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedProductServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedProductServiceServer) FindAllCategories(context.Context, *emptypb.Empty) (*FindAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllCategories not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}
func (UnimplementedProductServiceServer) testEmbeddedByValue()                        {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: ProductService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindMany(ctx, req.(*FindManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).ListStream(m, &grpc.GenericServerStream[ListRequest, ProductData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProductService_ListStreamServer = grpc.ServerStreamingServer[ProductData]

func _ProductService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ReserveInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ReserveInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ReserveInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ReserveInventory(ctx, req.(*ReserveInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ReleaseInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ReleaseInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ReleaseInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ReleaseInventory(ctx, req.(*ReleaseInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ActivateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ActivateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ActivateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ActivateProduct(ctx, req.(*ActivateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindAllCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindAllCategories(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _ProductService_FindById_Handler,
		},
		{
			MethodName: "FindMany",
			Handler:    _ProductService_FindMany_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProductService_List_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "ReserveInventory",
			Handler:    _ProductService_ReserveInventory_Handler,
		},
		{
			MethodName: "ReleaseInventory",
			Handler:    _ProductService_ReleaseInventory_Handler,
		},
		{
			MethodName: "ActivateProduct",
			Handler:    _ProductService_ActivateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _ProductService_UpdateStock_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _ProductService_CreateCategory_Handler,
		},
		{
			MethodName: "FindAllCategories",
			Handler:    _ProductService_FindAllCategories_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStream",
			Handler:       _ProductService_ListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
