// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/common.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OK struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OK) Reset()         { *m = OK{} }
func (m *OK) String() string { return proto.CompactTextString(m) }
func (*OK) ProtoMessage()    {}
func (*OK) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_997dcce0606fc3e6, []int{0}
}
func (m *OK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OK.Unmarshal(m, b)
}
func (m *OK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OK.Marshal(b, m, deterministic)
}
func (dst *OK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OK.Merge(dst, src)
}
func (m *OK) XXX_Size() int {
	return xxx_messageInfo_OK.Size(m)
}
func (m *OK) XXX_DiscardUnknown() {
	xxx_messageInfo_OK.DiscardUnknown(m)
}

var xxx_messageInfo_OK proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OK)(nil), "pb.OK")
}

func init() { proto.RegisterFile("pb/common.proto", fileDescriptor_common_997dcce0606fc3e6) }

var fileDescriptor_common_997dcce0606fc3e6 = []byte{
	// 56 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0xd2, 0x4f,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x62, 0xe1, 0x62, 0xf2, 0xf7, 0x4e, 0x62, 0x03, 0x0b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x2d, 0x65, 0xeb, 0x23, 0x00, 0x00, 0x00,
}
