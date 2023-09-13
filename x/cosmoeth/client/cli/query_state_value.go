package cli

import (
	"strconv"

	"CosmoEth/x/cosmoeth/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStateValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state-value [address] [slot]",
		Short: "Query state-value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqSlot := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStateValueRequest{

				Address: reqAddress,
				Slot:    reqSlot,
			}

			res, err := queryClient.StateValue(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
