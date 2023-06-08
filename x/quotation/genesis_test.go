package quotation_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "quotation-mvp/testutil/keeper"
	"quotation-mvp/testutil/nullify"
	"quotation-mvp/x/quotation"
	"quotation-mvp/x/quotation/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QuotationKeeper(t)
	quotation.InitGenesis(ctx, *k, genesisState)
	got := quotation.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
