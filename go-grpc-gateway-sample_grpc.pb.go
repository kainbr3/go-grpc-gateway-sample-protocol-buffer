// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: go-grpc-gateway-sample.proto

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

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// List all registered devices
	GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error)
	// Update a device&#8217;s state
	SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error)
	// Register a new device
	RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error) {
	out := new(Devices)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/GetAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/GetDeviceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/SwitchDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	// List all registered devices
	GetAllDevices(context.Context, *Empty) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(context.Context, *ID) (*Device, error)
	// Update a device&#8217;s state
	SwitchDevice(context.Context, *UpdateDevice) (*Device, error)
	// Register a new device
	RegisterDevice(context.Context, *Device) (*Device, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) GetAllDevices(context.Context, *Empty) (*Devices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDevices not implemented")
}
func (UnimplementedDeviceServiceServer) GetDeviceByID(context.Context, *ID) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByID not implemented")
}
func (UnimplementedDeviceServiceServer) SwitchDevice(context.Context, *UpdateDevice) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchDevice not implemented")
}
func (UnimplementedDeviceServiceServer) RegisterDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_GetAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/GetAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/GetDeviceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_SwitchDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/SwitchDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, req.(*UpdateDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDevices",
			Handler:    _DeviceService_GetAllDevices_Handler,
		},
		{
			MethodName: "GetDeviceByID",
			Handler:    _DeviceService_GetDeviceByID_Handler,
		},
		{
			MethodName: "SwitchDevice",
			Handler:    _DeviceService_SwitchDevice_Handler,
		},
		{
			MethodName: "RegisterDevice",
			Handler:    _DeviceService_RegisterDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-grpc-gateway-sample.proto",
}
