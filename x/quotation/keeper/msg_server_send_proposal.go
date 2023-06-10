package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"quotation-mvp/x/quotation/types"
)

func (k msgServer) SendProposal(goCtx context.Context, msg *types.MsgSendProposal) (*types.MsgSendProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	quotation, found := k.GetQuotation(ctx, msg.Tx.RqId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	p := types.Proposal{
		Id:       msg.Id,
		Envelope: msg.Envelope,
		Tx:       msg.Tx,
		Version:  msg.Version,
		Creator:  msg.Creator,
	}

	var ps []string

	ps = quotation.Tx.Proposals
	ps = append(ps, msg.Id)
	quotation.Tx.Proposals = ps
	k.SetQuotation(ctx, quotation)
	k.SetProposal(ctx, p)

	return &types.MsgSendProposalResponse{}, nil
}
