// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/foo.proto

package proto

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

var File_proto_foo_proto protoreflect.FileDescriptor

var file_proto_foo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x3d, 0x42, 0x08, 0x46, 0x6f, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x6e, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2d, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var file_proto_foo_proto_goTypes = []interface{}{}
var file_proto_foo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_foo_proto_init() }
func file_proto_foo_proto_init() {
	if File_proto_foo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_foo_proto_goTypes,
		DependencyIndexes: file_proto_foo_proto_depIdxs,
	}.Build()
	File_proto_foo_proto = out.File
	file_proto_foo_proto_rawDesc = nil
	file_proto_foo_proto_goTypes = nil
	file_proto_foo_proto_depIdxs = nil
}
