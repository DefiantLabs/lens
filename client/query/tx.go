package query

import (
	"fmt"

	"github.com/DefiantLabs/lens/client"
	"github.com/cosmos/cosmos-sdk/types/query"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
)

// TxRPC Get Transactions for the given block height.
// Other query options can be specified with the GetTxsEventRequest.
//
// RPC endpoint is defined in cosmos-sdk: proto/cosmos/tx/v1beta1/service.proto,
// See GetTxsEvent(GetTxsEventRequest) returns (GetTxsEventResponse)
func TxsAtHeightRPC(q *Query, height int64, codec client.Codec) (*txTypes.GetTxsEventResponse, error, error) {
	if q.Options.Pagination == nil {
		pagination := &query.PageRequest{Limit: 100}
		q.Options.Pagination = pagination
	}
	orderBy := txTypes.OrderBy_ORDER_BY_UNSPECIFIED

	req := &txTypes.GetTxsEventRequest{Pagination: q.Options.Pagination, OrderBy: orderBy, Query: "tx.height=" + fmt.Sprintf("%d", height)}
	return TxsRPC(q, req, codec)
}

// TxRPC Get Transactions for the given block height.
// Other query options can be specified with the GetTxsEventRequest.
//
// RPC endpoint is defined in cosmos-sdk: proto/cosmos/tx/v1beta1/service.proto,
// See GetTxsEvent(GetTxsEventRequest) returns (GetTxsEventResponse)
// This function returns 2 errors, one for query erroring which are usually fatal, the other for unpacking errors
// which can sometimes be skipped based on which TX message was not successfully unpacked
func TxsRPC(q *Query, req *txTypes.GetTxsEventRequest, codec client.Codec) (*txTypes.GetTxsEventResponse, error, error) {
	queryClient := txTypes.NewServiceClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	res, err := queryClient.GetTxsEvent(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	unpackError := &TXUnpackError{}
	unpackError.Errors = []string{}
	for _, tx := range res.GetTxs() {
		err := tx.UnpackInterfaces(codec.InterfaceRegistry)

		if err != nil {
			unpackError.Errors = append(unpackError.Errors, err.Error())
		}
	}

	if len(unpackError.Errors) != 0 {
		return res, unpackError, nil
	}

	return res, nil, nil
}
