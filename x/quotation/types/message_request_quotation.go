package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestQuotation = "request_quotation"

var _ sdk.Msg = &MsgRequestQuotation{}

func NewMsgRequestQuotation(creator string, payload *MsgRequestQuotation) *MsgRequestQuotation {
	return &MsgRequestQuotation{
		Creator:  creator,
		Id:       payload.Id,
		Tx:       payload.Tx,
		Envelope: payload.Envelope,
		Version:  payload.Version,
	}
}

func (msg *MsgRequestQuotation) Route() string {
	return RouterKey
}

func (msg *MsgRequestQuotation) Type() string {
	return TypeMsgRequestQuotation
}

func (msg *MsgRequestQuotation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestQuotation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestQuotation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
