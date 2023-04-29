// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: warehouse.proto

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

const (
	WarehouseService_CreateWarehouse_FullMethodName = "/warehouse.WarehouseService/CreateWarehouse"
	WarehouseService_UpdateWarehouse_FullMethodName = "/warehouse.WarehouseService/UpdateWarehouse"
	WarehouseService_DeleteWarehouse_FullMethodName = "/warehouse.WarehouseService/DeleteWarehouse"
)

// WarehouseServiceClient is the client API for WarehouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseServiceClient interface {
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error)
	UpdateWarehouse(ctx context.Context, in *UpdateWarehouseRequest, opts ...grpc.CallOption) (*UpdateWarehouseResponse, error)
	DeleteWarehouse(ctx context.Context, in *DeleteWarehouseRequest, opts ...grpc.CallOption) (*DeleteWarehouseResponse, error)
}

type warehouseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseServiceClient(cc grpc.ClientConnInterface) WarehouseServiceClient {
	return &warehouseServiceClient{cc}
}

func (c *warehouseServiceClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error) {
	out := new(CreateWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateWarehouse(ctx context.Context, in *UpdateWarehouseRequest, opts ...grpc.CallOption) (*UpdateWarehouseResponse, error) {
	out := new(UpdateWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) DeleteWarehouse(ctx context.Context, in *DeleteWarehouseRequest, opts ...grpc.CallOption) (*DeleteWarehouseResponse, error) {
	out := new(DeleteWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_DeleteWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServiceServer is the server API for WarehouseService service.
// All implementations must embed UnimplementedWarehouseServiceServer
// for forward compatibility
type WarehouseServiceServer interface {
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	UpdateWarehouse(context.Context, *UpdateWarehouseRequest) (*UpdateWarehouseResponse, error)
	DeleteWarehouse(context.Context, *DeleteWarehouseRequest) (*DeleteWarehouseResponse, error)
	mustEmbedUnimplementedWarehouseServiceServer()
}

// UnimplementedWarehouseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseServiceServer struct {
}

func (UnimplementedWarehouseServiceServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateWarehouse(context.Context, *UpdateWarehouseRequest) (*UpdateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWarehouse not implemented")
}
func (UnimplementedWarehouseServiceServer) DeleteWarehouse(context.Context, *DeleteWarehouseRequest) (*DeleteWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWarehouse not implemented")
}
func (UnimplementedWarehouseServiceServer) mustEmbedUnimplementedWarehouseServiceServer() {}

// UnsafeWarehouseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseServiceServer will
// result in compilation errors.
type UnsafeWarehouseServiceServer interface {
	mustEmbedUnimplementedWarehouseServiceServer()
}

func RegisterWarehouseServiceServer(s grpc.ServiceRegistrar, srv WarehouseServiceServer) {
	s.RegisterService(&WarehouseService_ServiceDesc, srv)
}

func _WarehouseService_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateWarehouse(ctx, req.(*UpdateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_DeleteWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).DeleteWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_DeleteWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).DeleteWarehouse(ctx, req.(*DeleteWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WarehouseService_ServiceDesc is the grpc.ServiceDesc for WarehouseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarehouseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.WarehouseService",
	HandlerType: (*WarehouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWarehouse",
			Handler:    _WarehouseService_CreateWarehouse_Handler,
		},
		{
			MethodName: "UpdateWarehouse",
			Handler:    _WarehouseService_UpdateWarehouse_Handler,
		},
		{
			MethodName: "DeleteWarehouse",
			Handler:    _WarehouseService_DeleteWarehouse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}
