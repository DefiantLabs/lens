package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lockup module sentinel errors
var (
	ErrNotLockOwner = sdkerrors.New(ModuleName, 1, "msg sender is not the owner of specified lock")
)
