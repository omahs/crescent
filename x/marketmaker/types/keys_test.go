package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"

	"github.com/crescent-network/crescent/v2/x/marketmaker/types"
)

var (
	addr1 = sdk.AccAddress(crypto.AddressHash([]byte("mm1")))
	addr2 = sdk.AccAddress(crypto.AddressHash([]byte("mm2")))
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetMarketMakerKey() {
	s.Require().Equal([]byte{0xc0, 0x14, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetMarketMakerKey(addr1, 1))
	s.Require().Equal([]byte{0xc0, 0x14, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2}, types.GetMarketMakerKey(addr1, 2))
	s.Require().Equal([]byte{0xc0, 0x14, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetMarketMakerKey(addr2, 1))
	s.Require().Equal([]byte{0xc0, 0x14, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2}, types.GetMarketMakerKey(addr2, 2))
}

func (s *keysTestSuite) TestGetMarketMakerByAddrPrefix() {
	s.Require().Equal([]byte{0xc0, 0x14, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa}, types.GetMarketMakerByAddrPrefix(addr1))
	s.Require().Equal([]byte{0xc0, 0x14, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5}, types.GetMarketMakerByAddrPrefix(addr2))
}

func (s *keysTestSuite) TestGetMarketMakerByPairIdPrefix() {
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetMarketMakerByPairIdPrefix(1))
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2}, types.GetMarketMakerByPairIdPrefix(2))
}

func (s *keysTestSuite) TestGetMarketMakerIndexByPairIdKey() {
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa}, types.GetMarketMakerIndexByPairIdKey(1, addr1))
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa}, types.GetMarketMakerIndexByPairIdKey(2, addr1))
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5}, types.GetMarketMakerIndexByPairIdKey(1, addr2))
	s.Require().Equal([]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5}, types.GetMarketMakerIndexByPairIdKey(2, addr2))
}

func (s *keysTestSuite) TestGetMarketMakerIndexByPairIdKey2() {
	testCases := []struct {
		mmAddr   sdk.AccAddress
		pairId   uint64
		expected []byte
	}{
		{
			addr1,
			1,
			[]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa},
		},
		{
			addr1,
			2,
			[]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa},
		},
		{
			addr2,
			1,
			[]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5},
		},
		{
			addr2,
			2,
			[]byte{0xc1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5},
		},
	}

	for _, tc := range testCases {
		key := types.GetMarketMakerIndexByPairIdKey(tc.pairId, tc.mmAddr)
		s.Require().Equal(tc.expected, key)

		pairId, mmAddr := types.ParseMarketMakerIndexByPairIdKey(key)
		s.Require().Equal(tc.pairId, pairId)
		s.Require().Equal(tc.mmAddr, mmAddr)
	}
}

func (s *keysTestSuite) TestGetIncentive() {
	s.Require().Equal([]byte{0xc5, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa}, types.GetIncentiveKey(addr1))
	s.Require().Equal([]byte{0xc5, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5}, types.GetIncentiveKey(addr2))
}

func (s *keysTestSuite) TestGetDepositKey() {
	testCases := []struct {
		mmAddr   sdk.AccAddress
		pairId   uint64
		expected []byte
	}{
		{
			addr1,
			1,
			[]byte{0xc2, 0x14, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			addr1,
			2,
			[]byte{0xc2, 0x14, 0xa6, 0xad, 0xc1, 0x49, 0x46, 0xc1, 0x1b, 0x25, 0x83, 0x23, 0xb4, 0x32, 0x29, 0x8e, 0xe2, 0xd8, 0x95, 0x3b, 0xee, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2},
		},
		{
			addr2,
			1,
			[]byte{0xc2, 0x14, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			addr2,
			2,
			[]byte{0xc2, 0x14, 0x4, 0xfd, 0x48, 0x29, 0x95, 0x87, 0x78, 0x78, 0x96, 0xa2, 0xa4, 0x33, 0xc1, 0xb3, 0x9f, 0xae, 0x90, 0x87, 0xce, 0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2},
		},
	}

	for _, tc := range testCases {
		key := types.GetDepositKey(tc.mmAddr, tc.pairId)
		s.Require().Equal(tc.expected, key)

		mmAddr, pairId := types.ParseDepositKey(key)
		s.Require().Equal(tc.pairId, pairId)
		s.Require().Equal(tc.mmAddr, mmAddr)
	}
}
