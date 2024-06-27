package query

import (
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func ContractsByCodeIDAtHeight(q *Query, codeID uint64, height int64) (*wasmTypes.QueryContractsByCodeResponse, error) {
	req := wasmTypes.QueryContractsByCodeRequest{
		CodeId: codeID,
	}
	queryClient := wasmTypes.NewQueryClient(q.Client)

	if height > 0 {
		q.Options.Height = height
	}

	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.ContractsByCode(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
