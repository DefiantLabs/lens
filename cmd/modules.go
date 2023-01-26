package cmd

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	osmosisEpochs "github.com/osmosis-labs/osmosis/v9/x/epochs"
	osmosisGamm "github.com/osmosis-labs/osmosis/v9/x/gamm"
	osmosisIncentives "github.com/osmosis-labs/osmosis/v9/x/incentives"
	osmosisLockup "github.com/osmosis-labs/osmosis/v9/x/lockup"
	osmosisMint "github.com/osmosis-labs/osmosis/v9/x/mint"
	osmosisPoolIncentives "github.com/osmosis-labs/osmosis/v9/x/pool-incentives"
	osmosisSuperfluid "github.com/osmosis-labs/osmosis/v9/x/superfluid"
	osmosisTokenFactory "github.com/osmosis-labs/osmosis/v9/x/tokenfactory"
	osmosisTxFees "github.com/osmosis-labs/osmosis/v9/x/txfees"
	tendermintLiquidity "github.com/tendermint/liquidity/x/liquidity"
)

// TODO: Import a bunch of custom modules like cosmwasm and osmosis
// Problem is SDK versioning. Need to find a fix for this.

var ModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	authz.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	// TODO: add osmosis governance proposal types here
	// TODO: add other proposal types here
	gov.NewAppModuleBasic(
		paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
	),
	crisis.AppModuleBasic{},
	distribution.AppModuleBasic{},
	feegrant.AppModuleBasic{},
	mint.AppModuleBasic{},
	params.AppModuleBasic{},
	slashing.AppModuleBasic{},
	staking.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	transfer.AppModuleBasic{},
	ibc.AppModuleBasic{},
	wasm.AppModuleBasic{},
	osmosisGamm.AppModuleBasic{},
	osmosisEpochs.AppModuleBasic{},
	osmosisIncentives.AppModuleBasic{},
	osmosisLockup.AppModuleBasic{},
	osmosisMint.AppModuleBasic{},
	osmosisPoolIncentives.AppModuleBasic{},
	osmosisSuperfluid.AppModuleBasic{},
	osmosisTokenFactory.AppModuleBasic{},
	osmosisTxFees.AppModuleBasic{},
	tendermintLiquidity.AppModuleBasic{},
}
