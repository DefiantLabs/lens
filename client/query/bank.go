package query

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"time"
)

// BalanceWithAddressRPC returns the balance of all coins for a single account.
func BalanceWithAddressRPC(q *Query, address string) (sdk.Coins, error) {
	var req *bankTypes.QueryAllBalancesRequest
	req = &bankTypes.QueryAllBalancesRequest{Address: address, Pagination: &q.Options.Pagination}
	queryClient := bankTypes.NewQueryClient(q.Client)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
	res, err := queryClient.AllBalances(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Balances, nil
}

// TotalSupplyRPC returns the supply of all coins
func TotalSupplyRPC(q *Query) (sdk.Coins, error) {
	req := &bankTypes.QueryTotalSupplyRequest{Pagination: &q.Options.Pagination}
	queryClient := bankTypes.NewQueryClient(q.Client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
	res, err := queryClient.TotalSupply(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Supply, nil
}

// DenomsMetadataRPC returns the metadata for all denoms
func DenomsMetadataRPC(q *Query) ([]bankTypes.Metadata, error) {
	req := &bankTypes.QueryDenomsMetadataRequest{Pagination: &q.Options.Pagination}
	queryClient := bankTypes.NewQueryClient(q.Client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
	res, err := queryClient.DenomsMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Metadatas, nil
}
