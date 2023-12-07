// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user/user.proto

package helloworld

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

// UserSearchClient is the client API for UserSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSearchClient interface {
	FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type userSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSearchClient(cc grpc.ClientConnInterface) UserSearchClient {
	return &userSearchClient{cc}
}

func (c *userSearchClient) FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.UserSearch/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSearchServer is the server API for UserSearch service.
// All implementations must embed UnimplementedUserSearchServer
// for forward compatibility
type UserSearchServer interface {
	FindUser(context.Context, *UserRequest) (*UserReply, error)
	mustEmbedUnimplementedUserSearchServer()
}

// UnimplementedUserSearchServer must be embedded to have forward compatible implementations.
type UnimplementedUserSearchServer struct {
}

func (UnimplementedUserSearchServer) FindUser(context.Context, *UserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedUserSearchServer) mustEmbedUnimplementedUserSearchServer() {}

// UnsafeUserSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSearchServer will
// result in compilation errors.
type UnsafeUserSearchServer interface {
	mustEmbedUnimplementedUserSearchServer()
}

func RegisterUserSearchServer(s grpc.ServiceRegistrar, srv UserSearchServer) {
	s.RegisterService(&UserSearch_ServiceDesc, srv)
}

func _UserSearch_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSearchServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSearch/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSearchServer).FindUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSearch_ServiceDesc is the grpc.ServiceDesc for UserSearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSearch",
	HandlerType: (*UserSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUser",
			Handler:    _UserSearch_FindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
