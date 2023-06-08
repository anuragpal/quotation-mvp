package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"quotation-mvp/x/quotation/types"
)

func (k Keeper) ShowQuotation(goCtx context.Context, req *types.QueryShowQuotationRequest) (*types.QueryShowQuotationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	quotation, found := k.GetQuotation(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowQuotationResponse{Quotation: &quotation}, nil
}
