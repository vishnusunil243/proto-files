// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: wishlist.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateWishlistRequest) Reset() {
	*x = CreateWishlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistRequest) ProtoMessage() {}

func (x *CreateWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishlistRequest.ProtoReflect.Descriptor instead.
func (*CreateWishlistRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWishlistRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddToWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductId uint32 `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *AddToWishlistRequest) Reset() {
	*x = AddToWishlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToWishlistRequest) ProtoMessage() {}

func (x *AddToWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToWishlistRequest.ProtoReflect.Descriptor instead.
func (*AddToWishlistRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{1}
}

func (x *AddToWishlistRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddToWishlistRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{2}
}

var File_wishlist_proto protoreflect.FileDescriptor

var file_wishlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x32, 0xa9, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x77, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x50,
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1e, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wishlist_proto_rawDescOnce sync.Once
	file_wishlist_proto_rawDescData = file_wishlist_proto_rawDesc
)

func file_wishlist_proto_rawDescGZIP() []byte {
	file_wishlist_proto_rawDescOnce.Do(func() {
		file_wishlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_wishlist_proto_rawDescData)
	})
	return file_wishlist_proto_rawDescData
}

var file_wishlist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wishlist_proto_goTypes = []interface{}{
	(*CreateWishlistRequest)(nil), // 0: wishlist.CreateWishlistRequest
	(*AddToWishlistRequest)(nil),  // 1: wishlist.AddToWishlistRequest
	(*NoParam)(nil),               // 2: wishlist.NoParam
}
var file_wishlist_proto_depIdxs = []int32{
	0, // 0: wishlist.WishlistService.CreateWishlist:input_type -> wishlist.CreateWishlistRequest
	1, // 1: wishlist.WishlistService.AddToWishlist:input_type -> wishlist.AddToWishlistRequest
	2, // 2: wishlist.WishlistService.CreateWishlist:output_type -> wishlist.NoParam
	0, // 3: wishlist.WishlistService.AddToWishlist:output_type -> wishlist.CreateWishlistRequest
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wishlist_proto_init() }
func file_wishlist_proto_init() {
	if File_wishlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wishlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWishlistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wishlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToWishlistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wishlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wishlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wishlist_proto_goTypes,
		DependencyIndexes: file_wishlist_proto_depIdxs,
		MessageInfos:      file_wishlist_proto_msgTypes,
	}.Build()
	File_wishlist_proto = out.File
	file_wishlist_proto_rawDesc = nil
	file_wishlist_proto_goTypes = nil
	file_wishlist_proto_depIdxs = nil
}
