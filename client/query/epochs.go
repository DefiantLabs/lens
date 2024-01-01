package query

import (
	"fmt"

	osmosisEpochs "github.com/osmosis-labs/osmosis/x/epochs/types"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func EpochsAtHeightRPC(q *Query, height int64) (*osmosisEpochs.QueryEpochsInfoResponse, error) {
	req := osmosisEpochs.QueryEpochsInfoRequest{}
	queryClient := osmosisEpochs.NewQueryClient(q.Client)

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

// BlockSearchEpochStartsLessThanHeightRPC searches for blocks with the epoch_start.epoch_number event with height less than the given height. This query only makes sense for Osmosis, which has the Epoch module emitting this event
// in the BeginBlock events.
func BlockSearchEpochStartsLessThanHeightRPC(q *Query, height int64, page int, perPage int) (*coretypes.ResultBlockSearch, error) {
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	resp, err := q.Client.RPCClient.BlockSearch(ctx, fmt.Sprintf("block.height<%d AND epoch_start.epoch_number EXISTS", height), &page, &perPage, "")
	if err != nil {
		return nil, err
	}

	return resp, nil
}
