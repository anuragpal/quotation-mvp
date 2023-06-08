package keeper

import (
	"quotation-mvp/x/quotation/types"
)

var _ types.QueryServer = Keeper{}
