// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/.proto

package proto

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

// WelcomeServiceClient is the client API for WelcomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WelcomeServiceClient interface {
	Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *CitiesArray, opts ...grpc.CallOption) (WelcomeService_ServerStreamClient, error)
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_BidirectionalStreamClient, error)
}

type welcomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWelcomeServiceClient(cc grpc.ClientConnInterface) WelcomeServiceClient {
	return &welcomeServiceClient{cc}
}

func (c *welcomeServiceClient) Welcome(ctx context.Context, in *WelcomeRequest, opts ...grpc.CallOption) (*WelcomeResponse, error) {
	out := new(WelcomeResponse)
	err := c.cc.Invoke(ctx, "/welcomeService.WelcomeService/Welcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *welcomeServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WelcomeService_ServiceDesc.Streams[0], "/welcomeService.WelcomeService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &welcomeServiceClientStreamClient{stream}
	return x, nil
}

type WelcomeService_ClientStreamClient interface {
	Send(*WelcomeRequest) error
	CloseAndRecv() (*CitiesArray, error)
	grpc.ClientStream
}

type welcomeServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *welcomeServiceClientStreamClient) Send(m *WelcomeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *welcomeServiceClientStreamClient) CloseAndRecv() (*CitiesArray, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CitiesArray)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *welcomeServiceClient) ServerStream(ctx context.Context, in *CitiesArray, opts ...grpc.CallOption) (WelcomeService_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WelcomeService_ServiceDesc.Streams[1], "/welcomeService.WelcomeService/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &welcomeServiceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WelcomeService_ServerStreamClient interface {
	Recv() (*WelcomeResponse, error)
	grpc.ClientStream
}

type welcomeServiceServerStreamClient struct {
	grpc.ClientStream
}

func (x *welcomeServiceServerStreamClient) Recv() (*WelcomeResponse, error) {
	m := new(WelcomeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *welcomeServiceClient) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WelcomeService_ServiceDesc.Streams[2], "/welcomeService.WelcomeService/BidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &welcomeServiceBidirectionalStreamClient{stream}
	return x, nil
}

type WelcomeService_BidirectionalStreamClient interface {
	Send(*WelcomeRequest) error
	Recv() (*WelcomeResponse, error)
	grpc.ClientStream
}

type welcomeServiceBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *welcomeServiceBidirectionalStreamClient) Send(m *WelcomeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *welcomeServiceBidirectionalStreamClient) Recv() (*WelcomeResponse, error) {
	m := new(WelcomeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WelcomeServiceServer is the server API for WelcomeService service.
// All implementations must embed UnimplementedWelcomeServiceServer
// for forward compatibility
type WelcomeServiceServer interface {
	Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error)
	ClientStream(WelcomeService_ClientStreamServer) error
	ServerStream(*CitiesArray, WelcomeService_ServerStreamServer) error
	BidirectionalStream(WelcomeService_BidirectionalStreamServer) error
	mustEmbedUnimplementedWelcomeServiceServer()
}

// UnimplementedWelcomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWelcomeServiceServer struct {
}

func (UnimplementedWelcomeServiceServer) Welcome(context.Context, *WelcomeRequest) (*WelcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}
func (UnimplementedWelcomeServiceServer) ClientStream(WelcomeService_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedWelcomeServiceServer) ServerStream(*CitiesArray, WelcomeService_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedWelcomeServiceServer) BidirectionalStream(WelcomeService_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}
func (UnimplementedWelcomeServiceServer) mustEmbedUnimplementedWelcomeServiceServer() {}

// UnsafeWelcomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WelcomeServiceServer will
// result in compilation errors.
type UnsafeWelcomeServiceServer interface {
	mustEmbedUnimplementedWelcomeServiceServer()
}

func RegisterWelcomeServiceServer(s grpc.ServiceRegistrar, srv WelcomeServiceServer) {
	s.RegisterService(&WelcomeService_ServiceDesc, srv)
}

func _WelcomeService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WelcomeServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/welcomeService.WelcomeService/Welcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WelcomeServiceServer).Welcome(ctx, req.(*WelcomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WelcomeService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WelcomeServiceServer).ClientStream(&welcomeServiceClientStreamServer{stream})
}

type WelcomeService_ClientStreamServer interface {
	SendAndClose(*CitiesArray) error
	Recv() (*WelcomeRequest, error)
	grpc.ServerStream
}

type welcomeServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *welcomeServiceClientStreamServer) SendAndClose(m *CitiesArray) error {
	return x.ServerStream.SendMsg(m)
}

func (x *welcomeServiceClientStreamServer) Recv() (*WelcomeRequest, error) {
	m := new(WelcomeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WelcomeService_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CitiesArray)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WelcomeServiceServer).ServerStream(m, &welcomeServiceServerStreamServer{stream})
}

type WelcomeService_ServerStreamServer interface {
	Send(*WelcomeResponse) error
	grpc.ServerStream
}

type welcomeServiceServerStreamServer struct {
	grpc.ServerStream
}

func (x *welcomeServiceServerStreamServer) Send(m *WelcomeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WelcomeService_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WelcomeServiceServer).BidirectionalStream(&welcomeServiceBidirectionalStreamServer{stream})
}

type WelcomeService_BidirectionalStreamServer interface {
	Send(*WelcomeResponse) error
	Recv() (*WelcomeRequest, error)
	grpc.ServerStream
}

type welcomeServiceBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *welcomeServiceBidirectionalStreamServer) Send(m *WelcomeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *welcomeServiceBidirectionalStreamServer) Recv() (*WelcomeRequest, error) {
	m := new(WelcomeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WelcomeService_ServiceDesc is the grpc.ServiceDesc for WelcomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WelcomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "welcomeService.WelcomeService",
	HandlerType: (*WelcomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _WelcomeService_Welcome_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _WelcomeService_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _WelcomeService_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _WelcomeService_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/.proto",
}
