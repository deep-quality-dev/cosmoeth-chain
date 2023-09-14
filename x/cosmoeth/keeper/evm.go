package keeper

import (
	"encoding/json"

	"cosmossdk.io/errors"

	"CosmoEth/utils"
	"CosmoEth/x/cosmoeth/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func (k Keeper) CheckStateValidity(ctx sdk.Context, stateInfo types.StateInfo) (bool, error) {
	var err error
	var value []byte

	if len(stateInfo.Value) != 0 {
		value, err = rlp.EncodeToBytes(common.FromHex(stateInfo.Value))
		if err != nil {
			return false, err
		}
	}

	return utils.VerifyProof(common.HexToHash(stateInfo.Root), stateInfo.Slot, value, stateInfo.Proofs)
}

func (k Keeper) AddStateInfo(ctx sdk.Context, address string, root string, height uint64, storageProofs []string) error {
	for _, proof := range storageProofs {
		storageProof := types.StorageProof{}
		err := json.Unmarshal([]byte(proof), &storageProof)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
		}

		err = k.SetStateInfoWithValidation(ctx, address, root, height, storageProof)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) SetStateInfoWithValidation(ctx sdk.Context, address string, root string, height uint64, storageProof types.StorageProof) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefix(types.StateKey))

	stateInfo := types.StateInfo{
		Address: address,
		Slot:    storageProof.Key,
		Height:  height,
		Value:   storageProof.Value,
		Root:    root,
		Proofs:  storageProof.Proof,
	}

	// check validity of submitted state info
	isValid, err := k.CheckStateValidity(ctx, stateInfo)
	if err != nil {
		return errors.Wrapf(err, "state validation failed")
	}
	if !isValid {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "state validation failed")
	}

	rawStateInfo, err := json.Marshal(stateInfo)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}

	// save state info into the store
	prefixStore.Set(append(types.KeyPrefix(address), types.KeyPrefix(storageProof.Key)...), []byte(rawStateInfo))
	return nil
}

func (k Keeper) GetStateInfo(ctx sdk.Context, address string, slot string) types.StateInfo {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefix(types.StateKey))
	rawInfo := prefixStore.Get(append(types.KeyPrefix(address), types.KeyPrefix(slot)...))

	// unmarshal raw state info to proto object
	stateInfo := types.StateInfo{}
	err := json.Unmarshal([]byte(rawInfo), &stateInfo)
	if err != nil {
		return types.StateInfo{}
	}

	return stateInfo
}
