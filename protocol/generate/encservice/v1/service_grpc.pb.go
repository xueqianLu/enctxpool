// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: encservice/v1/service.proto

package encservicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EncServiceClient is the client API for EncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncServiceClient interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	AddTx(ctx context.Context, in *AddTxRequest, opts ...grpc.CallOption) (*AddTxResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	Pending(ctx context.Context, in *PendingRequest, opts ...grpc.CallOption) (*PendingResponse, error)
}

type encServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncServiceClient(cc grpc.ClientConnInterface) EncServiceClient {
	return &encServiceClient{cc}
}

func (c *encServiceClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/encservice.v1.EncService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encServiceClient) AddTx(ctx context.Context, in *AddTxRequest, opts ...grpc.CallOption) (*AddTxResponse, error) {
	out := new(AddTxResponse)
	err := c.cc.Invoke(ctx, "/encservice.v1.EncService/AddTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/encservice.v1.EncService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encServiceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/encservice.v1.EncService/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encServiceClient) Pending(ctx context.Context, in *PendingRequest, opts ...grpc.CallOption) (*PendingResponse, error) {
	out := new(PendingResponse)
	err := c.cc.Invoke(ctx, "/encservice.v1.EncService/Pending", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncServiceServer is the server API for EncService service.
// All implementations must embed UnimplementedEncServiceServer
// for forward compatibility
type EncServiceServer interface {
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	AddTx(context.Context, *AddTxRequest) (*AddTxResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	Pending(context.Context, *PendingRequest) (*PendingResponse, error)
	mustEmbedUnimplementedEncServiceServer()
}

// UnimplementedEncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncServiceServer struct {
}

func (UnimplementedEncServiceServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedEncServiceServer) AddTx(context.Context, *AddTxRequest) (*AddTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTx not implemented")
}
func (UnimplementedEncServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedEncServiceServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedEncServiceServer) Pending(context.Context, *PendingRequest) (*PendingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pending not implemented")
}
func (UnimplementedEncServiceServer) mustEmbedUnimplementedEncServiceServer() {}

// UnsafeEncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncServiceServer will
// result in compilation errors.
type UnsafeEncServiceServer interface {
	mustEmbedUnimplementedEncServiceServer()
}

func RegisterEncServiceServer(s grpc.ServiceRegistrar, srv EncServiceServer) {
	s.RegisterService(&EncService_ServiceDesc, srv)
}

func _EncService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encservice.v1.EncService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncServiceServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncService_AddTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncServiceServer).AddTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encservice.v1.EncService/AddTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncServiceServer).AddTx(ctx, req.(*AddTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encservice.v1.EncService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncService_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncServiceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encservice.v1.EncService/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncServiceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncService_Pending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncServiceServer).Pending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encservice.v1.EncService/Pending",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncServiceServer).Pending(ctx, req.(*PendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncService_ServiceDesc is the grpc.ServiceDesc for EncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "encservice.v1.EncService",
	HandlerType: (*EncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _EncService_HealthCheck_Handler,
		},
		{
			MethodName: "AddTx",
			Handler:    _EncService_AddTx_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _EncService_Status_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _EncService_Reset_Handler,
		},
		{
			MethodName: "Pending",
			Handler:    _EncService_Pending_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encservice/v1/service.proto",
}
