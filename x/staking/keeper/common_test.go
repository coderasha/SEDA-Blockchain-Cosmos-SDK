package keeper_test

import (
	"math/big"
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/sedaapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	PKs = sedaapp.CreateTestPubKeys(500)
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

// createTestInput Returns a sedaapp with custom StakingKeeper
// to avoid messing with the hooks.
func createTestInput() (*codec.LegacyAmino, *sedaapp.SedaApp, sdk.Context) {
	app := sedaapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.StakingKeeper = keeper.NewKeeper(
		app.AppCodec(),
		app.GetKey(types.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(types.ModuleName),
	)
	return app.LegacyAmino(), app, ctx
}

// intended to be used with require/assert:  require.True(ValEq(...))
func ValEq(t *testing.T, exp, got types.Validator) (*testing.T, bool, string, types.Validator, types.Validator) {
	return t, exp.MinEqual(&got), "expected:\n%v\ngot:\n%v", exp, got
}

// generateAddresses generates numAddrs of normal AccAddrs and ValAddrs
func generateAddresses(app *sedaapp.SedaApp, ctx sdk.Context, numAddrs int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrDels := sedaapp.AddTestAddrsIncremental(app, ctx, numAddrs, sdk.NewInt(10000))
	addrVals := sedaapp.ConvertAddrsToValAddrs(addrDels)

	return addrDels, addrVals
}
