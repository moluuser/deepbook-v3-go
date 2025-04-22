package utils

import (
	"deepbook-v3-go/types"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/samber/lo"
)

const (
	FLOAT_SCALAR  = 1000000000
	MAX_TIMESTAMP = 1844674407370955161
	GAS_BUDGET    = 0.5 * 500000000 // Adjust based on benchmarking
	DEEP_SCALAR   = 1000000
)

type BalanceManagers map[string]types.BalanceManager

type DeepBookConfig struct {
	coins CoinMap
	pools PoolMap

	BalanceManagers BalanceManagers
	Address         string
	AdminCap        *string
	DeepbookPackageIds
}

func NewDeepBookConfig(
	env types.Environment,
	address string,
	adminCap *string,
	balanceManagers *BalanceManagers,
	coins *CoinMap,
	pools *PoolMap,
) *DeepBookConfig {
	bm := BalanceManagers{}
	if balanceManagers != nil {
		bm = *balanceManagers
	}

	config := &DeepBookConfig{
		Address:         string(utils.NormalizeSuiAddress(address)),
		AdminCap:        adminCap,
		BalanceManagers: bm,
	}

	if env == types.MAINNET {
		config.coins = *lo.Ternary(coins != nil, coins, &MainnetCoins)
		config.pools = *lo.Ternary(pools != nil, pools, &MainnetPools)
		config.DeepbookPackageIds = MainnetPackageIds
	} else {
		config.coins = *lo.Ternary(coins != nil, coins, &TestnetCoins)
		config.pools = *lo.Ternary(pools != nil, pools, &TestnetPools)
		config.DeepbookPackageIds = TestnetPackageIds
	}

	return config
}

func (config *DeepBookConfig) GetCoin(coin CoinNameType) *types.Coin {
	if coin, ok := config.coins[coin]; ok {
		return &coin
	} else {
		return nil
	}
}

func (config *DeepBookConfig) GetPool(pool PoolNameType) *types.Pool {
	if pool, ok := config.pools[pool]; ok {
		return &pool
	} else {
		return nil
	}
}

func (config *DeepBookConfig) GetBalanceManager(name string) *types.BalanceManager {
	if bm, ok := config.BalanceManagers[name]; ok {
		return &bm
	} else {
		return nil
	}
}
