// Code generated by protoc-gen-go. DO NOT EDIT.
// source: primitive.proto

/*
Package test_proto is a generated protocol buffer package.

It is generated from these files:
	primitive.proto
	with_any.proto

It has these top-level messages:
	Primitives
	SimpleMessage
	ContainerMessage
	WithAny
	TestMessage
*/
package test_proto

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

type Primitives_Corpus int32

const (
	Primitives_UNIVERSAL Primitives_Corpus = 0
	Primitives_WEB       Primitives_Corpus = 1
	Primitives_IMAGES    Primitives_Corpus = 2
	Primitives_LOCAL     Primitives_Corpus = 3
	Primitives_NEWS      Primitives_Corpus = 4
	Primitives_PRODUCTS  Primitives_Corpus = 5
	Primitives_VIDEO     Primitives_Corpus = 6
)

var Primitives_Corpus_name = map[int32]string{
	0: "UNIVERSAL",
	1: "WEB",
	2: "IMAGES",
	3: "LOCAL",
	4: "NEWS",
	5: "PRODUCTS",
	6: "VIDEO",
}
var Primitives_Corpus_value = map[string]int32{
	"UNIVERSAL": 0,
	"WEB":       1,
	"IMAGES":    2,
	"LOCAL":     3,
	"NEWS":      4,
	"PRODUCTS":  5,
	"VIDEO":     6,
}

func (x Primitives_Corpus) String() string {
	return proto.EnumName(Primitives_Corpus_name, int32(x))
}
func (Primitives_Corpus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Primitives struct {
	One      string                    `protobuf:"bytes,1,opt,name=one" json:"one,omitempty"`
	Two      int64                     `protobuf:"varint,2,opt,name=two" json:"two,omitempty"`
	Three    float64                   `protobuf:"fixed64,3,opt,name=three" json:"three,omitempty"`
	Four     bool                      `protobuf:"varint,4,opt,name=four" json:"four,omitempty"`
	Corpus   Primitives_Corpus         `protobuf:"varint,5,opt,name=corpus,enum=test_proto.Primitives_Corpus" json:"corpus,omitempty"`
	Rep      []int64                   `protobuf:"varint,6,rep,packed,name=rep" json:"rep,omitempty"`
	MapField map[int64]string          `protobuf:"bytes,7,rep,name=map_field,json=mapField" json:"map_field,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Sm       *SimpleMessage            `protobuf:"bytes,8,opt,name=sm" json:"sm,omitempty"`
	Rsm      []*SimpleMessage          `protobuf:"bytes,9,rep,name=rsm" json:"rsm,omitempty"`
	Msm      map[string]*SimpleMessage `protobuf:"bytes,10,rep,name=msm" json:"msm,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Primitives) Reset()                    { *m = Primitives{} }
func (m *Primitives) String() string            { return proto.CompactTextString(m) }
func (*Primitives) ProtoMessage()               {}
func (*Primitives) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Primitives) GetOne() string {
	if m != nil {
		return m.One
	}
	return ""
}

func (m *Primitives) GetTwo() int64 {
	if m != nil {
		return m.Two
	}
	return 0
}

func (m *Primitives) GetThree() float64 {
	if m != nil {
		return m.Three
	}
	return 0
}

func (m *Primitives) GetFour() bool {
	if m != nil {
		return m.Four
	}
	return false
}

func (m *Primitives) GetCorpus() Primitives_Corpus {
	if m != nil {
		return m.Corpus
	}
	return Primitives_UNIVERSAL
}

func (m *Primitives) GetRep() []int64 {
	if m != nil {
		return m.Rep
	}
	return nil
}

func (m *Primitives) GetMapField() map[int64]string {
	if m != nil {
		return m.MapField
	}
	return nil
}

func (m *Primitives) GetSm() *SimpleMessage {
	if m != nil {
		return m.Sm
	}
	return nil
}

func (m *Primitives) GetRsm() []*SimpleMessage {
	if m != nil {
		return m.Rsm
	}
	return nil
}

func (m *Primitives) GetMsm() map[string]*SimpleMessage {
	if m != nil {
		return m.Msm
	}
	return nil
}

type SimpleMessage struct {
	One int64 `protobuf:"varint,1,opt,name=one" json:"one,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SimpleMessage) GetOne() int64 {
	if m != nil {
		return m.One
	}
	return 0
}

type ContainerMessage struct {
	Cm  *SimpleMessage `protobuf:"bytes,1,opt,name=cm" json:"cm,omitempty"`
	Two int64          `protobuf:"varint,2,opt,name=two" json:"two,omitempty"`
}

func (m *ContainerMessage) Reset()                    { *m = ContainerMessage{} }
func (m *ContainerMessage) String() string            { return proto.CompactTextString(m) }
func (*ContainerMessage) ProtoMessage()               {}
func (*ContainerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContainerMessage) GetCm() *SimpleMessage {
	if m != nil {
		return m.Cm
	}
	return nil
}

func (m *ContainerMessage) GetTwo() int64 {
	if m != nil {
		return m.Two
	}
	return 0
}

func init() {
	proto.RegisterType((*Primitives)(nil), "test_proto.Primitives")
	proto.RegisterType((*SimpleMessage)(nil), "test_proto.SimpleMessage")
	proto.RegisterType((*ContainerMessage)(nil), "test_proto.ContainerMessage")
	proto.RegisterEnum("test_proto.Primitives_Corpus", Primitives_Corpus_name, Primitives_Corpus_value)
}

func init() { proto.RegisterFile("primitive.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x71, 0x9c, 0x66, 0xc9, 0x8d, 0x82, 0x75, 0xe2, 0xc1, 0x4c, 0x42, 0x84, 0x8a, 0x87,
	0x4c, 0x48, 0x45, 0x14, 0x21, 0x21, 0x78, 0x2a, 0x5d, 0x40, 0x95, 0xda, 0x75, 0x38, 0x6c, 0x93,
	0x78, 0x99, 0x42, 0xf1, 0x20, 0xa2, 0x6e, 0x22, 0xdb, 0x1d, 0xda, 0x87, 0xe5, 0xbb, 0x20, 0x3b,
	0xab, 0xda, 0x69, 0x5a, 0xdf, 0xfe, 0x97, 0xfc, 0xee, 0xfe, 0xe7, 0xbb, 0x83, 0xc7, 0x8d, 0xae,
	0x54, 0x65, 0xab, 0x2b, 0xd9, 0x6f, 0x74, 0x6d, 0x6b, 0x04, 0x2b, 0x8d, 0xbd, 0xf0, 0xba, 0xf7,
	0x2f, 0x04, 0x38, 0x59, 0xff, 0x37, 0xc8, 0x80, 0xd6, 0x4b, 0xc9, 0x49, 0x4a, 0xb2, 0x44, 0x38,
	0xe9, 0xbe, 0xd8, 0xbf, 0x35, 0x0f, 0x52, 0x92, 0x51, 0xe1, 0x24, 0x3e, 0x81, 0x8e, 0xfd, 0xad,
	0xa5, 0xe4, 0x34, 0x25, 0x19, 0x11, 0x6d, 0x80, 0x08, 0xe1, 0x65, 0xbd, 0xd2, 0x3c, 0x4c, 0x49,
	0x16, 0x0b, 0xaf, 0xf1, 0x1d, 0x44, 0xf3, 0x5a, 0x37, 0x2b, 0xc3, 0x3b, 0x29, 0xc9, 0x1e, 0x0d,
	0x9e, 0xf5, 0x37, 0xce, 0xfd, 0x8d, 0x6b, 0x7f, 0xe4, 0x21, 0x71, 0x03, 0x3b, 0x4b, 0x2d, 0x1b,
	0x1e, 0xa5, 0xd4, 0x59, 0x6a, 0xd9, 0xe0, 0x10, 0x12, 0x55, 0x36, 0x17, 0x97, 0x95, 0x5c, 0xfc,
	0xe4, 0x7b, 0x29, 0xcd, 0xf6, 0x07, 0x2f, 0xef, 0xa9, 0x35, 0x2d, 0x9b, 0xcf, 0x0e, 0xcb, 0x97,
	0x56, 0x5f, 0x8b, 0x58, 0xdd, 0x84, 0x78, 0x08, 0x81, 0x51, 0x3c, 0x4e, 0x49, 0xb6, 0x3f, 0x78,
	0xba, 0x9d, 0x5b, 0x54, 0xaa, 0x59, 0xc8, 0xa9, 0x34, 0xa6, 0xfc, 0x25, 0x45, 0x60, 0x14, 0xbe,
	0x02, 0xaa, 0x8d, 0xe2, 0x89, 0xf7, 0xd9, 0xc1, 0x3a, 0x0a, 0xdf, 0x00, 0x55, 0x46, 0x71, 0xf0,
	0xf0, 0xf3, 0xfb, 0x9a, 0x32, 0xaa, 0xed, 0xc7, 0xb1, 0x07, 0x1f, 0xa1, 0x7b, 0xab, 0x4b, 0xf7,
	0xe0, 0x3f, 0xf2, 0xda, 0x4f, 0x9d, 0x0a, 0x27, 0xdd, 0x8c, 0xaf, 0xca, 0xc5, 0x4a, 0xfa, 0xb9,
	0x27, 0xa2, 0x0d, 0x3e, 0x04, 0xef, 0xc9, 0xc1, 0x57, 0x88, 0xd7, 0xd5, 0xb6, 0xf3, 0x92, 0x36,
	0xef, 0xf5, 0x76, 0xde, 0xce, 0xe6, 0x37, 0x25, 0x7b, 0xdf, 0x21, 0x6a, 0x37, 0x80, 0x5d, 0x48,
	0x4e, 0x8f, 0xc7, 0x67, 0xb9, 0x28, 0x86, 0x13, 0xf6, 0x00, 0xf7, 0x80, 0x9e, 0xe7, 0x9f, 0x18,
	0x41, 0x80, 0x68, 0x3c, 0x1d, 0x7e, 0xc9, 0x0b, 0x16, 0x60, 0x02, 0x9d, 0xc9, 0x6c, 0x34, 0x9c,
	0x30, 0x8a, 0x31, 0x84, 0xc7, 0xf9, 0x79, 0xc1, 0x42, 0x7c, 0x08, 0xf1, 0x89, 0x98, 0x1d, 0x9d,
	0x8e, 0xbe, 0x15, 0xac, 0xe3, 0x90, 0xb3, 0xf1, 0x51, 0x3e, 0x63, 0x51, 0xef, 0x05, 0x74, 0x6f,
	0xf9, 0x6e, 0x5f, 0x18, 0xf5, 0x17, 0xd6, 0x9b, 0x01, 0x1b, 0xd5, 0x4b, 0x5b, 0x56, 0x4b, 0xa9,
	0xd7, 0xd4, 0x21, 0x04, 0x73, 0xe5, 0xa1, 0xdd, 0xdb, 0x9a, 0xab, 0xbb, 0x07, 0xfa, 0x23, 0xf2,
	0xe8, 0xdb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xdf, 0xf9, 0x11, 0xf9, 0x02, 0x00, 0x00,
}