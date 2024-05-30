package client

import (
	"time"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"

	// distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"

	// paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"

	// upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	ibcInterchainAccounts "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcClient "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	osmosisConcentratedLiquidity "github.com/osmosis-labs/osmosis/v25/x/concentrated-liquidity/clmodule"
	osmosisCosmWasmPool "github.com/osmosis-labs/osmosis/v25/x/cosmwasmpool/module"
	osmosisGamm "github.com/osmosis-labs/osmosis/v25/x/gamm"
	osmosisIncentives "github.com/osmosis-labs/osmosis/v25/x/incentives"
	osmosisLockup "github.com/osmosis-labs/osmosis/v25/x/lockup"
	osmosisMint "github.com/osmosis-labs/osmosis/v25/x/mint"
	osmosisPoolIncentives "github.com/osmosis-labs/osmosis/v25/x/pool-incentives"
	osmosisPoolManager "github.com/osmosis-labs/osmosis/v25/x/poolmanager/module"
	osmosisProtorev "github.com/osmosis-labs/osmosis/v25/x/protorev"
	osmosisSuperfluid "github.com/osmosis-labs/osmosis/v25/x/superfluid"
	osmosisTokenFactory "github.com/osmosis-labs/osmosis/v25/x/tokenfactory"
	osmosisTxFees "github.com/osmosis-labs/osmosis/v25/x/txfees"
	osmosisValsetPref "github.com/osmosis-labs/osmosis/v25/x/valset-pref/valpref-module"
	osmosisEpochs "github.com/osmosis-labs/osmosis/x/epochs"
	skipAuction "github.com/skip-mev/block-sdk/x/auction"
)

var (
	ModuleBasics = []module.AppModuleBasic{
		auth.AppModuleBasic{},
		authz.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		// TODO: add osmosis governance proposal types here
		// TODO: add other proposal types here
		gov.AppModuleBasic{},
		// gov.NewAppModuleBasic(
		// paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		// ),
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
		ibcClient.AppModuleBasic{},
		ibcInterchainAccounts.AppModuleBasic{},
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
		osmosisPoolManager.AppModuleBasic{},
		osmosisProtorev.AppModuleBasic{},
		osmosisConcentratedLiquidity.AppModuleBasic{},
		vesting.AppModuleBasic{},
		osmosisValsetPref.AppModuleBasic{},
		osmosisCosmWasmPool.AppModuleBasic{},
		skipAuction.AppModuleBasic{},
	}
)

type ChainClientConfig struct {
	Key            string                  `json:"key" yaml:"key"`
	ChainID        string                  `json:"chain-id" yaml:"chain-id"`
	RPCAddr        string                  `json:"rpc-addr" yaml:"rpc-addr"`
	GRPCAddr       string                  `json:"grpc-addr" yaml:"grpc-addr"`
	AccountPrefix  string                  `json:"account-prefix" yaml:"account-prefix"`
	KeyringBackend string                  `json:"keyring-backend" yaml:"keyring-backend"`
	GasAdjustment  float64                 `json:"gas-adjustment" yaml:"gas-adjustment"`
	GasPrices      string                  `json:"gas-prices" yaml:"gas-prices"`
	KeyDirectory   string                  `json:"key-directory" yaml:"key-directory"`
	Debug          bool                    `json:"debug" yaml:"debug"`
	Timeout        string                  `json:"timeout" yaml:"timeout"`
	OutputFormat   string                  `json:"output-format" yaml:"output-format"`
	SignModeStr    string                  `json:"sign-mode" yaml:"sign-mode"`
	Modules        []module.AppModuleBasic `json:"-" yaml:"-"`
}

func (ccc *ChainClientConfig) Validate() error {
	if _, err := time.ParseDuration(ccc.Timeout); err != nil {
		return err
	}
	return nil
}

func GetCosmosHubConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "default",
		ChainID:        "cosmoshub-4",
		RPCAddr:        "https://cosmoshub-4.technofractal.com:443",
		GRPCAddr:       "https://gprc.cosmoshub-4.technofractal.com:443",
		AccountPrefix:  "cosmos",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "0.01uatom",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetOsmosisConfig(keyHome string, debug bool) *ChainClientConfig {
	return &ChainClientConfig{
		Key:            "default",
		ChainID:        "osmosis-1",
		RPCAddr:        "https://osmosis-1.technofractal.com:443",
		GRPCAddr:       "https://gprc.osmosis-1.technofractal.com:443",
		AccountPrefix:  "osmo",
		KeyringBackend: "test",
		GasAdjustment:  1.2,
		GasPrices:      "0.01uosmo",
		KeyDirectory:   keyHome,
		Debug:          debug,
		Timeout:        "20s",
		OutputFormat:   "json",
		SignModeStr:    "direct",
	}
}

func GetTestClient() *ChainClient {
	homepath := "/tmp"
	cl, _ := NewChainClient(GetCosmosHubConfig(homepath, true), homepath, nil, nil)
	return cl
}
