// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rollapp/state_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StateInfoIndex is the data used for indexing and retrieving a StateInfo
// it updated and saved with every UpdateState in StateInfo.
// We use the this structure also for LatestStateInfoIndex which defines the rollapps' current (latest)
// index of the last UpdateState
type StateInfoIndex struct {
	// rollappId is the rollapp that the sequencer belongs to and asking to update
	// it used to identify the what rollapp a StateInfo belongs
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// latestStateInfoIndex is a sequential increasing number, updating on each
	// state update used for indexing to a specific state info
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *StateInfoIndex) Reset()         { *m = StateInfoIndex{} }
func (m *StateInfoIndex) String() string { return proto.CompactTextString(m) }
func (*StateInfoIndex) ProtoMessage()    {}
func (*StateInfoIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f682a28ae3061a0, []int{0}
}
func (m *StateInfoIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfoIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfoIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfoIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfoIndex.Merge(m, src)
}
func (m *StateInfoIndex) XXX_Size() int {
	return m.Size()
}
func (m *StateInfoIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfoIndex.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfoIndex proto.InternalMessageInfo

func (m *StateInfoIndex) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *StateInfoIndex) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// StateInfo defines a rollapps' state.
type StateInfo struct {
	// stateInfoIndex defines what rollapp the state belongs to
	// and in which index it can be referenced
	StateInfoIndex StateInfoIndex `protobuf:"bytes,1,opt,name=stateInfoIndex,proto3" json:"stateInfoIndex"`
	// sequencer is the bech32-encoded address of the sequencer sent the update
	Sequencer string `protobuf:"bytes,2,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	// startHeight is the block height of the first block in the batch
	StartHeight uint64 `protobuf:"varint,3,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	// numBlocks is the number of blocks included in this batch update
	NumBlocks uint64 `protobuf:"varint,4,opt,name=numBlocks,proto3" json:"numBlocks,omitempty"`
	// DAPath is the description of the location on the DA layer
	DAPath string `protobuf:"bytes,5,opt,name=DAPath,proto3" json:"DAPath,omitempty"`
	// version is the version of the rollapp
	Version uint64 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	// creationHeight is the height at which the UpdateState took place
	CreationHeight uint64 `protobuf:"varint,7,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
	// status is the status of the state update
	Status StateStatus `protobuf:"varint,8,opt,name=status,proto3,enum=dymensionxyz.dymension.rollapp.StateStatus" json:"status,omitempty"`
	// BDs is a list of block description objects (one per block)
	// the list must be ordered by height, starting from startHeight to startHeight+numBlocks-1
	BDs BlockDescriptors `protobuf:"bytes,9,opt,name=BDs,proto3" json:"BDs"`
}

func (m *StateInfo) Reset()         { *m = StateInfo{} }
func (m *StateInfo) String() string { return proto.CompactTextString(m) }
func (*StateInfo) ProtoMessage()    {}
func (*StateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f682a28ae3061a0, []int{1}
}
func (m *StateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfo.Merge(m, src)
}
func (m *StateInfo) XXX_Size() int {
	return m.Size()
}
func (m *StateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfo proto.InternalMessageInfo

func (m *StateInfo) GetStateInfoIndex() StateInfoIndex {
	if m != nil {
		return m.StateInfoIndex
	}
	return StateInfoIndex{}
}

func (m *StateInfo) GetSequencer() string {
	if m != nil {
		return m.Sequencer
	}
	return ""
}

func (m *StateInfo) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *StateInfo) GetNumBlocks() uint64 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *StateInfo) GetDAPath() string {
	if m != nil {
		return m.DAPath
	}
	return ""
}

func (m *StateInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StateInfo) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *StateInfo) GetStatus() StateStatus {
	if m != nil {
		return m.Status
	}
	return STATE_STATUS_UNSPECIFIED
}

func (m *StateInfo) GetBDs() BlockDescriptors {
	if m != nil {
		return m.BDs
	}
	return BlockDescriptors{}
}

// BlockHeightToFinalizationQueue defines a map from block height to list of states to finalized
type BlockHeightToFinalizationQueue struct {
	// finalizationHeight is the block height that the state should be finalized
	FinalizationHeight uint64 `protobuf:"varint,1,opt,name=finalizationHeight,proto3" json:"finalizationHeight,omitempty"`
	// finalizationQueue is a list of states that are waiting to be finalized
	// when the block height becomes finalizationHeight
	FinalizationQueue []StateInfoIndex `protobuf:"bytes,2,rep,name=finalizationQueue,proto3" json:"finalizationQueue"`
}

func (m *BlockHeightToFinalizationQueue) Reset()         { *m = BlockHeightToFinalizationQueue{} }
func (m *BlockHeightToFinalizationQueue) String() string { return proto.CompactTextString(m) }
func (*BlockHeightToFinalizationQueue) ProtoMessage()    {}
func (*BlockHeightToFinalizationQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f682a28ae3061a0, []int{2}
}
func (m *BlockHeightToFinalizationQueue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeightToFinalizationQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeightToFinalizationQueue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeightToFinalizationQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightToFinalizationQueue.Merge(m, src)
}
func (m *BlockHeightToFinalizationQueue) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeightToFinalizationQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightToFinalizationQueue.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightToFinalizationQueue proto.InternalMessageInfo

func (m *BlockHeightToFinalizationQueue) GetFinalizationHeight() uint64 {
	if m != nil {
		return m.FinalizationHeight
	}
	return 0
}

func (m *BlockHeightToFinalizationQueue) GetFinalizationQueue() []StateInfoIndex {
	if m != nil {
		return m.FinalizationQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*StateInfoIndex)(nil), "dymensionxyz.dymension.rollapp.StateInfoIndex")
	proto.RegisterType((*StateInfo)(nil), "dymensionxyz.dymension.rollapp.StateInfo")
	proto.RegisterType((*BlockHeightToFinalizationQueue)(nil), "dymensionxyz.dymension.rollapp.BlockHeightToFinalizationQueue")
}

func init() { proto.RegisterFile("rollapp/state_info.proto", fileDescriptor_0f682a28ae3061a0) }

var fileDescriptor_0f682a28ae3061a0 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x6d, 0xb6, 0xdd, 0xae, 0x93, 0x85, 0x82, 0x61, 0x91, 0x50, 0x24, 0x0e, 0x7d, 0x90, 0x82,
	0x90, 0xd1, 0xf5, 0x0b, 0xac, 0x45, 0xb6, 0xf8, 0xa2, 0xb3, 0x3e, 0x89, 0xb0, 0x4c, 0xa7, 0xe9,
	0x34, 0x38, 0x4d, 0xc6, 0x24, 0x23, 0xed, 0x7e, 0x85, 0x9f, 0xe2, 0x83, 0x1f, 0xb1, 0x8f, 0xfb,
	0xe8, 0x93, 0x48, 0xfb, 0x23, 0x92, 0x4c, 0xda, 0x69, 0xab, 0xa8, 0xec, 0xcb, 0x30, 0xf7, 0xdc,
	0x7b, 0x4e, 0xce, 0x3d, 0x33, 0x81, 0x58, 0xc9, 0x3c, 0x4f, 0x8a, 0x22, 0xd2, 0x26, 0x31, 0xec,
	0x8a, 0x8b, 0xa9, 0xa4, 0x85, 0x92, 0x46, 0x22, 0x32, 0x59, 0xce, 0x99, 0xd0, 0x5c, 0x8a, 0xc5,
	0xf2, 0x9a, 0x6e, 0x0b, 0xea, 0x09, 0xdd, 0xb3, 0x4c, 0x66, 0xd2, 0x8d, 0x46, 0xf6, 0xad, 0x62,
	0x75, 0xc9, 0x46, 0x6f, 0x9c, 0xcb, 0xf4, 0xe3, 0xd5, 0x84, 0xe9, 0x54, 0xf1, 0xc2, 0x48, 0xe5,
	0xfb, 0xdd, 0xfd, 0xf3, 0xec, 0xb3, 0xd4, 0x55, 0xaf, 0x37, 0x84, 0x9d, 0x4b, 0x8b, 0x8e, 0xc4,
	0x54, 0x8e, 0xc4, 0x84, 0x2d, 0xd0, 0x43, 0x18, 0xf8, 0xf9, 0xd1, 0x04, 0x83, 0x10, 0xf4, 0x83,
	0xb8, 0x06, 0xd0, 0x19, 0x3c, 0xe6, 0x76, 0x0c, 0x1f, 0x85, 0xa0, 0xdf, 0x8a, 0xab, 0xa2, 0xf7,
	0xb5, 0x09, 0x83, 0xad, 0x0c, 0xfa, 0x00, 0x3b, 0x7a, 0x4f, 0xd3, 0xc9, 0x9c, 0x9e, 0x53, 0xfa,
	0xf7, 0xf5, 0xe8, 0xbe, 0x93, 0x41, 0xeb, 0xe6, 0xc7, 0xa3, 0x46, 0x7c, 0xa0, 0x65, 0xfd, 0x69,
	0xf6, 0xa9, 0x64, 0x22, 0x65, 0xca, 0xb9, 0x08, 0xe2, 0x1a, 0x40, 0x21, 0x3c, 0xd5, 0x26, 0x51,
	0xe6, 0x82, 0xf1, 0x6c, 0x66, 0x70, 0xd3, 0xb9, 0xdc, 0x85, 0x2c, 0x5f, 0x94, 0xf3, 0x81, 0x8d,
	0x4a, 0xe3, 0x96, 0xeb, 0xd7, 0x00, 0x7a, 0x00, 0xdb, 0xc3, 0x17, 0x6f, 0x12, 0x33, 0xc3, 0xc7,
	0x4e, 0xda, 0x57, 0x08, 0xc3, 0x93, 0xcf, 0x4c, 0x59, 0xb7, 0xb8, 0xed, 0x38, 0x9b, 0x12, 0x3d,
	0x86, 0x9d, 0x54, 0xb1, 0xc4, 0x70, 0x29, 0xfc, 0xa1, 0x27, 0x6e, 0xe0, 0x00, 0x45, 0x2f, 0x61,
	0xbb, 0x4a, 0x1e, 0xdf, 0x0b, 0x41, 0xbf, 0x73, 0xfe, 0xe4, 0xbf, 0xd2, 0xb8, 0x74, 0x94, 0xd8,
	0x53, 0xd1, 0x05, 0x6c, 0x0e, 0x86, 0x1a, 0x07, 0x2e, 0xcf, 0xa7, 0xff, 0x52, 0x70, 0x3b, 0x0d,
	0xb7, 0xbf, 0x83, 0xf6, 0x89, 0x5a, 0x89, 0xde, 0x37, 0x00, 0x89, 0xeb, 0x57, 0xf6, 0xde, 0xc9,
	0x57, 0x5c, 0x24, 0x39, 0xbf, 0x76, 0x96, 0xdf, 0x96, 0xac, 0x64, 0x88, 0x42, 0x34, 0xdd, 0x01,
	0xfd, 0x76, 0xc0, 0x6d, 0xf7, 0x87, 0x0e, 0x1a, 0xc3, 0xfb, 0xd3, 0x43, 0x11, 0x7c, 0x14, 0x36,
	0xef, 0xfc, 0xe9, 0x7f, 0x97, 0x1b, 0xbc, 0xbe, 0x59, 0x11, 0x70, 0xbb, 0x22, 0xe0, 0xe7, 0x8a,
	0x80, 0x2f, 0x6b, 0xd2, 0xb8, 0x5d, 0x93, 0xc6, 0xf7, 0x35, 0x69, 0xbc, 0x7f, 0x96, 0x71, 0x33,
	0x2b, 0xc7, 0x34, 0x95, 0xf3, 0x68, 0xf7, 0xb0, 0xba, 0x88, 0x16, 0xd1, 0xe6, 0x26, 0x98, 0x65,
	0xc1, 0xf4, 0xb8, 0xed, 0xee, 0xc0, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xe9, 0x1c,
	0x90, 0x91, 0x03, 0x00, 0x00,
}

func (m *StateInfoIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfoIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfoIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.Version != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DAPath) > 0 {
		i -= len(m.DAPath)
		copy(dAtA[i:], m.DAPath)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.DAPath)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocks != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.StartHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Sequencer) > 0 {
		i -= len(m.Sequencer)
		copy(dAtA[i:], m.Sequencer)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.Sequencer)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.StateInfoIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockHeightToFinalizationQueue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeightToFinalizationQueue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeightToFinalizationQueue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalizationQueue) > 0 {
		for iNdEx := len(m.FinalizationQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FinalizationQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStateInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.FinalizationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.FinalizationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStateInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovStateInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StateInfoIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovStateInfo(uint64(m.Index))
	}
	return n
}

func (m *StateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StateInfoIndex.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	l = len(m.Sequencer)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.StartHeight))
	}
	if m.NumBlocks != 0 {
		n += 1 + sovStateInfo(uint64(m.NumBlocks))
	}
	l = len(m.DAPath)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovStateInfo(uint64(m.Version))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	if m.Status != 0 {
		n += 1 + sovStateInfo(uint64(m.Status))
	}
	l = m.BDs.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	return n
}

func (m *BlockHeightToFinalizationQueue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FinalizationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.FinalizationHeight))
	}
	if len(m.FinalizationQueue) > 0 {
		for _, e := range m.FinalizationQueue {
			l = e.Size()
			n += 1 + l + sovStateInfo(uint64(l))
		}
	}
	return n
}

func sovStateInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStateInfo(x uint64) (n int) {
	return sovStateInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StateInfoIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StateInfoIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfoIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateInfoIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StateStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeightToFinalizationQueue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizationHeight", wireType)
			}
			m.FinalizationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizationQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizationQueue = append(m.FinalizationQueue, StateInfoIndex{})
			if err := m.FinalizationQueue[len(m.FinalizationQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStateInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStateInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStateInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStateInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStateInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStateInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStateInfo = fmt.Errorf("proto: unexpected end of group")
)
