package quotation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"quotation-mvp/testutil/sample"
	quotationsimulation "quotation-mvp/x/quotation/simulation"
	"quotation-mvp/x/quotation/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = quotationsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestQuotation = "op_weight_msg_request_quotation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestQuotation int = 100

	opWeightMsgSendProposal = "op_weight_msg_send_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendProposal int = 100

	opWeightMsgAcceptOrRejectProposal = "op_weight_msg_accept_or_reject_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptOrRejectProposal int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	quotationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&quotationGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestQuotation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestQuotation, &weightMsgRequestQuotation, nil,
		func(_ *rand.Rand) {
			weightMsgRequestQuotation = defaultWeightMsgRequestQuotation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestQuotation,
		quotationsimulation.SimulateMsgRequestQuotation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendProposal, &weightMsgSendProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSendProposal = defaultWeightMsgSendProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendProposal,
		quotationsimulation.SimulateMsgSendProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptOrRejectProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptOrRejectProposal, &weightMsgAcceptOrRejectProposal, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptOrRejectProposal = defaultWeightMsgAcceptOrRejectProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptOrRejectProposal,
		quotationsimulation.SimulateMsgAcceptOrRejectProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
