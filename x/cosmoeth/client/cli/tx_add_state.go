package cli

import (
	"strconv"

	"CosmoEth/x/cosmoeth/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-state [address] [height] [root] [storage-proofs]",
		Short: "Broadcast message add-state",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argHeight, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			argRoot := args[2]
			argStorageProofs := strings.Split(args[3], ";")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddState(
				clientCtx.GetFromAddress().String(),
				argAddress,
				uint64(argHeight),
				argRoot,
				argStorageProofs,
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
