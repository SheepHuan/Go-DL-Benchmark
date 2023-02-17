// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: remote_terminal.proto

package terminal

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

// RemoteTerminalServiceClient is the client API for RemoteTerminalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteTerminalServiceClient interface {
	NewTerminal(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error)
	CloseTerminal(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error)
	ExecCommand(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error)
}

type remoteTerminalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteTerminalServiceClient(cc grpc.ClientConnInterface) RemoteTerminalServiceClient {
	return &remoteTerminalServiceClient{cc}
}

func (c *remoteTerminalServiceClient) NewTerminal(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error) {
	out := new(RemoteTerminalResponse)
	err := c.cc.Invoke(ctx, "/terminal.RemoteTerminalService/NewTerminal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteTerminalServiceClient) CloseTerminal(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error) {
	out := new(RemoteTerminalResponse)
	err := c.cc.Invoke(ctx, "/terminal.RemoteTerminalService/CloseTerminal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteTerminalServiceClient) ExecCommand(ctx context.Context, in *RemoteTerminalRequest, opts ...grpc.CallOption) (*RemoteTerminalResponse, error) {
	out := new(RemoteTerminalResponse)
	err := c.cc.Invoke(ctx, "/terminal.RemoteTerminalService/ExecCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteTerminalServiceServer is the server API for RemoteTerminalService service.
// All implementations must embed UnimplementedRemoteTerminalServiceServer
// for forward compatibility
type RemoteTerminalServiceServer interface {
	NewTerminal(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error)
	CloseTerminal(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error)
	ExecCommand(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error)
	mustEmbedUnimplementedRemoteTerminalServiceServer()
}

// UnimplementedRemoteTerminalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteTerminalServiceServer struct {
}

func (UnimplementedRemoteTerminalServiceServer) NewTerminal(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTerminal not implemented")
}
func (UnimplementedRemoteTerminalServiceServer) CloseTerminal(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTerminal not implemented")
}
func (UnimplementedRemoteTerminalServiceServer) ExecCommand(context.Context, *RemoteTerminalRequest) (*RemoteTerminalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecCommand not implemented")
}
func (UnimplementedRemoteTerminalServiceServer) mustEmbedUnimplementedRemoteTerminalServiceServer() {}

// UnsafeRemoteTerminalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteTerminalServiceServer will
// result in compilation errors.
type UnsafeRemoteTerminalServiceServer interface {
	mustEmbedUnimplementedRemoteTerminalServiceServer()
}

func RegisterRemoteTerminalServiceServer(s grpc.ServiceRegistrar, srv RemoteTerminalServiceServer) {
	s.RegisterService(&RemoteTerminalService_ServiceDesc, srv)
}

func _RemoteTerminalService_NewTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteTerminalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteTerminalServiceServer).NewTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terminal.RemoteTerminalService/NewTerminal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteTerminalServiceServer).NewTerminal(ctx, req.(*RemoteTerminalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteTerminalService_CloseTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteTerminalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteTerminalServiceServer).CloseTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terminal.RemoteTerminalService/CloseTerminal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteTerminalServiceServer).CloseTerminal(ctx, req.(*RemoteTerminalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteTerminalService_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteTerminalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteTerminalServiceServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terminal.RemoteTerminalService/ExecCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteTerminalServiceServer).ExecCommand(ctx, req.(*RemoteTerminalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteTerminalService_ServiceDesc is the grpc.ServiceDesc for RemoteTerminalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteTerminalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "terminal.RemoteTerminalService",
	HandlerType: (*RemoteTerminalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTerminal",
			Handler:    _RemoteTerminalService_NewTerminal_Handler,
		},
		{
			MethodName: "CloseTerminal",
			Handler:    _RemoteTerminalService_CloseTerminal_Handler,
		},
		{
			MethodName: "ExecCommand",
			Handler:    _RemoteTerminalService_ExecCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_terminal.proto",
}
