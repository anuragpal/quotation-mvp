package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "quotation-mvp/testutil/keeper"
	"quotation-mvp/x/quotation/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QuotationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
