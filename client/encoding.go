package client

import (
	osmosisOldGammTypes "github.com/DefiantLabs/lens/extra-codecs/osmosis/gamm/types"
	osmosisOldLockupTypes "github.com/DefiantLabs/lens/extra-codecs/osmosis/lockup/types"
	osmosisGammTypes "github.com/DefiantLabs/lens/osmosis/x/gamm/pool-models/balancer"
	tendermintLiquidityTypes "github.com/DefiantLabs/lens/tendermint/x/liquidity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type Codec struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeCodec(moduleBasics []module.AppModuleBasic) Codec {
	modBasic := module.NewBasicManager(moduleBasics...)
	encodingConfig := MakeCodecConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//Register older types that some clients may need to parse, but that are no longer in the Osmosis SDK
	osmosisOldGammTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	osmosisOldGammTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// osmosisOldLockupTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	osmosisOldLockupTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//Register
	tendermintLiquidityTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//osmosisGammTypes
	osmosisGammTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

func MakeCodecConfig() Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return Codec{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}
