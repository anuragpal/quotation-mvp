package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"quotation-mvp/x/quotation/types"
)

func (k msgServer) RequestQuotation(goCtx context.Context, msg *types.MsgRequestQuotation) (*types.MsgRequestQuotationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	rfq := types.Quotation{
		Id:       msg.Id,
		Envelope: msg.Envelope,
		Tx:       msg.Tx,
		Version:  msg.Version,
		Creator:  msg.Creator,
	}
	k.SetQuotation(ctx, rfq)

	return &types.MsgRequestQuotationResponse{}, nil
}
