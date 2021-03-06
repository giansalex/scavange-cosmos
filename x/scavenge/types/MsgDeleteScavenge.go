package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgDeleteScavenge
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgDeleteScavenge{}

// MsgDeleteScavenge - struct for unjailing jailed validator
type MsgDeleteScavenge struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`           // address of the scavenger creator
	SolutionHash string         `json:"solutionhash" yaml:"solutionhash"` // solutionhash of the scavenge
}

// NewMsgDeleteScavenge creates a new MsgDeleteScavenge instance
func NewMsgDeleteScavenge(creator sdk.AccAddress, solutionHash string) MsgDeleteScavenge {
	return MsgDeleteScavenge{
		Creator:      creator,
		SolutionHash: solutionHash,
	}
}

// DeleteScavengeConst is DeleteScavenge Constant
const DeleteScavengeConst = "DeleteScavenge"

// nolint
func (msg MsgDeleteScavenge) Route() string { return RouterKey }
func (msg MsgDeleteScavenge) Type() string  { return DeleteScavengeConst }
func (msg MsgDeleteScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDeleteScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionHash can't be empty")
	}
	return nil
}
