package evm

import (
	"context"

	feedertypes "CosmoEth/feeder/types"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func ConnectToEthereum() (*rpc.Client, *ethclient.Client, error) {
	// Replace with your Ethereum node's URL
	rpcCli, err := rpc.Dial("https://docs-demo.quiknode.pro")
	if err != nil {
		return nil, nil, err
	}

	ethCli := ethclient.NewClient(rpcCli)
	return rpcCli, ethCli, nil
}

func GetProof(ctx context.Context, addr string, keys []string) (*feedertypes.StorageProof, error) {

	rpcCli, ethCli, err := ConnectToEthereum()
	if err != nil {
		return nil, err
	}

	blockData, err := ethCli.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp feedertypes.StorageProof
	if err := rpcCli.CallContext(
		ctx,
		&resp,
		"eth_getProof",
		addr,
		keys,
		"latest",
	); err != nil {
		return nil, err
	}

	resp.StateRoot = blockData.Root()
	resp.Height = blockData.Header().Number
	return &resp, nil
}
