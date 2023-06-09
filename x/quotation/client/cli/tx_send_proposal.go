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

func CmdSendProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-proposal [payload]",
		Short: "Broadcast message send-proposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPayload := new(types.MsgSendProposal)
			err = json.Unmarshal([]byte(args[0]), argPayload)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendProposal(
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
