package types

import (
	"context"
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

	bankexported "cosmossdk.io/x/bank/exported"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper defines the expected staking keeper (noalias)
type StakingKeeper interface {
	ApplyAndReturnValidatorSetUpdates(context.Context) (updates []abci.ValidatorUpdate, err error)
}

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	NewAccount(context.Context, sdk.AccountI) sdk.AccountI
	SetAccount(context.Context, sdk.AccountI)
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object (noalias)
type GenesisAccountsIterator interface {
	IterateGenesisAccounts(
		cdc *codec.LegacyAmino,
		appGenesis map[string]json.RawMessage,
		cb func(sdk.AccountI) (stop bool),
	)
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object (noalias)
type GenesisBalancesIterator interface {
	IterateGenesisBalances(
		cdc codec.JSONCodec,
		appGenesis map[string]json.RawMessage,
		cb func(bankexported.GenesisBalance) (stop bool),
	)
}
