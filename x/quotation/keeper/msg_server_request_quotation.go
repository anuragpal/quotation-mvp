package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"quotation-mvp/x/quotation/types"
)

func (k msgServer) RequestQuotation(goCtx context.Context, msg *types.MsgRequestQuotation) (*types.MsgRequestQuotationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetQuotation(ctx, msg.Id)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Quotation with Id already exists.")
	}

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
