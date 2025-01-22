// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.1
// source: pkg/proto/asset/asset.proto

package asset

import (
	v1 "github.com/bazelbuild/remote-apis/build/bazel/remote/asset/v1"
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Asset_AssetType int32

const (
	Asset_BLOB      Asset_AssetType = 0
	Asset_DIRECTORY Asset_AssetType = 1
)

// Enum value maps for Asset_AssetType.
var (
	Asset_AssetType_name = map[int32]string{
		0: "BLOB",
		1: "DIRECTORY",
	}
	Asset_AssetType_value = map[string]int32{
		"BLOB":      0,
		"DIRECTORY": 1,
	}
)

func (x Asset_AssetType) Enum() *Asset_AssetType {
	p := new(Asset_AssetType)
	*p = x
	return p
}

func (x Asset_AssetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Asset_AssetType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_asset_asset_proto_enumTypes[0].Descriptor()
}

func (Asset_AssetType) Type() protoreflect.EnumType {
	return &file_pkg_proto_asset_asset_proto_enumTypes[0]
}

func (x Asset_AssetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Asset_AssetType.Descriptor instead.
func (Asset_AssetType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_asset_asset_proto_rawDescGZIP(), []int{1, 0}
}

type AssetReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uris       []string        `protobuf:"bytes,1,rep,name=uris,proto3" json:"uris,omitempty"`
	Qualifiers []*v1.Qualifier `protobuf:"bytes,2,rep,name=qualifiers,proto3" json:"qualifiers,omitempty"`
}

func (x *AssetReference) Reset() {
	*x = AssetReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_asset_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetReference) ProtoMessage() {}

func (x *AssetReference) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_asset_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetReference.ProtoReflect.Descriptor instead.
func (*AssetReference) Descriptor() ([]byte, []int) {
	return file_pkg_proto_asset_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AssetReference) GetUris() []string {
	if x != nil {
		return x.Uris
	}
	return nil
}

func (x *AssetReference) GetQualifiers() []*v1.Qualifier {
	if x != nil {
		return x.Qualifiers
	}
	return nil
}

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest      *v2.Digest             `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	ExpireAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Type        Asset_AssetType        `protobuf:"varint,4,opt,name=type,proto3,enum=buildbarn.asset.Asset_AssetType" json:"type,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_asset_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_asset_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_pkg_proto_asset_asset_proto_rawDescGZIP(), []int{1}
}

func (x *Asset) GetDigest() *v2.Digest {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *Asset) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *Asset) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Asset) GetType() Asset_AssetType {
	if x != nil {
		return x.Type
	}
	return Asset_BLOB
}

var File_pkg_proto_asset_asset_proto protoreflect.FileDescriptor

var file_pkg_proto_asset_asset_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x1a, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x69,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x69, 0x73, 0x12, 0x46, 0x0a,
	0x0a, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x9c, 0x02, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x3f, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x24,
	0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x4c, 0x4f, 0x42, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f,
	0x52, 0x59, 0x10, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_asset_asset_proto_rawDescOnce sync.Once
	file_pkg_proto_asset_asset_proto_rawDescData = file_pkg_proto_asset_asset_proto_rawDesc
)

func file_pkg_proto_asset_asset_proto_rawDescGZIP() []byte {
	file_pkg_proto_asset_asset_proto_rawDescOnce.Do(func() {
		file_pkg_proto_asset_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_asset_asset_proto_rawDescData)
	})
	return file_pkg_proto_asset_asset_proto_rawDescData
}

var file_pkg_proto_asset_asset_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_asset_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_asset_asset_proto_goTypes = []interface{}{
	(Asset_AssetType)(0),          // 0: buildbarn.asset.Asset.AssetType
	(*AssetReference)(nil),        // 1: buildbarn.asset.AssetReference
	(*Asset)(nil),                 // 2: buildbarn.asset.Asset
	(*v1.Qualifier)(nil),          // 3: build.bazel.remote.asset.v1.Qualifier
	(*v2.Digest)(nil),             // 4: build.bazel.remote.execution.v2.Digest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_pkg_proto_asset_asset_proto_depIdxs = []int32{
	3, // 0: buildbarn.asset.AssetReference.qualifiers:type_name -> build.bazel.remote.asset.v1.Qualifier
	4, // 1: buildbarn.asset.Asset.digest:type_name -> build.bazel.remote.execution.v2.Digest
	5, // 2: buildbarn.asset.Asset.expire_at:type_name -> google.protobuf.Timestamp
	5, // 3: buildbarn.asset.Asset.last_updated:type_name -> google.protobuf.Timestamp
	0, // 4: buildbarn.asset.Asset.type:type_name -> buildbarn.asset.Asset.AssetType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_asset_asset_proto_init() }
func file_pkg_proto_asset_asset_proto_init() {
	if File_pkg_proto_asset_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_asset_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetReference); i {
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
		file_pkg_proto_asset_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
			RawDescriptor: file_pkg_proto_asset_asset_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_asset_asset_proto_goTypes,
		DependencyIndexes: file_pkg_proto_asset_asset_proto_depIdxs,
		EnumInfos:         file_pkg_proto_asset_asset_proto_enumTypes,
		MessageInfos:      file_pkg_proto_asset_asset_proto_msgTypes,
	}.Build()
	File_pkg_proto_asset_asset_proto = out.File
	file_pkg_proto_asset_asset_proto_rawDesc = nil
	file_pkg_proto_asset_asset_proto_goTypes = nil
	file_pkg_proto_asset_asset_proto_depIdxs = nil
}
