package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"quotation-mvp/x/quotation/keeper"
	"quotation-mvp/x/quotation/types"
)

func SimulateMsgAcceptOrRejectProposal(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAcceptOrRejectProposal{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AcceptOrRejectProposal simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AcceptOrRejectProposal simulation not implemented"), nil, nil
	}
}
