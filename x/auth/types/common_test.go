package types_test

import (
	"github.com/cosmos/cosmos-sdk/sedaapp"
)

var (
	app                   = sedaapp.Setup(false)
	ecdc                  = sedaapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Marshaler, ecdc.Amino
)
