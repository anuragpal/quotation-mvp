package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"quotation-mvp/x/quotation/types"
)

func (k Keeper) SetQuotation(ctx sdk.Context, quotation types.Quotation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotationKeyPrefix))
	b := k.cdc.MustMarshal(&quotation)
	store.Set(types.QuotationKey(
		quotation.Id,
	), b)
}

func (k Keeper) GetQuotation(ctx sdk.Context, id string) (val types.Quotation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotationKeyPrefix))

	b := store.Get(types.QuotationKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
