// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: common/duration_settings.proto

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

type DurationSettings struct {
	state                              protoimpl.MessageState `protogen:"open.v1"`
	Duration                           int32                  `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	CreateTime                         int64                  `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateDate                         string                 `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	MeetingExpNoUserJoinedInMin        int32                  `protobuf:"varint,4,opt,name=meeting_exp_no_user_joined_in_min,json=meetingExpNoUserJoinedInMin,proto3" json:"meeting_exp_no_user_joined_in_min,omitempty"`
	MeetingExpLastUserLeftInMin        int32                  `protobuf:"varint,5,opt,name=meeting_exp_last_user_left_in_min,json=meetingExpLastUserLeftInMin,proto3" json:"meeting_exp_last_user_left_in_min,omitempty"`
	UserInactivityInspectTimeInMin     int32                  `protobuf:"varint,6,opt,name=user_inactivity_inspect_time_in_min,json=userInactivityInspectTimeInMin,proto3" json:"user_inactivity_inspect_time_in_min,omitempty"`
	UserInactivityThresholdInMin       int32                  `protobuf:"varint,7,opt,name=user_inactivity_threshold_in_min,json=userInactivityThresholdInMin,proto3" json:"user_inactivity_threshold_in_min,omitempty"`
	UserActivitySignResponseDelayInMin int32                  `protobuf:"varint,8,opt,name=user_activity_sign_response_delay_in_min,json=userActivitySignResponseDelayInMin,proto3" json:"user_activity_sign_response_delay_in_min,omitempty"`
	EndWhenNoMod                       bool                   `protobuf:"varint,9,opt,name=end_when_no_mod,json=endWhenNoMod,proto3" json:"end_when_no_mod,omitempty"`
	EndWhenNoModDelayInMin             int32                  `protobuf:"varint,10,opt,name=end_when_no_mod_delay_in_min,json=endWhenNoModDelayInMin,proto3" json:"end_when_no_mod_delay_in_min,omitempty"`
	LearningDashboardCleanupDelayInMin int32                  `protobuf:"varint,11,opt,name=learning_dashboard_cleanup_delay_in_min,json=learningDashboardCleanupDelayInMin,proto3" json:"learning_dashboard_cleanup_delay_in_min,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *DurationSettings) Reset() {
	*x = DurationSettings{}
	mi := &file_common_duration_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationSettings) ProtoMessage() {}

func (x *DurationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_common_duration_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationSettings.ProtoReflect.Descriptor instead.
func (*DurationSettings) Descriptor() ([]byte, []int) {
	return file_common_duration_settings_proto_rawDescGZIP(), []int{0}
}

func (x *DurationSettings) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *DurationSettings) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DurationSettings) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *DurationSettings) GetMeetingExpNoUserJoinedInMin() int32 {
	if x != nil {
		return x.MeetingExpNoUserJoinedInMin
	}
	return 0
}

func (x *DurationSettings) GetMeetingExpLastUserLeftInMin() int32 {
	if x != nil {
		return x.MeetingExpLastUserLeftInMin
	}
	return 0
}

func (x *DurationSettings) GetUserInactivityInspectTimeInMin() int32 {
	if x != nil {
		return x.UserInactivityInspectTimeInMin
	}
	return 0
}

func (x *DurationSettings) GetUserInactivityThresholdInMin() int32 {
	if x != nil {
		return x.UserInactivityThresholdInMin
	}
	return 0
}

func (x *DurationSettings) GetUserActivitySignResponseDelayInMin() int32 {
	if x != nil {
		return x.UserActivitySignResponseDelayInMin
	}
	return 0
}

func (x *DurationSettings) GetEndWhenNoMod() bool {
	if x != nil {
		return x.EndWhenNoMod
	}
	return false
}

func (x *DurationSettings) GetEndWhenNoModDelayInMin() int32 {
	if x != nil {
		return x.EndWhenNoModDelayInMin
	}
	return 0
}

func (x *DurationSettings) GetLearningDashboardCleanupDelayInMin() int32 {
	if x != nil {
		return x.LearningDashboardCleanupDelayInMin
	}
	return 0
}

var File_common_duration_settings_proto protoreflect.FileDescriptor

var file_common_duration_settings_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xa5, 0x05, 0x0a, 0x10, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a,
	0x21, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x5f, 0x6e, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x45, 0x78, 0x70, 0x4e, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64,
	0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x21, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x78, 0x70, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c,
	0x65, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x1b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x4c, 0x61, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x4b, 0x0a,
	0x23, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e,
	0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1e, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x4d,
	0x69, 0x6e, 0x12, 0x54, 0x0a, 0x28, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x22, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f,
	0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6e, 0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x57, 0x68, 0x65, 0x6e, 0x4e, 0x6f, 0x4d, 0x6f, 0x64, 0x12,
	0x3c, 0x0a, 0x1c, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6e, 0x6f, 0x5f, 0x6d,
	0x6f, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x65, 0x6e, 0x64, 0x57, 0x68, 0x65, 0x6e, 0x4e, 0x6f,
	0x4d, 0x6f, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x53, 0x0a,
	0x27, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x22,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x4d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_duration_settings_proto_rawDescOnce sync.Once
	file_common_duration_settings_proto_rawDescData = file_common_duration_settings_proto_rawDesc
)

func file_common_duration_settings_proto_rawDescGZIP() []byte {
	file_common_duration_settings_proto_rawDescOnce.Do(func() {
		file_common_duration_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_duration_settings_proto_rawDescData)
	})
	return file_common_duration_settings_proto_rawDescData
}

var file_common_duration_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_duration_settings_proto_goTypes = []any{
	(*DurationSettings)(nil), // 0: org.bigbluebutton.protos.DurationSettings
}
var file_common_duration_settings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_duration_settings_proto_init() }
func file_common_duration_settings_proto_init() {
	if File_common_duration_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_duration_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_duration_settings_proto_goTypes,
		DependencyIndexes: file_common_duration_settings_proto_depIdxs,
		MessageInfos:      file_common_duration_settings_proto_msgTypes,
	}.Build()
	File_common_duration_settings_proto = out.File
	file_common_duration_settings_proto_rawDesc = nil
	file_common_duration_settings_proto_goTypes = nil
	file_common_duration_settings_proto_depIdxs = nil
}
