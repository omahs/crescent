// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/amm/v1beta1/amm.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	PoolCreationFee    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee"`
	DefaultTickSpacing uint32                                   `protobuf:"varint,2,opt,name=default_tick_spacing,json=defaultTickSpacing,proto3" json:"default_tick_spacing,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfef6a2c44f2449, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type Pool struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MarketId       uint64 `protobuf:"varint,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Denom0         string `protobuf:"bytes,3,opt,name=denom0,proto3" json:"denom0,omitempty"`
	Denom1         string `protobuf:"bytes,4,opt,name=denom1,proto3" json:"denom1,omitempty"`
	TickSpacing    uint32 `protobuf:"varint,5,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty"`
	ReserveAddress string `protobuf:"bytes,6,opt,name=reserve_address,json=reserveAddress,proto3" json:"reserve_address,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfef6a2c44f2449, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

type PoolState struct {
	CurrentTick      int32                                  `protobuf:"varint,1,opt,name=current_tick,json=currentTick,proto3" json:"current_tick,omitempty"`
	CurrentSqrtPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_sqrt_price"`
	CurrentLiquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=current_liquidity,json=currentLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_liquidity"`
}

func (m *PoolState) Reset()         { *m = PoolState{} }
func (m *PoolState) String() string { return proto.CompactTextString(m) }
func (*PoolState) ProtoMessage()    {}
func (*PoolState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfef6a2c44f2449, []int{2}
}
func (m *PoolState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolState.Merge(m, src)
}
func (m *PoolState) XXX_Size() int {
	return m.Size()
}
func (m *PoolState) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolState.DiscardUnknown(m)
}

var xxx_messageInfo_PoolState proto.InternalMessageInfo

type Position struct {
	Id        uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PoolId    uint64                                 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	Owner     string                                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	LowerTick int32                                  `protobuf:"varint,4,opt,name=lower_tick,json=lowerTick,proto3" json:"lower_tick,omitempty"`
	UpperTick int32                                  `protobuf:"varint,5,opt,name=upper_tick,json=upperTick,proto3" json:"upper_tick,omitempty"`
	Liquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidity"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfef6a2c44f2449, []int{3}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

type TickInfo struct {
	GrossLiquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=gross_liquidity,json=grossLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"gross_liquidity"`
	NetLiquidity   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=net_liquidity,json=netLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"net_liquidity"`
}

func (m *TickInfo) Reset()         { *m = TickInfo{} }
func (m *TickInfo) String() string { return proto.CompactTextString(m) }
func (*TickInfo) ProtoMessage()    {}
func (*TickInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfef6a2c44f2449, []int{4}
}
func (m *TickInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickInfo.Merge(m, src)
}
func (m *TickInfo) XXX_Size() int {
	return m.Size()
}
func (m *TickInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TickInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TickInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "crescent.amm.v1beta1.Params")
	proto.RegisterType((*Pool)(nil), "crescent.amm.v1beta1.Pool")
	proto.RegisterType((*PoolState)(nil), "crescent.amm.v1beta1.PoolState")
	proto.RegisterType((*Position)(nil), "crescent.amm.v1beta1.Position")
	proto.RegisterType((*TickInfo)(nil), "crescent.amm.v1beta1.TickInfo")
}

func init() { proto.RegisterFile("crescent/amm/v1beta1/amm.proto", fileDescriptor_1dfef6a2c44f2449) }

var fileDescriptor_1dfef6a2c44f2449 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x8d, 0xdb, 0x24, 0x34, 0xb7, 0x2f, 0x6a, 0x45, 0x10, 0x8a, 0x3a, 0x0d, 0x59, 0x40, 0x36,
	0x9d, 0x69, 0xa8, 0xf8, 0x00, 0x5a, 0x84, 0x54, 0xa9, 0x8b, 0x30, 0x41, 0x42, 0x02, 0xa4, 0xd1,
	0x64, 0xe6, 0x36, 0x58, 0xc9, 0x8c, 0xa7, 0xb6, 0xd3, 0xd0, 0xbf, 0xe0, 0x3b, 0xd8, 0xb0, 0x60,
	0xc3, 0x27, 0x74, 0xd9, 0x0d, 0x12, 0x62, 0x51, 0xa0, 0x59, 0xf1, 0x17, 0xc8, 0xf6, 0xe4, 0x81,
	0xd8, 0xa0, 0xac, 0xc6, 0x3e, 0xc7, 0x3e, 0xf7, 0xde, 0x73, 0x46, 0x06, 0x27, 0x12, 0x28, 0x23,
	0x4c, 0x95, 0x17, 0x26, 0x89, 0x77, 0xde, 0xea, 0xa2, 0x0a, 0x5b, 0x7a, 0xed, 0x66, 0x82, 0x2b,
	0x4e, 0xab, 0x13, 0xde, 0xd5, 0x58, 0xce, 0x6f, 0x57, 0x7b, 0xbc, 0xc7, 0xcd, 0x01, 0x4f, 0xaf,
	0xec, 0xd9, 0x6d, 0x27, 0xe2, 0x32, 0xe1, 0xd2, 0xeb, 0x86, 0x12, 0xa7, 0x52, 0x11, 0x67, 0xa9,
	0xe5, 0x1b, 0x9f, 0x09, 0x94, 0xdb, 0xa1, 0x08, 0x13, 0x49, 0x47, 0xb0, 0x95, 0x71, 0x3e, 0x08,
	0x22, 0x81, 0xa1, 0x62, 0x3c, 0x0d, 0x4e, 0x11, 0x6b, 0xa4, 0xbe, 0xdc, 0x5c, 0x7d, 0x7c, 0xcf,
	0xb5, 0x32, 0xae, 0x96, 0x99, 0x54, 0x74, 0x8f, 0x38, 0x4b, 0x0f, 0xf7, 0x2f, 0xaf, 0x77, 0x0b,
	0x1f, 0x7f, 0xec, 0x36, 0x7b, 0x4c, 0xbd, 0x1b, 0x76, 0xdd, 0x88, 0x27, 0x5e, 0x5e, 0xd3, 0x7e,
	0xf6, 0x64, 0xdc, 0xf7, 0xd4, 0x45, 0x86, 0xd2, 0x5c, 0x90, 0xfe, 0xa6, 0xae, 0x72, 0x94, 0x17,
	0x79, 0x8e, 0x48, 0xf7, 0xa1, 0x1a, 0xe3, 0x69, 0x38, 0x1c, 0xa8, 0x40, 0xb1, 0xa8, 0x1f, 0xc8,
	0x2c, 0x8c, 0x58, 0xda, 0xab, 0x2d, 0xd5, 0x49, 0x73, 0xdd, 0xa7, 0x39, 0xf7, 0x92, 0x45, 0xfd,
	0x8e, 0x65, 0x1a, 0x9f, 0x08, 0x14, 0xdb, 0x9c, 0x0f, 0xe8, 0x06, 0x2c, 0xb1, 0xb8, 0x46, 0xea,
	0xa4, 0x59, 0xf4, 0x97, 0x58, 0x4c, 0xef, 0x43, 0x25, 0x09, 0x45, 0x1f, 0x55, 0xc0, 0x62, 0x73,
	0xbf, 0xe8, 0xaf, 0x58, 0xe0, 0x38, 0xa6, 0x77, 0xa0, 0x1c, 0x63, 0xca, 0x93, 0xfd, 0xda, 0x72,
	0x9d, 0x34, 0x2b, 0x7e, 0xbe, 0x9b, 0xe2, 0xad, 0x5a, 0x71, 0x0e, 0x6f, 0xd1, 0x07, 0xb0, 0xf6,
	0x57, 0x3f, 0x25, 0xd3, 0xcf, 0xaa, 0x9a, 0x35, 0x42, 0x1f, 0xc1, 0xa6, 0x40, 0x89, 0xe2, 0x1c,
	0x83, 0x30, 0x8e, 0x05, 0x4a, 0x59, 0x2b, 0x1b, 0x8d, 0x8d, 0x1c, 0x7e, 0x6a, 0xd1, 0xc6, 0x6f,
	0x02, 0x15, 0xdd, 0x71, 0x47, 0x85, 0x0a, 0xb5, 0x72, 0x34, 0x14, 0x02, 0x53, 0x3b, 0xb1, 0x19,
	0xa0, 0xe4, 0xaf, 0xe6, 0x98, 0x9e, 0x94, 0xbe, 0x05, 0x3a, 0x39, 0x22, 0xcf, 0x84, 0x0a, 0x32,
	0xc1, 0x22, 0x34, 0x23, 0x55, 0x0e, 0x5d, 0xed, 0xf9, 0xf7, 0xeb, 0xdd, 0x87, 0xff, 0xe1, 0xf9,
	0x33, 0x8c, 0xfc, 0xdb, 0xb9, 0x52, 0xe7, 0x4c, 0xa8, 0xb6, 0xd6, 0xa1, 0x6f, 0x60, 0x6b, 0xa2,
	0x3e, 0x60, 0x67, 0x43, 0x16, 0x33, 0x75, 0x61, 0x5d, 0x59, 0x58, 0xfc, 0x64, 0xa2, 0xd3, 0xf8,
	0x4a, 0x60, 0xa5, 0xcd, 0x25, 0xd3, 0xf9, 0xfe, 0x93, 0xd0, 0x5d, 0xb8, 0x65, 0xfe, 0xb2, 0x69,
	0x3e, 0x65, 0xbd, 0x3d, 0x8e, 0x69, 0x15, 0x4a, 0x7c, 0x94, 0xa2, 0xc8, 0xc3, 0xb1, 0x1b, 0xba,
	0x03, 0x30, 0xe0, 0x23, 0x14, 0xd6, 0xa7, 0xa2, 0xf1, 0xa9, 0x62, 0x10, 0xe3, 0xd2, 0x0e, 0xc0,
	0x30, 0xcb, 0x26, 0x74, 0xc9, 0xd2, 0x06, 0x31, 0xf4, 0x09, 0x54, 0x66, 0xe3, 0x95, 0x17, 0x1a,
	0x6f, 0x26, 0xd0, 0xf8, 0x42, 0x60, 0x45, 0xcb, 0x1e, 0xa7, 0xa7, 0x9c, 0xbe, 0x82, 0xcd, 0x9e,
	0xe0, 0x52, 0xce, 0xf9, 0x47, 0x16, 0x2a, 0xb0, 0x61, 0x64, 0xa6, 0xee, 0xd1, 0x0e, 0xac, 0xa7,
	0x38, 0x1f, 0xcb, 0x62, 0x99, 0xaf, 0xa5, 0x38, 0x8b, 0xe4, 0xf0, 0xc5, 0xe5, 0x2f, 0xa7, 0x70,
	0x79, 0xe3, 0x90, 0xab, 0x1b, 0x87, 0xfc, 0xbc, 0x71, 0xc8, 0x87, 0xb1, 0x53, 0xb8, 0x1a, 0x3b,
	0x85, 0x6f, 0x63, 0xa7, 0xf0, 0xfa, 0x60, 0x5e, 0x33, 0x7f, 0x5b, 0xf6, 0x52, 0x54, 0x23, 0x2e,
	0xfa, 0x53, 0xc0, 0x3b, 0x7f, 0xe2, 0xbd, 0x37, 0x2f, 0x92, 0x29, 0xd2, 0x2d, 0x9b, 0x07, 0xe4,
	0xe0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x58, 0x83, 0x67, 0xae, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultTickSpacing != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.DefaultTickSpacing))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAmm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReserveAddress) > 0 {
		i -= len(m.ReserveAddress)
		copy(dAtA[i:], m.ReserveAddress)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.ReserveAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.TickSpacing != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Denom1) > 0 {
		i -= len(m.Denom1)
		copy(dAtA[i:], m.Denom1)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.Denom1)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom0) > 0 {
		i -= len(m.Denom0)
		copy(dAtA[i:], m.Denom0)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.Denom0)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MarketId != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CurrentLiquidity.Size()
		i -= size
		if _, err := m.CurrentLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAmm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAmm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.CurrentTick != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.CurrentTick))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAmm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.UpperTick != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.UpperTick))
		i--
		dAtA[i] = 0x28
	}
	if m.LowerTick != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.LowerTick))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAmm(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PoolId != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAmm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TickInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NetLiquidity.Size()
		i -= size
		if _, err := m.NetLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAmm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.GrossLiquidity.Size()
		i -= size
		if _, err := m.GrossLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAmm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAmm(dAtA []byte, offset int, v uint64) int {
	offset -= sovAmm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovAmm(uint64(l))
		}
	}
	if m.DefaultTickSpacing != 0 {
		n += 1 + sovAmm(uint64(m.DefaultTickSpacing))
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAmm(uint64(m.Id))
	}
	if m.MarketId != 0 {
		n += 1 + sovAmm(uint64(m.MarketId))
	}
	l = len(m.Denom0)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	l = len(m.Denom1)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	if m.TickSpacing != 0 {
		n += 1 + sovAmm(uint64(m.TickSpacing))
	}
	l = len(m.ReserveAddress)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	return n
}

func (m *PoolState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentTick != 0 {
		n += 1 + sovAmm(uint64(m.CurrentTick))
	}
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovAmm(uint64(l))
	l = m.CurrentLiquidity.Size()
	n += 1 + l + sovAmm(uint64(l))
	return n
}

func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAmm(uint64(m.Id))
	}
	if m.PoolId != 0 {
		n += 1 + sovAmm(uint64(m.PoolId))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAmm(uint64(l))
	}
	if m.LowerTick != 0 {
		n += 1 + sovAmm(uint64(m.LowerTick))
	}
	if m.UpperTick != 0 {
		n += 1 + sovAmm(uint64(m.UpperTick))
	}
	l = m.Liquidity.Size()
	n += 1 + l + sovAmm(uint64(l))
	return n
}

func (m *TickInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GrossLiquidity.Size()
	n += 1 + l + sovAmm(uint64(l))
	l = m.NetLiquidity.Size()
	n += 1 + l + sovAmm(uint64(l))
	return n
}

func sovAmm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAmm(x uint64) (n int) {
	return sovAmm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultTickSpacing", wireType)
			}
			m.DefaultTickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultTickSpacing |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func (m *PoolState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
			return fmt.Errorf("proto: PoolState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			m.CurrentTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentTick |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerTick", wireType)
			}
			m.LowerTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LowerTick |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperTick", wireType)
			}
			m.UpperTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpperTick |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func (m *TickInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAmm
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
			return fmt.Errorf("proto: TickInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrossLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GrossLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAmm
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
				return ErrInvalidLengthAmm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAmm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAmm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAmm
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
func skipAmm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAmm
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
					return 0, ErrIntOverflowAmm
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
					return 0, ErrIntOverflowAmm
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
				return 0, ErrInvalidLengthAmm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAmm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAmm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAmm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAmm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAmm = fmt.Errorf("proto: unexpected end of group")
)
