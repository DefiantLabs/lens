package balancer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreateBalancerPool = "create_balancer_pool"
)

var _ sdk.Msg = &MsgCreateBalancerPool{}

func (msg MsgCreateBalancerPool) Route() string { return "" }
func (msg MsgCreateBalancerPool) Type() string  { return TypeMsgCreateBalancerPool }
func (msg MsgCreateBalancerPool) ValidateBasic() error {
	return nil
}
func (msg MsgCreateBalancerPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg MsgCreateBalancerPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
