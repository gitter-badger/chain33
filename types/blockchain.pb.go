// Code generated by protoc-gen-go.
// source: blockchain.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 区块头信息
// 	 version : 版本信息
// 	 parentHash :父哈希
// 	 txHash : 交易根哈希
// 	 stateHash :状态哈希
// 	 height : 区块高度
// 	 blockTime :区块产生时的时标
// 	 txCount : 区块上所有交易个数
// 	 difficulty :区块难度系数，
// 	 signature :交易签名
type Header struct {
	Version    int64      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte     `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte     `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte     `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64      `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64      `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	TxCount    int64      `protobuf:"varint,9,opt,name=txCount" json:"txCount,omitempty"`
	Hash       []byte     `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Difficulty uint32     `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Header) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Header) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Header) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Header) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *Header) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Header) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Header) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

//  参考Header解释
type Block struct {
	Version    int64          `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte         `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte         `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte         `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64          `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64          `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	Difficulty uint32         `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature     `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,7,rep,name=txs" json:"txs,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Block) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Block) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Block) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Block) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Block) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Block) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Blocks struct {
	Items []*Block `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Blocks) GetItems() []*Block {
	if m != nil {
		return m.Items
	}
	return nil
}

// 节点ID以及对应的Block
type BlockPid struct {
	Pid   string `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *BlockPid) Reset()                    { *m = BlockPid{} }
func (m *BlockPid) String() string            { return proto.CompactTextString(m) }
func (*BlockPid) ProtoMessage()               {}
func (*BlockPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BlockPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *BlockPid) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// resp
type BlockDetails struct {
	Items []*BlockDetail `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockDetails) Reset()                    { *m = BlockDetails{} }
func (m *BlockDetails) String() string            { return proto.CompactTextString(m) }
func (*BlockDetails) ProtoMessage()               {}
func (*BlockDetails) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BlockDetails) GetItems() []*BlockDetail {
	if m != nil {
		return m.Items
	}
	return nil
}

// resp
type Headers struct {
	Items []*Header `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Headers) Reset()                    { *m = Headers{} }
func (m *Headers) String() string            { return proto.CompactTextString(m) }
func (*Headers) ProtoMessage()               {}
func (*Headers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Headers) GetItems() []*Header {
	if m != nil {
		return m.Items
	}
	return nil
}

type HeadersPid struct {
	Pid     string   `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Headers *Headers `protobuf:"bytes,2,opt,name=headers" json:"headers,omitempty"`
}

func (m *HeadersPid) Reset()                    { *m = HeadersPid{} }
func (m *HeadersPid) String() string            { return proto.CompactTextString(m) }
func (*HeadersPid) ProtoMessage()               {}
func (*HeadersPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HeadersPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *HeadersPid) GetHeaders() *Headers {
	if m != nil {
		return m.Headers
	}
	return nil
}

// 区块视图
// 	 head : 区块头信息
// 	 txCount :区块上交易个数
// 	 txHashes : 区块上交易的哈希列表
type BlockOverview struct {
	Head     *Header  `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	TxCount  int64    `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
	TxHashes [][]byte `protobuf:"bytes,3,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
}

func (m *BlockOverview) Reset()                    { *m = BlockOverview{} }
func (m *BlockOverview) String() string            { return proto.CompactTextString(m) }
func (*BlockOverview) ProtoMessage()               {}
func (*BlockOverview) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BlockOverview) GetHead() *Header {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BlockOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *BlockOverview) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

// 区块详细信息
// 	 block : 区块信息
// 	 receipts :区块上所有交易的收据信息列表
type BlockDetail struct {
	Block    *Block         `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	Receipts []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *BlockDetail) Reset()                    { *m = BlockDetail{} }
func (m *BlockDetail) String() string            { return proto.CompactTextString(m) }
func (*BlockDetail) ProtoMessage()               {}
func (*BlockDetail) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *BlockDetail) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockDetail) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type Receipts struct {
	Receipts []*Receipt `protobuf:"bytes,1,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *Receipts) Reset()                    { *m = Receipts{} }
func (m *Receipts) String() string            { return proto.CompactTextString(m) }
func (*Receipts) ProtoMessage()               {}
func (*Receipts) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Receipts) GetReceipts() []*Receipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type ReceiptCheckTxList struct {
	Errs []string `protobuf:"bytes,1,rep,name=errs" json:"errs,omitempty"`
}

func (m *ReceiptCheckTxList) Reset()                    { *m = ReceiptCheckTxList{} }
func (m *ReceiptCheckTxList) String() string            { return proto.CompactTextString(m) }
func (*ReceiptCheckTxList) ProtoMessage()               {}
func (*ReceiptCheckTxList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ReceiptCheckTxList) GetErrs() []string {
	if m != nil {
		return m.Errs
	}
	return nil
}

// 区块链状态
// 	 currentHeight : 区块最新高度
// 	 mempoolSize :内存池大小
// 	 msgQueueSize : 消息队列大小
type ChainStatus struct {
	CurrentHeight int64 `protobuf:"varint,1,opt,name=currentHeight" json:"currentHeight,omitempty"`
	MempoolSize   int64 `protobuf:"varint,2,opt,name=mempoolSize" json:"mempoolSize,omitempty"`
	MsgQueueSize  int64 `protobuf:"varint,3,opt,name=msgQueueSize" json:"msgQueueSize,omitempty"`
}

func (m *ChainStatus) Reset()                    { *m = ChainStatus{} }
func (m *ChainStatus) String() string            { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()               {}
func (*ChainStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *ChainStatus) GetCurrentHeight() int64 {
	if m != nil {
		return m.CurrentHeight
	}
	return 0
}

func (m *ChainStatus) GetMempoolSize() int64 {
	if m != nil {
		return m.MempoolSize
	}
	return 0
}

func (m *ChainStatus) GetMsgQueueSize() int64 {
	if m != nil {
		return m.MsgQueueSize
	}
	return 0
}

// 获取区块信息
// 	 start : 获取区块的开始高度
// 	 end :获取区块的结束高度
// 	 Isdetail : 是否需要获取区块的详细信息
// 	 pid : peer列表
type ReqBlocks struct {
	Start    int64    `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End      int64    `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	Isdetail bool     `protobuf:"varint,3,opt,name=Isdetail" json:"Isdetail,omitempty"`
	Pid      []string `protobuf:"bytes,4,rep,name=pid" json:"pid,omitempty"`
}

func (m *ReqBlocks) Reset()                    { *m = ReqBlocks{} }
func (m *ReqBlocks) String() string            { return proto.CompactTextString(m) }
func (*ReqBlocks) ProtoMessage()               {}
func (*ReqBlocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *ReqBlocks) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ReqBlocks) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ReqBlocks) GetIsdetail() bool {
	if m != nil {
		return m.Isdetail
	}
	return false
}

func (m *ReqBlocks) GetPid() []string {
	if m != nil {
		return m.Pid
	}
	return nil
}

type MempoolSize struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *MempoolSize) Reset()                    { *m = MempoolSize{} }
func (m *MempoolSize) String() string            { return proto.CompactTextString(m) }
func (*MempoolSize) ProtoMessage()               {}
func (*MempoolSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *MempoolSize) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReplyBlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *ReplyBlockHeight) Reset()                    { *m = ReplyBlockHeight{} }
func (m *ReplyBlockHeight) String() string            { return proto.CompactTextString(m) }
func (*ReplyBlockHeight) ProtoMessage()               {}
func (*ReplyBlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *ReplyBlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// 区块体信息
// 	 txs : 区块上所有交易列表
// 	 receipts :区块上所有交易的收据信息列表
type BlockBody struct {
	Txs      []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
	Receipts []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *BlockBody) Reset()                    { *m = BlockBody{} }
func (m *BlockBody) String() string            { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()               {}
func (*BlockBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *BlockBody) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *BlockBody) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

//  区块追赶主链状态，用于判断本节点区块是否已经同步好
type IsCaughtUp struct {
	Iscaughtup bool `protobuf:"varint,1,opt,name=Iscaughtup" json:"Iscaughtup,omitempty"`
}

func (m *IsCaughtUp) Reset()                    { *m = IsCaughtUp{} }
func (m *IsCaughtUp) String() string            { return proto.CompactTextString(m) }
func (*IsCaughtUp) ProtoMessage()               {}
func (*IsCaughtUp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *IsCaughtUp) GetIscaughtup() bool {
	if m != nil {
		return m.Iscaughtup
	}
	return false
}

//  ntp时钟状态
type IsNtpClockSync struct {
	Isntpclocksync bool `protobuf:"varint,1,opt,name=isntpclocksync" json:"isntpclocksync,omitempty"`
}

func (m *IsNtpClockSync) Reset()                    { *m = IsNtpClockSync{} }
func (m *IsNtpClockSync) String() string            { return proto.CompactTextString(m) }
func (*IsNtpClockSync) ProtoMessage()               {}
func (*IsNtpClockSync) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *IsNtpClockSync) GetIsntpclocksync() bool {
	if m != nil {
		return m.Isntpclocksync
	}
	return false
}

type BlockChainQuery struct {
	Driver    string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	FuncName  string `protobuf:"bytes,2,opt,name=funcName" json:"funcName,omitempty"`
	StateHash []byte `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Param     []byte `protobuf:"bytes,4,opt,name=param,proto3" json:"param,omitempty"`
}

func (m *BlockChainQuery) Reset()                    { *m = BlockChainQuery{} }
func (m *BlockChainQuery) String() string            { return proto.CompactTextString(m) }
func (*BlockChainQuery) ProtoMessage()               {}
func (*BlockChainQuery) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *BlockChainQuery) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChainQuery) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *BlockChainQuery) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *BlockChainQuery) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "types.Header")
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*Blocks)(nil), "types.Blocks")
	proto.RegisterType((*BlockPid)(nil), "types.BlockPid")
	proto.RegisterType((*BlockDetails)(nil), "types.BlockDetails")
	proto.RegisterType((*Headers)(nil), "types.Headers")
	proto.RegisterType((*HeadersPid)(nil), "types.HeadersPid")
	proto.RegisterType((*BlockOverview)(nil), "types.BlockOverview")
	proto.RegisterType((*BlockDetail)(nil), "types.BlockDetail")
	proto.RegisterType((*Receipts)(nil), "types.Receipts")
	proto.RegisterType((*ReceiptCheckTxList)(nil), "types.ReceiptCheckTxList")
	proto.RegisterType((*ChainStatus)(nil), "types.ChainStatus")
	proto.RegisterType((*ReqBlocks)(nil), "types.ReqBlocks")
	proto.RegisterType((*MempoolSize)(nil), "types.MempoolSize")
	proto.RegisterType((*ReplyBlockHeight)(nil), "types.ReplyBlockHeight")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*IsCaughtUp)(nil), "types.IsCaughtUp")
	proto.RegisterType((*IsNtpClockSync)(nil), "types.IsNtpClockSync")
	proto.RegisterType((*BlockChainQuery)(nil), "types.BlockChainQuery")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x55, 0xdd, 0xae, 0xe3, 0x34,
	0x10, 0x56, 0x9a, 0xfe, 0x4e, 0x7f, 0x28, 0xd6, 0x0a, 0x45, 0x2b, 0x84, 0xba, 0x66, 0x85, 0xa2,
	0xd5, 0xaa, 0x17, 0x8b, 0x84, 0xf6, 0x12, 0x6d, 0xf7, 0xe2, 0x54, 0x82, 0x85, 0x75, 0x0f, 0x0f,
	0xe0, 0x93, 0xf8, 0x34, 0xd6, 0x69, 0x93, 0x60, 0x3b, 0xa5, 0xe1, 0x75, 0xb8, 0xe5, 0x21, 0x91,
	0xc7, 0x4e, 0x9b, 0x9c, 0x03, 0x48, 0x48, 0xdc, 0xec, 0x9d, 0x67, 0xbe, 0x19, 0xcf, 0x37, 0x5f,
	0xc6, 0x13, 0x58, 0xde, 0x1d, 0x8a, 0xe4, 0x21, 0xc9, 0xb8, 0xcc, 0xd7, 0xa5, 0x2a, 0x4c, 0x41,
	0x06, 0xa6, 0x2e, 0x85, 0x7e, 0xfe, 0xb9, 0x51, 0x3c, 0xd7, 0x3c, 0x31, 0xb2, 0xf0, 0x08, 0xfd,
	0xb3, 0x07, 0xc3, 0x1b, 0xc1, 0x53, 0xa1, 0x48, 0x04, 0xa3, 0x93, 0x50, 0x5a, 0x16, 0x79, 0x14,
	0xac, 0x82, 0x38, 0x64, 0x8d, 0x49, 0xbe, 0x02, 0x28, 0xb9, 0x12, 0xb9, 0xb9, 0xe1, 0x3a, 0x8b,
	0x7a, 0xab, 0x20, 0x9e, 0xb1, 0x96, 0x87, 0x7c, 0x01, 0x43, 0x73, 0x46, 0x2c, 0x44, 0xcc, 0x5b,
	0xe4, 0x4b, 0x98, 0x68, 0xc3, 0x8d, 0x40, 0xa8, 0x8f, 0xd0, 0xd5, 0x61, 0xb3, 0x32, 0x21, 0xf7,
	0x99, 0x89, 0x06, 0x58, 0xce, 0x5b, 0x36, 0x0b, 0x1b, 0xb8, 0x95, 0x47, 0x11, 0x0d, 0x11, 0xba,
	0x3a, 0x2c, 0x4b, 0x73, 0xde, 0x14, 0x55, 0x6e, 0xa2, 0x89, 0x63, 0xe9, 0x4d, 0x42, 0xa0, 0x9f,
	0xd9, 0x42, 0x80, 0x85, 0xf0, 0x6c, 0x99, 0xa7, 0xf2, 0xfe, 0x5e, 0x26, 0xd5, 0xc1, 0xd4, 0xd1,
	0x74, 0x15, 0xc4, 0x73, 0xd6, 0xf2, 0x90, 0x35, 0x4c, 0xb4, 0xdc, 0xe7, 0xdc, 0x54, 0x4a, 0x44,
	0xe3, 0x55, 0x10, 0x4f, 0xdf, 0x2c, 0xd7, 0x28, 0xd6, 0x7a, 0xd7, 0xf8, 0xd9, 0x35, 0x84, 0xfe,
	0xd1, 0x83, 0xc1, 0x3b, 0xcb, 0xe5, 0x13, 0x51, 0xeb, 0x7f, 0xee, 0x9f, 0xbc, 0x84, 0xd0, 0x9c,
	0x75, 0x34, 0x5a, 0x85, 0xf1, 0xf4, 0x0d, 0xf1, 0x91, 0xb7, 0xd7, 0xa9, 0x62, 0x16, 0xa6, 0xaf,
	0x61, 0x88, 0x22, 0x69, 0x42, 0x61, 0x20, 0x8d, 0x38, 0xea, 0x28, 0xc0, 0x8c, 0x99, 0xcf, 0x40,
	0x94, 0x39, 0x88, 0x7e, 0x0f, 0x63, 0xb4, 0x7f, 0x96, 0x29, 0x59, 0x42, 0x58, 0xca, 0x14, 0x15,
	0x9d, 0x30, 0x7b, 0xb4, 0x37, 0x60, 0x3b, 0x28, 0xe4, 0x93, 0x1b, 0x10, 0xa2, 0x6f, 0x61, 0x86,
	0xf6, 0x7b, 0x61, 0xb8, 0x3c, 0x68, 0x12, 0x77, 0xab, 0x92, 0x76, 0x8e, 0x8b, 0x69, 0x6a, 0xaf,
	0x61, 0xe4, 0xa6, 0x5f, 0x93, 0xaf, 0xbb, 0x49, 0x73, 0x9f, 0xe4, 0xe0, 0x26, 0xfe, 0x06, 0xc0,
	0xc7, 0xff, 0x3d, 0xdb, 0x18, 0x46, 0x99, 0xc3, 0x3d, 0xdf, 0x45, 0xe7, 0x1a, 0xcd, 0x1a, 0x98,
	0x66, 0x30, 0x47, 0x3e, 0x3f, 0x9d, 0x84, 0x3a, 0x49, 0xf1, 0x1b, 0x79, 0x01, 0x7d, 0x8b, 0xe1,
	0x6d, 0x4f, 0xca, 0x23, 0xd4, 0x9e, 0xfd, 0x5e, 0x77, 0xf6, 0x9f, 0xc3, 0xd8, 0x4d, 0x91, 0xd0,
	0x51, 0xb8, 0x0a, 0xe3, 0x19, 0xbb, 0xd8, 0x94, 0xc3, 0xb4, 0xd5, 0xf9, 0x55, 0xd0, 0xe0, 0x1f,
	0x05, 0x25, 0x6b, 0x18, 0x2b, 0x91, 0x08, 0x59, 0x1a, 0xdb, 0x47, 0x5b, 0x43, 0xe6, 0xdc, 0xef,
	0xb9, 0xe1, 0xec, 0x12, 0x43, 0xbf, 0x83, 0xb1, 0x07, 0x34, 0x79, 0xd5, 0xca, 0x75, 0x52, 0x2e,
	0xba, 0xb9, 0xad, 0xbc, 0x18, 0x88, 0x77, 0x6e, 0x32, 0x91, 0x3c, 0xdc, 0x9e, 0x7f, 0x90, 0x1a,
	0x1f, 0xb2, 0x50, 0xca, 0x65, 0x4f, 0x18, 0x9e, 0x69, 0x0d, 0xd3, 0x8d, 0x5d, 0x68, 0x3b, 0xc3,
	0x4d, 0xa5, 0xc9, 0x4b, 0x98, 0x27, 0x95, 0xc2, 0x27, 0xe5, 0x1e, 0x85, 0x7b, 0x83, 0x5d, 0x27,
	0x59, 0xc1, 0xf4, 0x28, 0x8e, 0x65, 0x51, 0x1c, 0x76, 0xf2, 0x77, 0xe1, 0x35, 0x6b, 0xbb, 0x08,
	0x85, 0xd9, 0x51, 0xef, 0x3f, 0x56, 0xa2, 0x12, 0x18, 0x12, 0x62, 0x48, 0xc7, 0x47, 0x39, 0x4c,
	0x98, 0xf8, 0xd5, 0x0f, 0xf4, 0x33, 0x18, 0x68, 0xc3, 0x55, 0x53, 0xd0, 0x19, 0x76, 0x10, 0x44,
	0x9e, 0xfa, 0x02, 0xf6, 0x68, 0x3f, 0xc8, 0x56, 0xa7, 0xa8, 0x38, 0x5e, 0x3a, 0x66, 0x17, 0xbb,
	0x19, 0x9b, 0x3e, 0xb6, 0x67, 0x8f, 0xf4, 0x05, 0x4c, 0x7f, 0x6c, 0xb1, 0x22, 0xd0, 0xd7, 0x96,
	0x8d, 0xab, 0x81, 0x67, 0xfa, 0x0a, 0x96, 0x4c, 0x94, 0x87, 0x1a, 0x79, 0xf8, 0xfe, 0xae, 0x3b,
	0x21, 0x68, 0xef, 0x04, 0xcb, 0x18, 0xc3, 0xde, 0x15, 0x69, 0xdd, 0x3c, 0xd9, 0xe0, 0x5f, 0x9f,
	0xec, 0x7f, 0xfe, 0xe2, 0xaf, 0x01, 0xb6, 0x7a, 0xc3, 0xab, 0x7d, 0x66, 0x7e, 0x29, 0xed, 0x9a,
	0xd9, 0xea, 0x04, 0xad, 0xaa, 0x44, 0x32, 0x63, 0xd6, 0xf2, 0xd0, 0xb7, 0xb0, 0xd8, 0xea, 0x0f,
	0xa6, 0xdc, 0x58, 0x56, 0xbb, 0x3a, 0x4f, 0xc8, 0x37, 0xb0, 0x90, 0x3a, 0x37, 0x65, 0x82, 0xb2,
	0xd6, 0x79, 0xe2, 0xb3, 0x1e, 0x79, 0x69, 0x0d, 0x9f, 0x61, 0x2b, 0xf8, 0xf1, 0x3f, 0x56, 0x42,
	0xd5, 0xb6, 0xeb, 0x54, 0xc9, 0x93, 0x50, 0xfe, 0xe1, 0x79, 0xcb, 0x4a, 0x7e, 0x5f, 0xe5, 0xc9,
	0x07, 0x7e, 0x74, 0x9f, 0x7a, 0xc2, 0x2e, 0x76, 0x77, 0xb7, 0x86, 0x8f, 0x77, 0xeb, 0x33, 0x18,
	0x94, 0x5c, 0xf1, 0xa3, 0xdf, 0xba, 0xce, 0xb8, 0x1b, 0xe2, 0x1f, 0xf2, 0xdb, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x61, 0x9f, 0x12, 0x31, 0x4f, 0x07, 0x00, 0x00,
}
