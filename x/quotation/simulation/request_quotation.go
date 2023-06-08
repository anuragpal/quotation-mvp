package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"quotation-mvp/x/quotation/keeper"
	"quotation-mvp/x/quotation/types"
)

func SimulateMsgRequestQuotation(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRequestQuotation{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RequestQuotation simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RequestQuotation simulation not implemented"), nil, nil
	}
}
