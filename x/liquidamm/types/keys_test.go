package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	utils "github.com/crescent-network/crescent/v5/types"
	"github.com/crescent-network/crescent/v5/x/liquidamm/types"
)

func TestGetPublicPositionKey(t *testing.T) {
	require.Equal(
		t, []byte{0x83, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetPublicPositionKey(1))
	require.Equal(
		t, []byte{0x83, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4}, types.GetPublicPositionKey(500))
	require.Equal(
		t, []byte{0x83, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10}, types.GetPublicPositionKey(10000))
}

func TestGetRewardsAuctionKey(t *testing.T) {
	require.Equal(
		t, []byte{0x85, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetRewardsAuctionKey(1, 1))
	require.Equal(
		t, []byte{0x85, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4}, types.GetRewardsAuctionKey(1, 500))
	require.Equal(
		t, []byte{0x85, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4}, types.GetRewardsAuctionKey(10000, 500))
}

func TestGetBidKey(t *testing.T) {
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0xc8, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetBidKey(1, 1, utils.TestAddress(100)))
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4,
			0x90, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetBidKey(1, 500, utils.TestAddress(200)))
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4,
			0xd8, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetBidKey(10000, 500, utils.TestAddress(300)))
}

func TestGetBidsByRewardsAuctionIteratorPrefix(t *testing.T) {
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetBidsByRewardsAuctionIteratorPrefix(1, 1))
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4}, types.GetBidsByRewardsAuctionIteratorPrefix(1, 500))
	require.Equal(
		t, []byte{0x86, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xf4}, types.GetBidsByRewardsAuctionIteratorPrefix(10000, 500))
}
