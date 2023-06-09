package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"quotation-mvp/x/quotation/types"
)

func (k msgServer) SendProposal(goCtx context.Context, msg *types.MsgSendProposal) (*types.MsgSendProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	p := types.Proposal{
		Id:       msg.Id,
		Envelope: msg.Envelope,
		Tx:       msg.Tx,
		Version:  msg.Version,
		Creator:  msg.Creator,
	}
	k.SetProposal(ctx, p)

	return &types.MsgSendProposalResponse{}, nil
}
