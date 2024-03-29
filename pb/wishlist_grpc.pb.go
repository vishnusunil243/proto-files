// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: wishlist.proto

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

// WishlistServiceClient is the client API for WishlistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishlistServiceClient interface {
	CreateWishlist(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistRequest, error)
	RemoveFromWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistRequest, error)
	GetAllWishlistItems(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (WishlistService_GetAllWishlistItemsClient, error)
}

type wishlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishlistServiceClient(cc grpc.ClientConnInterface) WishlistServiceClient {
	return &wishlistServiceClient{cc}
}

func (c *wishlistServiceClient) CreateWishlist(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/wishlist.WishlistService/CreateWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) AddToWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistRequest, error) {
	out := new(CreateWishlistRequest)
	err := c.cc.Invoke(ctx, "/wishlist.WishlistService/AddToWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) RemoveFromWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*CreateWishlistRequest, error) {
	out := new(CreateWishlistRequest)
	err := c.cc.Invoke(ctx, "/wishlist.WishlistService/RemoveFromWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) GetAllWishlistItems(ctx context.Context, in *CreateWishlistRequest, opts ...grpc.CallOption) (WishlistService_GetAllWishlistItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &WishlistService_ServiceDesc.Streams[0], "/wishlist.WishlistService/GetAllWishlistItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &wishlistServiceGetAllWishlistItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WishlistService_GetAllWishlistItemsClient interface {
	Recv() (*GetAllWishlistResponse, error)
	grpc.ClientStream
}

type wishlistServiceGetAllWishlistItemsClient struct {
	grpc.ClientStream
}

func (x *wishlistServiceGetAllWishlistItemsClient) Recv() (*GetAllWishlistResponse, error) {
	m := new(GetAllWishlistResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WishlistServiceServer is the server API for WishlistService service.
// All implementations must embed UnimplementedWishlistServiceServer
// for forward compatibility
type WishlistServiceServer interface {
	CreateWishlist(context.Context, *CreateWishlistRequest) (*emptypb.Empty, error)
	AddToWishlist(context.Context, *AddToWishlistRequest) (*CreateWishlistRequest, error)
	RemoveFromWishlist(context.Context, *AddToWishlistRequest) (*CreateWishlistRequest, error)
	GetAllWishlistItems(*CreateWishlistRequest, WishlistService_GetAllWishlistItemsServer) error
	mustEmbedUnimplementedWishlistServiceServer()
}

// UnimplementedWishlistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWishlistServiceServer struct {
}

func (UnimplementedWishlistServiceServer) CreateWishlist(context.Context, *CreateWishlistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWishlist not implemented")
}
func (UnimplementedWishlistServiceServer) AddToWishlist(context.Context, *AddToWishlistRequest) (*CreateWishlistRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWishlist not implemented")
}
func (UnimplementedWishlistServiceServer) RemoveFromWishlist(context.Context, *AddToWishlistRequest) (*CreateWishlistRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWishlist not implemented")
}
func (UnimplementedWishlistServiceServer) GetAllWishlistItems(*CreateWishlistRequest, WishlistService_GetAllWishlistItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllWishlistItems not implemented")
}
func (UnimplementedWishlistServiceServer) mustEmbedUnimplementedWishlistServiceServer() {}

// UnsafeWishlistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishlistServiceServer will
// result in compilation errors.
type UnsafeWishlistServiceServer interface {
	mustEmbedUnimplementedWishlistServiceServer()
}

func RegisterWishlistServiceServer(s grpc.ServiceRegistrar, srv WishlistServiceServer) {
	s.RegisterService(&WishlistService_ServiceDesc, srv)
}

func _WishlistService_CreateWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).CreateWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.WishlistService/CreateWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).CreateWishlist(ctx, req.(*CreateWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_AddToWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).AddToWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.WishlistService/AddToWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).AddToWishlist(ctx, req.(*AddToWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_RemoveFromWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).RemoveFromWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.WishlistService/RemoveFromWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).RemoveFromWishlist(ctx, req.(*AddToWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_GetAllWishlistItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateWishlistRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WishlistServiceServer).GetAllWishlistItems(m, &wishlistServiceGetAllWishlistItemsServer{stream})
}

type WishlistService_GetAllWishlistItemsServer interface {
	Send(*GetAllWishlistResponse) error
	grpc.ServerStream
}

type wishlistServiceGetAllWishlistItemsServer struct {
	grpc.ServerStream
}

func (x *wishlistServiceGetAllWishlistItemsServer) Send(m *GetAllWishlistResponse) error {
	return x.ServerStream.SendMsg(m)
}

// WishlistService_ServiceDesc is the grpc.ServiceDesc for WishlistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WishlistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wishlist.WishlistService",
	HandlerType: (*WishlistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWishlist",
			Handler:    _WishlistService_CreateWishlist_Handler,
		},
		{
			MethodName: "AddToWishlist",
			Handler:    _WishlistService_AddToWishlist_Handler,
		},
		{
			MethodName: "RemoveFromWishlist",
			Handler:    _WishlistService_RemoveFromWishlist_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllWishlistItems",
			Handler:       _WishlistService_GetAllWishlistItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wishlist.proto",
}
