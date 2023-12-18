// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/comet/comet.proto

package comet

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

// CometClient is the client API for Comet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CometClient interface {
	PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error)
}

type cometClient struct {
	cc grpc.ClientConnInterface
}

func NewCometClient(cc grpc.ClientConnInterface) CometClient {
	return &cometClient{cc}
}

func (c *cometClient) PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error) {
	out := new(PushMessageResp)
	err := c.cc.Invoke(ctx, "/comet.Comet/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CometServer is the server API for Comet service.
// All implementations must embed UnimplementedCometServer
// for forward compatibility
type CometServer interface {
	PushMessage(context.Context, *PushMessageReq) (*PushMessageResp, error)
	mustEmbedUnimplementedCometServer()
}

// UnimplementedCometServer must be embedded to have forward compatible implementations.
type UnimplementedCometServer struct {
}

func (UnimplementedCometServer) PushMessage(context.Context, *PushMessageReq) (*PushMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}
func (UnimplementedCometServer) mustEmbedUnimplementedCometServer() {}

// UnsafeCometServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CometServer will
// result in compilation errors.
type UnsafeCometServer interface {
	mustEmbedUnimplementedCometServer()
}

func RegisterCometServer(s grpc.ServiceRegistrar, srv CometServer) {
	s.RegisterService(&Comet_ServiceDesc, srv)
}

func _Comet_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CometServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comet.Comet/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CometServer).PushMessage(ctx, req.(*PushMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comet_ServiceDesc is the grpc.ServiceDesc for Comet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comet.Comet",
	HandlerType: (*CometServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMessage",
			Handler:    _Comet_PushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/comet/comet.proto",
}
