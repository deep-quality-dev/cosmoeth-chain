package keeper

import (
	"context"

	"CosmoEth/x/cosmoeth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StateValue(goCtx context.Context, req *types.QueryStateValueRequest) (*types.QueryStateValueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stateInfo := k.GetStateInfo(ctx, req.Address, req.Slot)

	resp := types.QueryStateValueResponse{
		Value:  stateInfo.Value,
		Height: stateInfo.Height,
	}

	return &resp, nil
}
