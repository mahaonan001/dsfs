// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: proto/noding.proto

package noding

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

type Hearting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Hearting) Reset() {
	*x = Hearting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_noding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hearting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hearting) ProtoMessage() {}

func (x *Hearting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_noding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hearting.ProtoReflect.Descriptor instead.
func (*Hearting) Descriptor() ([]byte, []int) {
	return file_proto_noding_proto_rawDescGZIP(), []int{0}
}

func (x *Hearting) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type HeartingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HeartingResponse) Reset() {
	*x = HeartingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_noding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartingResponse) ProtoMessage() {}

func (x *HeartingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_noding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartingResponse.ProtoReflect.Descriptor instead.
func (*HeartingResponse) Descriptor() ([]byte, []int) {
	return file_proto_noding_proto_rawDescGZIP(), []int{1}
}

func (x *HeartingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type WakeUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files  []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Number int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *WakeUp) Reset() {
	*x = WakeUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_noding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WakeUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WakeUp) ProtoMessage() {}

func (x *WakeUp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_noding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WakeUp.ProtoReflect.Descriptor instead.
func (*WakeUp) Descriptor() ([]byte, []int) {
	return file_proto_noding_proto_rawDescGZIP(), []int{2}
}

func (x *WakeUp) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *WakeUp) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type WakeUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WakeUpResponse) Reset() {
	*x = WakeUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_noding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WakeUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WakeUpResponse) ProtoMessage() {}

func (x *WakeUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_noding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WakeUpResponse.ProtoReflect.Descriptor instead.
func (*WakeUpResponse) Descriptor() ([]byte, []int) {
	return file_proto_noding_proto_rawDescGZIP(), []int{3}
}

func (x *WakeUpResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_noding_proto protoreflect.FileDescriptor

var file_proto_noding_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x1a, 0x0a, 0x08,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x2c, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x57, 0x61, 0x6b, 0x65, 0x55, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2a,
	0x0a, 0x0e, 0x57, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x71, 0x0a, 0x06, 0x4e, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x05, 0x48, 0x65, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e,
	0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x57,
	0x61, 0x6b, 0x65, 0x12, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x6b,
	0x65, 0x55, 0x70, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x6b,
	0x65, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_noding_proto_rawDescOnce sync.Once
	file_proto_noding_proto_rawDescData = file_proto_noding_proto_rawDesc
)

func file_proto_noding_proto_rawDescGZIP() []byte {
	file_proto_noding_proto_rawDescOnce.Do(func() {
		file_proto_noding_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_noding_proto_rawDescData)
	})
	return file_proto_noding_proto_rawDescData
}

var file_proto_noding_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_noding_proto_goTypes = []any{
	(*Hearting)(nil),         // 0: noding.Hearting
	(*HeartingResponse)(nil), // 1: noding.HeartingResponse
	(*WakeUp)(nil),           // 2: noding.WakeUp
	(*WakeUpResponse)(nil),   // 3: noding.WakeUpResponse
}
var file_proto_noding_proto_depIdxs = []int32{
	0, // 0: noding.Noding.Heart:input_type -> noding.Hearting
	2, // 1: noding.Noding.Wake:input_type -> noding.WakeUp
	1, // 2: noding.Noding.Heart:output_type -> noding.HeartingResponse
	3, // 3: noding.Noding.Wake:output_type -> noding.WakeUpResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_noding_proto_init() }
func file_proto_noding_proto_init() {
	if File_proto_noding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_noding_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Hearting); i {
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
		file_proto_noding_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HeartingResponse); i {
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
		file_proto_noding_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WakeUp); i {
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
		file_proto_noding_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WakeUpResponse); i {
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
			RawDescriptor: file_proto_noding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_noding_proto_goTypes,
		DependencyIndexes: file_proto_noding_proto_depIdxs,
		MessageInfos:      file_proto_noding_proto_msgTypes,
	}.Build()
	File_proto_noding_proto = out.File
	file_proto_noding_proto_rawDesc = nil
	file_proto_noding_proto_goTypes = nil
	file_proto_noding_proto_depIdxs = nil
}
