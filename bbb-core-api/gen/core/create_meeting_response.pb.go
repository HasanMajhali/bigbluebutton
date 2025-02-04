// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: core/create_meeting_response.proto

package core

import (
	common "github.com/bigbluebutton/bigbluebutton/bbb-core-api/gen/common"
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

type CreateMeetingResponse struct {
	state              protoimpl.MessageState     `protogen:"open.v1"`
	CreatedMeetingInfo *common.CreatedMeetingInfo `protobuf:"bytes,1,opt,name=created_meeting_info,json=createdMeetingInfo,proto3" json:"created_meeting_info,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CreateMeetingResponse) Reset() {
	*x = CreateMeetingResponse{}
	mi := &file_core_create_meeting_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMeetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeetingResponse) ProtoMessage() {}

func (x *CreateMeetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_create_meeting_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeetingResponse.ProtoReflect.Descriptor instead.
func (*CreateMeetingResponse) Descriptor() ([]byte, []int) {
	return file_core_create_meeting_response_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMeetingResponse) GetCreatedMeetingInfo() *common.CreatedMeetingInfo {
	if x != nil {
		return x.CreatedMeetingInfo
	}
	return nil
}

var File_core_create_meeting_response_proto protoreflect.FileDescriptor

var file_core_create_meeting_response_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75,
	0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x21,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x77, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x14, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62,
	0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_core_create_meeting_response_proto_rawDescOnce sync.Once
	file_core_create_meeting_response_proto_rawDescData = file_core_create_meeting_response_proto_rawDesc
)

func file_core_create_meeting_response_proto_rawDescGZIP() []byte {
	file_core_create_meeting_response_proto_rawDescOnce.Do(func() {
		file_core_create_meeting_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_create_meeting_response_proto_rawDescData)
	})
	return file_core_create_meeting_response_proto_rawDescData
}

var file_core_create_meeting_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_create_meeting_response_proto_goTypes = []any{
	(*CreateMeetingResponse)(nil),     // 0: org.bigbluebutton.protos.CreateMeetingResponse
	(*common.CreatedMeetingInfo)(nil), // 1: org.bigbluebutton.protos.CreatedMeetingInfo
}
var file_core_create_meeting_response_proto_depIdxs = []int32{
	1, // 0: org.bigbluebutton.protos.CreateMeetingResponse.created_meeting_info:type_name -> org.bigbluebutton.protos.CreatedMeetingInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_create_meeting_response_proto_init() }
func file_core_create_meeting_response_proto_init() {
	if File_core_create_meeting_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_create_meeting_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_create_meeting_response_proto_goTypes,
		DependencyIndexes: file_core_create_meeting_response_proto_depIdxs,
		MessageInfos:      file_core_create_meeting_response_proto_msgTypes,
	}.Build()
	File_core_create_meeting_response_proto = out.File
	file_core_create_meeting_response_proto_rawDesc = nil
	file_core_create_meeting_response_proto_goTypes = nil
	file_core_create_meeting_response_proto_depIdxs = nil
}
