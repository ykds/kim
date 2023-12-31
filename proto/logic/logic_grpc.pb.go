// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/logic/logic.proto

package logic

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

// LogicServiceClient is the client API for LogicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicServiceClient interface {
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
	HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatResp, error)
	DisConnect(ctx context.Context, in *DisConnectReq, opts ...grpc.CallOption) (*DisConnectResp, error)
}

type logicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicServiceClient(cc grpc.ClientConnInterface) LogicServiceClient {
	return &logicServiceClient{cc}
}

func (c *logicServiceClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/logic.LogicService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatResp, error) {
	out := new(HeartBeatResp)
	err := c.cc.Invoke(ctx, "/logic.LogicService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) DisConnect(ctx context.Context, in *DisConnectReq, opts ...grpc.CallOption) (*DisConnectResp, error) {
	out := new(DisConnectResp)
	err := c.cc.Invoke(ctx, "/logic.LogicService/DisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServiceServer is the server API for LogicService service.
// All implementations must embed UnimplementedLogicServiceServer
// for forward compatibility
type LogicServiceServer interface {
	Auth(context.Context, *AuthReq) (*AuthResp, error)
	HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatResp, error)
	DisConnect(context.Context, *DisConnectReq) (*DisConnectResp, error)
	mustEmbedUnimplementedLogicServiceServer()
}

// UnimplementedLogicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogicServiceServer struct {
}

func (UnimplementedLogicServiceServer) Auth(context.Context, *AuthReq) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedLogicServiceServer) HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedLogicServiceServer) DisConnect(context.Context, *DisConnectReq) (*DisConnectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisConnect not implemented")
}
func (UnimplementedLogicServiceServer) mustEmbedUnimplementedLogicServiceServer() {}

// UnsafeLogicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicServiceServer will
// result in compilation errors.
type UnsafeLogicServiceServer interface {
	mustEmbedUnimplementedLogicServiceServer()
}

func RegisterLogicServiceServer(s grpc.ServiceRegistrar, srv LogicServiceServer) {
	s.RegisterService(&LogicService_ServiceDesc, srv)
}

func _LogicService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.LogicService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.LogicService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).HeartBeat(ctx, req.(*HeartBeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_DisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).DisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.LogicService/DisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).DisConnect(ctx, req.(*DisConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LogicService_ServiceDesc is the grpc.ServiceDesc for LogicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logic.LogicService",
	HandlerType: (*LogicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _LogicService_Auth_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _LogicService_HeartBeat_Handler,
		},
		{
			MethodName: "DisConnect",
			Handler:    _LogicService_DisConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logic/logic.proto",
}
