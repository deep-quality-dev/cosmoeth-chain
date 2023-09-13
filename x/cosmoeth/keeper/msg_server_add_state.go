package keeper

import (
	"context"

	"CosmoEth/x/cosmoeth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddState(goCtx context.Context, msg *types.MsgAddState) (*types.MsgAddStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddStateInfo(ctx, msg.Address, msg.Root, msg.Height, msg.StorageProofs)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddStateResponse{}, nil
}
