// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: common/user.proto

package common

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

type User struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName        string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Role            string                 `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	IsPresenter     bool                   `protobuf:"varint,4,opt,name=is_presenter,json=isPresenter,proto3" json:"is_presenter,omitempty"`
	IsListeningOnly bool                   `protobuf:"varint,5,opt,name=is_listening_only,json=isListeningOnly,proto3" json:"is_listening_only,omitempty"`
	HasJoinedVoice  bool                   `protobuf:"varint,6,opt,name=has_joined_voice,json=hasJoinedVoice,proto3" json:"has_joined_voice,omitempty"`
	HasVideo        bool                   `protobuf:"varint,7,opt,name=has_video,json=hasVideo,proto3" json:"has_video,omitempty"`
	ClientType      string                 `protobuf:"bytes,8,opt,name=clientType,proto3" json:"clientType,omitempty"`
	CustomData      map[string]string      `protobuf:"bytes,9,rep,name=custom_data,json=customData,proto3" json:"custom_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_common_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetIsPresenter() bool {
	if x != nil {
		return x.IsPresenter
	}
	return false
}

func (x *User) GetIsListeningOnly() bool {
	if x != nil {
		return x.IsListeningOnly
	}
	return false
}

func (x *User) GetHasJoinedVoice() bool {
	if x != nil {
		return x.HasJoinedVoice
	}
	return false
}

func (x *User) GetHasVideo() bool {
	if x != nil {
		return x.HasVideo
	}
	return false
}

func (x *User) GetClientType() string {
	if x != nil {
		return x.ClientType
	}
	return ""
}

func (x *User) GetCustomData() map[string]string {
	if x != nil {
		return x.CustomData
	}
	return nil
}

var File_common_user_proto protoreflect.FileDescriptor

var file_common_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65,
	0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x96, 0x03,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x69, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x6c, 0x79, 0x12,
	0x28, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x4a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73,
	0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61,
	0x73, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_user_proto_rawDescOnce sync.Once
	file_common_user_proto_rawDescData = file_common_user_proto_rawDesc
)

func file_common_user_proto_rawDescGZIP() []byte {
	file_common_user_proto_rawDescOnce.Do(func() {
		file_common_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_user_proto_rawDescData)
	})
	return file_common_user_proto_rawDescData
}

var file_common_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_user_proto_goTypes = []any{
	(*User)(nil), // 0: org.bigbluebutton.protos.User
	nil,          // 1: org.bigbluebutton.protos.User.CustomDataEntry
}
var file_common_user_proto_depIdxs = []int32{
	1, // 0: org.bigbluebutton.protos.User.custom_data:type_name -> org.bigbluebutton.protos.User.CustomDataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_user_proto_init() }
func file_common_user_proto_init() {
	if File_common_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_user_proto_goTypes,
		DependencyIndexes: file_common_user_proto_depIdxs,
		MessageInfos:      file_common_user_proto_msgTypes,
	}.Build()
	File_common_user_proto = out.File
	file_common_user_proto_rawDesc = nil
	file_common_user_proto_goTypes = nil
	file_common_user_proto_depIdxs = nil
}
