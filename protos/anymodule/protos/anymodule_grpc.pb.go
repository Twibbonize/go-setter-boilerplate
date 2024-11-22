// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: protos/anymodule/anymodule.proto

package protos

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

// SetterClient is the client API for Setter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SetterClient interface {
	SeedOneByRandId(ctx context.Context, in *IngestRequestByRandId, opts ...grpc.CallOption) (*IngestStatus, error)
	SeedOneByUUID(ctx context.Context, in *IngestRequestByUUID, opts ...grpc.CallOption) (*IngestStatus, error)
	SeedMany(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestStatus, error)
	DeleteManyByAnyUUID(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*IngestStatus, error)
}

type setterClient struct {
	cc grpc.ClientConnInterface
}

func NewSetterClient(cc grpc.ClientConnInterface) SetterClient {
	return &setterClient{cc}
}

func (c *setterClient) SeedOneByRandId(ctx context.Context, in *IngestRequestByRandId, opts ...grpc.CallOption) (*IngestStatus, error) {
	out := new(IngestStatus)
	err := c.cc.Invoke(ctx, "/anymodule_proto.Setter/SeedOneByRandId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setterClient) SeedOneByUUID(ctx context.Context, in *IngestRequestByUUID, opts ...grpc.CallOption) (*IngestStatus, error) {
	out := new(IngestStatus)
	err := c.cc.Invoke(ctx, "/anymodule_proto.Setter/SeedOneByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setterClient) SeedMany(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestStatus, error) {
	out := new(IngestStatus)
	err := c.cc.Invoke(ctx, "/anymodule_proto.Setter/SeedMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setterClient) DeleteManyByAnyUUID(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*IngestStatus, error) {
	out := new(IngestStatus)
	err := c.cc.Invoke(ctx, "/anymodule_proto.Setter/DeleteManyByAnyUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetterServer is the server API for Setter service.
// All implementations must embed UnimplementedSetterServer
// for forward compatibility
type SetterServer interface {
	SeedOneByRandId(context.Context, *IngestRequestByRandId) (*IngestStatus, error)
	SeedOneByUUID(context.Context, *IngestRequestByUUID) (*IngestStatus, error)
	SeedMany(context.Context, *IngestRequest) (*IngestStatus, error)
	DeleteManyByAnyUUID(context.Context, *DeleteAllRequest) (*IngestStatus, error)
	mustEmbedUnimplementedSetterServer()
}

// UnimplementedSetterServer must be embedded to have forward compatible implementations.
type UnimplementedSetterServer struct {
}

func (UnimplementedSetterServer) SeedOneByRandId(context.Context, *IngestRequestByRandId) (*IngestStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeedOneByRandId not implemented")
}
func (UnimplementedSetterServer) SeedOneByUUID(context.Context, *IngestRequestByUUID) (*IngestStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeedOneByUUID not implemented")
}
func (UnimplementedSetterServer) SeedMany(context.Context, *IngestRequest) (*IngestStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeedMany not implemented")
}
func (UnimplementedSetterServer) DeleteManyByAnyUUID(context.Context, *DeleteAllRequest) (*IngestStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManyByAnyUUID not implemented")
}
func (UnimplementedSetterServer) mustEmbedUnimplementedSetterServer() {}

// UnsafeSetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SetterServer will
// result in compilation errors.
type UnsafeSetterServer interface {
	mustEmbedUnimplementedSetterServer()
}

func RegisterSetterServer(s grpc.ServiceRegistrar, srv SetterServer) {
	s.RegisterService(&Setter_ServiceDesc, srv)
}

func _Setter_SeedOneByRandId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequestByRandId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetterServer).SeedOneByRandId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anymodule_proto.Setter/SeedOneByRandId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetterServer).SeedOneByRandId(ctx, req.(*IngestRequestByRandId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setter_SeedOneByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequestByUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetterServer).SeedOneByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anymodule_proto.Setter/SeedOneByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetterServer).SeedOneByUUID(ctx, req.(*IngestRequestByUUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setter_SeedMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetterServer).SeedMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anymodule_proto.Setter/SeedMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetterServer).SeedMany(ctx, req.(*IngestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setter_DeleteManyByAnyUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetterServer).DeleteManyByAnyUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anymodule_proto.Setter/DeleteManyByAnyUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetterServer).DeleteManyByAnyUUID(ctx, req.(*DeleteAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Setter_ServiceDesc is the grpc.ServiceDesc for Setter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Setter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anymodule_proto.Setter",
	HandlerType: (*SetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeedOneByRandId",
			Handler:    _Setter_SeedOneByRandId_Handler,
		},
		{
			MethodName: "SeedOneByUUID",
			Handler:    _Setter_SeedOneByUUID_Handler,
		},
		{
			MethodName: "SeedMany",
			Handler:    _Setter_SeedMany_Handler,
		},
		{
			MethodName: "DeleteManyByAnyUUID",
			Handler:    _Setter_DeleteManyByAnyUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/anymodule/anymodule.proto",
}
