package query

import (
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/DefiantLabs/lens/client"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	osmosisProtorev "github.com/osmosis-labs/osmosis/v26/x/protorev/types"
	osmosisEpochs "github.com/osmosis-labs/osmosis/x/epochs/types"
)

type Query struct {
	Client  *client.ChainClient
	Options *QueryOptions
}

//Tx queries

// Tx returns the Tx and all contained messages/TxResponse.
func (q *Query) TxByHeight(cc client.Codec) (resp *txTypes.GetTxsEventResponse, unpackError error, queryError error) {
	resp, unpackError, queryError = TxsAtHeightRPC(q, q.Options.Height, cc)
	return
}

// Bank queries

// Balances returns the balance of all coins for a single account.
func (q *Query) Balances(address string) (*bankTypes.QueryAllBalancesResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BalanceWithAddressRPC(q, address)
}

// TotalSupply returns the supply of all coins
func (q *Query) TotalSupply() (*bankTypes.QueryTotalSupplyResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return TotalSupplyRPC(q)
}

// DenomsMetadata returns the metadata for all denoms
func (q *Query) DenomsMetadata() (*bankTypes.QueryDenomsMetadataResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return DenomsMetadataRPC(q)
}

// Staking queries

// Delegation returns the delegations to a particular validator
func (q *Query) Delegation(delegator, validator string) (*stakingTypes.DelegationResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return DelegationRPC(q, delegator, validator)
}

// Delegations returns all the delegations
func (q *Query) Delegations(delegator string) (*stakingTypes.QueryDelegatorDelegationsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return DelegationsRPC(q, delegator)
}

// ValidatorDelegations returns all the delegations for a validator
func (q *Query) ValidatorDelegations(validator string) (*stakingTypes.QueryValidatorDelegationsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return ValidatorDelegationssRPC(q, validator)
}

// Distribution queries

// DelegatorValidators returns the validators of a delegator
func (q *Query) DelegatorValidators(delegator string) (*distributionTypes.QueryDelegatorValidatorsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return DelegatorValidatorsRPC(q, delegator)
}

// Tendermint queries

// Block returns information about a block
func (q *Query) Block() (*coretypes.ResultBlock, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BlockRPC(q)
}

// BlockByHash returns information about a block by hash
func (q *Query) BlockByHash(hash string) (*coretypes.ResultBlock, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BlockByHashRPC(q, hash)
}

// BlockResults returns information about a block by hash
func (q *Query) BlockResults() (*coretypes.ResultBlockResults, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BlockResultsRPC(q)
}

// Status returns information about a node status
func (q *Query) Status() (*coretypes.ResultStatus, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return StatusRPC(q)
}

// ABCIInfo returns general information about the ABCI application
func (q *Query) ABCIInfo() (*coretypes.ResultABCIInfo, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return ABCIInfoRPC(q)
}

// ABCIQuery returns data from a particular path in the ABCI application
func (q *Query) ABCIQuery(path string, data string, prove bool) (*coretypes.ResultABCIQuery, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return ABCIQueryRPC(q, path, data, prove)
}

func (q *Query) EpochsAtHeight(height int64) (*osmosisEpochs.QueryEpochsInfoResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	resp, err := EpochsAtHeightRPC(q, height)
	return resp, err
}

func (q *Query) BlockSearchEpochStartsLessThanHeight(height int64) (*coretypes.ResultBlockSearch, error) {
	return BlockSearchEpochStartsLessThanHeightRPC(q, height, 1, 100)
}

func (q *Query) ProtorevDeveloperAccount() (*osmosisProtorev.QueryGetProtoRevDeveloperAccountResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	resp, err := ProtorevDeveloperAccountRPC(q)
	return resp, err
}

func (q *Query) ContractsByCodeIDAtHeight(codeID uint64, height int64) (*wasmTypes.QueryContractsByCodeResponse, error) {
	return ContractsByCodeIDAtHeight(q, codeID, height)
}
