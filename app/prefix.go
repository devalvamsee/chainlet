package app

import (
	cmdcfg "github.com/devalvamsee/chainlet/cmd/chainletd/config"
	ethcfg "github.com/evmos/ethermint/cmd/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetConfig() {
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	ethcfg.SetBip44CoinType(config)
	// Make sure address is compatible with ethereum
	config.SetAddressVerifier(VerifyAddressFormat)
	config.Seal()
}
