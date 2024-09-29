package client

import (
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
)

type Codec struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeCodec(moduleBasics []module.AppModuleBasic, addressPrefix string, valAddressPrefix string) Codec {
	modBasic := module.NewBasicManager(moduleBasics...)
	encodingConfig := MakeCodecConfig(addressPrefix, valAddressPrefix)
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//Register older types that some clients may need to parse, but that are no longer in the Osmosis SDK
	// osmosisOldGammTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	// osmosisOldGammTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// osmosisOldLockupTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	// osmosisOldLockupTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//Register
	// tendermintLiquidityTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

func MakeCodecConfig(accAddressPrefix string, valAddressPrefix string) Codec {
	interfaceRegistry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          address.NewBech32Codec(accAddressPrefix),
			ValidatorAddressCodec: address.NewBech32Codec(valAddressPrefix),
		},
	})

	if err != nil {
		panic(err)
	}

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return Codec{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}
