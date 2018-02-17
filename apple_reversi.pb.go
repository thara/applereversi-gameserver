// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apple_reversi.proto

/*
Package applereversi is a generated protocol buffer package.

It is generated from these files:
	apple_reversi.proto

It has these top-level messages:
	GameConfig
	Move
	Empty
*/
package applereversi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GameConfig_Color int32

const (
	GameConfig_BLACK GameConfig_Color = 0
	GameConfig_WHITE GameConfig_Color = 1
)

var GameConfig_Color_name = map[int32]string{
	0: "BLACK",
	1: "WHITE",
}
var GameConfig_Color_value = map[string]int32{
	"BLACK": 0,
	"WHITE": 1,
}

func (x GameConfig_Color) String() string {
	return proto.EnumName(GameConfig_Color_name, int32(x))
}
func (GameConfig_Color) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type GameConfig struct {
	Color GameConfig_Color `protobuf:"varint,1,opt,name=color,enum=applereversi.GameConfig_Color" json:"color,omitempty"`
}

func (m *GameConfig) Reset()                    { *m = GameConfig{} }
func (m *GameConfig) String() string            { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()               {}
func (*GameConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GameConfig) GetColor() GameConfig_Color {
	if m != nil {
		return m.Color
	}
	return GameConfig_BLACK
}

type Move struct {
	Row    int32 `protobuf:"varint,1,opt,name=row" json:"row,omitempty"`
	Column int32 `protobuf:"varint,2,opt,name=column" json:"column,omitempty"`
}

func (m *Move) Reset()                    { *m = Move{} }
func (m *Move) String() string            { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()               {}
func (*Move) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Move) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *Move) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*GameConfig)(nil), "applereversi.GameConfig")
	proto.RegisterType((*Move)(nil), "applereversi.Move")
	proto.RegisterType((*Empty)(nil), "applereversi.Empty")
	proto.RegisterEnum("applereversi.GameConfig_Color", GameConfig_Color_name, GameConfig_Color_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReversiAI service

type ReversiAIClient interface {
	Init(ctx context.Context, in *GameConfig, opts ...grpc.CallOption) (*Empty, error)
	SelectMove(ctx context.Context, opts ...grpc.CallOption) (ReversiAI_SelectMoveClient, error)
}

type reversiAIClient struct {
	cc *grpc.ClientConn
}

func NewReversiAIClient(cc *grpc.ClientConn) ReversiAIClient {
	return &reversiAIClient{cc}
}

func (c *reversiAIClient) Init(ctx context.Context, in *GameConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/applereversi.ReversiAI/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reversiAIClient) SelectMove(ctx context.Context, opts ...grpc.CallOption) (ReversiAI_SelectMoveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ReversiAI_serviceDesc.Streams[0], c.cc, "/applereversi.ReversiAI/SelectMove", opts...)
	if err != nil {
		return nil, err
	}
	x := &reversiAISelectMoveClient{stream}
	return x, nil
}

type ReversiAI_SelectMoveClient interface {
	Send(*Move) error
	Recv() (*Move, error)
	grpc.ClientStream
}

type reversiAISelectMoveClient struct {
	grpc.ClientStream
}

func (x *reversiAISelectMoveClient) Send(m *Move) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reversiAISelectMoveClient) Recv() (*Move, error) {
	m := new(Move)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ReversiAI service

type ReversiAIServer interface {
	Init(context.Context, *GameConfig) (*Empty, error)
	SelectMove(ReversiAI_SelectMoveServer) error
}

func RegisterReversiAIServer(s *grpc.Server, srv ReversiAIServer) {
	s.RegisterService(&_ReversiAI_serviceDesc, srv)
}

func _ReversiAI_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReversiAIServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applereversi.ReversiAI/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReversiAIServer).Init(ctx, req.(*GameConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReversiAI_SelectMove_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReversiAIServer).SelectMove(&reversiAISelectMoveServer{stream})
}

type ReversiAI_SelectMoveServer interface {
	Send(*Move) error
	Recv() (*Move, error)
	grpc.ServerStream
}

type reversiAISelectMoveServer struct {
	grpc.ServerStream
}

func (x *reversiAISelectMoveServer) Send(m *Move) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reversiAISelectMoveServer) Recv() (*Move, error) {
	m := new(Move)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ReversiAI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "applereversi.ReversiAI",
	HandlerType: (*ReversiAIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ReversiAI_Init_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SelectMove",
			Handler:       _ReversiAI_SelectMove_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apple_reversi.proto",
}

func init() { proto.RegisterFile("apple_reversi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0x28, 0xc8,
	0x49, 0x8d, 0x2f, 0x4a, 0x2d, 0x4b, 0x2d, 0x2a, 0xce, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x01, 0x0b, 0x42, 0xc5, 0x94, 0x12, 0xb9, 0xb8, 0xdc, 0x13, 0x73, 0x53, 0x9d, 0xf3, 0xf3,
	0xd2, 0x32, 0xd3, 0x85, 0x4c, 0xb8, 0x58, 0x93, 0xf3, 0x73, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xf8, 0x8c, 0xe4, 0xf4, 0x90, 0xd5, 0xea, 0x21, 0x14, 0xea, 0x39, 0x83, 0x54, 0x05, 0x41,
	0x14, 0x2b, 0xc9, 0x72, 0xb1, 0x82, 0xf9, 0x42, 0x9c, 0x5c, 0xac, 0x4e, 0x3e, 0x8e, 0xce, 0xde,
	0x02, 0x0c, 0x20, 0x66, 0xb8, 0x87, 0x67, 0x88, 0xab, 0x00, 0xa3, 0x92, 0x01, 0x17, 0x8b, 0x6f,
	0x7e, 0x59, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x51, 0x7e, 0x39, 0xd8, 0x68, 0xd6, 0x20, 0x10, 0x53,
	0x48, 0x8c, 0x8b, 0x2d, 0x39, 0x3f, 0xa7, 0x34, 0x37, 0x4f, 0x82, 0x09, 0x2c, 0x08, 0xe5, 0x29,
	0xb1, 0x73, 0xb1, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x1a, 0x35, 0x30, 0x72, 0x71, 0x06, 0x41, 0x6c,
	0x77, 0xf4, 0x14, 0x32, 0xe7, 0x62, 0xf1, 0xcc, 0xcb, 0x2c, 0x11, 0x92, 0xc0, 0xe5, 0x2c, 0x29,
	0x61, 0x54, 0x19, 0xb0, 0x21, 0x4a, 0x0c, 0x42, 0x56, 0x5c, 0x5c, 0xc1, 0xa9, 0x39, 0xa9, 0xc9,
	0x25, 0x60, 0x77, 0x08, 0xa1, 0x2a, 0x02, 0x89, 0x49, 0x61, 0x11, 0x53, 0x62, 0xd0, 0x60, 0x34,
	0x60, 0x4c, 0x62, 0x03, 0x87, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x73, 0xdf, 0xd9, 0xd6,
	0x4c, 0x01, 0x00, 0x00,
}
