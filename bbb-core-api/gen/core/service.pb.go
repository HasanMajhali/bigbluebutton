// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: core/service.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_core_service_proto protoreflect.FileDescriptor

var file_core_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75,
	0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x22,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xcb, 0x05, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x75, 0x0a, 0x10, 0x69, 0x73, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75,
	0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c,
	0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62,
	0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67,
	0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62,
	0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c,
	0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x32, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x7b,
	0x0a, 0x12, 0x69, 0x73, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49,
	0x6e, 0x55, 0x73, 0x65, 0x12, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c,
	0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x69,
	0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e,
	0x55, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6c, 0x75, 0x65, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_core_service_proto_goTypes = []any{
	(*MeetingRunningRequest)(nil),    // 0: org.bigbluebutton.protos.MeetingRunningRequest
	(*MeetingInfoRequest)(nil),       // 1: org.bigbluebutton.protos.MeetingInfoRequest
	(*ListMeetingsRequest)(nil),      // 2: org.bigbluebutton.protos.ListMeetingsRequest
	(*GetMeetingsStreamRequest)(nil), // 3: org.bigbluebutton.protos.GetMeetingsStreamRequest
	(*VoiceBridgeInUseRequest)(nil),  // 4: org.bigbluebutton.protos.VoiceBridgeInUseRequest
	(*CreateMeetingRequest)(nil),     // 5: org.bigbluebutton.protos.CreateMeetingRequest
	(*MeetingRunningResponse)(nil),   // 6: org.bigbluebutton.protos.MeetingRunningResponse
	(*MeetingInfoResponse)(nil),      // 7: org.bigbluebutton.protos.MeetingInfoResponse
	(*ListMeetingsResponse)(nil),     // 8: org.bigbluebutton.protos.ListMeetingsResponse
	(*VoiceBridgeInUseResponse)(nil), // 9: org.bigbluebutton.protos.VoiceBridgeInUseResponse
	(*CreateMeetingResponse)(nil),    // 10: org.bigbluebutton.protos.CreateMeetingResponse
}
var file_core_service_proto_depIdxs = []int32{
	0,  // 0: org.bigbluebutton.protos.CoreService.isMeetingRunning:input_type -> org.bigbluebutton.protos.MeetingRunningRequest
	1,  // 1: org.bigbluebutton.protos.CoreService.getMeetingInfo:input_type -> org.bigbluebutton.protos.MeetingInfoRequest
	2,  // 2: org.bigbluebutton.protos.CoreService.listMeetings:input_type -> org.bigbluebutton.protos.ListMeetingsRequest
	3,  // 3: org.bigbluebutton.protos.CoreService.getMeetingsStream:input_type -> org.bigbluebutton.protos.GetMeetingsStreamRequest
	4,  // 4: org.bigbluebutton.protos.CoreService.isVoiceBridgeInUse:input_type -> org.bigbluebutton.protos.VoiceBridgeInUseRequest
	5,  // 5: org.bigbluebutton.protos.CoreService.createMeeting:input_type -> org.bigbluebutton.protos.CreateMeetingRequest
	6,  // 6: org.bigbluebutton.protos.CoreService.isMeetingRunning:output_type -> org.bigbluebutton.protos.MeetingRunningResponse
	7,  // 7: org.bigbluebutton.protos.CoreService.getMeetingInfo:output_type -> org.bigbluebutton.protos.MeetingInfoResponse
	8,  // 8: org.bigbluebutton.protos.CoreService.listMeetings:output_type -> org.bigbluebutton.protos.ListMeetingsResponse
	7,  // 9: org.bigbluebutton.protos.CoreService.getMeetingsStream:output_type -> org.bigbluebutton.protos.MeetingInfoResponse
	9,  // 10: org.bigbluebutton.protos.CoreService.isVoiceBridgeInUse:output_type -> org.bigbluebutton.protos.VoiceBridgeInUseResponse
	10, // 11: org.bigbluebutton.protos.CoreService.createMeeting:output_type -> org.bigbluebutton.protos.CreateMeetingResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_service_proto_init() }
func file_core_service_proto_init() {
	if File_core_service_proto != nil {
		return
	}
	file_core_meeting_running_request_proto_init()
	file_core_meeting_running_response_proto_init()
	file_core_meeting_info_request_proto_init()
	file_core_meeting_info_response_proto_init()
	file_core_list_meetings_request_proto_init()
	file_core_list_meetings_response_proto_init()
	file_core_get_meetings_stream_request_proto_init()
	file_core_voice_bridge_in_use_request_proto_init()
	file_core_voice_bridge_in_use_response_proto_init()
	file_core_create_meeting_request_proto_init()
	file_core_create_meeting_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_service_proto_goTypes,
		DependencyIndexes: file_core_service_proto_depIdxs,
	}.Build()
	File_core_service_proto = out.File
	file_core_service_proto_rawDesc = nil
	file_core_service_proto_goTypes = nil
	file_core_service_proto_depIdxs = nil
}
