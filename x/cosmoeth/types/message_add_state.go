package types

import (
	"CosmoEth/utils"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddState = "add_state"

var _ sdk.Msg = &MsgAddState{}

func NewMsgAddState(creator string, address string, slots []string, values []string, storageProof string) *MsgAddState {
	return &MsgAddState{
		Creator:      creator,
		Address:      address,
		Slots:        slots,
		Values:       values,
		StorageProof: storageProof,
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

	if !utils.IsEthereumAddress(msg.Address) {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ethereum address")
	}

	for _, slot := range msg.Slots {
		if !utils.IsValidHexString(slot) {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid storage slot")
		}
	}

	for _, val := range msg.Values {
		if !utils.IsValidHexString(val) {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid state value")
		}
	}

	if len(msg.StorageProof) == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "empty storage proof")
	}
	return nil
}
