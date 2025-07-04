// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: some_service.proto

package someservicev1

import (
	_ "github.com/hughbliss/my_protobuf/go/pkg/acman/gen/permissions"
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

type SomeExampleMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SomeText field to provide some text
	SomeText string `protobuf:"bytes,1,opt,name=some_text,json=someText,proto3" json:"some_text,omitempty"`
	// SomeUnsignedIntegerValue field to provide some integer with unsigned value
	SomeUnsignedIntegerValue uint32 `protobuf:"varint,2,opt,name=some_unsigned_integer_value,json=someUnsignedIntegerValue,proto3" json:"some_unsigned_integer_value,omitempty"`
}

func (x *SomeExampleMethodRequest) Reset() {
	*x = SomeExampleMethodRequest{}
	mi := &file_some_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeExampleMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeExampleMethodRequest) ProtoMessage() {}

func (x *SomeExampleMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_some_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeExampleMethodRequest.ProtoReflect.Descriptor instead.
func (*SomeExampleMethodRequest) Descriptor() ([]byte, []int) {
	return file_some_service_proto_rawDescGZIP(), []int{0}
}

func (x *SomeExampleMethodRequest) GetSomeText() string {
	if x != nil {
		return x.SomeText
	}
	return ""
}

func (x *SomeExampleMethodRequest) GetSomeUnsignedIntegerValue() uint32 {
	if x != nil {
		return x.SomeUnsignedIntegerValue
	}
	return 0
}

type SomeExampleMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SomeResponse field provides some text will returned from service
	SomeResponse string `protobuf:"bytes,1,opt,name=some_response,json=someResponse,proto3" json:"some_response,omitempty"`
}

func (x *SomeExampleMethodResponse) Reset() {
	*x = SomeExampleMethodResponse{}
	mi := &file_some_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeExampleMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeExampleMethodResponse) ProtoMessage() {}

func (x *SomeExampleMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_some_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeExampleMethodResponse.ProtoReflect.Descriptor instead.
func (*SomeExampleMethodResponse) Descriptor() ([]byte, []int) {
	return file_some_service_proto_rawDescGZIP(), []int{1}
}

func (x *SomeExampleMethodResponse) GetSomeResponse() string {
	if x != nil {
		return x.SomeResponse
	}
	return ""
}

var File_some_service_proto protoreflect.FileDescriptor

var file_some_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x18,
	0x53, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6d,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x73, 0x6f, 0x6d, 0x65,
	0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x19, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaa, 0x03, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc2, 0x01, 0x0a, 0x11, 0x53, 0x6f, 0x6d, 0x65, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x2e, 0x73,
	0x6f, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x56, 0x8a, 0xb5, 0x18, 0x32, 0x0a, 0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x22, 0xd0, 0x94, 0xd0, 0xbe, 0xd1, 0x81, 0xd1,
	0x82, 0xd1, 0x83, 0xd0, 0xbf, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd1,
	0x80, 0xd0, 0xb8, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0xd5, 0x01, 0x0a, 0x14,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x8a, 0xb5, 0x18,
	0x3a, 0x0a, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x61, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x22, 0xd0, 0x94, 0xd0, 0xbe, 0xd1, 0x81, 0xd1,
	0x82, 0xd1, 0x83, 0xd0, 0xbf, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd1,
	0x80, 0xd0, 0xb8, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x75, 0x67, 0x68, 0x62, 0x6c, 0x69, 0x73, 0x73, 0x2f, 0x6d, 0x79, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_some_service_proto_rawDescOnce sync.Once
	file_some_service_proto_rawDescData = file_some_service_proto_rawDesc
)

func file_some_service_proto_rawDescGZIP() []byte {
	file_some_service_proto_rawDescOnce.Do(func() {
		file_some_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_some_service_proto_rawDescData)
	})
	return file_some_service_proto_rawDescData
}

var file_some_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_some_service_proto_goTypes = []any{
	(*SomeExampleMethodRequest)(nil),  // 0: some.service.v1.SomeExampleMethodRequest
	(*SomeExampleMethodResponse)(nil), // 1: some.service.v1.SomeExampleMethodResponse
}
var file_some_service_proto_depIdxs = []int32{
	0, // 0: some.service.v1.SomeService.SomeExampleMethod:input_type -> some.service.v1.SomeExampleMethodRequest
	0, // 1: some.service.v1.SomeService.AnotherExampleMethod:input_type -> some.service.v1.SomeExampleMethodRequest
	1, // 2: some.service.v1.SomeService.SomeExampleMethod:output_type -> some.service.v1.SomeExampleMethodResponse
	1, // 3: some.service.v1.SomeService.AnotherExampleMethod:output_type -> some.service.v1.SomeExampleMethodResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_some_service_proto_init() }
func file_some_service_proto_init() {
	if File_some_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_some_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_some_service_proto_goTypes,
		DependencyIndexes: file_some_service_proto_depIdxs,
		MessageInfos:      file_some_service_proto_msgTypes,
	}.Build()
	File_some_service_proto = out.File
	file_some_service_proto_rawDesc = nil
	file_some_service_proto_goTypes = nil
	file_some_service_proto_depIdxs = nil
}
