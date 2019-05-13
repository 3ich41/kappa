// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conf.proto

package conf_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FetchRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{0}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Conf struct {
	Username             string    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Buttons              []*Button `protobuf:"bytes,2,rep,name=buttons,proto3" json:"buttons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Conf) Reset()         { *m = Conf{} }
func (m *Conf) String() string { return proto.CompactTextString(m) }
func (*Conf) ProtoMessage()    {}
func (*Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{1}
}

func (m *Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conf.Unmarshal(m, b)
}
func (m *Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conf.Marshal(b, m, deterministic)
}
func (m *Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conf.Merge(m, src)
}
func (m *Conf) XXX_Size() int {
	return xxx_messageInfo_Conf.Size(m)
}
func (m *Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_Conf proto.InternalMessageInfo

func (m *Conf) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Conf) GetButtons() []*Button {
	if m != nil {
		return m.Buttons
	}
	return nil
}

type Button struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Button) Reset()         { *m = Button{} }
func (m *Button) String() string { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()    {}
func (*Button) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{2}
}

func (m *Button) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Button.Unmarshal(m, b)
}
func (m *Button) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Button.Marshal(b, m, deterministic)
}
func (m *Button) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Button.Merge(m, src)
}
func (m *Button) XXX_Size() int {
	return xxx_messageInfo_Button.Size(m)
}
func (m *Button) XXX_DiscardUnknown() {
	xxx_messageInfo_Button.DiscardUnknown(m)
}

var xxx_messageInfo_Button proto.InternalMessageInfo

func (m *Button) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Button) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*FetchRequest)(nil), "conf_grpc.FetchRequest")
	proto.RegisterType((*Conf)(nil), "conf_grpc.Conf")
	proto.RegisterType((*Button)(nil), "conf_grpc.Button")
}

func init() { proto.RegisterFile("conf.proto", fileDescriptor_0b6ecbfc68e85c65) }

var fileDescriptor_0b6ecbfc68e85c65 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0xcf, 0x4b,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xe3, 0xd3, 0x8b, 0x0a, 0x92, 0x95,
	0xb4, 0xb8, 0x78, 0xdc, 0x52, 0x4b, 0x92, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xe0, 0x7c, 0x25, 0x7f, 0x2e, 0x16, 0xe7, 0xfc, 0xbc, 0x34, 0x7c, 0x6a, 0x84, 0xb4,
	0xb9, 0xd8, 0x93, 0x4a, 0x4b, 0x4a, 0xf2, 0xf3, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d,
	0x04, 0xf5, 0xe0, 0x96, 0xe9, 0x39, 0x81, 0x65, 0x82, 0x60, 0x2a, 0x94, 0x8c, 0xb8, 0xd8, 0x20,
	0x42, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x50, 0xe3, 0xc0, 0x6c, 0x21, 0x11, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x63, 0xe4, 0xc4, 0xc5, 0x0d,
	0x72, 0x84, 0x47, 0x62, 0x5e, 0x4a, 0x4e, 0x6a, 0x91, 0x90, 0x31, 0x17, 0xbb, 0x7b, 0x6a, 0x09,
	0xd8, 0x59, 0xe2, 0x48, 0x36, 0x21, 0xfb, 0x49, 0x8a, 0x1f, 0x49, 0x02, 0xa4, 0x32, 0x89, 0x0d,
	0x1c, 0x0c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x78, 0xef, 0x75, 0x14, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfHandlerClient is the client API for ConfHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfHandlerClient interface {
	GetConf(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Conf, error)
}

type confHandlerClient struct {
	cc *grpc.ClientConn
}

func NewConfHandlerClient(cc *grpc.ClientConn) ConfHandlerClient {
	return &confHandlerClient{cc}
}

func (c *confHandlerClient) GetConf(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Conf, error) {
	out := new(Conf)
	err := c.cc.Invoke(ctx, "/conf_grpc.ConfHandler/GetConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfHandlerServer is the server API for ConfHandler service.
type ConfHandlerServer interface {
	GetConf(context.Context, *FetchRequest) (*Conf, error)
}

func RegisterConfHandlerServer(s *grpc.Server, srv ConfHandlerServer) {
	s.RegisterService(&_ConfHandler_serviceDesc, srv)
}

func _ConfHandler_GetConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfHandlerServer).GetConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf_grpc.ConfHandler/GetConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfHandlerServer).GetConf(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conf_grpc.ConfHandler",
	HandlerType: (*ConfHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConf",
			Handler:    _ConfHandler_GetConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conf.proto",
}
