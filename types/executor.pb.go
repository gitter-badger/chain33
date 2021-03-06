// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Genesis struct {
	Isrun bool `protobuf:"varint,1,opt,name=isrun" json:"isrun,omitempty"`
}

func (m *Genesis) Reset()                    { *m = Genesis{} }
func (m *Genesis) String() string            { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()               {}
func (*Genesis) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Genesis) GetIsrun() bool {
	if m != nil {
		return m.Isrun
	}
	return false
}

type ExecTxList struct {
	StateHash  []byte         `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,2,rep,name=txs" json:"txs,omitempty"`
	BlockTime  int64          `protobuf:"varint,3,opt,name=blockTime" json:"blockTime,omitempty"`
	Height     int64          `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Difficulty uint64         `protobuf:"varint,5,opt,name=difficulty" json:"difficulty,omitempty"`
	IsMempool  bool           `protobuf:"varint,6,opt,name=isMempool" json:"isMempool,omitempty"`
}

func (m *ExecTxList) Reset()                    { *m = ExecTxList{} }
func (m *ExecTxList) String() string            { return proto.CompactTextString(m) }
func (*ExecTxList) ProtoMessage()               {}
func (*ExecTxList) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ExecTxList) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ExecTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *ExecTxList) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ExecTxList) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ExecTxList) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *ExecTxList) GetIsMempool() bool {
	if m != nil {
		return m.IsMempool
	}
	return false
}

type Query struct {
	Execer   []byte `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	FuncName string `protobuf:"bytes,2,opt,name=funcName" json:"funcName,omitempty"`
	Payload  []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Query) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Query) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *Query) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateTxIn struct {
	Execer     []byte `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	ActionName string `protobuf:"bytes,2,opt,name=actionName" json:"actionName,omitempty"`
	Payload    []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *CreateTxIn) Reset()                    { *m = CreateTxIn{} }
func (m *CreateTxIn) String() string            { return proto.CompactTextString(m) }
func (*CreateTxIn) ProtoMessage()               {}
func (*CreateTxIn) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *CreateTxIn) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *CreateTxIn) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *CreateTxIn) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// 配置修改部分
type ArrayConfig struct {
	Value []string `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
}

func (m *ArrayConfig) Reset()                    { *m = ArrayConfig{} }
func (m *ArrayConfig) String() string            { return proto.CompactTextString(m) }
func (*ArrayConfig) ProtoMessage()               {}
func (*ArrayConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ArrayConfig) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringConfig struct {
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *StringConfig) Reset()                    { *m = StringConfig{} }
func (m *StringConfig) String() string            { return proto.CompactTextString(m) }
func (*StringConfig) ProtoMessage()               {}
func (*StringConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *StringConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Int32Config struct {
	Value int32 `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
}

func (m *Int32Config) Reset()                    { *m = Int32Config{} }
func (m *Int32Config) String() string            { return proto.CompactTextString(m) }
func (*Int32Config) ProtoMessage()               {}
func (*Int32Config) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *Int32Config) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfigItem struct {
	Key  string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*ConfigItem_Arr
	//	*ConfigItem_Str
	//	*ConfigItem_Int
	Value isConfigItem_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,11,opt,name=Ty" json:"Ty,omitempty"`
}

func (m *ConfigItem) Reset()                    { *m = ConfigItem{} }
func (m *ConfigItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()               {}
func (*ConfigItem) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type isConfigItem_Value interface {
	isConfigItem_Value()
}

type ConfigItem_Arr struct {
	Arr *ArrayConfig `protobuf:"bytes,3,opt,name=arr,oneof"`
}
type ConfigItem_Str struct {
	Str *StringConfig `protobuf:"bytes,4,opt,name=str,oneof"`
}
type ConfigItem_Int struct {
	Int *Int32Config `protobuf:"bytes,5,opt,name=int,oneof"`
}

func (*ConfigItem_Arr) isConfigItem_Value() {}
func (*ConfigItem_Str) isConfigItem_Value() {}
func (*ConfigItem_Int) isConfigItem_Value() {}

func (m *ConfigItem) GetValue() isConfigItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ConfigItem) GetArr() *ArrayConfig {
	if x, ok := m.GetValue().(*ConfigItem_Arr); ok {
		return x.Arr
	}
	return nil
}

func (m *ConfigItem) GetStr() *StringConfig {
	if x, ok := m.GetValue().(*ConfigItem_Str); ok {
		return x.Str
	}
	return nil
}

func (m *ConfigItem) GetInt() *Int32Config {
	if x, ok := m.GetValue().(*ConfigItem_Int); ok {
		return x.Int
	}
	return nil
}

func (m *ConfigItem) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigItem_OneofMarshaler, _ConfigItem_OneofUnmarshaler, _ConfigItem_OneofSizer, []interface{}{
		(*ConfigItem_Arr)(nil),
		(*ConfigItem_Str)(nil),
		(*ConfigItem_Int)(nil),
	}
}

func _ConfigItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigItem)
	// value
	switch x := m.Value.(type) {
	case *ConfigItem_Arr:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Arr); err != nil {
			return err
		}
	case *ConfigItem_Str:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Str); err != nil {
			return err
		}
	case *ConfigItem_Int:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigItem.Value has unexpected type %T", x)
	}
	return nil
}

func _ConfigItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigItem)
	switch tag {
	case 3: // value.arr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayConfig)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Arr{msg}
		return true, err
	case 4: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringConfig)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Str{msg}
		return true, err
	case 5: // value.int
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int32Config)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Int{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigItem)
	// value
	switch x := m.Value.(type) {
	case *ConfigItem_Arr:
		s := proto.Size(x.Arr)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigItem_Str:
		s := proto.Size(x.Str)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigItem_Int:
		s := proto.Size(x.Int)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ModifyConfig struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Op    string `protobuf:"bytes,3,opt,name=op" json:"op,omitempty"`
	Addr  string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *ModifyConfig) Reset()                    { *m = ModifyConfig{} }
func (m *ModifyConfig) String() string            { return proto.CompactTextString(m) }
func (*ModifyConfig) ProtoMessage()               {}
func (*ModifyConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *ModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ModifyConfig) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ReceiptConfig struct {
	Prev    *ConfigItem `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *ConfigItem `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptConfig) Reset()                    { *m = ReceiptConfig{} }
func (m *ReceiptConfig) String() string            { return proto.CompactTextString(m) }
func (*ReceiptConfig) ProtoMessage()               {}
func (*ReceiptConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *ReceiptConfig) GetPrev() *ConfigItem {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptConfig) GetCurrent() *ConfigItem {
	if m != nil {
		return m.Current
	}
	return nil
}

type ReplyConfig struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ReplyConfig) Reset()                    { *m = ReplyConfig{} }
func (m *ReplyConfig) String() string            { return proto.CompactTextString(m) }
func (*ReplyConfig) ProtoMessage()               {}
func (*ReplyConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *ReplyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HistoryCertStore struct {
	Rootcerts         [][]byte `protobuf:"bytes,1,rep,name=rootcerts,proto3" json:"rootcerts,omitempty"`
	IntermediateCerts [][]byte `protobuf:"bytes,2,rep,name=intermediateCerts,proto3" json:"intermediateCerts,omitempty"`
	RevocationList    [][]byte `protobuf:"bytes,3,rep,name=revocationList,proto3" json:"revocationList,omitempty"`
	CurHeigth         int64    `protobuf:"varint,4,opt,name=curHeigth" json:"curHeigth,omitempty"`
	NxtHeight         int64    `protobuf:"varint,5,opt,name=nxtHeight" json:"nxtHeight,omitempty"`
}

func (m *HistoryCertStore) Reset()                    { *m = HistoryCertStore{} }
func (m *HistoryCertStore) String() string            { return proto.CompactTextString(m) }
func (*HistoryCertStore) ProtoMessage()               {}
func (*HistoryCertStore) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *HistoryCertStore) GetRootcerts() [][]byte {
	if m != nil {
		return m.Rootcerts
	}
	return nil
}

func (m *HistoryCertStore) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *HistoryCertStore) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *HistoryCertStore) GetCurHeigth() int64 {
	if m != nil {
		return m.CurHeigth
	}
	return 0
}

func (m *HistoryCertStore) GetNxtHeight() int64 {
	if m != nil {
		return m.NxtHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Genesis)(nil), "types.Genesis")
	proto.RegisterType((*ExecTxList)(nil), "types.ExecTxList")
	proto.RegisterType((*Query)(nil), "types.Query")
	proto.RegisterType((*CreateTxIn)(nil), "types.CreateTxIn")
	proto.RegisterType((*ArrayConfig)(nil), "types.ArrayConfig")
	proto.RegisterType((*StringConfig)(nil), "types.StringConfig")
	proto.RegisterType((*Int32Config)(nil), "types.Int32Config")
	proto.RegisterType((*ConfigItem)(nil), "types.ConfigItem")
	proto.RegisterType((*ModifyConfig)(nil), "types.ModifyConfig")
	proto.RegisterType((*ReceiptConfig)(nil), "types.ReceiptConfig")
	proto.RegisterType((*ReplyConfig)(nil), "types.ReplyConfig")
	proto.RegisterType((*HistoryCertStore)(nil), "types.HistoryCertStore")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x61, 0x6b, 0x13, 0x41,
	0x10, 0xf5, 0x72, 0x49, 0xd3, 0x4c, 0x62, 0x69, 0x57, 0x91, 0xa3, 0x48, 0x0d, 0xd7, 0x5a, 0x03,
	0x4a, 0x84, 0x04, 0x7f, 0x80, 0x0d, 0x62, 0x0a, 0x56, 0xf0, 0x1a, 0xbf, 0xf4, 0x83, 0xb0, 0xdd,
	0x4c, 0x92, 0xa5, 0xc9, 0xee, 0xb1, 0x3b, 0x57, 0x72, 0x7f, 0x4f, 0xfc, 0x61, 0xb2, 0x9b, 0x4b,
	0xee, 0xb0, 0x55, 0xf0, 0xd3, 0xdd, 0xbc, 0x79, 0xbc, 0x99, 0x37, 0x73, 0x73, 0x70, 0x80, 0x6b,
	0x14, 0x19, 0x69, 0xd3, 0x4f, 0x8d, 0x26, 0xcd, 0x1a, 0x94, 0xa7, 0x68, 0x8f, 0x8f, 0xc8, 0x70,
	0x65, 0xb9, 0x20, 0xa9, 0xd5, 0x26, 0x13, 0xbf, 0x82, 0xe6, 0x67, 0x54, 0x68, 0xa5, 0x65, 0xcf,
	0xa1, 0x21, 0xad, 0xc9, 0x54, 0x14, 0x74, 0x83, 0xde, 0x7e, 0xb2, 0x09, 0xe2, 0x5f, 0x01, 0xc0,
	0xa7, 0x35, 0x8a, 0xc9, 0xfa, 0x8b, 0xb4, 0xc4, 0x5e, 0x42, 0xcb, 0x12, 0x27, 0x1c, 0x73, 0xbb,
	0xf0, 0xc4, 0x4e, 0x52, 0x02, 0xec, 0x0c, 0x42, 0x5a, 0xdb, 0xa8, 0xd6, 0x0d, 0x7b, 0xed, 0x01,
	0xeb, 0xfb, 0xaa, 0xfd, 0x49, 0x59, 0x34, 0x71, 0x69, 0xa7, 0x71, 0xbb, 0xd4, 0xe2, 0x6e, 0x22,
	0x57, 0x18, 0x85, 0xdd, 0xa0, 0x17, 0x26, 0x25, 0xc0, 0x5e, 0xc0, 0xde, 0x02, 0xe5, 0x7c, 0x41,
	0x51, 0xdd, 0xa7, 0x8a, 0x88, 0x9d, 0x00, 0x4c, 0xe5, 0x6c, 0x26, 0x45, 0xb6, 0xa4, 0x3c, 0x6a,
	0x74, 0x83, 0x5e, 0x3d, 0xa9, 0x20, 0x4e, 0x55, 0xda, 0x2b, 0x5c, 0xa5, 0x5a, 0x2f, 0xa3, 0x3d,
	0x6f, 0xa1, 0x04, 0xe2, 0xef, 0xd0, 0xf8, 0x96, 0xa1, 0xc9, 0x9d, 0xbc, 0x1b, 0x0e, 0x9a, 0xa2,
	0xfb, 0x22, 0x62, 0xc7, 0xb0, 0x3f, 0xcb, 0x94, 0xf8, 0xca, 0x57, 0x18, 0xd5, 0xba, 0x41, 0xaf,
	0x95, 0xec, 0x62, 0x16, 0x41, 0x33, 0xe5, 0xf9, 0x52, 0xf3, 0xa9, 0x6f, 0xb7, 0x93, 0x6c, 0xc3,
	0xf8, 0x07, 0xc0, 0xc8, 0x20, 0x27, 0x9c, 0xac, 0x2f, 0xd5, 0x5f, 0xb5, 0x4f, 0x00, 0x36, 0xfe,
	0x2b, 0xea, 0x15, 0xe4, 0x1f, 0xfa, 0xa7, 0xd0, 0xfe, 0x68, 0x0c, 0xcf, 0x47, 0x5a, 0xcd, 0xe4,
	0xdc, 0xad, 0xe8, 0x9e, 0x2f, 0x33, 0x37, 0xb5, 0xb0, 0xd7, 0x4a, 0x36, 0x41, 0x7c, 0x06, 0x9d,
	0x6b, 0x32, 0x52, 0xcd, 0x1f, 0xb2, 0x82, 0x92, 0x75, 0x0a, 0xed, 0x4b, 0x45, 0xc3, 0xc1, 0x63,
	0xa4, 0xc6, 0x96, 0xe4, 0xb6, 0xbd, 0x21, 0x5c, 0x12, 0xae, 0xd8, 0x21, 0x84, 0x77, 0x98, 0x7b,
	0x37, 0xad, 0xc4, 0xbd, 0x32, 0x06, 0x75, 0x3e, 0x9d, 0x9a, 0xc2, 0x84, 0x7f, 0x67, 0xe7, 0x10,
	0x72, 0x63, 0xbc, 0x50, 0xb9, 0xf5, 0x4a, 0xdb, 0xe3, 0x27, 0x89, 0x23, 0xb0, 0x37, 0x10, 0x5a,
	0x32, 0x7e, 0xad, 0xed, 0xc1, 0xb3, 0x82, 0x57, 0xed, 0xdc, 0x11, 0x2d, 0x79, 0x41, 0xa9, 0xc8,
	0xef, 0xb8, 0x14, 0xac, 0x34, 0xef, 0x78, 0x52, 0x11, 0x3b, 0x80, 0xda, 0x24, 0x8f, 0xda, 0xde,
	0x40, 0x6d, 0x92, 0x5f, 0x34, 0x0b, 0x4f, 0xf1, 0x0d, 0x74, 0xae, 0xf4, 0x54, 0xce, 0xb6, 0x73,
	0x7b, 0xe8, 0x63, 0x67, 0xbf, 0x56, 0x99, 0x91, 0x13, 0xd4, 0x69, 0x31, 0xb6, 0x9a, 0x4e, 0x77,
	0x6e, 0xeb, 0xa5, 0xdb, 0x58, 0xc0, 0xd3, 0x04, 0x05, 0xca, 0x94, 0x0a, 0xf1, 0xd7, 0x50, 0x4f,
	0x0d, 0xde, 0x7b, 0xf5, 0xf6, 0xe0, 0xa8, 0x68, 0xb7, 0x9c, 0x62, 0xe2, 0xd3, 0xec, 0x2d, 0x34,
	0x45, 0x66, 0x0c, 0x2a, 0xf2, 0x35, 0x1f, 0x65, 0x6e, 0x19, 0xf1, 0x07, 0x68, 0x27, 0x98, 0x2e,
	0xff, 0xb3, 0xff, 0xf8, 0x67, 0x00, 0x87, 0x63, 0x69, 0x49, 0x9b, 0x7c, 0x84, 0x86, 0xae, 0x49,
	0x1b, 0x74, 0x87, 0x61, 0xb4, 0x26, 0x81, 0x86, 0x6c, 0x14, 0x74, 0x43, 0x77, 0xb2, 0x3b, 0x80,
	0xbd, 0x83, 0x23, 0xa9, 0x08, 0xcd, 0x0a, 0xa7, 0x92, 0x13, 0x8e, 0x3c, 0xab, 0xe6, 0x59, 0x0f,
	0x13, 0xec, 0x1c, 0x0e, 0x0c, 0xde, 0x6b, 0xc1, 0xdd, 0xb7, 0xeb, 0x7e, 0x08, 0xfe, 0x4b, 0xec,
	0x24, 0x7f, 0xa0, 0xae, 0xa6, 0xc8, 0xcc, 0x18, 0xe5, 0x9c, 0x16, 0xc5, 0x1d, 0x97, 0x80, 0xcb,
	0xaa, 0x35, 0x8d, 0x37, 0x57, 0xde, 0xd8, 0x64, 0x77, 0xc0, 0xc5, 0xd9, 0x4d, 0x3c, 0x97, 0xb4,
	0xe4, 0xb7, 0xfd, 0xe1, 0xb0, 0x2f, 0xd4, 0x7b, 0xb1, 0xe0, 0x52, 0x0d, 0x87, 0xbb, 0xa7, 0x9f,
	0xda, 0xed, 0x9e, 0xff, 0x7f, 0x0d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x8a, 0x71, 0x62,
	0xeb, 0x04, 0x00, 0x00,
}
