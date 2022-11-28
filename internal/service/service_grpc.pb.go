// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package service

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

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCServiceClient interface {
	// BlockIP creates an entry for a new blocked ip.
	BlockIP(ctx context.Context, in *BlockIPRequest, opts ...grpc.CallOption) (*BlockIPResponse, error)
	// IsBlockedIP determines if ip is currently blocked from interacting with Twitch resources.
	IsBlockedIP(ctx context.Context, in *IsBlockedIPRequest, opts ...grpc.CallOption) (*IsBlockedIPResponse, error)
	// UnblockIP deletes entry for ip currently being blocked.
	UnblockIP(ctx context.Context, in *UnblockIPRequest, opts ...grpc.CallOption) (*UnblockIPResponse, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) BlockIP(ctx context.Context, in *BlockIPRequest, opts ...grpc.CallOption) (*BlockIPResponse, error) {
	out := new(BlockIPResponse)
	err := c.cc.Invoke(ctx, "/service.RPCService/BlockIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) IsBlockedIP(ctx context.Context, in *IsBlockedIPRequest, opts ...grpc.CallOption) (*IsBlockedIPResponse, error) {
	out := new(IsBlockedIPResponse)
	err := c.cc.Invoke(ctx, "/service.RPCService/IsBlockedIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) UnblockIP(ctx context.Context, in *UnblockIPRequest, opts ...grpc.CallOption) (*UnblockIPResponse, error) {
	out := new(UnblockIPResponse)
	err := c.cc.Invoke(ctx, "/service.RPCService/UnblockIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
// All implementations must embed UnimplementedRPCServiceServer
// for forward compatibility
type RPCServiceServer interface {
	// BlockIP creates an entry for a new blocked ip.
	BlockIP(context.Context, *BlockIPRequest) (*BlockIPResponse, error)
	// IsBlockedIP determines if ip is currently blocked from interacting with Twitch resources.
	IsBlockedIP(context.Context, *IsBlockedIPRequest) (*IsBlockedIPResponse, error)
	// UnblockIP deletes entry for ip currently being blocked.
	UnblockIP(context.Context, *UnblockIPRequest) (*UnblockIPResponse, error)
	mustEmbedUnimplementedRPCServiceServer()
}

// UnimplementedRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (UnimplementedRPCServiceServer) BlockIP(context.Context, *BlockIPRequest) (*BlockIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockIP not implemented")
}
func (UnimplementedRPCServiceServer) IsBlockedIP(context.Context, *IsBlockedIPRequest) (*IsBlockedIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlockedIP not implemented")
}
func (UnimplementedRPCServiceServer) UnblockIP(context.Context, *UnblockIPRequest) (*UnblockIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockIP not implemented")
}
func (UnimplementedRPCServiceServer) mustEmbedUnimplementedRPCServiceServer() {}

// UnsafeRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServiceServer will
// result in compilation errors.
type UnsafeRPCServiceServer interface {
	mustEmbedUnimplementedRPCServiceServer()
}

func RegisterRPCServiceServer(s grpc.ServiceRegistrar, srv RPCServiceServer) {
	s.RegisterService(&RPCService_ServiceDesc, srv)
}

func _RPCService_BlockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).BlockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RPCService/BlockIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).BlockIP(ctx, req.(*BlockIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_IsBlockedIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBlockedIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).IsBlockedIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RPCService/IsBlockedIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).IsBlockedIP(ctx, req.(*IsBlockedIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_UnblockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnblockIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).UnblockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RPCService/UnblockIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).UnblockIP(ctx, req.(*UnblockIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCService_ServiceDesc is the grpc.ServiceDesc for RPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockIP",
			Handler:    _RPCService_BlockIP_Handler,
		},
		{
			MethodName: "IsBlockedIP",
			Handler:    _RPCService_IsBlockedIP_Handler,
		},
		{
			MethodName: "UnblockIP",
			Handler:    _RPCService_UnblockIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
