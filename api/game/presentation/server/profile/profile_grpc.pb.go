// プロフィール

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: profile.proto

package profile

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

const (
	Profile_Get_FullMethodName    = "/proto.Profile/Get"
	Profile_Create_FullMethodName = "/proto.Profile/Create"
	Profile_Update_FullMethodName = "/proto.Profile/Update"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	Get(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error)
	Create(ctx context.Context, in *ProfileCreateRequest, opts ...grpc.CallOption) (*ProfileCreateResponse, error)
	Update(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileUpdateResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Get(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error) {
	out := new(ProfileGetResponse)
	err := c.cc.Invoke(ctx, Profile_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Create(ctx context.Context, in *ProfileCreateRequest, opts ...grpc.CallOption) (*ProfileCreateResponse, error) {
	out := new(ProfileCreateResponse)
	err := c.cc.Invoke(ctx, Profile_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Update(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileUpdateResponse, error) {
	out := new(ProfileUpdateResponse)
	err := c.cc.Invoke(ctx, Profile_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	Get(context.Context, *ProfileGetRequest) (*ProfileGetResponse, error)
	Create(context.Context, *ProfileCreateRequest) (*ProfileCreateResponse, error)
	Update(context.Context, *ProfileUpdateRequest) (*ProfileUpdateResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) Get(context.Context, *ProfileGetRequest) (*ProfileGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProfileServer) Create(context.Context, *ProfileCreateRequest) (*ProfileCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfileServer) Update(context.Context, *ProfileUpdateRequest) (*ProfileUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Get(ctx, req.(*ProfileGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Create(ctx, req.(*ProfileCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Update(ctx, req.(*ProfileUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Profile_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Profile_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Profile_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
