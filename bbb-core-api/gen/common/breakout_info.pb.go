// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: common/breakout_info.proto

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

type BreakoutInfo struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	IsBreakout      bool                   `protobuf:"varint,1,opt,name=is_breakout,json=isBreakout,proto3" json:"is_breakout,omitempty"`
	ParentMeetingId string                 `protobuf:"bytes,2,opt,name=parent_meeting_id,json=parentMeetingId,proto3" json:"parent_meeting_id,omitempty"`
	Sequence        int32                  `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	FreeJoin        bool                   `protobuf:"varint,4,opt,name=free_join,json=freeJoin,proto3" json:"free_join,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *BreakoutInfo) Reset() {
	*x = BreakoutInfo{}
	mi := &file_common_breakout_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BreakoutInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutInfo) ProtoMessage() {}

func (x *BreakoutInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_breakout_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutInfo.ProtoReflect.Descriptor instead.
func (*BreakoutInfo) Descriptor() ([]byte, []int) {
	return file_common_breakout_info_proto_rawDescGZIP(), []int{0}
}

func (x *BreakoutInfo) GetIsBreakout() bool {
	if x != nil {
		return x.IsBreakout
	}
	return false
}

func (x *BreakoutInfo) GetParentMeetingId() string {
	if x != nil {
		return x.ParentMeetingId
	}
	return ""
}

func (x *BreakoutInfo) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *BreakoutInfo) GetFreeJoin() bool {
	if x != nil {
		return x.FreeJoin
	}
	return false
}

var File_common_breakout_info_proto protoreflect.FileDescriptor

var file_common_breakout_info_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72,
	0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x6f, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_breakout_info_proto_rawDescOnce sync.Once
	file_common_breakout_info_proto_rawDescData = file_common_breakout_info_proto_rawDesc
)

func file_common_breakout_info_proto_rawDescGZIP() []byte {
	file_common_breakout_info_proto_rawDescOnce.Do(func() {
		file_common_breakout_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_breakout_info_proto_rawDescData)
	})
	return file_common_breakout_info_proto_rawDescData
}

var file_common_breakout_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_breakout_info_proto_goTypes = []any{
	(*BreakoutInfo)(nil), // 0: org.bigbluebutton.protos.BreakoutInfo
}
var file_common_breakout_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_breakout_info_proto_init() }
func file_common_breakout_info_proto_init() {
	if File_common_breakout_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_breakout_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_breakout_info_proto_goTypes,
		DependencyIndexes: file_common_breakout_info_proto_depIdxs,
		MessageInfos:      file_common_breakout_info_proto_msgTypes,
	}.Build()
	File_common_breakout_info_proto = out.File
	file_common_breakout_info_proto_rawDesc = nil
	file_common_breakout_info_proto_goTypes = nil
	file_common_breakout_info_proto_depIdxs = nil
}
