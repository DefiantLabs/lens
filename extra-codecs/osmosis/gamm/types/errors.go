package types

// x/gamm module sentinel errors
var (
// ErrPoolNotFound      = sdkerrors.New(ModuleName, 1, "pool not found")
// ErrPoolAlreadyExist  = sdkerrors.New(ModuleName, 2, "pool already exist")
// ErrPoolLocked        = sdkerrors.New(ModuleName, 3, "pool is locked")
// ErrTooFewPoolAssets  = sdkerrors.New(ModuleName, 4, "pool should have at least 2 assets, as they must be swapping between at least two assets")
// ErrTooManyPoolAssets = sdkerrors.New(ModuleName, 5, "pool has too many assets (currently capped at 8 assets per pool)")
// ErrLimitMaxAmount    = sdkerrors.New(ModuleName, 6, "calculated amount is larger than max amount")
// ErrLimitMinAmount    = sdkerrors.New(ModuleName, 7, "calculated amount is lesser than min amount")
// ErrInvalidMathApprox = sdkerrors.New(ModuleName, 8, "invalid calculated result")

// ErrEmptyRoutes              = sdkerrors.New(ModuleName, 21, "routes not defined")
// ErrEmptyPoolAssets          = sdkerrors.New(ModuleName, 22, "PoolAssets not defined")
// ErrNegativeSwapFee = sdkerrors.New(ModuleName, 23, "swap fee is negative")
// // ErrNegativeExitFee          = sdkerrors.New(ModuleName, 24, "exit fee is negative")
// ErrTooMuchSwapFee = sdkerrors.New(ModuleName, 25, "swap fee should be lesser than 1 (100%)")

// ErrTooMuchExitFee           = sdkerrors.New(ModuleName, 26, "exit fee should be lesser than 1 (100%)")
// ErrNotPositiveWeight        = sdkerrors.New(ModuleName, 27, "token weight should be greater than 0")
// ErrWeightTooLarge           = sdkerrors.New(ModuleName, 28, "user specified token weight should be less than 2^20")
// ErrNotPositiveCriteria      = sdkerrors.New(ModuleName, 29, "min out amount or max in amount should be positive")
// ErrNotPositiveRequireAmount = sdkerrors.New(ModuleName, 30, "required amount should be positive")
// ErrTooManyTokensOut         = sdkerrors.New(ModuleName, 31, "tx is trying to get more tokens out of the pool than exist")

// ErrPoolParamsInvalidDenom     = sdkerrors.New(ModuleName, 50, "pool params' LBP params has an invalid denomination")
// ErrPoolParamsInvalidNumDenoms = sdkerrors.New(ModuleName, 51, "pool params' LBP doesn't have same number of params as underlying pool")
)
