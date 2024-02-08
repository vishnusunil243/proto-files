// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
	AdminLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
	SuperAdminLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
	GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllUsersClient, error)
	GetAllAdmins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllAdminsClient, error)
	AddAdmin(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdminLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/AdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SuperAdminLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/SuperAdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/user.UserService/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetAllUsersClient interface {
	Recv() (*UserSignupResponse, error)
	grpc.ClientStream
}

type userServiceGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceGetAllUsersClient) Recv() (*UserSignupResponse, error) {
	m := new(UserSignupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetAllAdmins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllAdminsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/user.UserService/GetAllAdmins", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetAllAdminsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetAllAdminsClient interface {
	Recv() (*UserSignupResponse, error)
	grpc.ClientStream
}

type userServiceGetAllAdminsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetAllAdminsClient) Recv() (*UserSignupResponse, error) {
	m := new(UserSignupResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) AddAdmin(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/AddAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserSignup(context.Context, *UserSignupRequest) (*UserSignupResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error)
	AdminLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error)
	SuperAdminLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error)
	GetAllUsers(*emptypb.Empty, UserService_GetAllUsersServer) error
	GetAllAdmins(*emptypb.Empty, UserService_GetAllAdminsServer) error
	AddAdmin(context.Context, *UserSignupRequest) (*UserSignupResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserSignup(context.Context, *UserSignupRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) AdminLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedUserServiceServer) SuperAdminLogin(context.Context, *UserLoginRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperAdminLogin not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(*emptypb.Empty, UserService_GetAllUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) GetAllAdmins(*emptypb.Empty, UserService_GetAllAdminsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAdmins not implemented")
}
func (UnimplementedUserServiceServer) AddAdmin(context.Context, *UserSignupRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignup(ctx, req.(*UserSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdminLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SuperAdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SuperAdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SuperAdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SuperAdminLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetAllUsers(m, &userServiceGetAllUsersServer{stream})
}

type UserService_GetAllUsersServer interface {
	Send(*UserSignupResponse) error
	grpc.ServerStream
}

type userServiceGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceGetAllUsersServer) Send(m *UserSignupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_GetAllAdmins_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetAllAdmins(m, &userServiceGetAllAdminsServer{stream})
}

type UserService_GetAllAdminsServer interface {
	Send(*UserSignupResponse) error
	grpc.ServerStream
}

type userServiceGetAllAdminsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetAllAdminsServer) Send(m *UserSignupResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AddAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAdmin(ctx, req.(*UserSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignup",
			Handler:    _UserService_UserSignup_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _UserService_AdminLogin_Handler,
		},
		{
			MethodName: "SuperAdminLogin",
			Handler:    _UserService_SuperAdminLogin_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _UserService_AddAdmin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllUsers",
			Handler:       _UserService_GetAllUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllAdmins",
			Handler:       _UserService_GetAllAdmins_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
