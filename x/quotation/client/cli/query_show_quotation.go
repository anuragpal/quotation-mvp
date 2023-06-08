package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"quotation-mvp/x/quotation/types"
)

var _ = strconv.Itoa(0)

func CmdShowQuotation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-quotation [id]",
		Short: "Query show-quotation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowQuotationRequest{

				Id: reqId,
			}

			res, err := queryClient.ShowQuotation(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
