package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestQuotation{}, "quotation/RequestQuotation", nil)
	cdc.RegisterConcrete(&MsgSendProposal{}, "quotation/SendProposal", nil)
	cdc.RegisterConcrete(&MsgAcceptOrRejectProposal{}, "quotation/AcceptOrRejectProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestQuotation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptOrRejectProposal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
