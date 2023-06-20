// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/liquidfarming/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type LiquidFarmCreateProposal struct {
	Title        string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string                                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PoolId       uint64                                 `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	LowerPrice   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=lower_price,json=lowerPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"lower_price"`
	UpperPrice   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=upper_price,json=upperPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"upper_price"`
	MinBidAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=min_bid_amount,json=minBidAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bid_amount"`
	FeeRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate"`
}

func (m *LiquidFarmCreateProposal) Reset()      { *m = LiquidFarmCreateProposal{} }
func (*LiquidFarmCreateProposal) ProtoMessage() {}
func (*LiquidFarmCreateProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a6ebab9c2ba4f8, []int{0}
}
func (m *LiquidFarmCreateProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidFarmCreateProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidFarmCreateProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidFarmCreateProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidFarmCreateProposal.Merge(m, src)
}
func (m *LiquidFarmCreateProposal) XXX_Size() int {
	return m.Size()
}
func (m *LiquidFarmCreateProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidFarmCreateProposal.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidFarmCreateProposal proto.InternalMessageInfo

type LiquidFarmParameterChangeProposal struct {
	Title       string                      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Changes     []LiquidFarmParameterChange `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes"`
}

func (m *LiquidFarmParameterChangeProposal) Reset()      { *m = LiquidFarmParameterChangeProposal{} }
func (*LiquidFarmParameterChangeProposal) ProtoMessage() {}
func (*LiquidFarmParameterChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a6ebab9c2ba4f8, []int{1}
}
func (m *LiquidFarmParameterChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidFarmParameterChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidFarmParameterChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidFarmParameterChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidFarmParameterChangeProposal.Merge(m, src)
}
func (m *LiquidFarmParameterChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *LiquidFarmParameterChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidFarmParameterChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidFarmParameterChangeProposal proto.InternalMessageInfo

type LiquidFarmParameterChange struct {
	LiquidFarmId uint64                                 `protobuf:"varint,1,opt,name=liquid_farm_id,json=liquidFarmId,proto3" json:"liquid_farm_id,omitempty"`
	MinBidAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_bid_amount,json=minBidAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bid_amount"`
	FeeRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate"`
}

func (m *LiquidFarmParameterChange) Reset()         { *m = LiquidFarmParameterChange{} }
func (m *LiquidFarmParameterChange) String() string { return proto.CompactTextString(m) }
func (*LiquidFarmParameterChange) ProtoMessage()    {}
func (*LiquidFarmParameterChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_03a6ebab9c2ba4f8, []int{2}
}
func (m *LiquidFarmParameterChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidFarmParameterChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidFarmParameterChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidFarmParameterChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidFarmParameterChange.Merge(m, src)
}
func (m *LiquidFarmParameterChange) XXX_Size() int {
	return m.Size()
}
func (m *LiquidFarmParameterChange) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidFarmParameterChange.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidFarmParameterChange proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LiquidFarmCreateProposal)(nil), "crescent.liquidfarming.v1beta1.LiquidFarmCreateProposal")
	proto.RegisterType((*LiquidFarmParameterChangeProposal)(nil), "crescent.liquidfarming.v1beta1.LiquidFarmParameterChangeProposal")
	proto.RegisterType((*LiquidFarmParameterChange)(nil), "crescent.liquidfarming.v1beta1.LiquidFarmParameterChange")
}

func init() {
	proto.RegisterFile("crescent/liquidfarming/v1beta1/proposal.proto", fileDescriptor_03a6ebab9c2ba4f8)
}

var fileDescriptor_03a6ebab9c2ba4f8 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x75, 0x9a, 0xc0, 0xa5, 0xea, 0x60, 0x55, 0xc2, 0x30, 0x38, 0xa1, 0x42, 0x28,
	0x4b, 0x6c, 0x15, 0xc4, 0x00, 0x12, 0x03, 0x29, 0x42, 0x8a, 0x84, 0x44, 0x64, 0xb1, 0x00, 0x83,
	0x75, 0xb1, 0xdf, 0xb8, 0xa7, 0xfa, 0xfe, 0x70, 0x77, 0x69, 0xe1, 0x5b, 0x30, 0x32, 0xf2, 0x21,
	0x18, 0xf9, 0x00, 0x19, 0x3b, 0x22, 0x86, 0x0a, 0x92, 0x2f, 0x82, 0xee, 0xe2, 0x86, 0x82, 0x28,
	0x12, 0x4d, 0x27, 0xfb, 0xce, 0xcf, 0xf3, 0x7b, 0x5f, 0xdf, 0xf3, 0x1e, 0xee, 0xe7, 0x0a, 0x74,
	0x0e, 0xdc, 0x24, 0x15, 0x7d, 0x3b, 0xa5, 0xc5, 0x84, 0x28, 0x46, 0x79, 0x99, 0x1c, 0xed, 0x8d,
	0xc1, 0x90, 0xbd, 0x44, 0x2a, 0x21, 0x85, 0x26, 0x55, 0x2c, 0x95, 0x30, 0x22, 0x88, 0xce, 0xe4,
	0xf1, 0x6f, 0xf2, 0xb8, 0x96, 0xdf, 0xda, 0x29, 0x45, 0x29, 0x9c, 0x34, 0xb1, 0x6f, 0x4b, 0xd7,
	0xee, 0x67, 0x1f, 0x87, 0xcf, 0x9d, 0xfe, 0x19, 0x51, 0x6c, 0x5f, 0x01, 0x31, 0x30, 0xaa, 0xc1,
	0xc1, 0x0e, 0xde, 0x34, 0xd4, 0x54, 0x10, 0xa2, 0x2e, 0xea, 0x5d, 0x4f, 0x97, 0x8b, 0xa0, 0x8b,
	0xdb, 0x05, 0xe8, 0x5c, 0x51, 0x69, 0xa8, 0xe0, 0xe1, 0x86, 0xfb, 0x76, 0x7e, 0x2b, 0xb8, 0x81,
	0x5b, 0x52, 0x88, 0x2a, 0xa3, 0x45, 0xe8, 0x77, 0x51, 0xaf, 0x91, 0x36, 0xed, 0x72, 0x58, 0x04,
	0x2f, 0x70, 0xbb, 0x12, 0xc7, 0xa0, 0x32, 0xa9, 0x68, 0x0e, 0x61, 0xc3, 0x5a, 0x07, 0xf1, 0xec,
	0xb4, 0xe3, 0x7d, 0x3b, 0xed, 0xdc, 0x2d, 0xa9, 0x39, 0x98, 0x8e, 0xe3, 0x5c, 0xb0, 0x24, 0x17,
	0x9a, 0x09, 0x5d, 0x3f, 0xfa, 0xba, 0x38, 0x4c, 0xcc, 0x7b, 0x09, 0x3a, 0x7e, 0x0a, 0x79, 0x8a,
	0x1d, 0x62, 0x64, 0x09, 0x16, 0x38, 0x95, 0x72, 0x05, 0xdc, 0xbc, 0x1c, 0xd0, 0x21, 0x96, 0xc0,
	0x97, 0x78, 0x9b, 0x51, 0x9e, 0x8d, 0x69, 0x91, 0x11, 0x26, 0xa6, 0xdc, 0x84, 0xcd, 0xff, 0x66,
	0x0e, 0xb9, 0x49, 0xb7, 0x18, 0xe5, 0x03, 0x5a, 0x3c, 0x71, 0x8c, 0x60, 0x88, 0xaf, 0x4d, 0x00,
	0x32, 0x45, 0x0c, 0x84, 0xad, 0x4b, 0xf5, 0xd8, 0x9a, 0x00, 0xa4, 0xc4, 0xc0, 0xa3, 0xc6, 0xc7,
	0x4f, 0x1d, 0x6f, 0xf7, 0x0b, 0xc2, 0xb7, 0x7f, 0xc5, 0x36, 0x22, 0x8a, 0x30, 0x30, 0xa0, 0xf6,
	0x0f, 0x08, 0x2f, 0xd7, 0xcf, 0xef, 0x15, 0x6e, 0xe5, 0x8e, 0xa4, 0x43, 0xbf, 0xeb, 0xf7, 0xda,
	0xf7, 0x1e, 0xc6, 0xff, 0x1e, 0xae, 0xf8, 0xc2, 0x5e, 0x06, 0x0d, 0xfb, 0xa3, 0xe9, 0x19, 0xaf,
	0x6e, 0x7f, 0x8e, 0xf0, 0xcd, 0x0b, 0x2d, 0xc1, 0x1d, 0xbc, 0xbd, 0xac, 0x92, 0xd9, 0x32, 0x76,
	0x8a, 0x90, 0x9b, 0xa2, 0xad, 0x6a, 0x65, 0x19, 0x16, 0x7f, 0x49, 0x6a, 0xe3, 0x8a, 0x93, 0xf2,
	0xd7, 0x4a, 0x6a, 0xf0, 0x66, 0xf6, 0x23, 0xf2, 0x66, 0xf3, 0x08, 0x9d, 0xcc, 0x23, 0xf4, 0x7d,
	0x1e, 0xa1, 0x0f, 0x8b, 0xc8, 0x3b, 0x59, 0x44, 0xde, 0xd7, 0x45, 0xe4, 0xbd, 0x7e, 0x7c, 0x1e,
	0x57, 0x1f, 0x6e, 0x9f, 0x83, 0x39, 0x16, 0xea, 0x70, 0xb5, 0x91, 0x1c, 0x3d, 0x48, 0xde, 0xfd,
	0x71, 0xfd, 0x5d, 0xa5, 0x71, 0xd3, 0x5d, 0xdf, 0xfb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x63,
	0xbe, 0x25, 0x38, 0x25, 0x04, 0x00, 0x00,
}

func (m *LiquidFarmCreateProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidFarmCreateProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidFarmCreateProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.MinBidAmount.Size()
		i -= size
		if _, err := m.MinBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.UpperPrice.Size()
		i -= size
		if _, err := m.UpperPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LowerPrice.Size()
		i -= size
		if _, err := m.LowerPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PoolId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidFarmParameterChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidFarmParameterChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidFarmParameterChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidFarmParameterChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidFarmParameterChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidFarmParameterChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinBidAmount.Size()
		i -= size
		if _, err := m.MinBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LiquidFarmId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.LiquidFarmId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidFarmCreateProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovProposal(uint64(m.PoolId))
	}
	l = m.LowerPrice.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.UpperPrice.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.MinBidAmount.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.FeeRate.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *LiquidFarmParameterChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *LiquidFarmParameterChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LiquidFarmId != 0 {
		n += 1 + sovProposal(uint64(m.LiquidFarmId))
	}
	l = m.MinBidAmount.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.FeeRate.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidFarmCreateProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: LiquidFarmCreateProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidFarmCreateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LowerPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UpperPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *LiquidFarmParameterChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: LiquidFarmParameterChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidFarmParameterChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, LiquidFarmParameterChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *LiquidFarmParameterChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: LiquidFarmParameterChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidFarmParameterChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidFarmId", wireType)
			}
			m.LiquidFarmId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidFarmId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
