// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: booking.proto

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

// FPTBookingClient is the client API for FPTBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FPTBookingClient interface {
	CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	FindBooking(ctx context.Context, in *FindBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	ListBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
}

type fPTBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewFPTBookingClient(cc grpc.ClientConnInterface) FPTBookingClient {
	return &fPTBookingClient{cc}
}

func (c *fPTBookingClient) CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) FindBooking(ctx context.Context, in *FindBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/FindBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) ListBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/ListBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FPTBookingServer is the server API for FPTBooking service.
// All implementations must embed UnimplementedFPTBookingServer
// for forward compatibility
type FPTBookingServer interface {
	CreateBooking(context.Context, *Booking) (*Booking, error)
	FindBooking(context.Context, *FindBookingRequest) (*Booking, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*Booking, error)
	ListBooking(context.Context, *Booking) (*Booking, error)
	mustEmbedUnimplementedFPTBookingServer()
}

// UnimplementedFPTBookingServer must be embedded to have forward compatible implementations.
type UnimplementedFPTBookingServer struct {
}

func (UnimplementedFPTBookingServer) CreateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedFPTBookingServer) FindBooking(context.Context, *FindBookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBooking not implemented")
}
func (UnimplementedFPTBookingServer) CancelBooking(context.Context, *CancelBookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedFPTBookingServer) ListBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooking not implemented")
}
func (UnimplementedFPTBookingServer) mustEmbedUnimplementedFPTBookingServer() {}

// UnsafeFPTBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FPTBookingServer will
// result in compilation errors.
type UnsafeFPTBookingServer interface {
	mustEmbedUnimplementedFPTBookingServer()
}

func RegisterFPTBookingServer(s grpc.ServiceRegistrar, srv FPTBookingServer) {
	s.RegisterService(&FPTBooking_ServiceDesc, srv)
}

func _FPTBooking_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).CreateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_FindBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).FindBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/FindBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).FindBooking(ctx, req.(*FindBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_ListBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).ListBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/ListBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).ListBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

// FPTBooking_ServiceDesc is the grpc.ServiceDesc for FPTBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FPTBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.FPTBooking",
	HandlerType: (*FPTBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _FPTBooking_CreateBooking_Handler,
		},
		{
			MethodName: "FindBooking",
			Handler:    _FPTBooking_FindBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _FPTBooking_CancelBooking_Handler,
		},
		{
			MethodName: "ListBooking",
			Handler:    _FPTBooking_ListBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
