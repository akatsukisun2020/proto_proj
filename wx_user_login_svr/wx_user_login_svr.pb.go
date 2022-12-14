// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: wx_user_login_svr.proto

package wx_user_login_svr

import (
	context "context"
	user "github.com/akatsukisun2020/proto_proj/user"
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

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 通过用户id获取用户信息
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_user_login_svr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_wx_user_login_svr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_wx_user_login_svr_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode  int32        `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg   string       `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	UserInfo *user.WXUser `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 用户信息
}

func (x *GetUserInfoRsp) Reset() {
	*x = GetUserInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wx_user_login_svr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRsp) ProtoMessage() {}

func (x *GetUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_wx_user_login_svr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRsp.ProtoReflect.Descriptor instead.
func (*GetUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_wx_user_login_svr_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRsp) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *GetUserInfoRsp) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *GetUserInfoRsp) GetUserInfo() *user.WXUser {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_wx_user_login_svr_proto protoreflect.FileDescriptor

var file_wx_user_login_svr_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x76, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x78, 0x75, 0x73, 0x65,
	0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x76, 0x72, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x57, 0x58, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0x90, 0x01, 0x0a, 0x12, 0x57, 0x58, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x76, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x7a, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x77, 0x78, 0x75,
	0x73, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x76, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x77, 0x78, 0x75,
	0x73, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x76, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x22, 0x20, 0x2f, 0x77, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x73, 0x76, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x61, 0x74, 0x73, 0x75, 0x6b, 0x69, 0x73, 0x75,
	0x6e, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x2f, 0x77, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73,
	0x76, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wx_user_login_svr_proto_rawDescOnce sync.Once
	file_wx_user_login_svr_proto_rawDescData = file_wx_user_login_svr_proto_rawDesc
)

func file_wx_user_login_svr_proto_rawDescGZIP() []byte {
	file_wx_user_login_svr_proto_rawDescOnce.Do(func() {
		file_wx_user_login_svr_proto_rawDescData = protoimpl.X.CompressGZIP(file_wx_user_login_svr_proto_rawDescData)
	})
	return file_wx_user_login_svr_proto_rawDescData
}

var file_wx_user_login_svr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wx_user_login_svr_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil), // 0: wxuserloginsvr.GetUserInfoReq
	(*GetUserInfoRsp)(nil), // 1: wxuserloginsvr.GetUserInfoRsp
	(*user.WXUser)(nil),    // 2: user.WXUser
}
var file_wx_user_login_svr_proto_depIdxs = []int32{
	2, // 0: wxuserloginsvr.GetUserInfoRsp.user_info:type_name -> user.WXUser
	0, // 1: wxuserloginsvr.WXUserLoginSvrHttp.GetUserInfo:input_type -> wxuserloginsvr.GetUserInfoReq
	1, // 2: wxuserloginsvr.WXUserLoginSvrHttp.GetUserInfo:output_type -> wxuserloginsvr.GetUserInfoRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wx_user_login_svr_proto_init() }
func file_wx_user_login_svr_proto_init() {
	if File_wx_user_login_svr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wx_user_login_svr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_wx_user_login_svr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRsp); i {
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
			RawDescriptor: file_wx_user_login_svr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wx_user_login_svr_proto_goTypes,
		DependencyIndexes: file_wx_user_login_svr_proto_depIdxs,
		MessageInfos:      file_wx_user_login_svr_proto_msgTypes,
	}.Build()
	File_wx_user_login_svr_proto = out.File
	file_wx_user_login_svr_proto_rawDesc = nil
	file_wx_user_login_svr_proto_goTypes = nil
	file_wx_user_login_svr_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WXUserLoginSvrHttpClient is the client API for WXUserLoginSvrHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WXUserLoginSvrHttpClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error)
}

type wXUserLoginSvrHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewWXUserLoginSvrHttpClient(cc grpc.ClientConnInterface) WXUserLoginSvrHttpClient {
	return &wXUserLoginSvrHttpClient{cc}
}

func (c *wXUserLoginSvrHttpClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error) {
	out := new(GetUserInfoRsp)
	err := c.cc.Invoke(ctx, "/wxuserloginsvr.WXUserLoginSvrHttp/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WXUserLoginSvrHttpServer is the server API for WXUserLoginSvrHttp service.
type WXUserLoginSvrHttpServer interface {
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error)
}

// UnimplementedWXUserLoginSvrHttpServer can be embedded to have forward compatible implementations.
type UnimplementedWXUserLoginSvrHttpServer struct {
}

func (*UnimplementedWXUserLoginSvrHttpServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterWXUserLoginSvrHttpServer(s *grpc.Server, srv WXUserLoginSvrHttpServer) {
	s.RegisterService(&_WXUserLoginSvrHttp_serviceDesc, srv)
}

func _WXUserLoginSvrHttp_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WXUserLoginSvrHttpServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxuserloginsvr.WXUserLoginSvrHttp/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WXUserLoginSvrHttpServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WXUserLoginSvrHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wxuserloginsvr.WXUserLoginSvrHttp",
	HandlerType: (*WXUserLoginSvrHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _WXUserLoginSvrHttp_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wx_user_login_svr.proto",
}
