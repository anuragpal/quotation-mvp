package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"quotation-mvp/x/quotation/types"
)

func (k Keeper) ListProposal(goCtx context.Context, req *types.QueryListProposalRequest) (*types.QueryListProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var proposals []types.Proposal

	store := ctx.KVStore(k.storeKey)
	qStore := prefix.NewStore(store, types.KeyPrefix(types.ProposalKeyPrefix))

	pageRes, err := query.Paginate(qStore, req.Pagination, func(key []byte, value []byte) error {
		var proposal types.Proposal
		if err := k.cdc.Unmarshal(value, &proposal); err != nil {
			return err
		}

		proposals = append(proposals, proposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListProposalResponse{Proposal: proposals, Pagination: pageRes}, nil
}
