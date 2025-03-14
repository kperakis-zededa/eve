// Copyright(c) 2023 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: info/patch_envelope.proto

package info

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

type EVE_PATCH_ENVELOPE_STATE int32

const (
	EVE_PATCH_ENVELOPE_STATE_PATCH_UNKOWN EVE_PATCH_ENVELOPE_STATE = 0
	// There is an error with config, during download,
	//
	//	verification failed, or decryption failure
	EVE_PATCH_ENVELOPE_STATE_PATCH_ERROR EVE_PATCH_ENVELOPE_STATE = 1
	// Configuration received but no downloads started
	EVE_PATCH_ENVELOPE_STATE_PATCH_RECEIVED EVE_PATCH_ENVELOPE_STATE = 2
	// Artifact/Volume download started
	// One or more of the artifacts are being downloaded
	EVE_PATCH_ENVELOPE_STATE_PATCH_DOWNLOADING EVE_PATCH_ENVELOPE_STATE = 3
	// All downloads finished, verified and added to
	// content tree
	EVE_PATCH_ENVELOPE_STATE_PATCH_DOWNLOADED EVE_PATCH_ENVELOPE_STATE = 4
	// Patch envelope ready for application instances
	// application instances will still not be
	// allowed to fetch the patch envelope contents
	EVE_PATCH_ENVELOPE_STATE_PATCH_READY EVE_PATCH_ENVELOPE_STATE = 5
	// Application instances are now allowed to fetch
	// contents
	EVE_PATCH_ENVELOPE_STATE_PATCH_ACTIVE EVE_PATCH_ENVELOPE_STATE = 6
)

// Enum value maps for EVE_PATCH_ENVELOPE_STATE.
var (
	EVE_PATCH_ENVELOPE_STATE_name = map[int32]string{
		0: "PATCH_UNKOWN",
		1: "PATCH_ERROR",
		2: "PATCH_RECEIVED",
		3: "PATCH_DOWNLOADING",
		4: "PATCH_DOWNLOADED",
		5: "PATCH_READY",
		6: "PATCH_ACTIVE",
	}
	EVE_PATCH_ENVELOPE_STATE_value = map[string]int32{
		"PATCH_UNKOWN":      0,
		"PATCH_ERROR":       1,
		"PATCH_RECEIVED":    2,
		"PATCH_DOWNLOADING": 3,
		"PATCH_DOWNLOADED":  4,
		"PATCH_READY":       5,
		"PATCH_ACTIVE":      6,
	}
)

func (x EVE_PATCH_ENVELOPE_STATE) Enum() *EVE_PATCH_ENVELOPE_STATE {
	p := new(EVE_PATCH_ENVELOPE_STATE)
	*p = x
	return p
}

func (x EVE_PATCH_ENVELOPE_STATE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVE_PATCH_ENVELOPE_STATE) Descriptor() protoreflect.EnumDescriptor {
	return file_info_patch_envelope_proto_enumTypes[0].Descriptor()
}

func (EVE_PATCH_ENVELOPE_STATE) Type() protoreflect.EnumType {
	return &file_info_patch_envelope_proto_enumTypes[0]
}

func (x EVE_PATCH_ENVELOPE_STATE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVE_PATCH_ENVELOPE_STATE.Descriptor instead.
func (EVE_PATCH_ENVELOPE_STATE) EnumDescriptor() ([]byte, []int) {
	return file_info_patch_envelope_proto_rawDescGZIP(), []int{0}
}

type ZInfoPatchEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id      string                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Version string                   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	State   EVE_PATCH_ENVELOPE_STATE `protobuf:"varint,4,opt,name=state,proto3,enum=org.lfedge.eve.info.EVE_PATCH_ENVELOPE_STATE" json:"state,omitempty"`
	// Size of PatchEnvelope in bytes
	Size   uint64   `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Errors []string `protobuf:"bytes,6,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ZInfoPatchEnvelope) Reset() {
	*x = ZInfoPatchEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_patch_envelope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZInfoPatchEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZInfoPatchEnvelope) ProtoMessage() {}

func (x *ZInfoPatchEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_info_patch_envelope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZInfoPatchEnvelope.ProtoReflect.Descriptor instead.
func (*ZInfoPatchEnvelope) Descriptor() ([]byte, []int) {
	return file_info_patch_envelope_proto_rawDescGZIP(), []int{0}
}

func (x *ZInfoPatchEnvelope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ZInfoPatchEnvelope) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ZInfoPatchEnvelope) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ZInfoPatchEnvelope) GetState() EVE_PATCH_ENVELOPE_STATE {
	if x != nil {
		return x.State
	}
	return EVE_PATCH_ENVELOPE_STATE_PATCH_UNKOWN
}

func (x *ZInfoPatchEnvelope) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ZInfoPatchEnvelope) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_info_patch_envelope_proto protoreflect.FileDescriptor

var file_info_patch_envelope_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0xc3, 0x01, 0x0a, 0x12, 0x5a, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x56, 0x45, 0x5f, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0xa1, 0x01, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x5f, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x4e, 0x4b,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c,
	0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x06, 0x42, 0x39, 0x0a, 0x13, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66,
	0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_patch_envelope_proto_rawDescOnce sync.Once
	file_info_patch_envelope_proto_rawDescData = file_info_patch_envelope_proto_rawDesc
)

func file_info_patch_envelope_proto_rawDescGZIP() []byte {
	file_info_patch_envelope_proto_rawDescOnce.Do(func() {
		file_info_patch_envelope_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_patch_envelope_proto_rawDescData)
	})
	return file_info_patch_envelope_proto_rawDescData
}

var file_info_patch_envelope_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_info_patch_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_info_patch_envelope_proto_goTypes = []interface{}{
	(EVE_PATCH_ENVELOPE_STATE)(0), // 0: org.lfedge.eve.info.EVE_PATCH_ENVELOPE_STATE
	(*ZInfoPatchEnvelope)(nil),    // 1: org.lfedge.eve.info.ZInfoPatchEnvelope
}
var file_info_patch_envelope_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.info.ZInfoPatchEnvelope.state:type_name -> org.lfedge.eve.info.EVE_PATCH_ENVELOPE_STATE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_info_patch_envelope_proto_init() }
func file_info_patch_envelope_proto_init() {
	if File_info_patch_envelope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_patch_envelope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZInfoPatchEnvelope); i {
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
			RawDescriptor: file_info_patch_envelope_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_info_patch_envelope_proto_goTypes,
		DependencyIndexes: file_info_patch_envelope_proto_depIdxs,
		EnumInfos:         file_info_patch_envelope_proto_enumTypes,
		MessageInfos:      file_info_patch_envelope_proto_msgTypes,
	}.Build()
	File_info_patch_envelope_proto = out.File
	file_info_patch_envelope_proto_rawDesc = nil
	file_info_patch_envelope_proto_goTypes = nil
	file_info_patch_envelope_proto_depIdxs = nil
}
