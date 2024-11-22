// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: protos/anymodule/anymodule.proto

package protos

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

type EnumStatus int32

const (
	EnumStatus_FAIL      EnumStatus = 0
	EnumStatus_SUCCESS   EnumStatus = 1
	EnumStatus_BLANK     EnumStatus = 2
	EnumStatus_LAST_PAGE EnumStatus = 3
)

// Enum value maps for EnumStatus.
var (
	EnumStatus_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
		2: "BLANK",
		3: "LAST_PAGE",
	}
	EnumStatus_value = map[string]int32{
		"FAIL":      0,
		"SUCCESS":   1,
		"BLANK":     2,
		"LAST_PAGE": 3,
	}
)

func (x EnumStatus) Enum() *EnumStatus {
	p := new(EnumStatus)
	*p = x
	return p
}

func (x EnumStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_anymodule_anymodule_proto_enumTypes[0].Descriptor()
}

func (EnumStatus) Type() protoreflect.EnumType {
	return &file_protos_anymodule_anymodule_proto_enumTypes[0]
}

func (x EnumStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumStatus.Descriptor instead.
func (EnumStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{0}
}

type IngestRequestByRandId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandId string `protobuf:"bytes,1,opt,name=RandId,proto3" json:"RandId,omitempty"`
}

func (x *IngestRequestByRandId) Reset() {
	*x = IngestRequestByRandId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_anymodule_anymodule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestRequestByRandId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestRequestByRandId) ProtoMessage() {}

func (x *IngestRequestByRandId) ProtoReflect() protoreflect.Message {
	mi := &file_protos_anymodule_anymodule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestRequestByRandId.ProtoReflect.Descriptor instead.
func (*IngestRequestByRandId) Descriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{0}
}

func (x *IngestRequestByRandId) GetRandId() string {
	if x != nil {
		return x.RandId
	}
	return ""
}

type IngestRequestByUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *IngestRequestByUUID) Reset() {
	*x = IngestRequestByUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_anymodule_anymodule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestRequestByUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestRequestByUUID) ProtoMessage() {}

func (x *IngestRequestByUUID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_anymodule_anymodule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestRequestByUUID.ProtoReflect.Descriptor instead.
func (*IngestRequestByUUID) Descriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{1}
}

func (x *IngestRequestByUUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *DeleteAllRequest) Reset() {
	*x = DeleteAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_anymodule_anymodule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllRequest) ProtoMessage() {}

func (x *DeleteAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_anymodule_anymodule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllRequest) Descriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAllRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type IngestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyUUID         string `protobuf:"bytes,1,opt,name=anyUUID,proto3" json:"anyUUID,omitempty"`
	RetrievedLength int64  `protobuf:"varint,2,opt,name=RetrievedLength,proto3" json:"RetrievedLength,omitempty"`
	LastObjectIdHex string `protobuf:"bytes,3,opt,name=LastObjectIdHex,proto3" json:"LastObjectIdHex,omitempty"`
	ValidLastUUID   string `protobuf:"bytes,4,opt,name=ValidLastUUID,proto3" json:"ValidLastUUID,omitempty"`
}

func (x *IngestRequest) Reset() {
	*x = IngestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_anymodule_anymodule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestRequest) ProtoMessage() {}

func (x *IngestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_anymodule_anymodule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestRequest.ProtoReflect.Descriptor instead.
func (*IngestRequest) Descriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{3}
}

func (x *IngestRequest) GetAnyUUID() string {
	if x != nil {
		return x.AnyUUID
	}
	return ""
}

func (x *IngestRequest) GetRetrievedLength() int64 {
	if x != nil {
		return x.RetrievedLength
	}
	return 0
}

func (x *IngestRequest) GetLastObjectIdHex() string {
	if x != nil {
		return x.LastObjectIdHex
	}
	return ""
}

func (x *IngestRequest) GetValidLastUUID() string {
	if x != nil {
		return x.ValidLastUUID
	}
	return ""
}

type IngestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *IngestStatus `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *IngestStatus) Reset() {
	*x = IngestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_anymodule_anymodule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestStatus) ProtoMessage() {}

func (x *IngestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protos_anymodule_anymodule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestStatus.ProtoReflect.Descriptor instead.
func (*IngestStatus) Descriptor() ([]byte, []int) {
	return file_protos_anymodule_anymodule_proto_rawDescGZIP(), []int{4}
}

func (x *IngestStatus) GetStatus() *IngestStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_protos_anymodule_anymodule_proto protoreflect.FileDescriptor

var file_protos_anymodule_anymodule_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x79,
	0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6e, 0x79, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x28, 0x0a,
	0x0f, 0x4c, 0x61, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x48, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4c, 0x61, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x48, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x4c, 0x61, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x45, 0x0a,
	0x0c, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x3d, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41,
	0x4e, 0x4b, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x47,
	0x45, 0x10, 0x03, 0x32, 0xe4, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x5a,
	0x0a, 0x0f, 0x53, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x26, 0x2e, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x61, 0x6e, 0x79, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x53, 0x65,
	0x65, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x61, 0x6e,
	0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49,
	0x44, 0x1a, 0x1d, 0x2e, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x53, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x1e,
	0x2e, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x59, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x41,
	0x6e, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x6e, 0x79, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_anymodule_anymodule_proto_rawDescOnce sync.Once
	file_protos_anymodule_anymodule_proto_rawDescData = file_protos_anymodule_anymodule_proto_rawDesc
)

func file_protos_anymodule_anymodule_proto_rawDescGZIP() []byte {
	file_protos_anymodule_anymodule_proto_rawDescOnce.Do(func() {
		file_protos_anymodule_anymodule_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_anymodule_anymodule_proto_rawDescData)
	})
	return file_protos_anymodule_anymodule_proto_rawDescData
}

var file_protos_anymodule_anymodule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_anymodule_anymodule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_anymodule_anymodule_proto_goTypes = []any{
	(EnumStatus)(0),               // 0: anymodule_proto.EnumStatus
	(*IngestRequestByRandId)(nil), // 1: anymodule_proto.IngestRequestByRandId
	(*IngestRequestByUUID)(nil),   // 2: anymodule_proto.IngestRequestByUUID
	(*DeleteAllRequest)(nil),      // 3: anymodule_proto.DeleteAllRequest
	(*IngestRequest)(nil),         // 4: anymodule_proto.IngestRequest
	(*IngestStatus)(nil),          // 5: anymodule_proto.IngestStatus
}
var file_protos_anymodule_anymodule_proto_depIdxs = []int32{
	5, // 0: anymodule_proto.IngestStatus.Status:type_name -> anymodule_proto.IngestStatus
	1, // 1: anymodule_proto.Setter.SeedOneByRandId:input_type -> anymodule_proto.IngestRequestByRandId
	2, // 2: anymodule_proto.Setter.SeedOneByUUID:input_type -> anymodule_proto.IngestRequestByUUID
	4, // 3: anymodule_proto.Setter.SeedMany:input_type -> anymodule_proto.IngestRequest
	3, // 4: anymodule_proto.Setter.DeleteManyByAnyUUID:input_type -> anymodule_proto.DeleteAllRequest
	5, // 5: anymodule_proto.Setter.SeedOneByRandId:output_type -> anymodule_proto.IngestStatus
	5, // 6: anymodule_proto.Setter.SeedOneByUUID:output_type -> anymodule_proto.IngestStatus
	5, // 7: anymodule_proto.Setter.SeedMany:output_type -> anymodule_proto.IngestStatus
	5, // 8: anymodule_proto.Setter.DeleteManyByAnyUUID:output_type -> anymodule_proto.IngestStatus
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_anymodule_anymodule_proto_init() }
func file_protos_anymodule_anymodule_proto_init() {
	if File_protos_anymodule_anymodule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_anymodule_anymodule_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IngestRequestByRandId); i {
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
		file_protos_anymodule_anymodule_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IngestRequestByUUID); i {
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
		file_protos_anymodule_anymodule_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAllRequest); i {
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
		file_protos_anymodule_anymodule_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IngestRequest); i {
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
		file_protos_anymodule_anymodule_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IngestStatus); i {
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
			RawDescriptor: file_protos_anymodule_anymodule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_anymodule_anymodule_proto_goTypes,
		DependencyIndexes: file_protos_anymodule_anymodule_proto_depIdxs,
		EnumInfos:         file_protos_anymodule_anymodule_proto_enumTypes,
		MessageInfos:      file_protos_anymodule_anymodule_proto_msgTypes,
	}.Build()
	File_protos_anymodule_anymodule_proto = out.File
	file_protos_anymodule_anymodule_proto_rawDesc = nil
	file_protos_anymodule_anymodule_proto_goTypes = nil
	file_protos_anymodule_anymodule_proto_depIdxs = nil
}
