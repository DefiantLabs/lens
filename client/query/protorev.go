package query

import (
	osmosisProtorev "github.com/osmosis-labs/osmosis/v26/x/protorev/types"
)

func ProtorevDeveloperAccountRPC(q *Query) (*osmosisProtorev.QueryGetProtoRevDeveloperAccountResponse, error) {
	req := osmosisProtorev.QueryGetProtoRevDeveloperAccountRequest{}
	queryClient := osmosisProtorev.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.GetProtoRevDeveloperAccount(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
