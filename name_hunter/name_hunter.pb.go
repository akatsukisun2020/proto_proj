// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: name_hunter.proto

package namehunter

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

type RandomNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book  string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`    // 著作
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 数量
}

func (x *RandomNameReq) Reset() {
	*x = RandomNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_name_hunter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNameReq) ProtoMessage() {}

func (x *RandomNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_name_hunter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNameReq.ProtoReflect.Descriptor instead.
func (*RandomNameReq) Descriptor() ([]byte, []int) {
	return file_name_hunter_proto_rawDescGZIP(), []int{0}
}

func (x *RandomNameReq) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *RandomNameReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SelName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         // 名字
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`       // 标题
	Author   string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`     // 作者
	Book     string `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`         // 书名
	Dynasty  string `protobuf:"bytes,5,opt,name=dynasty,proto3" json:"dynasty,omitempty"`   // 朝代
	Sentence string `protobuf:"bytes,6,opt,name=sentence,proto3" json:"sentence,omitempty"` // 诗句
}

func (x *SelName) Reset() {
	*x = SelName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_name_hunter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelName) ProtoMessage() {}

func (x *SelName) ProtoReflect() protoreflect.Message {
	mi := &file_name_hunter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelName.ProtoReflect.Descriptor instead.
func (*SelName) Descriptor() ([]byte, []int) {
	return file_name_hunter_proto_rawDescGZIP(), []int{1}
}

func (x *SelName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SelName) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SelName) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *SelName) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *SelName) GetDynasty() string {
	if x != nil {
		return x.Dynasty
	}
	return ""
}

func (x *SelName) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type RandomNameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []*SelName `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"` // 名字
}

func (x *RandomNameRsp) Reset() {
	*x = RandomNameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_name_hunter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNameRsp) ProtoMessage() {}

func (x *RandomNameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_name_hunter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNameRsp.ProtoReflect.Descriptor instead.
func (*RandomNameRsp) Descriptor() ([]byte, []int) {
	return file_name_hunter_proto_rawDescGZIP(), []int{2}
}

func (x *RandomNameRsp) GetNames() []*SelName {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_name_hunter_proto protoreflect.FileDescriptor

var file_name_hunter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a,
	0x0d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x79,
	0x6e, 0x61, 0x73, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x79, 0x6e,
	0x61, 0x73, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x3a, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x29, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x78, 0x0a, 0x0e,
	0x4e, 0x61, 0x6d, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x66,
	0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x68, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x6e, 0x61, 0x6d,
	0x65, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x61, 0x74, 0x73, 0x75, 0x6b, 0x69, 0x73, 0x75, 0x6e,
	0x32, 0x30, 0x32, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_name_hunter_proto_rawDescOnce sync.Once
	file_name_hunter_proto_rawDescData = file_name_hunter_proto_rawDesc
)

func file_name_hunter_proto_rawDescGZIP() []byte {
	file_name_hunter_proto_rawDescOnce.Do(func() {
		file_name_hunter_proto_rawDescData = protoimpl.X.CompressGZIP(file_name_hunter_proto_rawDescData)
	})
	return file_name_hunter_proto_rawDescData
}

var file_name_hunter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_name_hunter_proto_goTypes = []interface{}{
	(*RandomNameReq)(nil), // 0: namehunter.RandomNameReq
	(*SelName)(nil),       // 1: namehunter.SelName
	(*RandomNameRsp)(nil), // 2: namehunter.RandomNameRsp
}
var file_name_hunter_proto_depIdxs = []int32{
	1, // 0: namehunter.RandomNameRsp.names:type_name -> namehunter.SelName
	0, // 1: namehunter.NameHunterHttp.RandomName:input_type -> namehunter.RandomNameReq
	2, // 2: namehunter.NameHunterHttp.RandomName:output_type -> namehunter.RandomNameRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_name_hunter_proto_init() }
func file_name_hunter_proto_init() {
	if File_name_hunter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_name_hunter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNameReq); i {
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
		file_name_hunter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelName); i {
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
		file_name_hunter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNameRsp); i {
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
			RawDescriptor: file_name_hunter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_name_hunter_proto_goTypes,
		DependencyIndexes: file_name_hunter_proto_depIdxs,
		MessageInfos:      file_name_hunter_proto_msgTypes,
	}.Build()
	File_name_hunter_proto = out.File
	file_name_hunter_proto_rawDesc = nil
	file_name_hunter_proto_goTypes = nil
	file_name_hunter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NameHunterHttpClient is the client API for NameHunterHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameHunterHttpClient interface {
	// 根据古诗词随机生成名字的接口
	RandomName(ctx context.Context, in *RandomNameReq, opts ...grpc.CallOption) (*RandomNameRsp, error)
}

type nameHunterHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewNameHunterHttpClient(cc grpc.ClientConnInterface) NameHunterHttpClient {
	return &nameHunterHttpClient{cc}
}

func (c *nameHunterHttpClient) RandomName(ctx context.Context, in *RandomNameReq, opts ...grpc.CallOption) (*RandomNameRsp, error) {
	out := new(RandomNameRsp)
	err := c.cc.Invoke(ctx, "/namehunter.NameHunterHttp/RandomName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameHunterHttpServer is the server API for NameHunterHttp service.
type NameHunterHttpServer interface {
	// 根据古诗词随机生成名字的接口
	RandomName(context.Context, *RandomNameReq) (*RandomNameRsp, error)
}

// UnimplementedNameHunterHttpServer can be embedded to have forward compatible implementations.
type UnimplementedNameHunterHttpServer struct {
}

func (*UnimplementedNameHunterHttpServer) RandomName(context.Context, *RandomNameReq) (*RandomNameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomName not implemented")
}

func RegisterNameHunterHttpServer(s *grpc.Server, srv NameHunterHttpServer) {
	s.RegisterService(&_NameHunterHttp_serviceDesc, srv)
}

func _NameHunterHttp_RandomName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameHunterHttpServer).RandomName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namehunter.NameHunterHttp/RandomName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameHunterHttpServer).RandomName(ctx, req.(*RandomNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NameHunterHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "namehunter.NameHunterHttp",
	HandlerType: (*NameHunterHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RandomName",
			Handler:    _NameHunterHttp_RandomName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "name_hunter.proto",
}