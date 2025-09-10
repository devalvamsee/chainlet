package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

const (
	codeErrIbcCltDenomEmpty = uint32(iota) + 2 // NOTE: code 1 is reserved for internal errors
	codeErrIbcCltDenomInvalid
)

// x/chainlet module sentinel errors
var (
	ErrIbcCltDenomEmpty   = errors.Register(ModuleName, codeErrIbcCltDenomEmpty, "ibc clt denom is not set")
	ErrIbcCltDenomInvalid = errors.Register(ModuleName, codeErrIbcCltDenomInvalid, "ibc clt denom is invalid")
	// this line is used by starport scaffolding # ibc/errors
)
