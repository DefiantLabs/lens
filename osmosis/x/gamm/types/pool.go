package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

// PoolI defines an interface for pools that hold tokens.
type PoolI interface {
	proto.Message

	GetAddress() sdk.AccAddress
	String() string

	GetId() uint64
	GetPoolSwapFee() sdk.Dec
	GetPoolExitFee() sdk.Dec
	GetTotalWeight() sdk.Int
	GetTotalShares() sdk.Coin
	AddTotalShares(amt sdk.Int)
	SubTotalShares(amt sdk.Int)
	GetPoolAsset(denom string) (PoolAsset, error)
	// UpdatePoolAssetBalance updates the balances for
	// the token with denomination coin.denom
	UpdatePoolAssetBalance(coin sdk.Coin) error
	// UpdatePoolAssetBalances calls UpdatePoolAssetBalance
	// on each constituent coin.
	UpdatePoolAssetBalances(coins sdk.Coins) error
	GetPoolAssets(denoms ...string) ([]PoolAsset, error)
	GetAllPoolAssets() []PoolAsset
	PokeTokenWeights(blockTime time.Time)
	GetTokenWeight(denom string) (sdk.Int, error)
	GetTokenBalance(denom string) (sdk.Int, error)
	NumAssets() int
	IsActive(curBlockTime time.Time) bool
}

var (
	MaxUserSpecifiedWeight    sdk.Int = sdk.NewIntFromUint64(1 << 20)
	GuaranteedWeightPrecision int64   = 1 << 30
)
