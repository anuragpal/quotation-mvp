package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptOrRejectProposal = "accept_or_reject_proposal"

var _ sdk.Msg = &MsgAcceptOrRejectProposal{}

func NewMsgAcceptOrRejectProposal(creator string, payload *MsgAcceptOrRejectProposal) *MsgAcceptOrRejectProposal {
	return &MsgAcceptOrRejectProposal{
		Creator: creator,
		Id:       payload.Id,
		Tx:       payload.Tx,
		Envelope: payload.Envelope,
		Version:  payload.Version,
	}
}

func (msg *MsgAcceptOrRejectProposal) Route() string {
	return RouterKey
}

func (msg *MsgAcceptOrRejectProposal) Type() string {
	return TypeMsgAcceptOrRejectProposal
}

func (msg *MsgAcceptOrRejectProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptOrRejectProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptOrRejectProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
