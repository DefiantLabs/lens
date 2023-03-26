package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgLockTokens        = "lock_tokens"
	TypeMsgBeginUnlockingAll = "begin_unlocking_all"
	TypeMsgUnlockTokens      = "unlock_tokens"
	TypeMsgBeginUnlocking    = "begin_unlocking"
	TypeMsgUnlockPeriodLock  = "unlock_period_lock"
)

var _ sdk.Msg = &MsgUnlockPeriodLock{}

// NewMsgUnlockPeriodLock creates a message to begin unlock tokens of a specific lockid
func NewMsgUnlockPeriodLock(owner sdk.AccAddress, id uint64) *MsgUnlockPeriodLock {
	return &MsgUnlockPeriodLock{
		Owner: owner.String(),
		ID:    id,
	}
}

func (m MsgUnlockPeriodLock) Route() string { return RouterKey }
func (m MsgUnlockPeriodLock) Type() string  { return TypeMsgUnlockPeriodLock }
func (m MsgUnlockPeriodLock) ValidateBasic() error {
	return nil
}
func (m MsgUnlockPeriodLock) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgUnlockPeriodLock) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}
