// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/admin/service/v1/admin.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUserRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_admin_proto_rawDescGZIP(), []int{1}
}

type ListUserReply_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Age      int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *ListUserReply_User) Reset() {
	*x = ListUserReply_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply_User) ProtoMessage() {}

func (x *ListUserReply_User) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply_User.ProtoReflect.Descriptor instead.
func (*ListUserReply_User) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_admin_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListUserReply_User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListUserReply_User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ListUserReply_User) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *ListUserReply_User) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_api_admin_service_v1_admin_proto protoreflect.FileDescriptor

var file_api_admin_service_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x1a, 0x5d, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x32, 0x78, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x6f, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x1f, 0x5a,
	0x1d, 0x63, 0x61, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_admin_service_v1_admin_proto_rawDescOnce sync.Once
	file_api_admin_service_v1_admin_proto_rawDescData = file_api_admin_service_v1_admin_proto_rawDesc
)

func file_api_admin_service_v1_admin_proto_rawDescGZIP() []byte {
	file_api_admin_service_v1_admin_proto_rawDescOnce.Do(func() {
		file_api_admin_service_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_service_v1_admin_proto_rawDescData)
	})
	return file_api_admin_service_v1_admin_proto_rawDescData
}

var file_api_admin_service_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_admin_service_v1_admin_proto_goTypes = []interface{}{
	(*ListUserRequest)(nil),    // 0: api.admin.service.v1.ListUserRequest
	(*ListUserReply)(nil),      // 1: api.admin.service.v1.ListUserReply
	(*ListUserReply_User)(nil), // 2: api.admin.service.v1.ListUserReply.User
}
var file_api_admin_service_v1_admin_proto_depIdxs = []int32{
	0, // 0: api.admin.service.v1.Admin.ListUser:input_type -> api.admin.service.v1.ListUserRequest
	1, // 1: api.admin.service.v1.Admin.ListUser:output_type -> api.admin.service.v1.ListUserReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_admin_service_v1_admin_proto_init() }
func file_api_admin_service_v1_admin_proto_init() {
	if File_api_admin_service_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_service_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_api_admin_service_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
		file_api_admin_service_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply_User); i {
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
			RawDescriptor: file_api_admin_service_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_admin_service_v1_admin_proto_goTypes,
		DependencyIndexes: file_api_admin_service_v1_admin_proto_depIdxs,
		MessageInfos:      file_api_admin_service_v1_admin_proto_msgTypes,
	}.Build()
	File_api_admin_service_v1_admin_proto = out.File
	file_api_admin_service_v1_admin_proto_rawDesc = nil
	file_api_admin_service_v1_admin_proto_goTypes = nil
	file_api_admin_service_v1_admin_proto_depIdxs = nil
}
