package transactions

import (
	"deepbook-v3-go/utils"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/transaction"
)

type BalanceManagerContract struct {
	config utils.DeepBookConfig
}

func NewBalanceManagerContract(
	config utils.DeepBookConfig,
) *BalanceManagerContract {
	return &BalanceManagerContract{
		config: config,
	}
}

func (bmc *BalanceManagerContract) CreateAndShareBalanceManager(tx transaction.Transaction) {
	manager := tx.MoveCall(
		models.SuiAddress(bmc.config.DEEPBOOK_PACKAGE_ID),
		"balance_manager",
		"new",
		[]transaction.TypeTag{},
		[]transaction.Argument{},
	)

	pkgIdBytes, err := transaction.ConvertSuiAddressStringToBytes(models.SuiAddress(bmc.config.DEEPBOOK_PACKAGE_ID))
	if err != nil {
		panic(err)
	}

	tx.MoveCall(
		"0x2",
		"transfer",
		"public_share_object",
		[]transaction.TypeTag{
			{
				Struct: &transaction.StructTag{
					Address: *pkgIdBytes,
					Module:  "balance_manager",
					Name:    "BalanceManager",
				},
			},
		},
		[]transaction.Argument{manager},
	)

	return
}
