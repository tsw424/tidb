// Code generated by protoc-gen-go.
// source: errorpb.proto
// DO NOT EDIT!

/*
Package errorpb is a generated protocol buffer package.

It is generated from these files:
	errorpb.proto

It has these top-level messages:
	NotLeader
	RegionNotFound
	KeyNotInRegion
	StaleEpoch
	Error
*/
package errorpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type NotLeader struct {
	RegionId         *uint64 `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	LeaderStoreId    *uint64 `protobuf:"varint,2,opt,name=leader_store_id" json:"leader_store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NotLeader) Reset()                    { *m = NotLeader{} }
func (m *NotLeader) String() string            { return proto.CompactTextString(m) }
func (*NotLeader) ProtoMessage()               {}
func (*NotLeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NotLeader) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *NotLeader) GetLeaderStoreId() uint64 {
	if m != nil && m.LeaderStoreId != nil {
		return *m.LeaderStoreId
	}
	return 0
}

type RegionNotFound struct {
	RegionId         *uint64 `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegionNotFound) Reset()                    { *m = RegionNotFound{} }
func (m *RegionNotFound) String() string            { return proto.CompactTextString(m) }
func (*RegionNotFound) ProtoMessage()               {}
func (*RegionNotFound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegionNotFound) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

type KeyNotInRegion struct {
	Key              []byte  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	RegionId         *uint64 `protobuf:"varint,2,opt,name=region_id" json:"region_id,omitempty"`
	StartKey         []byte  `protobuf:"bytes,3,opt,name=start_key" json:"start_key,omitempty"`
	EndKey           []byte  `protobuf:"bytes,4,opt,name=end_key" json:"end_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyNotInRegion) Reset()                    { *m = KeyNotInRegion{} }
func (m *KeyNotInRegion) String() string            { return proto.CompactTextString(m) }
func (*KeyNotInRegion) ProtoMessage()               {}
func (*KeyNotInRegion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyNotInRegion) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyNotInRegion) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *KeyNotInRegion) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *KeyNotInRegion) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

type StaleEpoch struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StaleEpoch) Reset()                    { *m = StaleEpoch{} }
func (m *StaleEpoch) String() string            { return proto.CompactTextString(m) }
func (*StaleEpoch) ProtoMessage()               {}
func (*StaleEpoch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Error struct {
	Message          *string         `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	NotLeader        *NotLeader      `protobuf:"bytes,2,opt,name=not_leader" json:"not_leader,omitempty"`
	RegionNotFound   *RegionNotFound `protobuf:"bytes,3,opt,name=region_not_found" json:"region_not_found,omitempty"`
	KeyNotInRegion   *KeyNotInRegion `protobuf:"bytes,4,opt,name=key_not_in_region" json:"key_not_in_region,omitempty"`
	StaleEpoch       *StaleEpoch     `protobuf:"bytes,5,opt,name=stale_epoch" json:"stale_epoch,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Error) GetNotLeader() *NotLeader {
	if m != nil {
		return m.NotLeader
	}
	return nil
}

func (m *Error) GetRegionNotFound() *RegionNotFound {
	if m != nil {
		return m.RegionNotFound
	}
	return nil
}

func (m *Error) GetKeyNotInRegion() *KeyNotInRegion {
	if m != nil {
		return m.KeyNotInRegion
	}
	return nil
}

func (m *Error) GetStaleEpoch() *StaleEpoch {
	if m != nil {
		return m.StaleEpoch
	}
	return nil
}

func init() {
	proto.RegisterType((*NotLeader)(nil), "errorpb.NotLeader")
	proto.RegisterType((*RegionNotFound)(nil), "errorpb.RegionNotFound")
	proto.RegisterType((*KeyNotInRegion)(nil), "errorpb.KeyNotInRegion")
	proto.RegisterType((*StaleEpoch)(nil), "errorpb.StaleEpoch")
	proto.RegisterType((*Error)(nil), "errorpb.Error")
}

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0x81, 0x90, 0x4e, 0x11, 0x64, 0x3d, 0xc0, 0xd1, 0xac, 0x89, 0xe1, 0x44, 0x62,
	0x2f, 0xfe, 0x02, 0x4c, 0x8c, 0xc6, 0x83, 0x7a, 0xdf, 0x54, 0x3b, 0x62, 0x23, 0xee, 0x34, 0xb3,
	0xeb, 0x81, 0x9f, 0xe9, 0x3f, 0xb2, 0x33, 0x45, 0xd0, 0xc6, 0xe3, 0xce, 0xfb, 0xde, 0xcc, 0x7b,
	0x0b, 0xc7, 0xc8, 0x4c, 0x5c, 0x3f, 0x2f, 0x6b, 0xa6, 0x48, 0x66, 0xb8, 0x7b, 0xda, 0x2b, 0x48,
	0xef, 0x29, 0xde, 0x61, 0x51, 0x22, 0x9b, 0x29, 0xa4, 0x8c, 0xeb, 0x8a, 0xbc, 0xab, 0xca, 0x79,
	0x72, 0x96, 0x2c, 0xfa, 0x66, 0x06, 0x93, 0x8d, 0x8a, 0x2e, 0x44, 0x62, 0x14, 0xe1, 0x48, 0x04,
	0x7b, 0x0e, 0xe3, 0x07, 0x65, 0x1b, 0xfb, 0x35, 0x7d, 0xfa, 0xf2, 0x1f, 0xb7, 0x7d, 0x82, 0xf1,
	0x2d, 0x6e, 0x1b, 0xe2, 0xc6, 0xb7, 0xb0, 0xc9, 0xa0, 0xf7, 0x8e, 0x5b, 0x95, 0x47, 0x7f, 0x1d,
	0xba, 0x56, 0x46, 0x21, 0x16, 0x1c, 0x9d, 0x50, 0x3d, 0xa5, 0x26, 0x30, 0x44, 0x5f, 0xea, 0xa0,
	0x2f, 0x03, 0x3b, 0x02, 0x78, 0x8c, 0xc5, 0x06, 0x57, 0x35, 0xbd, 0xbc, 0xd9, 0xaf, 0x04, 0x06,
	0x2b, 0x69, 0x23, 0xe0, 0x07, 0x86, 0x50, 0xac, 0x51, 0xf7, 0xa7, 0xe6, 0x02, 0xc0, 0x53, 0x74,
	0x6d, 0x01, 0x3d, 0x90, 0xe5, 0x66, 0xf9, 0xf3, 0x13, 0x87, 0xde, 0x97, 0x70, 0xb2, 0xcb, 0x21,
	0xf8, 0xab, 0xb4, 0xd1, 0xdb, 0x59, 0x3e, 0xdb, 0xd3, 0x9d, 0xb2, 0x39, 0x4c, 0x9b, 0x40, 0xca,
	0x57, 0xde, 0xb5, 0x6e, 0x8d, 0xf7, 0xdb, 0xd3, 0xe9, 0xbe, 0x80, 0x2c, 0x48, 0x6e, 0x87, 0x12,
	0x7c, 0x3e, 0x50, 0xfa, 0x74, 0x4f, 0x1f, 0x3a, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x19, 0x25,
	0x10, 0x74, 0xae, 0x01, 0x00, 0x00,
}