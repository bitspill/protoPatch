// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch.proto

/*
Package patch is a generated protocol buffer package.

It is generated from these files:
	patch.proto

It has these top-level messages:
	ProtoPatch
	ProtoOp
	ProtoStep
*/
package patch

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoAction int32

const (
	ProtoAction_ActionInvalid    ProtoAction = 0
	ProtoAction_ActionReplace    ProtoAction = 1
	ProtoAction_ActionAppend     ProtoAction = 2
	ProtoAction_ActionRemove     ProtoAction = 3
	ProtoAction_ActionRemoveOne  ProtoAction = 4
	ProtoAction_ActionReplaceOne ProtoAction = 5
	ProtoAction_ActionAppendOne  ProtoAction = 6
	ProtoAction_ActionStrPatch   ProtoAction = 7
	ProtoAction_ActionStepInto   ProtoAction = 8
)

var ProtoAction_name = map[int32]string{
	0: "ActionInvalid",
	1: "ActionReplace",
	2: "ActionAppend",
	3: "ActionRemove",
	4: "ActionRemoveOne",
	5: "ActionReplaceOne",
	6: "ActionAppendOne",
	7: "ActionStrPatch",
	8: "ActionStepInto",
}
var ProtoAction_value = map[string]int32{
	"ActionInvalid":    0,
	"ActionReplace":    1,
	"ActionAppend":     2,
	"ActionRemove":     3,
	"ActionRemoveOne":  4,
	"ActionReplaceOne": 5,
	"ActionAppendOne":  6,
	"ActionStrPatch":   7,
	"ActionStepInto":   8,
}

func (x ProtoAction) String() string {
	return proto.EnumName(ProtoAction_name, int32(x))
}
func (ProtoAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProtoPatch struct {
	NewValues *google_protobuf.Any `protobuf:"bytes,1,opt,name=newValues" json:"newValues,omitempty"`
	Ops       []*ProtoOp           `protobuf:"bytes,2,rep,name=ops" json:"ops,omitempty"`
}

func (m *ProtoPatch) Reset()                    { *m = ProtoPatch{} }
func (m *ProtoPatch) String() string            { return proto.CompactTextString(m) }
func (*ProtoPatch) ProtoMessage()               {}
func (*ProtoPatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoPatch) GetNewValues() *google_protobuf.Any {
	if m != nil {
		return m.NewValues
	}
	return nil
}

func (m *ProtoPatch) GetOps() []*ProtoOp {
	if m != nil {
		return m.Ops
	}
	return nil
}

type ProtoOp struct {
	Path []*ProtoStep `protobuf:"bytes,1,rep,name=Path" json:"Path,omitempty"`
}

func (m *ProtoOp) Reset()                    { *m = ProtoOp{} }
func (m *ProtoOp) String() string            { return proto.CompactTextString(m) }
func (*ProtoOp) ProtoMessage()               {}
func (*ProtoOp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoOp) GetPath() []*ProtoStep {
	if m != nil {
		return m.Path
	}
	return nil
}

type ProtoStep struct {
	Tag      int32       `protobuf:"varint,1,opt,name=Tag" json:"Tag,omitempty"`
	Name     string      `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	JsonName string      `protobuf:"bytes,3,opt,name=JsonName" json:"JsonName,omitempty"`
	Action   ProtoAction `protobuf:"varint,4,opt,name=Action,enum=patch.ProtoAction" json:"Action,omitempty"`
	SrcIndex int32       `protobuf:"varint,5,opt,name=SrcIndex" json:"SrcIndex,omitempty"`
	DstIndex int32       `protobuf:"varint,6,opt,name=DstIndex" json:"DstIndex,omitempty"`
	// Types that are valid to be assigned to MapKey:
	//	*ProtoStep_StrMapKey
	//	*ProtoStep_IntMapKey
	//	*ProtoStep_UIntMapKey
	//	*ProtoStep_BoolMapKey
	MapKey isProtoStep_MapKey `protobuf_oneof:"MapKey"`
}

func (m *ProtoStep) Reset()                    { *m = ProtoStep{} }
func (m *ProtoStep) String() string            { return proto.CompactTextString(m) }
func (*ProtoStep) ProtoMessage()               {}
func (*ProtoStep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isProtoStep_MapKey interface {
	isProtoStep_MapKey()
}

type ProtoStep_StrMapKey struct {
	StrMapKey string `protobuf:"bytes,7,opt,name=StrMapKey,oneof"`
}
type ProtoStep_IntMapKey struct {
	IntMapKey int64 `protobuf:"varint,8,opt,name=IntMapKey,oneof"`
}
type ProtoStep_UIntMapKey struct {
	UIntMapKey uint64 `protobuf:"varint,9,opt,name=UIntMapKey,oneof"`
}
type ProtoStep_BoolMapKey struct {
	BoolMapKey bool `protobuf:"varint,10,opt,name=BoolMapKey,oneof"`
}

func (*ProtoStep_StrMapKey) isProtoStep_MapKey()  {}
func (*ProtoStep_IntMapKey) isProtoStep_MapKey()  {}
func (*ProtoStep_UIntMapKey) isProtoStep_MapKey() {}
func (*ProtoStep_BoolMapKey) isProtoStep_MapKey() {}

func (m *ProtoStep) GetMapKey() isProtoStep_MapKey {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *ProtoStep) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *ProtoStep) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoStep) GetJsonName() string {
	if m != nil {
		return m.JsonName
	}
	return ""
}

func (m *ProtoStep) GetAction() ProtoAction {
	if m != nil {
		return m.Action
	}
	return ProtoAction_ActionInvalid
}

func (m *ProtoStep) GetSrcIndex() int32 {
	if m != nil {
		return m.SrcIndex
	}
	return 0
}

func (m *ProtoStep) GetDstIndex() int32 {
	if m != nil {
		return m.DstIndex
	}
	return 0
}

func (m *ProtoStep) GetStrMapKey() string {
	if x, ok := m.GetMapKey().(*ProtoStep_StrMapKey); ok {
		return x.StrMapKey
	}
	return ""
}

func (m *ProtoStep) GetIntMapKey() int64 {
	if x, ok := m.GetMapKey().(*ProtoStep_IntMapKey); ok {
		return x.IntMapKey
	}
	return 0
}

func (m *ProtoStep) GetUIntMapKey() uint64 {
	if x, ok := m.GetMapKey().(*ProtoStep_UIntMapKey); ok {
		return x.UIntMapKey
	}
	return 0
}

func (m *ProtoStep) GetBoolMapKey() bool {
	if x, ok := m.GetMapKey().(*ProtoStep_BoolMapKey); ok {
		return x.BoolMapKey
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProtoStep) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProtoStep_OneofMarshaler, _ProtoStep_OneofUnmarshaler, _ProtoStep_OneofSizer, []interface{}{
		(*ProtoStep_StrMapKey)(nil),
		(*ProtoStep_IntMapKey)(nil),
		(*ProtoStep_UIntMapKey)(nil),
		(*ProtoStep_BoolMapKey)(nil),
	}
}

func _ProtoStep_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProtoStep)
	// MapKey
	switch x := m.MapKey.(type) {
	case *ProtoStep_StrMapKey:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StrMapKey)
	case *ProtoStep_IntMapKey:
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntMapKey))
	case *ProtoStep_UIntMapKey:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.UIntMapKey))
	case *ProtoStep_BoolMapKey:
		t := uint64(0)
		if x.BoolMapKey {
			t = 1
		}
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ProtoStep.MapKey has unexpected type %T", x)
	}
	return nil
}

func _ProtoStep_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProtoStep)
	switch tag {
	case 7: // MapKey.StrMapKey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MapKey = &ProtoStep_StrMapKey{x}
		return true, err
	case 8: // MapKey.IntMapKey
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MapKey = &ProtoStep_IntMapKey{int64(x)}
		return true, err
	case 9: // MapKey.UIntMapKey
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MapKey = &ProtoStep_UIntMapKey{x}
		return true, err
	case 10: // MapKey.BoolMapKey
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MapKey = &ProtoStep_BoolMapKey{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ProtoStep_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProtoStep)
	// MapKey
	switch x := m.MapKey.(type) {
	case *ProtoStep_StrMapKey:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StrMapKey)))
		n += len(x.StrMapKey)
	case *ProtoStep_IntMapKey:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntMapKey))
	case *ProtoStep_UIntMapKey:
		n += proto.SizeVarint(9<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.UIntMapKey))
	case *ProtoStep_BoolMapKey:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ProtoPatch)(nil), "patch.ProtoPatch")
	proto.RegisterType((*ProtoOp)(nil), "patch.ProtoOp")
	proto.RegisterType((*ProtoStep)(nil), "patch.ProtoStep")
	proto.RegisterEnum("patch.ProtoAction", ProtoAction_name, ProtoAction_value)
}

func init() { proto.RegisterFile("patch.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x26, 0x4d, 0x93, 0x57, 0x28, 0xe6, 0xb1, 0x43, 0xd8, 0x01, 0x59, 0x15, 0x87,
	0x68, 0x87, 0x54, 0x2a, 0x9f, 0xa0, 0x13, 0x87, 0x15, 0x04, 0xab, 0x5c, 0xe0, 0xee, 0xb6, 0xa6,
	0x9b, 0x94, 0xd9, 0x56, 0xe2, 0x0d, 0xfa, 0xe9, 0xf6, 0xd5, 0x90, 0xed, 0x2c, 0x09, 0xb7, 0xf7,
	0x7e, 0xff, 0x9f, 0xdf, 0xb3, 0x25, 0xc3, 0xcc, 0x08, 0x7b, 0xb8, 0x2b, 0x4d, 0xad, 0xad, 0xc6,
	0x89, 0x6f, 0x2e, 0xdf, 0x9f, 0xb4, 0x3e, 0x55, 0x72, 0xe9, 0xe1, 0xfe, 0xf1, 0xf7, 0x52, 0xa8,
	0x73, 0x30, 0x16, 0x7b, 0x80, 0xad, 0x2b, 0xb6, 0x4e, 0xc4, 0x15, 0x64, 0x4a, 0xfe, 0xf9, 0x25,
	0xaa, 0x47, 0xd9, 0xe4, 0x84, 0x91, 0x62, 0xb6, 0xba, 0x28, 0xc3, 0xe1, 0xf2, 0xe5, 0x70, 0xb9,
	0x56, 0x67, 0xde, 0x6b, 0xc8, 0x20, 0xd2, 0xa6, 0xc9, 0xc7, 0x2c, 0x2a, 0x66, 0xab, 0x79, 0x19,
	0xd6, 0xfb, 0x99, 0xb7, 0x86, 0xbb, 0x68, 0xb1, 0x84, 0x69, 0xdb, 0xe3, 0x47, 0x88, 0xb7, 0xc2,
	0xde, 0xe5, 0xc4, 0xdb, 0x74, 0x68, 0xef, 0xac, 0x34, 0xdc, 0xa7, 0x8b, 0xe7, 0x31, 0x64, 0x1d,
	0x43, 0x0a, 0xd1, 0x0f, 0x71, 0xf2, 0xd7, 0x99, 0x70, 0x57, 0x22, 0x42, 0xfc, 0x5d, 0x3c, 0xc8,
	0x7c, 0xcc, 0x48, 0x91, 0x71, 0x5f, 0xe3, 0x25, 0xa4, 0x5f, 0x1a, 0xad, 0x3c, 0x8f, 0x3c, 0xef,
	0x7a, 0xbc, 0x82, 0x64, 0x7d, 0xb0, 0xf7, 0x5a, 0xe5, 0x31, 0x23, 0xc5, 0x7c, 0x85, 0xc3, 0xbd,
	0x21, 0xe1, 0xad, 0xe1, 0xe6, 0xec, 0xea, 0xc3, 0x46, 0x1d, 0xe5, 0xdf, 0x7c, 0xe2, 0x57, 0x76,
	0xbd, 0xcb, 0x3e, 0x37, 0x36, 0x64, 0x49, 0xc8, 0x5e, 0x7a, 0xfc, 0x00, 0xd9, 0xce, 0xd6, 0xdf,
	0x84, 0xf9, 0x2a, 0xcf, 0xf9, 0xd4, 0x5d, 0xe0, 0x66, 0xc4, 0x7b, 0xe4, 0xf2, 0x8d, 0xb2, 0x6d,
	0x9e, 0x32, 0x52, 0x44, 0x2e, 0xef, 0x10, 0x32, 0x80, 0x9f, 0xbd, 0x90, 0x31, 0x52, 0xc4, 0x37,
	0x23, 0x3e, 0x60, 0xce, 0xb8, 0xd6, 0xba, 0x6a, 0x0d, 0x60, 0xa4, 0x48, 0x9d, 0xd1, 0xb3, 0xeb,
	0x14, 0x92, 0x50, 0x5d, 0x3d, 0x13, 0x98, 0x0d, 0x5e, 0x87, 0x6f, 0xe1, 0x75, 0xa8, 0x36, 0xea,
	0x49, 0x54, 0xf7, 0x47, 0x3a, 0xea, 0x11, 0x97, 0xa6, 0x12, 0x07, 0x49, 0x09, 0x52, 0x78, 0x15,
	0xd0, 0xda, 0x18, 0xa9, 0x8e, 0x74, 0xdc, 0x13, 0x2e, 0x1f, 0xf4, 0x93, 0xa4, 0x11, 0xbe, 0x83,
	0x37, 0x43, 0x72, 0xab, 0x24, 0x8d, 0xf1, 0x02, 0xe8, 0x7f, 0xb3, 0x1c, 0x9d, 0xf4, 0x6a, 0x18,
	0xe7, 0x60, 0x82, 0x08, 0xf3, 0x00, 0x77, 0xb6, 0xf6, 0x9f, 0x8e, 0x4e, 0x87, 0x4c, 0x9a, 0x8d,
	0xb2, 0x9a, 0xa6, 0xfb, 0xc4, 0xff, 0xb7, 0x4f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x97,
	0x22, 0x9f, 0xd0, 0x02, 0x00, 0x00,
}
