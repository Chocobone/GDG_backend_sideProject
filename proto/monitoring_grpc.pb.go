// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: monitoring.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Bitcoin_Service_SayHello_FullMethodName = "/monitoring.Bitcoin_Service/SayHello"
)

// Bitcoin_ServiceClient is the client API for Bitcoin_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Bitcoin_ServiceClient interface {
	SayHello(ctx context.Context, in *BitcoinRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BitcoinResponse], error)
}

type bitcoin_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBitcoin_ServiceClient(cc grpc.ClientConnInterface) Bitcoin_ServiceClient {
	return &bitcoin_ServiceClient{cc}
}

func (c *bitcoin_ServiceClient) SayHello(ctx context.Context, in *BitcoinRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BitcoinResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Bitcoin_Service_ServiceDesc.Streams[0], Bitcoin_Service_SayHello_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BitcoinRequest, BitcoinResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bitcoin_Service_SayHelloClient = grpc.ServerStreamingClient[BitcoinResponse]

// Bitcoin_ServiceServer is the server API for Bitcoin_Service service.
// All implementations must embed UnimplementedBitcoin_ServiceServer
// for forward compatibility.
type Bitcoin_ServiceServer interface {
	SayHello(*BitcoinRequest, grpc.ServerStreamingServer[BitcoinResponse]) error
	mustEmbedUnimplementedBitcoin_ServiceServer()
}

// UnimplementedBitcoin_ServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBitcoin_ServiceServer struct{}

func (UnimplementedBitcoin_ServiceServer) SayHello(*BitcoinRequest, grpc.ServerStreamingServer[BitcoinResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBitcoin_ServiceServer) mustEmbedUnimplementedBitcoin_ServiceServer() {}
func (UnimplementedBitcoin_ServiceServer) testEmbeddedByValue()                         {}

// UnsafeBitcoin_ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Bitcoin_ServiceServer will
// result in compilation errors.
type UnsafeBitcoin_ServiceServer interface {
	mustEmbedUnimplementedBitcoin_ServiceServer()
}

func RegisterBitcoin_ServiceServer(s grpc.ServiceRegistrar, srv Bitcoin_ServiceServer) {
	// If the following call pancis, it indicates UnimplementedBitcoin_ServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bitcoin_Service_ServiceDesc, srv)
}

func _Bitcoin_Service_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BitcoinRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Bitcoin_ServiceServer).SayHello(m, &grpc.GenericServerStream[BitcoinRequest, BitcoinResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bitcoin_Service_SayHelloServer = grpc.ServerStreamingServer[BitcoinResponse]

// Bitcoin_Service_ServiceDesc is the grpc.ServiceDesc for Bitcoin_Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bitcoin_Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Bitcoin_Service",
	HandlerType: (*Bitcoin_ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Bitcoin_Service_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "monitoring.proto",
}
