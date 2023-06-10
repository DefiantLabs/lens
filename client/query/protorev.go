package query

import (
	protorevTypes "github.com/DefiantLabs/lens/osmosis/x/protorev/types"
)

func ProtorevDeveloperAccountRPC(q *Query) (*protorevTypes.QueryGetProtoRevDeveloperAccountResponse, error) {
	req := protorevTypes.QueryGetProtoRevDeveloperAccountRequest{}
	queryClient := protorevTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.GetProtoRevDeveloperAccount(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
