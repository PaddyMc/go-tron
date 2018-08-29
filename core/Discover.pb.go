// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core 
//import "github.com/tronprotocol/grpc-gateway/core"

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

type Endpoint struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodeId               []byte   `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Endpoint `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Version              int32     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{1}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Echo                 int32     `protobuf:"varint,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{2}
}
func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (dst *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(dst, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TargetId             []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindNeighbours) Reset()         { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()    {}
func (*FindNeighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{3}
}
func (m *FindNeighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNeighbours.Unmarshal(m, b)
}
func (m *FindNeighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNeighbours.Marshal(b, m, deterministic)
}
func (dst *FindNeighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNeighbours.Merge(dst, src)
}
func (m *FindNeighbours) XXX_Size() int {
	return xxx_messageInfo_FindNeighbours.Size(m)
}
func (m *FindNeighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNeighbours.DiscardUnknown(m)
}

var xxx_messageInfo_FindNeighbours proto.InternalMessageInfo

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From                 *Endpoint   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Neighbours           []*Endpoint `protobuf:"bytes,2,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	Timestamp            int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Neighbours) Reset()         { *m = Neighbours{} }
func (m *Neighbours) String() string { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()    {}
func (*Neighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{4}
}
func (m *Neighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbours.Unmarshal(m, b)
}
func (m *Neighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbours.Marshal(b, m, deterministic)
}
func (dst *Neighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbours.Merge(dst, src)
}
func (m *Neighbours) XXX_Size() int {
	return xxx_messageInfo_Neighbours.Size(m)
}
func (m *Neighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbours.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbours proto.InternalMessageInfo

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BackupMessage struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupMessage) Reset()         { *m = BackupMessage{} }
func (m *BackupMessage) String() string { return proto.CompactTextString(m) }
func (*BackupMessage) ProtoMessage()    {}
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_8a0ee7b57601823c, []int{5}
}
func (m *BackupMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupMessage.Unmarshal(m, b)
}
func (m *BackupMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupMessage.Marshal(b, m, deterministic)
}
func (dst *BackupMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMessage.Merge(dst, src)
}
func (m *BackupMessage) XXX_Size() int {
	return xxx_messageInfo_BackupMessage.Size(m)
}
func (m *BackupMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMessage proto.InternalMessageInfo

func (m *BackupMessage) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *BackupMessage) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
	proto.RegisterType((*BackupMessage)(nil), "protocol.BackupMessage")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor_Discover_8a0ee7b57601823c) }

var fileDescriptor_Discover_8a0ee7b57601823c = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xc9, 0x9f, 0xf6, 0xcd, 0x3b, 0xed, 0xfb, 0x0a, 0x2b, 0x48, 0x10, 0x0f, 0x21, 0x07,
	0xa9, 0x07, 0x13, 0xa8, 0x1f, 0x40, 0x28, 0x5a, 0xe8, 0x41, 0x29, 0x39, 0x7a, 0xdb, 0x26, 0xdb,
	0xed, 0x62, 0xb3, 0x13, 0x76, 0xb7, 0x95, 0x7e, 0x01, 0xef, 0x7e, 0x63, 0xc9, 0xda, 0xad, 0x7f,
	0x50, 0xaa, 0x9e, 0x32, 0x4f, 0xe6, 0x49, 0x7e, 0xcf, 0x0c, 0x03, 0x87, 0x25, 0x2a, 0x96, 0x5f,
	0x09, 0x5d, 0xe2, 0x9a, 0xa9, 0xac, 0x51, 0x68, 0x90, 0x44, 0xf6, 0x51, 0xe2, 0x32, 0x9d, 0x42,
	0x74, 0x2d, 0xab, 0x06, 0x85, 0x34, 0x24, 0x86, 0x3f, 0xb4, 0xaa, 0x14, 0xd3, 0x3a, 0xf6, 0x12,
	0x6f, 0xd0, 0x2f, 0x9c, 0x24, 0x04, 0xc2, 0x06, 0x95, 0x89, 0xfd, 0xc4, 0x1b, 0x74, 0x0a, 0x5b,
	0x93, 0x23, 0xe8, 0x4a, 0xac, 0xd8, 0xa4, 0x8a, 0x03, 0x6b, 0xde, 0xaa, 0xf4, 0xc9, 0x83, 0xde,
	0x54, 0x48, 0x7e, 0xc3, 0xb4, 0xa6, 0x9c, 0x91, 0x53, 0x08, 0xe7, 0x0a, 0x6b, 0xfb, 0xcb, 0xde,
	0x90, 0x64, 0x0e, 0x9d, 0x39, 0x6e, 0x61, 0xfb, 0x24, 0x05, 0xdf, 0xa0, 0x25, 0x7c, 0xee, 0xf2,
	0x0d, 0xb6, 0x09, 0xd7, 0x4c, 0x69, 0x81, 0xd2, 0x42, 0x3b, 0x85, 0x93, 0xe4, 0x04, 0xfe, 0x1a,
	0x51, 0x33, 0x6d, 0x68, 0xdd, 0xc4, 0x61, 0xe2, 0x0d, 0x82, 0xe2, 0xf5, 0x45, 0xca, 0xa1, 0x37,
	0xc5, 0x9f, 0x47, 0x22, 0x10, 0xb2, 0x72, 0x81, 0x6e, 0xec, 0xb6, 0x7e, 0x0f, 0x0a, 0x3e, 0x82,
	0x14, 0xfc, 0x1f, 0x0b, 0x59, 0xdd, 0x32, 0xc1, 0x17, 0x33, 0x5c, 0x29, 0xfd, 0x6d, 0xd6, 0x31,
	0x44, 0x86, 0x2a, 0xce, 0xcc, 0xa4, 0xb2, 0xbc, 0x7e, 0xb1, 0xd3, 0x7b, 0x98, 0x8f, 0x1e, 0xc0,
	0x2f, 0x80, 0x43, 0x00, 0xb9, 0xfb, 0x2a, 0xf6, 0x93, 0xe0, 0x0b, 0xf7, 0x1b, 0xd7, 0x9e, 0x20,
	0x97, 0xf0, 0x6f, 0x44, 0xcb, 0xfb, 0x55, 0xe3, 0xf6, 0x4c, 0x20, 0x9c, 0x2f, 0x29, 0xb7, 0x51,
	0xa2, 0xc2, 0xd6, 0xed, 0x9c, 0x8d, 0x12, 0xa8, 0x84, 0xd9, 0x6c, 0xf7, 0xba, 0xd3, 0xa3, 0x31,
	0x1c, 0xa0, 0xe2, 0x99, 0x51, 0x28, 0x5f, 0x82, 0xe8, 0x51, 0xe4, 0x2e, 0xf7, 0xee, 0x8c, 0x0b,
	0xb3, 0x58, 0xcd, 0xb2, 0x12, 0xeb, 0xbc, 0x75, 0xb8, 0xa4, 0x39, 0x57, 0x4d, 0x79, 0xce, 0xa9,
	0x61, 0x0f, 0x74, 0x93, 0xb7, 0xd7, 0x3e, 0xeb, 0xda, 0xde, 0xc5, 0x73, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4a, 0x9a, 0x1b, 0xf5, 0xfc, 0x02, 0x00, 0x00,
}
