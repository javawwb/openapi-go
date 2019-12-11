// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/easyopsapis/openapi-go/gerr/message.proto

package gerr

import (
	fmt "fmt"
	codes "github.com/easyops-cn/giraffe-micro/codes"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// EasyOps 状态信息定义
type Message struct {
	// 错误码
	Code codes.Code `protobuf:"varint,1,opt,name=code,proto3,enum=codes.Code" json:"code,omitempty"`
	// 面向程序员的错误信息, 应为全英文
	CodeExplain string `protobuf:"bytes,2,opt,name=codeExplain,proto3" json:"codeExplain,omitempty"`
	// 面向终端用户的错误信息, 应用业务语言描述
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// 兼容旧接口
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	// 兼容旧接口
	Msg                  string   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_423c02171752337f, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetCode() codes.Code {
	if m != nil {
		return m.Code
	}
	return codes.Code_OK
}

func (m *Message) GetCodeExplain() string {
	if m != nil {
		return m.CodeExplain
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "gerr.Message")
}

func init() {
	proto.RegisterFile("github.com/easyopsapis/openapi-go/gerr/message.proto", fileDescriptor_423c02171752337f)
}

var fileDescriptor_423c02171752337f = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x05, 0x21,
	0x14, 0x85, 0xb1, 0x37, 0xaf, 0x47, 0x3e, 0x88, 0x90, 0x16, 0xd2, 0xa6, 0xa1, 0xd5, 0x40, 0xa8,
	0x50, 0xed, 0xda, 0x15, 0x2d, 0xdb, 0xcc, 0xb2, 0x9d, 0xe3, 0xdc, 0x31, 0x21, 0xe7, 0xca, 0x75,
	0x82, 0xfa, 0x13, 0xfd, 0xe6, 0xd0, 0x29, 0x68, 0xd1, 0xe2, 0x6d, 0x8e, 0x9e, 0xef, 0x78, 0x44,
	0x2f, 0xbf, 0xf3, 0x61, 0x79, 0x7d, 0x1f, 0xb4, 0xc3, 0x68, 0xc0, 0xe6, 0x4f, 0x4c, 0xd9, 0xa6,
	0x90, 0x0d, 0x26, 0x98, 0x6d, 0x0a, 0xca, 0xa3, 0xf1, 0x40, 0x64, 0x22, 0xe4, 0x6c, 0x3d, 0xe8,
	0x44, 0xb8, 0xa0, 0x68, 0x0a, 0xbb, 0xf8, 0xa7, 0xab, 0xdc, 0x6c, 0x7c, 0x20, 0x3b, 0x4d, 0xa0,
	0x62, 0x70, 0x84, 0xc6, 0xe1, 0x08, 0xb9, 0xea, 0xda, 0xbd, 0xfa, 0x62, 0x7c, 0xf7, 0xbc, 0xde,
	0x26, 0x2e, 0x79, 0x53, 0x12, 0xc9, 0x5a, 0xd6, 0x9d, 0xde, 0xec, 0x75, 0x3d, 0xac, 0x1f, 0x71,
	0x84, 0xbe, 0x06, 0xa2, 0xe5, 0xfb, 0xb2, 0x3e, 0x7d, 0xa4, 0x37, 0x1b, 0x66, 0x79, 0xd4, 0xb2,
	0xee, 0xa4, 0xff, 0x8b, 0x84, 0xe4, 0xbb, 0x9f, 0xb7, 0xc9, 0x4d, 0x4d, 0x7f, 0xad, 0x38, 0xe7,
	0x5b, 0x20, 0x42, 0x92, 0x4d, 0xe5, 0xab, 0x11, 0x67, 0x7c, 0x13, 0xb3, 0x97, 0xdb, 0xca, 0xca,
	0xf6, 0x41, 0xbd, 0x5c, 0x1f, 0x36, 0x84, 0xfb, 0x22, 0xc3, 0x71, 0xfd, 0xc6, 0xed, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x3a, 0x40, 0xa4, 0x3a, 0x01, 0x00, 0x00,
}
