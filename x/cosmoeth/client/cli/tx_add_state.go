package cli

import (
	"strconv"

	"CosmoEth/x/cosmoeth/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdAddState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-state [address] [slots] [values] [storage-proof]",
		Short: "Broadcast message add-state",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argSlots := strings.Split(args[1], listSeparator)
			argValues := strings.Split(args[2], listSeparator)
			argStorageProof := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddState(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argSlots,
				argValues,
				argStorageProof,
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
