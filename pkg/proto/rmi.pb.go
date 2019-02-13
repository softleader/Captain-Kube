// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rmi.proto

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

type RmiRequest struct {
	Images               []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	Verbose              bool     `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Timeout              int64    `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Color                bool     `protobuf:"varint,5,opt,name=color,proto3" json:"color,omitempty"`
	Force                bool     `protobuf:"varint,6,opt,name=force,proto3" json:"force,omitempty"`
	DryRun               bool     `protobuf:"varint,7,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmiRequest) Reset()         { *m = RmiRequest{} }
func (m *RmiRequest) String() string { return proto.CompactTextString(m) }
func (*RmiRequest) ProtoMessage()    {}
func (*RmiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_581e02de60ed53de, []int{0}
}

func (m *RmiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmiRequest.Unmarshal(m, b)
}
func (m *RmiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmiRequest.Marshal(b, m, deterministic)
}
func (m *RmiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmiRequest.Merge(m, src)
}
func (m *RmiRequest) XXX_Size() int {
	return xxx_messageInfo_RmiRequest.Size(m)
}
func (m *RmiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RmiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RmiRequest proto.InternalMessageInfo

func (m *RmiRequest) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *RmiRequest) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *RmiRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RmiRequest) GetColor() bool {
	if m != nil {
		return m.Color
	}
	return false
}

func (m *RmiRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *RmiRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func init() {
	proto.RegisterType((*RmiRequest)(nil), "softleader.captainkube.v2.RmiRequest")
}

func init() { proto.RegisterFile("rmi.proto", fileDescriptor_581e02de60ed53de) }

var fileDescriptor_581e02de60ed53de = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x85, 0xba, 0x70, 0x95, 0x18, 0x2c, 0x84, 0x0c, 0x93, 0xc5, 0x94, 0xc9, 0x43,
	0x58, 0x98, 0xd9, 0x58, 0x3d, 0xb2, 0x39, 0xee, 0x15, 0x59, 0xd4, 0xbd, 0xe0, 0x8f, 0x20, 0xfe,
	0x1c, 0xbf, 0x0d, 0xc5, 0x49, 0x06, 0x86, 0x8c, 0xcf, 0xfb, 0x3e, 0x77, 0xba, 0x83, 0x9b, 0x18,
	0xbc, 0x1e, 0x22, 0x65, 0x12, 0x0f, 0x89, 0x4e, 0xf9, 0x8c, 0xf6, 0x88, 0x51, 0x3b, 0x3b, 0x64,
	0xeb, 0x2f, 0x9f, 0xa5, 0x47, 0x3d, 0x76, 0x8f, 0x07, 0x1f, 0xec, 0x07, 0xce, 0xde, 0xd3, 0x2f,
	0x03, 0x30, 0xc1, 0x1b, 0xfc, 0x2a, 0x98, 0xb2, 0x78, 0x01, 0x5e, 0xdb, 0x24, 0x99, 0x6a, 0xda,
	0x43, 0xa7, 0xf4, 0xe6, 0x1e, 0xfd, 0x36, 0x89, 0x66, 0xf1, 0x85, 0x84, 0xfd, 0x88, 0xb1, 0xa7,
	0x84, 0xb2, 0x51, 0xac, 0xbd, 0x36, 0x2b, 0x4e, 0x4d, 0xf6, 0x01, 0xa9, 0x64, 0x79, 0xa5, 0x58,
	0xdb, 0x98, 0x15, 0xc5, 0x1d, 0xec, 0x1c, 0x9d, 0x29, 0xca, 0x5d, 0x9d, 0x98, 0x61, 0x4a, 0x4f,
	0x14, 0x1d, 0x4a, 0x3e, 0xa7, 0x15, 0xc4, 0x3d, 0xf0, 0x63, 0xfc, 0x31, 0xe5, 0x22, 0xf7, 0x35,
	0x5e, 0xe8, 0xb5, 0x03, 0x95, 0xbf, 0xb5, 0xa3, 0xb0, 0x7d, 0xe9, 0xfb, 0xed, 0x7f, 0xee, 0x79,
	0xfd, 0xfd, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xe6, 0xa5, 0xbd, 0x30, 0x01, 0x00, 0x00,
}