// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package textualv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/msg/textual/v1/textual.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_cosmos_msg_textual_v1_textual_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         11110009,
		Name:          "cosmos.msg.textual.v1.expert_custom_renderer",
		Tag:           "bytes,11110009,opt,name=expert_custom_renderer",
		Filename:      "cosmos/msg/textual/v1/textual.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// expert_custom_renderer is an informative identifier to reference the
	// algorithm used to generate the custom textual representation of the
	// protobuf message where this annotation is applied. We recommend to use a
	// short, versioned name as this identifier, e.g. "replace_with_username_v1".
	// We also recommend providing a human-readable description as protobuf
	// comments on this annotation, for example a short specification or a link
	// to the relevant documentation.
	//
	// Also see the section on Custom Message Renderers in ADR-050.
	//
	// optional string expert_custom_renderer = 11110009;
	E_ExpertCustomRenderer = &file_cosmos_msg_textual_v1_textual_proto_extTypes[0]
)

var File_cosmos_msg_textual_v1_textual_proto protoreflect.FileDescriptor

var file_cosmos_msg_textual_v1_textual_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x74, 0x65, 0x78,
	0x74, 0x75, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x58,
	0x0a, 0x16, 0x65, 0x78, 0x70, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf9, 0x8c, 0xa6, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x70, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0xd2, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x74, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64,
	0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x6d, 0x73, 0x67, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x54, 0xaa, 0x02,
	0x15, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x4d, 0x73, 0x67, 0x5c, 0x54, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x21, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x4d, 0x73, 0x67, 0x5c, 0x54, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x4d, 0x73, 0x67,
	0x3a, 0x3a, 0x54, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cosmos_msg_textual_v1_textual_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_cosmos_msg_textual_v1_textual_proto_depIdxs = []int32{
	0, // 0: cosmos.msg.textual.v1.expert_custom_renderer:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_msg_textual_v1_textual_proto_init() }
func file_cosmos_msg_textual_v1_textual_proto_init() {
	if File_cosmos_msg_textual_v1_textual_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_msg_textual_v1_textual_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_msg_textual_v1_textual_proto_goTypes,
		DependencyIndexes: file_cosmos_msg_textual_v1_textual_proto_depIdxs,
		ExtensionInfos:    file_cosmos_msg_textual_v1_textual_proto_extTypes,
	}.Build()
	File_cosmos_msg_textual_v1_textual_proto = out.File
	file_cosmos_msg_textual_v1_textual_proto_rawDesc = nil
	file_cosmos_msg_textual_v1_textual_proto_goTypes = nil
	file_cosmos_msg_textual_v1_textual_proto_depIdxs = nil
}
