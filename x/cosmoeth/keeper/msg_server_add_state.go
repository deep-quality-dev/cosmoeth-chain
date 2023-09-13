package keeper

import (
	"context"

	"CosmoEth/x/cosmoeth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddState(goCtx context.Context, msg *types.MsgAddState) (*types.MsgAddStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddStateResponse{}, nil
}
