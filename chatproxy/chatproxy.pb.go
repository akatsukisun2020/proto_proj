// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: chatproxy.proto

package chatproxy

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChatCompletionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatCompletionReq) Reset() {
	*x = ChatCompletionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionReq) ProtoMessage() {}

func (x *ChatCompletionReq) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionReq.ProtoReflect.Descriptor instead.
func (*ChatCompletionReq) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{0}
}

func (x *ChatCompletionReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatCompletionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatCompletionRsp) Reset() {
	*x = ChatCompletionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionRsp) ProtoMessage() {}

func (x *ChatCompletionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionRsp.ProtoReflect.Descriptor instead.
func (*ChatCompletionRsp) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{1}
}

// ChatCompletionMessage ...
type ChatCompletionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatCompletionMessage) Reset() {
	*x = ChatCompletionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionMessage) ProtoMessage() {}

func (x *ChatCompletionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionMessage.ProtoReflect.Descriptor instead.
func (*ChatCompletionMessage) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{2}
}

func (x *ChatCompletionMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatCompletionMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ChatStreamCompletionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model                  string                   `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`                        // 模型
	MaxToken               int32                    `protobuf:"varint,2,opt,name=max_token,json=maxToken,proto3" json:"max_token,omitempty"` // 最大token数量
	ChatCompletionMessages []*ChatCompletionMessage `protobuf:"bytes,3,rep,name=chat_completion_messages,json=chatCompletionMessages,proto3" json:"chat_completion_messages,omitempty"`
}

func (x *ChatStreamCompletionReq) Reset() {
	*x = ChatStreamCompletionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatStreamCompletionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatStreamCompletionReq) ProtoMessage() {}

func (x *ChatStreamCompletionReq) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatStreamCompletionReq.ProtoReflect.Descriptor instead.
func (*ChatStreamCompletionReq) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{3}
}

func (x *ChatStreamCompletionReq) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatStreamCompletionReq) GetMaxToken() int32 {
	if x != nil {
		return x.MaxToken
	}
	return 0
}

func (x *ChatStreamCompletionReq) GetChatCompletionMessages() []*ChatCompletionMessage {
	if x != nil {
		return x.ChatCompletionMessages
	}
	return nil
}

type ChatCompletionStreamChoiceDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatCompletionStreamChoiceDelta) Reset() {
	*x = ChatCompletionStreamChoiceDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionStreamChoiceDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionStreamChoiceDelta) ProtoMessage() {}

func (x *ChatCompletionStreamChoiceDelta) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionStreamChoiceDelta.ProtoReflect.Descriptor instead.
func (*ChatCompletionStreamChoiceDelta) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{4}
}

func (x *ChatCompletionStreamChoiceDelta) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ChatCompletionStreamChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index         int32                            `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Delta         *ChatCompletionStreamChoiceDelta `protobuf:"bytes,2,opt,name=delta,proto3" json:"delta,omitempty"`
	FinishReasion string                           `protobuf:"bytes,3,opt,name=finish_reasion,json=finishReasion,proto3" json:"finish_reasion,omitempty"`
}

func (x *ChatCompletionStreamChoice) Reset() {
	*x = ChatCompletionStreamChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionStreamChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionStreamChoice) ProtoMessage() {}

func (x *ChatCompletionStreamChoice) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionStreamChoice.ProtoReflect.Descriptor instead.
func (*ChatCompletionStreamChoice) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{5}
}

func (x *ChatCompletionStreamChoice) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChatCompletionStreamChoice) GetDelta() *ChatCompletionStreamChoiceDelta {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *ChatCompletionStreamChoice) GetFinishReasion() string {
	if x != nil {
		return x.FinishReasion
	}
	return ""
}

type ChatStreamCompletionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string                        `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Created int64                         `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Model   string                        `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Choices []*ChatCompletionStreamChoice `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"`
}

func (x *ChatStreamCompletionRsp) Reset() {
	*x = ChatStreamCompletionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatStreamCompletionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatStreamCompletionRsp) ProtoMessage() {}

func (x *ChatStreamCompletionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chatproxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatStreamCompletionRsp.ProtoReflect.Descriptor instead.
func (*ChatStreamCompletionRsp) Descriptor() ([]byte, []int) {
	return file_chatproxy_proto_rawDescGZIP(), []int{6}
}

func (x *ChatStreamCompletionRsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatStreamCompletionRsp) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ChatStreamCompletionRsp) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *ChatStreamCompletionRsp) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatStreamCompletionRsp) GetChoices() []*ChatCompletionStreamChoice {
	if x != nil {
		return x.Choices
	}
	return nil
}

var File_chatproxy_proto protoreflect.FileDescriptor

var file_chatproxy_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x45,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5a, 0x0a, 0x18, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x16, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x3b, 0x0a, 0x1f, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x9b, 0x01,
	0x0a, 0x1a, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x40, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x17,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x3f, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x32, 0x95, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74,
	0x74, 0x70, 0x12, 0x73, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x22, 0x21, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x61, 0x74, 0x73, 0x75, 0x6b, 0x69, 0x73,
	0x75, 0x6e, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chatproxy_proto_rawDescOnce sync.Once
	file_chatproxy_proto_rawDescData = file_chatproxy_proto_rawDesc
)

func file_chatproxy_proto_rawDescGZIP() []byte {
	file_chatproxy_proto_rawDescOnce.Do(func() {
		file_chatproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatproxy_proto_rawDescData)
	})
	return file_chatproxy_proto_rawDescData
}

var file_chatproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chatproxy_proto_goTypes = []interface{}{
	(*ChatCompletionReq)(nil),               // 0: chatproxy.ChatCompletionReq
	(*ChatCompletionRsp)(nil),               // 1: chatproxy.ChatCompletionRsp
	(*ChatCompletionMessage)(nil),           // 2: chatproxy.ChatCompletionMessage
	(*ChatStreamCompletionReq)(nil),         // 3: chatproxy.ChatStreamCompletionReq
	(*ChatCompletionStreamChoiceDelta)(nil), // 4: chatproxy.ChatCompletionStreamChoiceDelta
	(*ChatCompletionStreamChoice)(nil),      // 5: chatproxy.ChatCompletionStreamChoice
	(*ChatStreamCompletionRsp)(nil),         // 6: chatproxy.ChatStreamCompletionRsp
}
var file_chatproxy_proto_depIdxs = []int32{
	2, // 0: chatproxy.ChatStreamCompletionReq.chat_completion_messages:type_name -> chatproxy.ChatCompletionMessage
	4, // 1: chatproxy.ChatCompletionStreamChoice.delta:type_name -> chatproxy.ChatCompletionStreamChoiceDelta
	5, // 2: chatproxy.ChatStreamCompletionRsp.choices:type_name -> chatproxy.ChatCompletionStreamChoice
	0, // 3: chatproxy.ChatproxyHttp.ChatCompletion:input_type -> chatproxy.ChatCompletionReq
	3, // 4: chatproxy.ChatproxyHttp.ChatStreamCompletion:input_type -> chatproxy.ChatStreamCompletionReq
	1, // 5: chatproxy.ChatproxyHttp.ChatCompletion:output_type -> chatproxy.ChatCompletionRsp
	6, // 6: chatproxy.ChatproxyHttp.ChatStreamCompletion:output_type -> chatproxy.ChatStreamCompletionRsp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chatproxy_proto_init() }
func file_chatproxy_proto_init() {
	if File_chatproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionReq); i {
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
		file_chatproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionRsp); i {
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
		file_chatproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionMessage); i {
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
		file_chatproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatStreamCompletionReq); i {
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
		file_chatproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionStreamChoiceDelta); i {
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
		file_chatproxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionStreamChoice); i {
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
		file_chatproxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatStreamCompletionRsp); i {
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
			RawDescriptor: file_chatproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatproxy_proto_goTypes,
		DependencyIndexes: file_chatproxy_proto_depIdxs,
		MessageInfos:      file_chatproxy_proto_msgTypes,
	}.Build()
	File_chatproxy_proto = out.File
	file_chatproxy_proto_rawDesc = nil
	file_chatproxy_proto_goTypes = nil
	file_chatproxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatproxyHttpClient is the client API for ChatproxyHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatproxyHttpClient interface {
	ChatCompletion(ctx context.Context, in *ChatCompletionReq, opts ...grpc.CallOption) (*ChatCompletionRsp, error)
	ChatStreamCompletion(ctx context.Context, in *ChatStreamCompletionReq, opts ...grpc.CallOption) (ChatproxyHttp_ChatStreamCompletionClient, error)
}

type chatproxyHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewChatproxyHttpClient(cc grpc.ClientConnInterface) ChatproxyHttpClient {
	return &chatproxyHttpClient{cc}
}

func (c *chatproxyHttpClient) ChatCompletion(ctx context.Context, in *ChatCompletionReq, opts ...grpc.CallOption) (*ChatCompletionRsp, error) {
	out := new(ChatCompletionRsp)
	err := c.cc.Invoke(ctx, "/chatproxy.ChatproxyHttp/ChatCompletion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatproxyHttpClient) ChatStreamCompletion(ctx context.Context, in *ChatStreamCompletionReq, opts ...grpc.CallOption) (ChatproxyHttp_ChatStreamCompletionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatproxyHttp_serviceDesc.Streams[0], "/chatproxy.ChatproxyHttp/ChatStreamCompletion", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatproxyHttpChatStreamCompletionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatproxyHttp_ChatStreamCompletionClient interface {
	Recv() (*ChatStreamCompletionRsp, error)
	grpc.ClientStream
}

type chatproxyHttpChatStreamCompletionClient struct {
	grpc.ClientStream
}

func (x *chatproxyHttpChatStreamCompletionClient) Recv() (*ChatStreamCompletionRsp, error) {
	m := new(ChatStreamCompletionRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatproxyHttpServer is the server API for ChatproxyHttp service.
type ChatproxyHttpServer interface {
	ChatCompletion(context.Context, *ChatCompletionReq) (*ChatCompletionRsp, error)
	ChatStreamCompletion(*ChatStreamCompletionReq, ChatproxyHttp_ChatStreamCompletionServer) error
}

// UnimplementedChatproxyHttpServer can be embedded to have forward compatible implementations.
type UnimplementedChatproxyHttpServer struct {
}

func (*UnimplementedChatproxyHttpServer) ChatCompletion(context.Context, *ChatCompletionReq) (*ChatCompletionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatCompletion not implemented")
}
func (*UnimplementedChatproxyHttpServer) ChatStreamCompletion(*ChatStreamCompletionReq, ChatproxyHttp_ChatStreamCompletionServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatStreamCompletion not implemented")
}

func RegisterChatproxyHttpServer(s *grpc.Server, srv ChatproxyHttpServer) {
	s.RegisterService(&_ChatproxyHttp_serviceDesc, srv)
}

func _ChatproxyHttp_ChatCompletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatCompletionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatproxyHttpServer).ChatCompletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatproxy.ChatproxyHttp/ChatCompletion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatproxyHttpServer).ChatCompletion(ctx, req.(*ChatCompletionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatproxyHttp_ChatStreamCompletion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatStreamCompletionReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatproxyHttpServer).ChatStreamCompletion(m, &chatproxyHttpChatStreamCompletionServer{stream})
}

type ChatproxyHttp_ChatStreamCompletionServer interface {
	Send(*ChatStreamCompletionRsp) error
	grpc.ServerStream
}

type chatproxyHttpChatStreamCompletionServer struct {
	grpc.ServerStream
}

func (x *chatproxyHttpChatStreamCompletionServer) Send(m *ChatStreamCompletionRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatproxyHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatproxy.ChatproxyHttp",
	HandlerType: (*ChatproxyHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatCompletion",
			Handler:    _ChatproxyHttp_ChatCompletion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatStreamCompletion",
			Handler:       _ChatproxyHttp_ChatStreamCompletion_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatproxy.proto",
}
