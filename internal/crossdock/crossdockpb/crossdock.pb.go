// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.5.1
// source: internal/crossdock/crossdockpb/crossdock.proto

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package crossdockpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beep string `protobuf:"bytes,1,opt,name=beep,proto3" json:"beep,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_internal_crossdock_crossdockpb_crossdock_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetBeep() string {
	if x != nil {
		return x.Beep
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boop string `protobuf:"bytes,1,opt,name=boop,proto3" json:"boop,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_internal_crossdock_crossdockpb_crossdock_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetBoop() string {
	if x != nil {
		return x.Boop
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_internal_crossdock_crossdockpb_crossdock_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_internal_crossdock_crossdockpb_crossdock_proto protoreflect.FileDescriptor

var file_internal_crossdock_crossdockpb_crossdock_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x64, 0x6f, 0x63, 0x6b, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x70, 0x62,
	0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x79, 0x61, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x22,
	0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x65, 0x65, 0x70, 0x22, 0x1a, 0x0a, 0x04, 0x50,
	0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x70, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x58, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x50,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x23, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x79, 0x61,
	0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x23, 0x2e, 0x75, 0x62,
	0x65, 0x72, 0x2e, 0x79, 0x61, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x64, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_crossdock_crossdockpb_crossdock_proto_rawDescOnce sync.Once
	file_internal_crossdock_crossdockpb_crossdock_proto_rawDescData = file_internal_crossdock_crossdockpb_crossdock_proto_rawDesc
)

func file_internal_crossdock_crossdockpb_crossdock_proto_rawDescGZIP() []byte {
	file_internal_crossdock_crossdockpb_crossdock_proto_rawDescOnce.Do(func() {
		file_internal_crossdock_crossdockpb_crossdock_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_crossdock_crossdockpb_crossdock_proto_rawDescData)
	})
	return file_internal_crossdock_crossdockpb_crossdock_proto_rawDescData
}

var file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_crossdock_crossdockpb_crossdock_proto_goTypes = []interface{}{
	(*Ping)(nil),  // 0: uber.yarpc.internal.crossdock.Ping
	(*Pong)(nil),  // 1: uber.yarpc.internal.crossdock.Pong
	(*Token)(nil), // 2: uber.yarpc.internal.crossdock.Token
}
var file_internal_crossdock_crossdockpb_crossdock_proto_depIdxs = []int32{
	0, // 0: uber.yarpc.internal.crossdock.Echo.Echo:input_type -> uber.yarpc.internal.crossdock.Ping
	1, // 1: uber.yarpc.internal.crossdock.Echo.Echo:output_type -> uber.yarpc.internal.crossdock.Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_crossdock_crossdockpb_crossdock_proto_init() }
func file_internal_crossdock_crossdockpb_crossdock_proto_init() {
	if File_internal_crossdock_crossdockpb_crossdock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_internal_crossdock_crossdockpb_crossdock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_crossdock_crossdockpb_crossdock_proto_goTypes,
		DependencyIndexes: file_internal_crossdock_crossdockpb_crossdock_proto_depIdxs,
		MessageInfos:      file_internal_crossdock_crossdockpb_crossdock_proto_msgTypes,
	}.Build()
	File_internal_crossdock_crossdockpb_crossdock_proto = out.File
	file_internal_crossdock_crossdockpb_crossdock_proto_rawDesc = nil
	file_internal_crossdock_crossdockpb_crossdock_proto_goTypes = nil
	file_internal_crossdock_crossdockpb_crossdock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/uber.yarpc.internal.crossdock.Echo/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Echo(context.Context, *Ping) (*Pong, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Echo(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.yarpc.internal.crossdock.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uber.yarpc.internal.crossdock.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/crossdock/crossdockpb/crossdock.proto",
}
