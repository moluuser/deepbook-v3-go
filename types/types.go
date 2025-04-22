package types

import "github.com/block-vision/sui-go-sdk/transaction"

type Environment string

const (
	MAINNET Environment = "mainnet"
	TESTNET Environment = "testnet"
)

type BalanceManager struct {
	Address  string
	TradeCap *string
}

type Coin struct {
	Address string
	Type    string
	Scalar  uint64
}

type Pool struct {
	Address   string
	BaseCoin  string
	QuoteCoin string
}

type OrderType string

const (
	NO_RESTRICTION      OrderType = "NoRestriction"
	IMMEDIATE_OR_CANCEL OrderType = "ImmediateOrCancel"
	FILL_OR_KILL        OrderType = "FillOrKill"
	POST_ONLY           OrderType = "PostOnly"
)

type SelfMatchingOptions string

const (
	SELF_MATCHING_ALLOWED SelfMatchingOptions = "SelfMatchingAllowed"
	CANCEL_TAKER          SelfMatchingOptions = "CancelTaker"
	CANCEL_MAKER          SelfMatchingOptions = "CancelMaker"
)

type PlaceLimitOrderParams struct {
	PoolKey            string
	BalanceManagerKey  string
	ClientOrderId      string
	Price              float64
	Quantity           float64
	IsBid              bool
	Expiration         *float64
	OrderType          *OrderType
	SelfMatchingOption *SelfMatchingOptions
	PayWithDeep        *bool
}

type PlaceMarketOrderParams struct {
	PoolKey            string
	BalanceManagerKey  string
	ClientOrderId      string
	Quantity           float64
	IsBid              bool
	SelfMatchingOption *SelfMatchingOptions
	PayWithDeep        *bool
}

type ProposalParams struct {
	PoolKey           string
	BalanceManagerKey string
	TakerFee          float64
	MakerFee          float64
	StakeRequired     float64
}

type SwapParams struct {
	PoolKey    string
	Amount     float64
	DeepAmount float64
	MinOut     float64
	DeepCoin   *transaction.Argument
	BaseCoin   *transaction.Argument
	QuoteCoin  *transaction.Argument
}

type CreatePoolAdminParams struct {
	BaseCoinKey  string
	QuoteCoinKey string
	TickSize     float64
	LotSize      float64
	MinSize      float64
	Whitelisted  bool
	StablePool   bool
	DeepCoin     *transaction.Argument
	BaseCoin     *transaction.Argument
}
