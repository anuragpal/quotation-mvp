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

func (k Keeper) ListQuotation(goCtx context.Context, req *types.QueryListQuotationRequest) (*types.QueryListQuotationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var quotations []types.Quotation
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	qStore := prefix.NewStore(store, types.KeyPrefix(types.QuotationKeyPrefix))

	pageRes, err := query.Paginate(qStore, req.Pagination, func(key []byte, value []byte) error {
		var quotation types.Quotation
		if err := k.cdc.Unmarshal(value, &quotation); err != nil {
			return err
		}

		quotations = append(quotations, quotation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListQuotationResponse{Quotation: quotations, Pagination: pageRes}, nil
}
