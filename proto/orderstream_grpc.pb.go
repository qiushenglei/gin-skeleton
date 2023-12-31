// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: orderstream.proto

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

// OrderStreamServerClient is the client API for OrderStreamServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderStreamServerClient interface {
	// 获取列表
	FindAll(ctx context.Context, opts ...grpc.CallOption) (OrderStreamServer_FindAllClient, error)
	Insert(ctx context.Context, opts ...grpc.CallOption) (OrderStreamServer_InsertClient, error)
}

type orderStreamServerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderStreamServerClient(cc grpc.ClientConnInterface) OrderStreamServerClient {
	return &orderStreamServerClient{cc}
}

func (c *orderStreamServerClient) FindAll(ctx context.Context, opts ...grpc.CallOption) (OrderStreamServer_FindAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderStreamServer_ServiceDesc.Streams[0], "/pdfiles.user.OrderStreamServer/FindAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderStreamServerFindAllClient{stream}
	return x, nil
}

type OrderStreamServer_FindAllClient interface {
	Send(*FindAllStreamRequest) error
	Recv() (*FindAllStreamReply, error)
	grpc.ClientStream
}

type orderStreamServerFindAllClient struct {
	grpc.ClientStream
}

func (x *orderStreamServerFindAllClient) Send(m *FindAllStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderStreamServerFindAllClient) Recv() (*FindAllStreamReply, error) {
	m := new(FindAllStreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderStreamServerClient) Insert(ctx context.Context, opts ...grpc.CallOption) (OrderStreamServer_InsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderStreamServer_ServiceDesc.Streams[1], "/pdfiles.user.OrderStreamServer/Insert", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderStreamServerInsertClient{stream}
	return x, nil
}

type OrderStreamServer_InsertClient interface {
	Send(*InsertStreamRequest) error
	Recv() (*InsertStreamReply, error)
	grpc.ClientStream
}

type orderStreamServerInsertClient struct {
	grpc.ClientStream
}

func (x *orderStreamServerInsertClient) Send(m *InsertStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderStreamServerInsertClient) Recv() (*InsertStreamReply, error) {
	m := new(InsertStreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderStreamServerServer is the server API for OrderStreamServer service.
// All implementations must embed UnimplementedOrderStreamServerServer
// for forward compatibility
type OrderStreamServerServer interface {
	// 获取列表
	FindAll(OrderStreamServer_FindAllServer) error
	Insert(OrderStreamServer_InsertServer) error
	mustEmbedUnimplementedOrderStreamServerServer()
}

// UnimplementedOrderStreamServerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderStreamServerServer struct {
}

func (UnimplementedOrderStreamServerServer) FindAll(OrderStreamServer_FindAllServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedOrderStreamServerServer) Insert(OrderStreamServer_InsertServer) error {
	return status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedOrderStreamServerServer) mustEmbedUnimplementedOrderStreamServerServer() {}

// UnsafeOrderStreamServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderStreamServerServer will
// result in compilation errors.
type UnsafeOrderStreamServerServer interface {
	mustEmbedUnimplementedOrderStreamServerServer()
}

func RegisterOrderStreamServerServer(s grpc.ServiceRegistrar, srv OrderStreamServerServer) {
	s.RegisterService(&OrderStreamServer_ServiceDesc, srv)
}

func _OrderStreamServer_FindAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderStreamServerServer).FindAll(&orderStreamServerFindAllServer{stream})
}

type OrderStreamServer_FindAllServer interface {
	Send(*FindAllStreamReply) error
	Recv() (*FindAllStreamRequest, error)
	grpc.ServerStream
}

type orderStreamServerFindAllServer struct {
	grpc.ServerStream
}

func (x *orderStreamServerFindAllServer) Send(m *FindAllStreamReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderStreamServerFindAllServer) Recv() (*FindAllStreamRequest, error) {
	m := new(FindAllStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderStreamServer_Insert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderStreamServerServer).Insert(&orderStreamServerInsertServer{stream})
}

type OrderStreamServer_InsertServer interface {
	Send(*InsertStreamReply) error
	Recv() (*InsertStreamRequest, error)
	grpc.ServerStream
}

type orderStreamServerInsertServer struct {
	grpc.ServerStream
}

func (x *orderStreamServerInsertServer) Send(m *InsertStreamReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderStreamServerInsertServer) Recv() (*InsertStreamRequest, error) {
	m := new(InsertStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderStreamServer_ServiceDesc is the grpc.ServiceDesc for OrderStreamServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderStreamServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfiles.user.OrderStreamServer",
	HandlerType: (*OrderStreamServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAll",
			Handler:       _OrderStreamServer_FindAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Insert",
			Handler:       _OrderStreamServer_Insert_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderstream.proto",
}
