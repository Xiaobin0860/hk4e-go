// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: HomeVerifyData.proto

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

type HomeVerifyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aid             string                        `protobuf:"bytes,7,opt,name=aid,proto3" json:"aid,omitempty"`
	Timestamp       uint32                        `protobuf:"fixed32,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Uid             uint32                        `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	ArrangementData *HomeSceneArrangementMuipData `protobuf:"bytes,9,opt,name=arrangement_data,json=arrangementData,proto3" json:"arrangement_data,omitempty"`
	Region          string                        `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Token           string                        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	HomeInfo        *HomeVerifySceneData          `protobuf:"bytes,6,opt,name=home_info,json=homeInfo,proto3" json:"home_info,omitempty"`
	Lang            LanguageType                  `protobuf:"varint,8,opt,name=lang,proto3,enum=proto.LanguageType" json:"lang,omitempty"`
}

func (x *HomeVerifyData) Reset() {
	*x = HomeVerifyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeVerifyData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeVerifyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeVerifyData) ProtoMessage() {}

func (x *HomeVerifyData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeVerifyData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeVerifyData.ProtoReflect.Descriptor instead.
func (*HomeVerifyData) Descriptor() ([]byte, []int) {
	return file_HomeVerifyData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeVerifyData) GetAid() string {
	if x != nil {
		return x.Aid
	}
	return ""
}

func (x *HomeVerifyData) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HomeVerifyData) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *HomeVerifyData) GetArrangementData() *HomeSceneArrangementMuipData {
	if x != nil {
		return x.ArrangementData
	}
	return nil
}

func (x *HomeVerifyData) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HomeVerifyData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *HomeVerifyData) GetHomeInfo() *HomeVerifySceneData {
	if x != nil {
		return x.HomeInfo
	}
	return nil
}

func (x *HomeVerifyData) GetLang() LanguageType {
	if x != nil {
		return x.Lang
	}
	return LanguageType_LANGUAGE_TYPE_NONE
}

var File_HomeVerifyData_proto protoreflect.FileDescriptor

var file_HomeVerifyData_proto_rawDesc = []byte{
	0x0a, 0x14, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x48,
	0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb2, 0x02, 0x0a, 0x0e, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x07, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x10, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeVerifyData_proto_rawDescOnce sync.Once
	file_HomeVerifyData_proto_rawDescData = file_HomeVerifyData_proto_rawDesc
)

func file_HomeVerifyData_proto_rawDescGZIP() []byte {
	file_HomeVerifyData_proto_rawDescOnce.Do(func() {
		file_HomeVerifyData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeVerifyData_proto_rawDescData)
	})
	return file_HomeVerifyData_proto_rawDescData
}

var file_HomeVerifyData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeVerifyData_proto_goTypes = []interface{}{
	(*HomeVerifyData)(nil),               // 0: proto.HomeVerifyData
	(*HomeSceneArrangementMuipData)(nil), // 1: proto.HomeSceneArrangementMuipData
	(*HomeVerifySceneData)(nil),          // 2: proto.HomeVerifySceneData
	(LanguageType)(0),                    // 3: proto.LanguageType
}
var file_HomeVerifyData_proto_depIdxs = []int32{
	1, // 0: proto.HomeVerifyData.arrangement_data:type_name -> proto.HomeSceneArrangementMuipData
	2, // 1: proto.HomeVerifyData.home_info:type_name -> proto.HomeVerifySceneData
	3, // 2: proto.HomeVerifyData.lang:type_name -> proto.LanguageType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_HomeVerifyData_proto_init() }
func file_HomeVerifyData_proto_init() {
	if File_HomeVerifyData_proto != nil {
		return
	}
	file_HomeSceneArrangementMuipData_proto_init()
	file_HomeVerifySceneData_proto_init()
	file_LanguageType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeVerifyData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeVerifyData); i {
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
			RawDescriptor: file_HomeVerifyData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeVerifyData_proto_goTypes,
		DependencyIndexes: file_HomeVerifyData_proto_depIdxs,
		MessageInfos:      file_HomeVerifyData_proto_msgTypes,
	}.Build()
	File_HomeVerifyData_proto = out.File
	file_HomeVerifyData_proto_rawDesc = nil
	file_HomeVerifyData_proto_goTypes = nil
	file_HomeVerifyData_proto_depIdxs = nil
}