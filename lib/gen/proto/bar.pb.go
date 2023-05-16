// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/bar.proto

package proto

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

type MoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Money       *Money `protobuf:"bytes,1,opt,name=money,proto3" json:"money,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MoneyRequest) Reset() {
	*x = MoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoneyRequest) ProtoMessage() {}

func (x *MoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoneyRequest.ProtoReflect.Descriptor instead.
func (*MoneyRequest) Descriptor() ([]byte, []int) {
	return file_proto_bar_proto_rawDescGZIP(), []int{0}
}

func (x *MoneyRequest) GetMoney() *Money {
	if x != nil {
		return x.Money
	}
	return nil
}

func (x *MoneyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_bar_proto protoreflect.FileDescriptor

var file_proto_bar_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6b, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x62, 0x42, 0x08, 0x42, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x6e, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x50, 0x62, 0xca, 0x02, 0x02,
	0x50, 0x62, 0xe2, 0x02, 0x0e, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bar_proto_rawDescOnce sync.Once
	file_proto_bar_proto_rawDescData = file_proto_bar_proto_rawDesc
)

func file_proto_bar_proto_rawDescGZIP() []byte {
	file_proto_bar_proto_rawDescOnce.Do(func() {
		file_proto_bar_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bar_proto_rawDescData)
	})
	return file_proto_bar_proto_rawDescData
}

var file_proto_bar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_bar_proto_goTypes = []interface{}{
	(*MoneyRequest)(nil), // 0: pb.MoneyRequest
	(*Money)(nil),        // 1: pb.Money
}
var file_proto_bar_proto_depIdxs = []int32{
	1, // 0: pb.MoneyRequest.money:type_name -> pb.Money
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_bar_proto_init() }
func file_proto_bar_proto_init() {
	if File_proto_bar_proto != nil {
		return
	}
	file_proto_money_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_bar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoneyRequest); i {
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
			RawDescriptor: file_proto_bar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bar_proto_goTypes,
		DependencyIndexes: file_proto_bar_proto_depIdxs,
		MessageInfos:      file_proto_bar_proto_msgTypes,
	}.Build()
	File_proto_bar_proto = out.File
	file_proto_bar_proto_rawDesc = nil
	file_proto_bar_proto_goTypes = nil
	file_proto_bar_proto_depIdxs = nil
}