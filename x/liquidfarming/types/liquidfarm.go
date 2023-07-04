package types

import (
	"fmt"
	"regexp"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	utils "github.com/crescent-network/crescent/v5/types"
)

const (
	ShareDenomPrefix = "lfshare"
)

var (
	shareDenomRe = regexp.MustCompile(`^lfshare([1-9]\d*)$`)
)

// DeriveBidReserveAddress creates the reserve address for bids
// with the given liquid farm id.
func DeriveBidReserveAddress(liquidFarmId uint64) sdk.AccAddress {
	return address.Module(ModuleName, []byte(fmt.Sprintf("BidReserveAddress/%d", liquidFarmId)))
}

// NewLiquidFarm returns a new LiquidFarm.
func NewLiquidFarm(
	liquidFarmId, poolId uint64, lowerTick, upperTick int32, minBidAmt sdk.Int, feeRate sdk.Dec) LiquidFarm {
	return LiquidFarm{
		Id:                   liquidFarmId,
		PoolId:               poolId,
		LowerTick:            lowerTick,
		UpperTick:            upperTick,
		BidReserveAddress:    DeriveBidReserveAddress(liquidFarmId).String(),
		MinBidAmount:         minBidAmt,
		FeeRate:              feeRate,
		LastRewardsAuctionId: 0,
	}
}

// Validate validates LiquidFarm.
func (liquidFarm LiquidFarm) Validate() error {
	if liquidFarm.Id == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if liquidFarm.PoolId == 0 {
		return fmt.Errorf("pool id must not be 0")
	}
	if liquidFarm.LowerTick >= liquidFarm.UpperTick {
		return fmt.Errorf("lower tick must be lower than upper tick")
	}
	if _, err := sdk.AccAddressFromBech32(liquidFarm.BidReserveAddress); err != nil {
		return fmt.Errorf("invalid bid reserve address %w", err)
	}
	if liquidFarm.MinBidAmount.IsNegative() {
		return fmt.Errorf("minimum bid amount must not be negative: %s", liquidFarm.MinBidAmount)
	}
	if liquidFarm.FeeRate.IsNegative() {
		return fmt.Errorf("fee rate must not be negative: %s", liquidFarm.FeeRate)
	}
	if liquidFarm.FeeRate.GT(utils.OneDec) {
		return fmt.Errorf("fee rate must not be greater than 1: %s", liquidFarm.FeeRate)
	}
	return nil
}

// ShareDenom returns a unique liquid farm share denom.
func ShareDenom(liquidFarmId uint64) string {
	return fmt.Sprintf("%s%d", ShareDenomPrefix, liquidFarmId)
}

// ParseShareDenom parses a liquid farm share denom and returns the liquid farm's id.
func ParseShareDenom(denom string) (liquidFarmId uint64, err error) {
	chunks := shareDenomRe.FindStringSubmatch(denom)
	if len(chunks) == 0 {
		return 0, fmt.Errorf("invalid share denom: %s", denom)
	}
	liquidFarmId, err = strconv.ParseUint(chunks[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid liquid farm id: %s", chunks[1])
	}
	return liquidFarmId, nil
}

func CalculateMintRate(totalLiquidity, shareSupply sdk.Int) sdk.Dec {
	return shareSupply.ToDec().QuoTruncate(totalLiquidity.ToDec())
}

// CalculateMintedShareAmount calculates minting liquid farm share amount.
// mintedShareAmt = shareSupply * (addedLiquidity / totalLiquidity)
func CalculateMintedShareAmount(addedLiquidity, totalLiquidity, shareSupply sdk.Int) sdk.Int {
	if shareSupply.IsZero() { // initial minting
		return addedLiquidity
	}
	return shareSupply.Mul(addedLiquidity).Quo(totalLiquidity)
}

func CalculateBurnRate(shareSupply ,totalLiquidity, prevWinningBidShareAmt sdk.Int) sdk.Dec {
	return totalLiquidity.ToDec().QuoTruncate(shareSupply.Add(prevWinningBidShareAmt).ToDec())
}

// CalculateRemovedLiquidity calculates liquidity amount to be removed when
// burning liquid farm share.
// removedLiquidity = totalLiquidity * (burnedShareAmt / (shareSupply + prevWinningBidShareAmt))
func CalculateRemovedLiquidity(
	burnedShareAmt, shareSupply, totalLiquidity, prevWinningBidShareAmt sdk.Int) sdk.Int {
	if burnedShareAmt.Equal(shareSupply) { // last one to unfarm
		return totalLiquidity
	}
	return totalLiquidity.Mul(burnedShareAmt).Quo(shareSupply.Add(prevWinningBidShareAmt))
}

// DeductFees deducts fees from rewards by the fee rate.
func DeductFees(rewards sdk.Coins, feeRate sdk.Dec) (deductedRewards sdk.Coins, fees sdk.Coins) {
	deductedRewards, _ = sdk.NewDecCoinsFromCoins(rewards...).MulDecTruncate(utils.OneDec.Sub(feeRate)).TruncateDecimal()
	fees = rewards.Sub(deductedRewards)
	return
}
