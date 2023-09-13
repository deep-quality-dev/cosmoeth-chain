package cosmoeth_test

import (
	"testing"

	keepertest "CosmoEth/testutil/keeper"
	"CosmoEth/testutil/nullify"
	"CosmoEth/x/cosmoeth"
	"CosmoEth/x/cosmoeth/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoethKeeper(t)
	cosmoeth.InitGenesis(ctx, *k, genesisState)
	got := cosmoeth.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
