// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chart.proto

package tw_com_softleader

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
	FileSize             int64    `protobuf:"varint,3,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
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
	SkipSslValidation    bool     `protobuf:"varint,5,opt,name=skipSslValidation,proto3" json:"skipSslValidation,omitempty"`
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
	proto.RegisterType((*Chart)(nil), "tw.com.softleader.Chart")
	proto.RegisterType((*Tiller)(nil), "tw.com.softleader.Tiller")
}

func init() { proto.RegisterFile("chart.proto", fileDescriptor_b2b6c44db8db459e) }

var fileDescriptor_b2b6c44db8db459e = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0x04, 0x21,
	0x10, 0x86, 0xc3, 0x9d, 0x77, 0xde, 0xa1, 0xcd, 0x51, 0x11, 0x2b, 0x72, 0x15, 0x85, 0xd9, 0xc6,
	0x47, 0xb0, 0xb7, 0xe0, 0x8c, 0x89, 0x25, 0xc2, 0x5c, 0x24, 0xb2, 0x0c, 0x81, 0xd9, 0x6c, 0xe2,
	0xf3, 0xf8, 0xa0, 0x86, 0x35, 0x6c, 0x63, 0xf9, 0xe5, 0xfb, 0xe7, 0x9f, 0xc9, 0xf0, 0x3b, 0xf7,
	0x69, 0x0b, 0x0d, 0xb9, 0x20, 0xa1, 0x38, 0xd1, 0x3c, 0x38, 0x1c, 0x87, 0x8a, 0x57, 0x8a, 0x60,
	0x3d, 0x94, 0xf3, 0x3b, 0xdf, 0x3d, 0xb7, 0x84, 0x78, 0xe0, 0x87, 0x6b, 0x88, 0xf0, 0x62, 0x47,
	0x90, 0x4c, 0x31, 0x7d, 0x34, 0x2b, 0x0b, 0xc9, 0x6f, 0x1d, 0x26, 0x82, 0x44, 0x72, 0xa3, 0x98,
	0xbe, 0x37, 0x1d, 0xfb, 0xd4, 0x25, 0x7c, 0x83, 0xdc, 0x2a, 0xa6, 0xb7, 0x66, 0xe5, 0xf3, 0x0f,
	0xe3, 0xfb, 0xd7, 0x10, 0x23, 0x94, 0x16, 0x83, 0xe4, 0x33, 0x86, 0x44, 0xbd, 0xbc, 0x73, 0x73,
	0x53, 0x85, 0x92, 0xda, 0xe2, 0xcd, 0x9f, 0xeb, 0xdc, 0x5c, 0xb6, 0xb5, 0xce, 0x58, 0xfc, 0x52,
	0x7f, 0x34, 0x2b, 0xb7, 0xa3, 0xac, 0x73, 0x38, 0x25, 0x92, 0x37, 0x8b, 0xea, 0x28, 0x1e, 0xf9,
	0xa9, 0x7e, 0x85, 0x7c, 0xa9, 0xf1, 0xcd, 0xc6, 0xe0, 0x2d, 0x05, 0x4c, 0x72, 0xa7, 0x98, 0x3e,
	0x98, 0xff, 0xe2, 0x63, 0xbf, 0xfc, 0xe6, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x64, 0x9e, 0xcd,
	0xe7, 0x2a, 0x01, 0x00, 0x00,
}
