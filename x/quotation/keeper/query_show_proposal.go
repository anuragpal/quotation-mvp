package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"quotation-mvp/x/quotation/types"
)

func (k Keeper) ShowProposal(goCtx context.Context, req *types.QueryShowProposalRequest) (*types.QueryShowProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, found := k.GetProposal(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowProposalResponse{Proposal: &proposal}, nil
}
