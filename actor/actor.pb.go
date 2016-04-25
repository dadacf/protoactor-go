// Code generated by protoc-gen-go.
// source: actor.proto
// DO NOT EDIT!

/*
Package actor is a generated protocol buffer package.

It is generated from these files:
	actor.proto

It has these top-level messages:
	PID
	Restarting
	Stopping
	Stopped
	PoisonPill
	Started
*/
package actor

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

type PID struct {
	Host string `protobuf:"bytes,1,opt,name=Host,json=host" json:"Host,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *PID) Reset()                    { *m = PID{} }
func (m *PID) String() string            { return proto.CompactTextString(m) }
func (*PID) ProtoMessage()               {}
func (*PID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// user messages
type Restarting struct {
}

func (m *Restarting) Reset()                    { *m = Restarting{} }
func (m *Restarting) String() string            { return proto.CompactTextString(m) }
func (*Restarting) ProtoMessage()               {}
func (*Restarting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Stopping struct {
}

func (m *Stopping) Reset()                    { *m = Stopping{} }
func (m *Stopping) String() string            { return proto.CompactTextString(m) }
func (*Stopping) ProtoMessage()               {}
func (*Stopping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Stopped struct {
}

func (m *Stopped) Reset()                    { *m = Stopped{} }
func (m *Stopped) String() string            { return proto.CompactTextString(m) }
func (*Stopped) ProtoMessage()               {}
func (*Stopped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PoisonPill struct {
}

func (m *PoisonPill) Reset()                    { *m = PoisonPill{} }
func (m *PoisonPill) String() string            { return proto.CompactTextString(m) }
func (*PoisonPill) ProtoMessage()               {}
func (*PoisonPill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Started struct {
}

func (m *Started) Reset()                    { *m = Started{} }
func (m *Started) String() string            { return proto.CompactTextString(m) }
func (*Started) ProtoMessage()               {}
func (*Started) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*PID)(nil), "actor.PID")
	proto.RegisterType((*Restarting)(nil), "actor.Restarting")
	proto.RegisterType((*Stopping)(nil), "actor.Stopping")
	proto.RegisterType((*Stopped)(nil), "actor.Stopped")
	proto.RegisterType((*PoisonPill)(nil), "actor.PoisonPill")
	proto.RegisterType((*Started)(nil), "actor.Started")
}

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x2e, 0xc9,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x34, 0xb9, 0x98, 0x03,
	0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0x3c, 0xf2, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x58, 0x32, 0x80, 0x6c, 0x21, 0x3e, 0x2e, 0x26, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0x08, 0x53,
	0x66, 0x8a, 0x12, 0x0f, 0x17, 0x57, 0x50, 0x6a, 0x71, 0x49, 0x62, 0x51, 0x49, 0x66, 0x5e, 0xba,
	0x12, 0x17, 0x17, 0x47, 0x70, 0x49, 0x7e, 0x41, 0x01, 0x88, 0xcd, 0xc9, 0xc5, 0x0e, 0x66, 0xa7,
	0x82, 0x15, 0x05, 0xe4, 0x67, 0x16, 0xe7, 0xe7, 0x05, 0x64, 0xe6, 0xe4, 0x40, 0x24, 0x80, 0x1a,
	0x52, 0x53, 0x92, 0xd8, 0xc0, 0xd6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x88, 0x9f,
	0x3e, 0x85, 0x00, 0x00, 0x00,
}
