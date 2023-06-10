package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"quotation-mvp/x/quotation/types"
)

func (k msgServer) AcceptOrRejectProposal(goCtx context.Context, msg *types.MsgAcceptOrRejectProposal) (*types.MsgAcceptOrRejectProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	if(msg.Tx.Action) {
		proposal.Tx.Status = 1
		quotation, found := k.GetQuotation(ctx, proposal.Tx.RqId)
		if !found {
			return nil, sdkerrors.ErrKeyNotFound
		}
		quotation.Tx.Status = 1
		k.SetQuotation(ctx, quotation)
	} else {
		proposal.Tx.Status = 2
	}

	k.SetProposal(ctx, proposal)

	return &types.MsgAcceptOrRejectProposalResponse{}, nil
}
