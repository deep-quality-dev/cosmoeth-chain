package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

const TypeMsgAddState = "add_state"

var _ sdk.Msg = &MsgAddState{}

func NewMsgAddState(creator string, address string, height uint64, storageProofs []string) *MsgAddState {
	return &MsgAddState{
		Creator:       creator,
		Address:       address,
		Height:        height,
		StorageProofs: storageProofs,
	}
}

func (msg *MsgAddState) Route() string {
	return RouterKey
}

func (msg *MsgAddState) Type() string {
	return TypeMsgAddState
}

func (msg *MsgAddState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !common.IsHexAddress(msg.Address) {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ethereum address")
	}

	if msg.Height == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "zero block number")
	}

	if len(msg.StorageProofs) == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "empty storage proofs")
	}
	return nil
}
