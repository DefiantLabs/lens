package query

import (
	epochsTypes "github.com/DefiantLabs/lens/osmosis/x/epochs/types"
)

func EpochsAtHeightRPC(q *Query, height int64) (*epochsTypes.QueryEpochsInfoResponse, error) {
	req := epochsTypes.QueryEpochsInfoRequest{}
	queryClient := epochsTypes.NewQueryClient(q.Client)

	if height > 0 {
		q.Options.Height = height
	}

	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.EpochInfos(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
