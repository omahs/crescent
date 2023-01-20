// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/bootstrap/v1beta1/proposal.proto

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

type BootstrapProposal struct {
	// title specifies the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description specifies the description of the proposal
	Description    string                                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ReturnAddress  string                                  `protobuf:"bytes,3,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
	OfferCoin      types.Coin                              `protobuf:"bytes,4,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
	QuoteCoinDenom string                                  `protobuf:"bytes,5,opt,name=quote_coin_denom,json=quoteCoinDenom,proto3" json:"quote_coin_denom,omitempty"`
	MinPrice       *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=min_price,json=minPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_price,omitempty"`
	MaxPrice       *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=max_price,json=maxPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_price,omitempty"`
	PairId         uint64                                  `protobuf:"varint,8,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	PoolId         uint64                                  `protobuf:"varint,9,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	InitialOrders  []*InitialOrder                         `protobuf:"bytes,10,rep,name=initial_orders,json=initialOrders,proto3" json:"initial_orders,omitempty"`
}

func (m *BootstrapProposal) Reset()      { *m = BootstrapProposal{} }
func (*BootstrapProposal) ProtoMessage() {}
func (*BootstrapProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91229c2c9a1c550, []int{0}
}
func (m *BootstrapProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BootstrapProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BootstrapProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapProposal.Merge(m, src)
}
func (m *BootstrapProposal) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapProposal.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapProposal proto.InternalMessageInfo

type InitialOrder struct {
	OfferCoin types.Coin                             `protobuf:"bytes,1,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
	Price     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
}

func (m *InitialOrder) Reset()         { *m = InitialOrder{} }
func (m *InitialOrder) String() string { return proto.CompactTextString(m) }
func (*InitialOrder) ProtoMessage()    {}
func (*InitialOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91229c2c9a1c550, []int{1}
}
func (m *InitialOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitialOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitialOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitialOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialOrder.Merge(m, src)
}
func (m *InitialOrder) XXX_Size() int {
	return m.Size()
}
func (m *InitialOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialOrder.DiscardUnknown(m)
}

var xxx_messageInfo_InitialOrder proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BootstrapProposal)(nil), "crescent.bootstrap.v1beta1.BootstrapProposal")
	proto.RegisterType((*InitialOrder)(nil), "crescent.bootstrap.v1beta1.InitialOrder")
}

func init() {
	proto.RegisterFile("crescent/bootstrap/v1beta1/proposal.proto", fileDescriptor_a91229c2c9a1c550)
}

var fileDescriptor_a91229c2c9a1c550 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0xc9, 0x47, 0x93, 0x0b, 0x8d, 0xc0, 0xaa, 0x84, 0x9b, 0xc1, 0x8e, 0x2a, 0x81, 0x4c,
	0xa5, 0xda, 0x6a, 0x61, 0x81, 0x01, 0x89, 0x10, 0x09, 0x65, 0x6a, 0x15, 0x36, 0x16, 0xcb, 0x1f,
	0x97, 0x70, 0xaa, 0x7d, 0xcf, 0xdc, 0x5d, 0x4a, 0xf8, 0x0f, 0x18, 0x19, 0x99, 0x50, 0xfe, 0x9c,
	0x8e, 0x1d, 0x11, 0x43, 0x85, 0x92, 0x85, 0xff, 0x81, 0x05, 0xdd, 0x47, 0x42, 0x84, 0x04, 0x52,
	0x3b, 0x25, 0xef, 0xf7, 0xf1, 0xde, 0xdd, 0xef, 0x9e, 0xd1, 0xe3, 0x8c, 0x61, 0x9e, 0x61, 0x2a,
	0xa2, 0x14, 0x40, 0x70, 0xc1, 0x92, 0x2a, 0xba, 0x38, 0x4e, 0xb1, 0x48, 0x8e, 0xa3, 0x8a, 0x41,
	0x05, 0x3c, 0x29, 0xc2, 0x8a, 0x81, 0x00, 0xa7, 0xb7, 0x96, 0x86, 0x1b, 0x69, 0x68, 0xa4, 0xbd,
	0xbd, 0x29, 0x4c, 0x41, 0xc9, 0x22, 0xf9, 0x4f, 0x3b, 0x7a, 0xfb, 0x19, 0xf0, 0x12, 0x78, 0xac,
	0x09, 0x5d, 0x18, 0xca, 0xd3, 0x55, 0x94, 0x26, 0x1c, 0x6f, 0x06, 0x66, 0x40, 0xa8, 0xe1, 0x0f,
	0xff, 0x73, 0xae, 0x3f, 0xe3, 0xb5, 0xd6, 0x9f, 0x02, 0x4c, 0x0b, 0x1c, 0xa9, 0x2a, 0x9d, 0x4d,
	0x22, 0x41, 0x4a, 0xcc, 0x45, 0x52, 0x1a, 0xc1, 0xc1, 0xaf, 0x1a, 0xba, 0x3f, 0x58, 0x9b, 0xce,
	0xcc, 0xad, 0x9c, 0x3d, 0xd4, 0x10, 0x44, 0x14, 0xd8, 0xb5, 0xfb, 0x76, 0xd0, 0x1e, 0xeb, 0xc2,
	0xe9, 0xa3, 0x4e, 0x8e, 0x79, 0xc6, 0x48, 0x25, 0x08, 0x50, 0xf7, 0x8e, 0xe2, 0xb6, 0x21, 0xe7,
	0x21, 0xea, 0x32, 0x2c, 0x66, 0x8c, 0xc6, 0x49, 0x9e, 0x33, 0xcc, 0xb9, 0x5b, 0x53, 0xa2, 0x5d,
	0x8d, 0xbe, 0xd4, 0xa0, 0xf3, 0x02, 0x21, 0x98, 0x4c, 0x30, 0x8b, 0xe5, 0xad, 0xdc, 0x7a, 0xdf,
	0x0e, 0x3a, 0x27, 0xfb, 0xa1, 0x09, 0x41, 0x5e, 0x7b, 0x1d, 0x5e, 0xf8, 0x0a, 0x08, 0x1d, 0xd4,
	0x2f, 0xaf, 0x7d, 0x6b, 0xdc, 0x56, 0x16, 0x09, 0x38, 0x01, 0xba, 0xf7, 0x7e, 0x06, 0x02, 0x2b,
	0x7f, 0x9c, 0x63, 0x0a, 0xa5, 0xdb, 0x50, 0x83, 0xba, 0x0a, 0x97, 0xa2, 0xa1, 0x44, 0x9d, 0xd7,
	0xa8, 0x5d, 0x12, 0x1a, 0x57, 0x8c, 0x64, 0xd8, 0x6d, 0x4a, 0xc9, 0xe0, 0xf0, 0xfb, 0xb5, 0xff,
	0x68, 0x4a, 0xc4, 0xbb, 0x59, 0x1a, 0x66, 0x50, 0x9a, 0xec, 0xcd, 0xcf, 0x11, 0xcf, 0xcf, 0x23,
	0xf1, 0xb1, 0xc2, 0x3c, 0x1c, 0xe2, 0x6c, 0xdc, 0x2a, 0x09, 0x3d, 0x93, 0x5e, 0xd5, 0x28, 0x99,
	0x9b, 0x46, 0x3b, 0xb7, 0x68, 0x94, 0xcc, 0x75, 0xa3, 0x07, 0x68, 0xa7, 0x4a, 0x08, 0x8b, 0x49,
	0xee, 0xb6, 0xfa, 0x76, 0x50, 0x1f, 0x37, 0x65, 0x39, 0xca, 0x15, 0x01, 0x50, 0x48, 0xa2, 0x6d,
	0x08, 0x80, 0x62, 0x94, 0x3b, 0xa7, 0xa8, 0x4b, 0x28, 0x11, 0x24, 0x29, 0x62, 0x60, 0x39, 0x66,
	0xdc, 0x45, 0xfd, 0x5a, 0xd0, 0x39, 0x09, 0xc2, 0x7f, 0x6f, 0x5d, 0x38, 0xd2, 0x8e, 0x53, 0x69,
	0x18, 0xef, 0x92, 0xad, 0x8a, 0x3f, 0x6f, 0x7d, 0x5a, 0xf8, 0xd6, 0x97, 0x85, 0x6f, 0x1d, 0x7c,
	0xb5, 0xd1, 0xdd, 0x6d, 0xe5, 0x5f, 0x2f, 0x63, 0xdf, 0xf8, 0x65, 0x86, 0xa8, 0xa1, 0x23, 0x52,
	0xcb, 0x31, 0x08, 0x25, 0x7f, 0x83, 0x98, 0xb4, 0x59, 0x1f, 0xf0, 0xe7, 0xc2, 0xb7, 0x06, 0x6f,
	0x2e, 0x97, 0x9e, 0x7d, 0xb5, 0xf4, 0xec, 0x1f, 0x4b, 0xcf, 0xfe, 0xbc, 0xf2, 0xac, 0xab, 0x95,
	0x67, 0x7d, 0x5b, 0x79, 0xd6, 0xdb, 0x67, 0xdb, 0x2d, 0x4d, 0x0e, 0x47, 0x14, 0x8b, 0x0f, 0xc0,
	0xce, 0x37, 0x40, 0x74, 0xf1, 0x34, 0x9a, 0x6f, 0x7d, 0x26, 0x6a, 0x52, 0xda, 0x54, 0xab, 0xff,
	0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x7d, 0x40, 0x02, 0xe1, 0x03, 0x00, 0x00,
}

func (m *BootstrapProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InitialOrders) > 0 {
		for iNdEx := len(m.InitialOrders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialOrders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x48
	}
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxPrice != nil {
		{
			size := m.MaxPrice.Size()
			i -= size
			if _, err := m.MaxPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.MinPrice != nil {
		{
			size := m.MinPrice.Size()
			i -= size
			if _, err := m.MinPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.QuoteCoinDenom) > 0 {
		i -= len(m.QuoteCoinDenom)
		copy(dAtA[i:], m.QuoteCoinDenom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.QuoteCoinDenom)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ReturnAddress) > 0 {
		i -= len(m.ReturnAddress)
		copy(dAtA[i:], m.ReturnAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ReturnAddress)))
		i--
		dAtA[i] = 0x1a
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

func (m *InitialOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitialOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitialOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *BootstrapProposal) Size() (n int) {
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
	l = len(m.ReturnAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = len(m.QuoteCoinDenom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.MinPrice != nil {
		l = m.MinPrice.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.MaxPrice != nil {
		l = m.MaxPrice.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	if m.PoolId != 0 {
		n += 1 + sovProposal(uint64(m.PoolId))
	}
	if len(m.InitialOrders) > 0 {
		for _, e := range m.InitialOrders {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *InitialOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OfferCoin.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BootstrapProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BootstrapProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnAddress", wireType)
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
			m.ReturnAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
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
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteCoinDenom", wireType)
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
			m.QuoteCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.MinPrice = &v
			if err := m.MinPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.MaxPrice = &v
			if err := m.MaxPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
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
		case 9:
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialOrders", wireType)
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
			m.InitialOrders = append(m.InitialOrders, &InitialOrder{})
			if err := m.InitialOrders[len(m.InitialOrders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *InitialOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: InitialOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitialOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
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
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
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
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
