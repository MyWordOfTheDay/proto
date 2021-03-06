// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: mywordoftheday/v1alpha1/mywordoftheday.proto

package mywordoftheday

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{0}
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{1}
}

type AddWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The word to add
	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *AddWordRequest) Reset() {
	*x = AddWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordRequest) ProtoMessage() {}

func (x *AddWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordRequest.ProtoReflect.Descriptor instead.
func (*AddWordRequest) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{2}
}

func (x *AddWordRequest) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

type AddWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The added word
	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *AddWordResponse) Reset() {
	*x = AddWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWordResponse) ProtoMessage() {}

func (x *AddWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWordResponse.ProtoReflect.Descriptor instead.
func (*AddWordResponse) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{3}
}

func (x *AddWordResponse) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

type ListWordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListWordsRequest) Reset() {
	*x = ListWordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordsRequest) ProtoMessage() {}

func (x *ListWordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordsRequest.ProtoReflect.Descriptor instead.
func (*ListWordsRequest) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{4}
}

type ListWordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of words
	Words []*Word `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *ListWordsResponse) Reset() {
	*x = ListWordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordsResponse) ProtoMessage() {}

func (x *ListWordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordsResponse.ProtoReflect.Descriptor instead.
func (*ListWordsResponse) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{5}
}

func (x *ListWordsResponse) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

type DeleteWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the change.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWordRequest) Reset() {
	*x = DeleteWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordRequest) ProtoMessage() {}

func (x *DeleteWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordRequest.ProtoReflect.Descriptor instead.
func (*DeleteWordRequest) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWordRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The deleted word
	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *DeleteWordResponse) Reset() {
	*x = DeleteWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWordResponse) ProtoMessage() {}

func (x *DeleteWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWordResponse.ProtoReflect.Descriptor instead.
func (*DeleteWordResponse) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteWordResponse) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

type RandomWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RandomWordRequest) Reset() {
	*x = RandomWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomWordRequest) ProtoMessage() {}

func (x *RandomWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomWordRequest.ProtoReflect.Descriptor instead.
func (*RandomWordRequest) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{8}
}

type RandomWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *RandomWordResponse) Reset() {
	*x = RandomWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomWordResponse) ProtoMessage() {}

func (x *RandomWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomWordResponse.ProtoReflect.Descriptor instead.
func (*RandomWordResponse) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{9}
}

func (x *RandomWordResponse) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the word.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The word to store
	Word string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	// A custom definition for the word
	CustomDefinition string `protobuf:"bytes,3,opt,name=custom_definition,json=customDefinition,proto3" json:"custom_definition,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP(), []int{10}
}

func (x *Word) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Word) GetCustomDefinition() string {
	if x != nil {
		return x.CustomDefinition
	}
	return ""
}

var File_mywordoftheday_v1alpha1_mywordoftheday_proto protoreflect.FileDescriptor

var file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x57, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x44, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x32, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0d, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x06, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x47, 0x0a, 0x12, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f,
	0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x69, 0x0a, 0x04, 0x57,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x31, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x01, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9f, 0x05, 0x0a, 0x15, 0x4d, 0x79, 0x57, 0x6f, 0x72,
	0x64, 0x4f, 0x66, 0x54, 0x68, 0x65, 0x44, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x7f, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x29, 0x2e,
	0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x7a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e, 0x6d,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66,
	0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x7b, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x29, 0x2e, 0x6d, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66,
	0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x2e, 0x6d, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66,
	0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x2a, 0x2e, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72,
	0x64, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x42, 0x8f, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66,
	0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x6d, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x6f, 0x66, 0x74, 0x68, 0x65, 0x64, 0x61,
	0x79, 0x92, 0x41, 0x46, 0x12, 0x20, 0x0a, 0x12, 0x4d, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x4f, 0x66,
	0x54, 0x68, 0x65, 0x44, 0x61, 0x79, 0x20, 0x41, 0x50, 0x49, 0x32, 0x0a, 0x31, 0x2e, 0x30, 0x2d,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescOnce sync.Once
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescData = file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDesc
)

func file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescGZIP() []byte {
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescOnce.Do(func() {
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescData = protoimpl.X.CompressGZIP(file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescData)
	})
	return file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDescData
}

var file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mywordoftheday_v1alpha1_mywordoftheday_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil),   // 0: mywordoftheday.v1alpha1.HeartbeatRequest
	(*HeartbeatResponse)(nil),  // 1: mywordoftheday.v1alpha1.HeartbeatResponse
	(*AddWordRequest)(nil),     // 2: mywordoftheday.v1alpha1.AddWordRequest
	(*AddWordResponse)(nil),    // 3: mywordoftheday.v1alpha1.AddWordResponse
	(*ListWordsRequest)(nil),   // 4: mywordoftheday.v1alpha1.ListWordsRequest
	(*ListWordsResponse)(nil),  // 5: mywordoftheday.v1alpha1.ListWordsResponse
	(*DeleteWordRequest)(nil),  // 6: mywordoftheday.v1alpha1.DeleteWordRequest
	(*DeleteWordResponse)(nil), // 7: mywordoftheday.v1alpha1.DeleteWordResponse
	(*RandomWordRequest)(nil),  // 8: mywordoftheday.v1alpha1.RandomWordRequest
	(*RandomWordResponse)(nil), // 9: mywordoftheday.v1alpha1.RandomWordResponse
	(*Word)(nil),               // 10: mywordoftheday.v1alpha1.Word
}
var file_mywordoftheday_v1alpha1_mywordoftheday_proto_depIdxs = []int32{
	10, // 0: mywordoftheday.v1alpha1.AddWordRequest.word:type_name -> mywordoftheday.v1alpha1.Word
	10, // 1: mywordoftheday.v1alpha1.AddWordResponse.word:type_name -> mywordoftheday.v1alpha1.Word
	10, // 2: mywordoftheday.v1alpha1.ListWordsResponse.words:type_name -> mywordoftheday.v1alpha1.Word
	10, // 3: mywordoftheday.v1alpha1.DeleteWordResponse.word:type_name -> mywordoftheday.v1alpha1.Word
	10, // 4: mywordoftheday.v1alpha1.RandomWordResponse.word:type_name -> mywordoftheday.v1alpha1.Word
	0,  // 5: mywordoftheday.v1alpha1.MyWordOfTheDayService.Heartbeat:input_type -> mywordoftheday.v1alpha1.HeartbeatRequest
	2,  // 6: mywordoftheday.v1alpha1.MyWordOfTheDayService.AddWord:input_type -> mywordoftheday.v1alpha1.AddWordRequest
	4,  // 7: mywordoftheday.v1alpha1.MyWordOfTheDayService.ListWords:input_type -> mywordoftheday.v1alpha1.ListWordsRequest
	6,  // 8: mywordoftheday.v1alpha1.MyWordOfTheDayService.DeleteWord:input_type -> mywordoftheday.v1alpha1.DeleteWordRequest
	8,  // 9: mywordoftheday.v1alpha1.MyWordOfTheDayService.RandomWord:input_type -> mywordoftheday.v1alpha1.RandomWordRequest
	1,  // 10: mywordoftheday.v1alpha1.MyWordOfTheDayService.Heartbeat:output_type -> mywordoftheday.v1alpha1.HeartbeatResponse
	3,  // 11: mywordoftheday.v1alpha1.MyWordOfTheDayService.AddWord:output_type -> mywordoftheday.v1alpha1.AddWordResponse
	5,  // 12: mywordoftheday.v1alpha1.MyWordOfTheDayService.ListWords:output_type -> mywordoftheday.v1alpha1.ListWordsResponse
	7,  // 13: mywordoftheday.v1alpha1.MyWordOfTheDayService.DeleteWord:output_type -> mywordoftheday.v1alpha1.DeleteWordResponse
	9,  // 14: mywordoftheday.v1alpha1.MyWordOfTheDayService.RandomWord:output_type -> mywordoftheday.v1alpha1.RandomWordResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_mywordoftheday_v1alpha1_mywordoftheday_proto_init() }
func file_mywordoftheday_v1alpha1_mywordoftheday_proto_init() {
	if File_mywordoftheday_v1alpha1_mywordoftheday_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordRequest); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWordResponse); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordsRequest); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordsResponse); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordRequest); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWordResponse); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomWordRequest); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomWordResponse); i {
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
		file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
			RawDescriptor: file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mywordoftheday_v1alpha1_mywordoftheday_proto_goTypes,
		DependencyIndexes: file_mywordoftheday_v1alpha1_mywordoftheday_proto_depIdxs,
		MessageInfos:      file_mywordoftheday_v1alpha1_mywordoftheday_proto_msgTypes,
	}.Build()
	File_mywordoftheday_v1alpha1_mywordoftheday_proto = out.File
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_rawDesc = nil
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_goTypes = nil
	file_mywordoftheday_v1alpha1_mywordoftheday_proto_depIdxs = nil
}
