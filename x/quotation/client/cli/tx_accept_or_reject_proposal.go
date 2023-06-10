package cli

import (
	"strconv"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"quotation-mvp/x/quotation/types"
)

var _ = strconv.Itoa(0)

func CmdAcceptOrRejectProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-or-reject-proposal [payload]",
		Short: "Broadcast message accept-or-reject-proposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPayload := new(types.MsgAcceptOrRejectProposal)
			err = json.Unmarshal([]byte(args[0]), argPayload)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptOrRejectProposal(
				clientCtx.GetFromAddress().String(),
				argPayload,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
