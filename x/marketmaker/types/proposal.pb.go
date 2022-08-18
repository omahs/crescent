// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/marketmaker/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MarketMakerProposal struct {
	// title specifies the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description specifies the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// set the market makers to eligible, refund deposit
	Inclusions []MarketMakerHandle `protobuf:"bytes,3,rep,name=inclusions,proto3" json:"inclusions" yaml:"inclusions"`
	// delete existing eligible market makers
	Exclusions []MarketMakerHandle `protobuf:"bytes,4,rep,name=exclusions,proto3" json:"exclusions" yaml:"exclusions"`
	// delete the not eligible market makers, refund deposit
	Rejections []MarketMakerHandle `protobuf:"bytes,5,rep,name=rejections,proto3" json:"rejections" yaml:"rejections"`
	// distribute claimable incentive to eligible market makers
	Distributions []IncentiveDistribution `protobuf:"bytes,6,rep,name=distributions,proto3" json:"distributions" yaml:"distributions"`
}

func (m *MarketMakerProposal) Reset()      { *m = MarketMakerProposal{} }
func (*MarketMakerProposal) ProtoMessage() {}
func (*MarketMakerProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_762a5d4b14c9db68, []int{0}
}
func (m *MarketMakerProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMakerProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMakerProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMakerProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMakerProposal.Merge(m, src)
}
func (m *MarketMakerProposal) XXX_Size() int {
	return m.Size()
}
func (m *MarketMakerProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMakerProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMakerProposal proto.InternalMessageInfo

type MarketMakerHandle struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	PairId  uint64 `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
}

func (m *MarketMakerHandle) Reset()         { *m = MarketMakerHandle{} }
func (m *MarketMakerHandle) String() string { return proto.CompactTextString(m) }
func (*MarketMakerHandle) ProtoMessage()    {}
func (*MarketMakerHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_762a5d4b14c9db68, []int{1}
}
func (m *MarketMakerHandle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMakerHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMakerHandle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMakerHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMakerHandle.Merge(m, src)
}
func (m *MarketMakerHandle) XXX_Size() int {
	return m.Size()
}
func (m *MarketMakerHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMakerHandle.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMakerHandle proto.InternalMessageInfo

type IncentiveDistribution struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	PairId  uint64                                   `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	Amount  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *IncentiveDistribution) Reset()         { *m = IncentiveDistribution{} }
func (m *IncentiveDistribution) String() string { return proto.CompactTextString(m) }
func (*IncentiveDistribution) ProtoMessage()    {}
func (*IncentiveDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_762a5d4b14c9db68, []int{2}
}
func (m *IncentiveDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveDistribution.Merge(m, src)
}
func (m *IncentiveDistribution) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveDistribution proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarketMakerProposal)(nil), "crescent.marketmaker.v1beta1.MarketMakerProposal")
	proto.RegisterType((*MarketMakerHandle)(nil), "crescent.marketmaker.v1beta1.MarketMakerHandle")
	proto.RegisterType((*IncentiveDistribution)(nil), "crescent.marketmaker.v1beta1.IncentiveDistribution")
}

func init() {
	proto.RegisterFile("crescent/marketmaker/v1beta1/proposal.proto", fileDescriptor_762a5d4b14c9db68)
}

var fileDescriptor_762a5d4b14c9db68 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x31, 0x6f, 0x13, 0x31,
	0x18, 0x86, 0xef, 0x68, 0x9a, 0x82, 0xab, 0x22, 0xf5, 0x08, 0x52, 0x52, 0x55, 0x77, 0x91, 0xa7,
	0x48, 0xa5, 0x67, 0x95, 0x2e, 0xa8, 0x1b, 0x81, 0x81, 0x0e, 0x95, 0xd0, 0x8d, 0x2c, 0x95, 0xef,
	0xce, 0x04, 0x37, 0x77, 0xf6, 0xc9, 0x76, 0xda, 0xf4, 0x1f, 0x30, 0x32, 0x30, 0x30, 0x76, 0xe6,
	0x97, 0x74, 0xec, 0xc8, 0x14, 0x50, 0xb2, 0x30, 0xe7, 0x17, 0xa0, 0xb3, 0xdd, 0xc4, 0x11, 0x28,
	0x53, 0x99, 0x72, 0xdf, 0xbd, 0xef, 0xf7, 0x3e, 0xb6, 0xe3, 0xef, 0xc0, 0x41, 0x26, 0x88, 0xcc,
	0x08, 0x53, 0xa8, 0xc4, 0x62, 0x48, 0x54, 0x89, 0x87, 0x44, 0xa0, 0xcb, 0xa3, 0x94, 0x28, 0x7c,
	0x84, 0x2a, 0xc1, 0x2b, 0x2e, 0x71, 0x11, 0x57, 0x82, 0x2b, 0x1e, 0xec, 0xdf, 0x9b, 0x63, 0xc7,
	0x1c, 0x5b, 0xf3, 0x5e, 0x6b, 0xc0, 0x07, 0x5c, 0x1b, 0x51, 0xfd, 0x64, 0x7a, 0xf6, 0x3a, 0x19,
	0x97, 0x25, 0x97, 0xe7, 0x46, 0x30, 0x85, 0x95, 0x42, 0x53, 0xa1, 0x14, 0x4b, 0xb2, 0x40, 0x66,
	0x9c, 0x32, 0xab, 0xc7, 0x6b, 0xd7, 0xe6, 0x2e, 0xc1, 0xf8, 0xa3, 0x01, 0xe7, 0x83, 0x82, 0x20,
	0x5d, 0xa5, 0xa3, 0x8f, 0x48, 0xd1, 0x92, 0x48, 0x85, 0xcb, 0xca, 0x18, 0xe0, 0xd7, 0x06, 0x78,
	0x76, 0xa6, 0xdb, 0xce, 0xea, 0xb6, 0xf7, 0x76, 0x77, 0x41, 0x0b, 0x6c, 0x2a, 0xaa, 0x0a, 0xd2,
	0xf6, 0xbb, 0x7e, 0xef, 0x49, 0x62, 0x8a, 0xa0, 0x0b, 0xb6, 0x73, 0x22, 0x33, 0x41, 0x2b, 0x45,
	0x39, 0x6b, 0x3f, 0xd2, 0x9a, 0xfb, 0x2a, 0xb8, 0x00, 0x80, 0xb2, 0xac, 0x18, 0x49, 0xca, 0x99,
	0x6c, 0x6f, 0x74, 0x37, 0x7a, 0xdb, 0x2f, 0x51, 0xbc, 0xee, 0x90, 0x62, 0x07, 0xff, 0x0e, 0xb3,
	0xbc, 0x20, 0xfd, 0xce, 0xed, 0x24, 0xf2, 0xe6, 0x93, 0x68, 0xf7, 0x1a, 0x97, 0xc5, 0x09, 0x5c,
	0x06, 0xc2, 0xc4, 0x49, 0xaf, 0x59, 0x64, 0xbc, 0x60, 0x35, 0x1e, 0x84, 0xb5, 0x0c, 0x84, 0x89,
	0x93, 0x5e, 0xb3, 0x04, 0xb9, 0x20, 0x99, 0xd2, 0xac, 0xcd, 0x07, 0x61, 0x2d, 0x03, 0x61, 0xe2,
	0xa4, 0x07, 0x57, 0x60, 0x27, 0xa7, 0x52, 0x09, 0x9a, 0x8e, 0x0c, 0xae, 0xa9, 0x71, 0xc7, 0xeb,
	0x71, 0xa7, 0xac, 0x96, 0xe8, 0x25, 0x79, 0xeb, 0xf4, 0xf6, 0xf7, 0x2d, 0xb2, 0x65, 0x90, 0x2b,
	0xb9, 0x30, 0x59, 0xe5, 0x9c, 0x3c, 0xfe, 0x7c, 0x13, 0x79, 0xdf, 0x6e, 0x22, 0x0f, 0x8e, 0xc1,
	0xee, 0x5f, 0xcb, 0x0f, 0x5e, 0x80, 0x2d, 0x9c, 0xe7, 0x82, 0x48, 0x69, 0x6e, 0x45, 0x3f, 0x98,
	0x4f, 0xa2, 0xa7, 0x26, 0xd8, 0x0a, 0x30, 0xb9, 0xb7, 0x04, 0x07, 0x60, 0xab, 0xc2, 0x54, 0x9c,
	0xd3, 0x5c, 0xdf, 0x93, 0x86, 0xeb, 0xb6, 0x02, 0x4c, 0x9a, 0xf5, 0xd3, 0x69, 0x6e, 0xc8, 0xbf,
	0x6b, 0xf2, 0xdc, 0x07, 0xcf, 0xff, 0xb9, 0x95, 0xff, 0x88, 0x0f, 0x14, 0x68, 0xe2, 0x92, 0x8f,
	0x98, 0xb2, 0x37, 0xb6, 0x13, 0xdb, 0xa9, 0xac, 0xe7, 0x70, 0x71, 0xc2, 0x6f, 0x38, 0x65, 0xfd,
	0xd7, 0xf6, 0x40, 0x77, 0x2c, 0x58, 0xb7, 0xc1, 0xef, 0x3f, 0xa3, 0xde, 0x80, 0xaa, 0x4f, 0xa3,
	0x34, 0xce, 0x78, 0x69, 0x67, 0xda, 0xfe, 0x1c, 0xca, 0x7c, 0x88, 0xd4, 0x75, 0x45, 0xa4, 0x4e,
	0x90, 0x89, 0x65, 0x2d, 0x37, 0xdd, 0x4f, 0x6e, 0xa7, 0xa1, 0x7f, 0x37, 0x0d, 0xfd, 0x5f, 0xd3,
	0xd0, 0xff, 0x32, 0x0b, 0xbd, 0xbb, 0x59, 0xe8, 0xfd, 0x98, 0x85, 0xde, 0x87, 0x57, 0x6e, 0xaa,
	0xfd, 0xfb, 0x0f, 0x19, 0x51, 0x57, 0x5c, 0x0c, 0x17, 0x2f, 0xd0, 0x78, 0xe5, 0x73, 0xa0, 0x59,
	0x69, 0x53, 0x0f, 0xf8, 0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0xda, 0x87, 0x5d, 0xcf,
	0x04, 0x00, 0x00,
}

func (m *MarketMakerProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMakerProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMakerProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Distributions) > 0 {
		for iNdEx := len(m.Distributions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Distributions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Rejections) > 0 {
		for iNdEx := len(m.Rejections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Exclusions) > 0 {
		for iNdEx := len(m.Exclusions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Exclusions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Inclusions) > 0 {
		for iNdEx := len(m.Inclusions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inclusions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *MarketMakerHandle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMakerHandle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMakerHandle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *MarketMakerProposal) Size() (n int) {
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
	if len(m.Inclusions) > 0 {
		for _, e := range m.Inclusions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Exclusions) > 0 {
		for _, e := range m.Exclusions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Rejections) > 0 {
		for _, e := range m.Rejections {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Distributions) > 0 {
		for _, e := range m.Distributions {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *MarketMakerHandle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	return n
}

func (m *IncentiveDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketMakerProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MarketMakerProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMakerProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Inclusions", wireType)
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
			m.Inclusions = append(m.Inclusions, MarketMakerHandle{})
			if err := m.Inclusions[len(m.Inclusions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exclusions", wireType)
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
			m.Exclusions = append(m.Exclusions, MarketMakerHandle{})
			if err := m.Exclusions[len(m.Exclusions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejections", wireType)
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
			m.Rejections = append(m.Rejections, MarketMakerHandle{})
			if err := m.Rejections[len(m.Rejections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distributions", wireType)
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
			m.Distributions = append(m.Distributions, IncentiveDistribution{})
			if err := m.Distributions[len(m.Distributions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MarketMakerHandle) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MarketMakerHandle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMakerHandle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *IncentiveDistribution) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IncentiveDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
