package keeper_test

import (
	"testing"

	testkeeper "CosmoEth/testutil/keeper"
	"CosmoEth/x/cosmoeth/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmoethKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
