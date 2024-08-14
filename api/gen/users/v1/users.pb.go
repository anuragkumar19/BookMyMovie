// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: users/v1/users.proto

package usersv1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_ROLE_UNSPECIFIED  Role = 0
	Role_ROLE_REGULAR_USER Role = 1
	Role_ROLE_ADMIN        Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_REGULAR_USER",
		2: "ROLE_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED":  0,
		"ROLE_REGULAR_USER": 1,
		"ROLE_ADMIN":        2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_users_v1_users_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_users_v1_users_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{0}
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{0}
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role      Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=users.v1.Role" json:"role,omitempty"`
	Dob       *date.Date             `protobuf:"bytes,5,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	Version   int32                  `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_v1_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_v1_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_users_v1_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserInfoResponse) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *GetUserInfoResponse) GetDob() *date.Date {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *GetUserInfoResponse) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GetUserInfoResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_users_v1_users_proto protoreflect.FileDescriptor

var file_users_v1_users_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x28, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x64, 0x6f, 0x62, 0x2a, 0x43, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x52, 0x45,
	0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x32, 0x5a, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x79, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x79, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_v1_users_proto_rawDescOnce sync.Once
	file_users_v1_users_proto_rawDescData = file_users_v1_users_proto_rawDesc
)

func file_users_v1_users_proto_rawDescGZIP() []byte {
	file_users_v1_users_proto_rawDescOnce.Do(func() {
		file_users_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_v1_users_proto_rawDescData)
	})
	return file_users_v1_users_proto_rawDescData
}

var file_users_v1_users_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_users_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_users_v1_users_proto_goTypes = []any{
	(Role)(0),                     // 0: users.v1.Role
	(*GetUserInfoRequest)(nil),    // 1: users.v1.GetUserInfoRequest
	(*GetUserInfoResponse)(nil),   // 2: users.v1.GetUserInfoResponse
	(*date.Date)(nil),             // 3: google.type.Date
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_users_v1_users_proto_depIdxs = []int32{
	0, // 0: users.v1.GetUserInfoResponse.role:type_name -> users.v1.Role
	3, // 1: users.v1.GetUserInfoResponse.dob:type_name -> google.type.Date
	4, // 2: users.v1.GetUserInfoResponse.created_at:type_name -> google.protobuf.Timestamp
	1, // 3: users.v1.UsersService.GetUserInfo:input_type -> users.v1.GetUserInfoRequest
	2, // 4: users.v1.UsersService.GetUserInfo:output_type -> users.v1.GetUserInfoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_users_v1_users_proto_init() }
func file_users_v1_users_proto_init() {
	if File_users_v1_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_v1_users_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoRequest); i {
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
		file_users_v1_users_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoResponse); i {
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
	file_users_v1_users_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_users_v1_users_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_v1_users_proto_goTypes,
		DependencyIndexes: file_users_v1_users_proto_depIdxs,
		EnumInfos:         file_users_v1_users_proto_enumTypes,
		MessageInfos:      file_users_v1_users_proto_msgTypes,
	}.Build()
	File_users_v1_users_proto = out.File
	file_users_v1_users_proto_rawDesc = nil
	file_users_v1_users_proto_goTypes = nil
	file_users_v1_users_proto_depIdxs = nil
}
