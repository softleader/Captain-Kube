// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rmc.proto

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

type RmcRequest struct {
	Chart                *Chart   `protobuf:"bytes,1,opt,name=chart,proto3" json:"chart,omitempty"`
	Set                  []string `protobuf:"bytes,2,rep,name=set,proto3" json:"set,omitempty"`
	Constraint           string   `protobuf:"bytes,3,opt,name=constraint,proto3" json:"constraint,omitempty"`
	Verbose              bool     `protobuf:"varint,4,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Timeout              int64    `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Color                bool     `protobuf:"varint,6,opt,name=color,proto3" json:"color,omitempty"`
	Force                bool     `protobuf:"varint,7,opt,name=force,proto3" json:"force,omitempty"`
	DryRun               bool     `protobuf:"varint,8,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmcRequest) Reset()         { *m = RmcRequest{} }
func (m *RmcRequest) String() string { return proto.CompactTextString(m) }
func (*RmcRequest) ProtoMessage()    {}
func (*RmcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35e41f379f3a71cc, []int{0}
}

func (m *RmcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmcRequest.Unmarshal(m, b)
}
func (m *RmcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmcRequest.Marshal(b, m, deterministic)
}
func (m *RmcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmcRequest.Merge(m, src)
}
func (m *RmcRequest) XXX_Size() int {
	return xxx_messageInfo_RmcRequest.Size(m)
}
func (m *RmcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RmcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RmcRequest proto.InternalMessageInfo

func (m *RmcRequest) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *RmcRequest) GetSet() []string {
	if m != nil {
		return m.Set
	}
	return nil
}

func (m *RmcRequest) GetConstraint() string {
	if m != nil {
		return m.Constraint
	}
	return ""
}

func (m *RmcRequest) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *RmcRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RmcRequest) GetColor() bool {
	if m != nil {
		return m.Color
	}
	return false
}

func (m *RmcRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *RmcRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func init() {
	proto.RegisterType((*RmcRequest)(nil), "softleader.captainkube.v2.RmcRequest")
}

func init() { proto.RegisterFile("rmc.proto", fileDescriptor_35e41f379f3a71cc) }

var fileDescriptor_35e41f379f3a71cc = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xb5, 0xdd, 0xed, 0x2c, 0x88, 0x04, 0x91, 0xe8, 0x41, 0x82, 0xa7, 0x9c, 0x72,
	0xa8, 0xe0, 0x03, 0xe8, 0x1b, 0xe4, 0xe8, 0x2d, 0xcd, 0xce, 0x62, 0x71, 0xdb, 0x59, 0x93, 0xe9,
	0x8a, 0x2f, 0xed, 0x33, 0x48, 0x53, 0x17, 0xdc, 0x43, 0x6f, 0xf9, 0xbe, 0xff, 0xcf, 0xc0, 0x0c,
	0xd4, 0xb1, 0x0f, 0xf6, 0x10, 0x89, 0x49, 0xde, 0x25, 0xda, 0xf1, 0x1e, 0xfd, 0x16, 0xa3, 0x0d,
	0xfe, 0xc0, 0xbe, 0x1b, 0x3e, 0xc6, 0x16, 0xed, 0xb1, 0xb9, 0xdf, 0x84, 0x77, 0x1f, 0x79, 0xee,
	0x3d, 0xfe, 0x08, 0x00, 0xd7, 0x07, 0x87, 0x9f, 0x23, 0x26, 0x96, 0xcf, 0x50, 0xe6, 0x54, 0x09,
	0x2d, 0xcc, 0xa6, 0xd1, 0x76, 0x71, 0x8c, 0x7d, 0x9d, 0x7a, 0x6e, 0xae, 0xcb, 0x6b, 0x28, 0x12,
	0xb2, 0xba, 0xd0, 0x85, 0xa9, 0xdd, 0xf4, 0x94, 0x0f, 0x00, 0x81, 0x86, 0xc4, 0xd1, 0x77, 0x03,
	0xab, 0x42, 0x0b, 0x53, 0xbb, 0x7f, 0x46, 0x2a, 0x58, 0x1d, 0x31, 0xb6, 0x94, 0x50, 0x5d, 0x6a,
	0x61, 0xd6, 0xee, 0x84, 0x53, 0xc2, 0x5d, 0x8f, 0x34, 0xb2, 0x2a, 0xb5, 0x30, 0x85, 0x3b, 0xa1,
	0xbc, 0x81, 0x32, 0xd0, 0x9e, 0xa2, 0xaa, 0xf2, 0x8f, 0x19, 0x26, 0xbb, 0xa3, 0x18, 0x50, 0xad,
	0x66, 0x9b, 0x41, 0xde, 0x42, 0xb5, 0x8d, 0xdf, 0x6e, 0x1c, 0xd4, 0x3a, 0xeb, 0x3f, 0x7a, 0x69,
	0x40, 0xf3, 0x97, 0x0d, 0xd4, 0x2f, 0xaf, 0xf6, 0x76, 0x75, 0xce, 0x6d, 0x95, 0x6f, 0xf5, 0xf4,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x59, 0xd7, 0x32, 0x55, 0x60, 0x01, 0x00, 0x00,
}