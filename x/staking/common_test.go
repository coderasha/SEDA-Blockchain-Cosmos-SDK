package staking_test

import (
	"math/big"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/sedaapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

// nolint:deadcode,unused,varcheck
var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())
	priv2 = secp256k1.GenPrivKey()
	addr2 = sdk.AccAddress(priv2.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())

	commissionRates = types.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	PKs = sedaapp.CreateTestPubKeys(500)
)

// getBaseSedaappWithCustomKeeper Returns a sedaapp with custom StakingKeeper
// to avoid messing with the hooks.
func getBaseSedaappWithCustomKeeper() (*codec.LegacyAmino, *sedaapp.SedaApp, sdk.Context) {
	app := sedaapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	appCodec := app.AppCodec()

	app.StakingKeeper = keeper.NewKeeper(
		appCodec,
		app.GetKey(types.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(types.ModuleName),
	)
	app.StakingKeeper.SetParams(ctx, types.DefaultParams())

	return codec.NewLegacyAmino(), app, ctx
}

// generateAddresses generates numAddrs of normal AccAddrs and ValAddrs
func generateAddresses(app *sedaapp.SedaApp, ctx sdk.Context, numAddrs int, accAmount sdk.Int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrDels := sedaapp.AddTestAddrsIncremental(app, ctx, numAddrs, accAmount)
	addrVals := sedaapp.ConvertAddrsToValAddrs(addrDels)

	return addrDels, addrVals
}
