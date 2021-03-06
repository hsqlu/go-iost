// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/block/pb/block.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pb1 "github.com/iost-official/go-iost/core/tx/pb"
	pb "github.com/iost-official/go-iost/crypto/pb"
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

type BlockType int32

const (
	BlockType_NORMAL   BlockType = 0
	BlockType_ONLYHASH BlockType = 1
)

var BlockType_name = map[int32]string{
	0: "NORMAL",
	1: "ONLYHASH",
}

var BlockType_value = map[string]int32{
	"NORMAL":   0,
	"ONLYHASH": 1,
}

func (x BlockType) String() string {
	return proto.EnumName(BlockType_name, int32(x))
}

func (BlockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dc6664e18d413fc7, []int{0}
}

type BlockHead struct {
	Version              int64    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxMerkleHash         []byte   `protobuf:"bytes,3,opt,name=txMerkleHash,proto3" json:"txMerkleHash,omitempty"`
	TxReceiptMerkleHash  []byte   `protobuf:"bytes,4,opt,name=txReceiptMerkleHash,proto3" json:"txReceiptMerkleHash,omitempty"`
	Info                 []byte   `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Number               int64    `protobuf:"varint,6,opt,name=number,proto3" json:"number,omitempty"`
	Witness              string   `protobuf:"bytes,7,opt,name=witness,proto3" json:"witness,omitempty"`
	Time                 int64    `protobuf:"varint,8,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHead) Reset()         { *m = BlockHead{} }
func (m *BlockHead) String() string { return proto.CompactTextString(m) }
func (*BlockHead) ProtoMessage()    {}
func (*BlockHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc6664e18d413fc7, []int{0}
}

func (m *BlockHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHead.Unmarshal(m, b)
}
func (m *BlockHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHead.Marshal(b, m, deterministic)
}
func (m *BlockHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHead.Merge(m, src)
}
func (m *BlockHead) XXX_Size() int {
	return xxx_messageInfo_BlockHead.Size(m)
}
func (m *BlockHead) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHead.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHead proto.InternalMessageInfo

func (m *BlockHead) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHead) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHead) GetTxMerkleHash() []byte {
	if m != nil {
		return m.TxMerkleHash
	}
	return nil
}

func (m *BlockHead) GetTxReceiptMerkleHash() []byte {
	if m != nil {
		return m.TxReceiptMerkleHash
	}
	return nil
}

func (m *BlockHead) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *BlockHead) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHead) GetWitness() string {
	if m != nil {
		return m.Witness
	}
	return ""
}

func (m *BlockHead) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Block struct {
	Head                 *BlockHead       `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Sign                 *pb.Signature    `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Txs                  []*pb1.Tx        `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	Receipts             []*pb1.TxReceipt `protobuf:"bytes,4,rep,name=receipts,proto3" json:"receipts,omitempty"`
	TxHashes             [][]byte         `protobuf:"bytes,5,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
	ReceiptHashes        [][]byte         `protobuf:"bytes,6,rep,name=receiptHashes,proto3" json:"receiptHashes,omitempty"`
	BlockType            BlockType        `protobuf:"varint,7,opt,name=blockType,proto3,enum=blockpb.BlockType" json:"blockType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc6664e18d413fc7, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHead() *BlockHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Block) GetSign() *pb.Signature {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Block) GetTxs() []*pb1.Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *Block) GetReceipts() []*pb1.TxReceipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *Block) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

func (m *Block) GetReceiptHashes() [][]byte {
	if m != nil {
		return m.ReceiptHashes
	}
	return nil
}

func (m *Block) GetBlockType() BlockType {
	if m != nil {
		return m.BlockType
	}
	return BlockType_NORMAL
}

func init() {
	proto.RegisterEnum("blockpb.BlockType", BlockType_name, BlockType_value)
	proto.RegisterType((*BlockHead)(nil), "blockpb.BlockHead")
	proto.RegisterType((*Block)(nil), "blockpb.Block")
}

func init() { proto.RegisterFile("core/block/pb/block.proto", fileDescriptor_dc6664e18d413fc7) }

var fileDescriptor_dc6664e18d413fc7 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0xc9, 0x92, 0xa6, 0xe9, 0xd7, 0x02, 0xd5, 0x37, 0x09, 0x79, 0xbd, 0x40, 0x51, 0x35,
	0x50, 0x04, 0x5a, 0x52, 0x95, 0x27, 0xd8, 0xae, 0x7a, 0xb1, 0x3f, 0x92, 0xb7, 0x1b, 0xb8, 0x4b,
	0x32, 0xb7, 0xb5, 0xd6, 0xc6, 0x91, 0xed, 0x42, 0xf6, 0x1a, 0xbc, 0x2b, 0xf7, 0xc8, 0x5f, 0xd2,
	0xb2, 0xa2, 0xdd, 0xf9, 0x9c, 0xf3, 0x4b, 0xe4, 0x73, 0x64, 0x38, 0x2b, 0x95, 0x16, 0x59, 0xb1,
	0x51, 0xe5, 0x53, 0x56, 0x17, 0xed, 0x21, 0xad, 0xb5, 0xb2, 0x0a, 0xfb, 0x24, 0xea, 0x62, 0x72,
	0x56, 0xea, 0xe7, 0xda, 0x2a, 0x97, 0x1b, 0xb9, 0xaa, 0x72, 0xbb, 0xd3, 0xa2, 0x65, 0x26, 0xa7,
	0xf4, 0xb9, 0x6d, 0x5c, 0x66, 0x9b, 0xd6, 0x9c, 0xfe, 0xf1, 0x60, 0x70, 0xe5, 0xbe, 0x5d, 0x88,
	0xfc, 0x11, 0x19, 0xf4, 0x7f, 0x0a, 0x6d, 0xa4, 0xaa, 0x98, 0x17, 0x7b, 0x89, 0xcf, 0xf7, 0x12,
	0x3f, 0x02, 0xd4, 0xb9, 0x16, 0x95, 0x5d, 0xe4, 0x66, 0xcd, 0x4e, 0x62, 0x2f, 0x19, 0xf1, 0x17,
	0x0e, 0x4e, 0x61, 0x64, 0x9b, 0x1b, 0xa1, 0x9f, 0x36, 0x82, 0x08, 0x9f, 0x88, 0x23, 0x0f, 0x67,
	0x70, 0x6a, 0x1b, 0x2e, 0x4a, 0x21, 0x6b, 0xfb, 0x02, 0x0d, 0x08, 0x7d, 0x2d, 0x42, 0x84, 0x40,
	0x56, 0x4b, 0xc5, 0x7a, 0x84, 0xd0, 0x19, 0x3f, 0x40, 0x58, 0xed, 0xb6, 0x85, 0xd0, 0x2c, 0xa4,
	0x2b, 0x76, 0xca, 0xdd, 0xfd, 0x97, 0xb4, 0x95, 0x30, 0x86, 0xf5, 0x63, 0x2f, 0x19, 0xf0, 0xbd,
	0x74, 0x7f, 0xb1, 0x72, 0x2b, 0x58, 0x44, 0x3c, 0x9d, 0xa7, 0xbf, 0x4f, 0xa0, 0x47, 0xbd, 0xf1,
	0x33, 0x04, 0x6b, 0x91, 0x3f, 0x52, 0xe1, 0xe1, 0x1c, 0xd3, 0x6e, 0xc9, 0xf4, 0xb0, 0x0a, 0xa7,
	0x1c, 0xcf, 0x21, 0x70, 0x8b, 0x52, 0xf7, 0xe1, 0x7c, 0x9c, 0x1a, 0xb9, 0xaa, 0x8b, 0xf4, 0x7e,
	0x3f, 0x32, 0xa7, 0x14, 0x27, 0xe0, 0xdb, 0xc6, 0x30, 0x3f, 0xf6, 0x93, 0xe1, 0x3c, 0x4a, 0x6d,
	0x53, 0x17, 0xe9, 0x43, 0xc3, 0x9d, 0x89, 0x5f, 0x21, 0xd2, 0x6d, 0x45, 0xc3, 0x02, 0x02, 0xde,
	0x1f, 0x80, 0xd6, 0xe7, 0x07, 0x00, 0x27, 0x10, 0xd9, 0xc6, 0x8d, 0x20, 0x0c, 0xeb, 0xc5, 0x7e,
	0x32, 0xe2, 0x07, 0x8d, 0xe7, 0xf0, 0xb6, 0xe3, 0x3a, 0x20, 0x24, 0xe0, 0xd8, 0xc4, 0x19, 0x0c,
	0xa8, 0xcb, 0xc3, 0x73, 0x2d, 0x68, 0x92, 0x77, 0xff, 0xb7, 0x73, 0x09, 0xff, 0x07, 0x7d, 0xf9,
	0xd4, 0xbd, 0x05, 0x27, 0x10, 0x20, 0xbc, 0xbd, 0xe3, 0x37, 0x97, 0xd7, 0xe3, 0x37, 0x38, 0x82,
	0xe8, 0xee, 0xf6, 0xfa, 0xfb, 0xe2, 0xf2, 0x7e, 0x31, 0xf6, 0xae, 0x66, 0x3f, 0xd2, 0x95, 0xb4,
	0xeb, 0x5d, 0x91, 0x96, 0x6a, 0x9b, 0x49, 0x65, 0xec, 0x85, 0x5a, 0x2e, 0x65, 0x29, 0xf3, 0x4d,
	0xb6, 0x52, 0x17, 0xce, 0xc8, 0x8e, 0x9e, 0x6a, 0x11, 0xd2, 0x63, 0xfb, 0xf6, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x0f, 0xa7, 0xd1, 0xc2, 0x02, 0x00, 0x00,
}
