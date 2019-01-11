// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chart.proto

// 依照 https://cloud.google.com/apis/design/naming_convention 規範

package captainkube_v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Chart struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	FileSize             int64    `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chart) Reset()         { *m = Chart{} }
func (m *Chart) String() string { return proto.CompactTextString(m) }
func (*Chart) ProtoMessage()    {}
func (*Chart) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2b6c44db8db459e, []int{0}
}

func (m *Chart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chart.Unmarshal(m, b)
}
func (m *Chart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chart.Marshal(b, m, deterministic)
}
func (m *Chart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chart.Merge(m, src)
}
func (m *Chart) XXX_Size() int {
	return xxx_messageInfo_Chart.Size(m)
}
func (m *Chart) XXX_DiscardUnknown() {
	xxx_messageInfo_Chart.DiscardUnknown(m)
}

var xxx_messageInfo_Chart proto.InternalMessageInfo

func (m *Chart) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Chart) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chart) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

// Tiller is the Helm Server
type Tiller struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Account              string   `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	SkipSslValidation    bool     `protobuf:"varint,5,opt,name=skip_ssl_validation,json=skipSslValidation,proto3" json:"skip_ssl_validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tiller) Reset()         { *m = Tiller{} }
func (m *Tiller) String() string { return proto.CompactTextString(m) }
func (*Tiller) ProtoMessage()    {}
func (*Tiller) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2b6c44db8db459e, []int{1}
}

func (m *Tiller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tiller.Unmarshal(m, b)
}
func (m *Tiller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tiller.Marshal(b, m, deterministic)
}
func (m *Tiller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tiller.Merge(m, src)
}
func (m *Tiller) XXX_Size() int {
	return xxx_messageInfo_Tiller.Size(m)
}
func (m *Tiller) XXX_DiscardUnknown() {
	xxx_messageInfo_Tiller.DiscardUnknown(m)
}

var xxx_messageInfo_Tiller proto.InternalMessageInfo

func (m *Tiller) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Tiller) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Tiller) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Tiller) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Tiller) GetSkipSslValidation() bool {
	if m != nil {
		return m.SkipSslValidation
	}
	return false
}

func init() {
	proto.RegisterType((*Chart)(nil), "softleader.captainkube.v2.Chart")
	proto.RegisterType((*Tiller)(nil), "softleader.captainkube.v2.Tiller")
}

func init() { proto.RegisterFile("chart.proto", fileDescriptor_b2b6c44db8db459e) }

var fileDescriptor_b2b6c44db8db459e = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0x03, 0x21,
	0x18, 0x86, 0x43, 0x6b, 0x6b, 0x0f, 0x8d, 0x89, 0xb8, 0xa0, 0x2e, 0xa4, 0xd3, 0x4d, 0x0c, 0xf5,
	0x1f, 0xe8, 0xee, 0x70, 0x35, 0x0e, 0x5d, 0x2e, 0x94, 0xfb, 0x1a, 0x49, 0x29, 0x10, 0xf8, 0xae,
	0x4d, 0xfa, 0x83, 0xfc, 0x9d, 0x86, 0x6b, 0xb8, 0xc4, 0xa1, 0xe3, 0x93, 0xe7, 0xe5, 0xe5, 0xcb,
	0x4b, 0xef, 0xf4, 0x8f, 0x8a, 0x28, 0x43, 0xf4, 0xe8, 0xd9, 0x73, 0xf2, 0x3b, 0xb4, 0xa0, 0x3a,
	0x88, 0x52, 0xab, 0x80, 0xca, 0xb8, 0x7d, 0xbf, 0x05, 0x79, 0x5c, 0x2d, 0x37, 0x74, 0xf6, 0x91,
	0x93, 0xec, 0x85, 0x2e, 0x76, 0xc6, 0xc2, 0xa7, 0x3a, 0x00, 0x27, 0x82, 0xd4, 0x55, 0x33, 0x32,
	0xe3, 0xf4, 0x56, 0x7b, 0x87, 0xe0, 0x90, 0x4f, 0x04, 0xa9, 0xef, 0x9b, 0x82, 0xec, 0x95, 0x56,
	0x39, 0xd5, 0x26, 0x73, 0x06, 0x3e, 0x15, 0xa4, 0x9e, 0x5e, 0x9e, 0xad, 0xcd, 0x19, 0x96, 0xbf,
	0x84, 0xce, 0xbf, 0x8c, 0xb5, 0x10, 0x73, 0x3b, 0xb8, 0x2e, 0x78, 0xe3, 0xb0, 0xb4, 0x17, 0xce,
	0xae, 0x4f, 0x10, 0x5d, 0xfe, 0x79, 0x72, 0x71, 0x85, 0xb3, 0x0b, 0x2a, 0xa5, 0x93, 0x8f, 0xdd,
	0x50, 0x5f, 0x35, 0x23, 0xe7, 0xab, 0x94, 0xd6, 0xbe, 0x77, 0xc8, 0x6f, 0x06, 0x55, 0x90, 0x49,
	0xfa, 0x94, 0xf6, 0x26, 0xb4, 0x29, 0xd9, 0xf6, 0xa8, 0xac, 0xe9, 0x14, 0x1a, 0xef, 0xf8, 0x4c,
	0x90, 0x7a, 0xd1, 0x3c, 0x66, 0xb5, 0x4e, 0xf6, 0x7b, 0x14, 0xef, 0x2b, 0x2a, 0xf0, 0x24, 0xb5,
	0x3f, 0xc8, 0xab, 0x43, 0x6d, 0x1e, 0xfe, 0xf3, 0x76, 0x3e, 0x4c, 0xfb, 0xf6, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x33, 0x67, 0x9b, 0x69, 0x01, 0x00, 0x00,
}
