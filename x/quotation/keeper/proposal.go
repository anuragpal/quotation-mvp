package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"quotation-mvp/x/quotation/types"
)

func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))
	b := k.cdc.MustMarshal(&proposal)
	store.Set(types.ProposalKey(
		proposal.Id,
	), b)
}

func (k Keeper) GetProposal(ctx sdk.Context, id string) (val types.Proposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))

	b := store.Get(types.ProposalKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
