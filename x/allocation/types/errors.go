package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	ErrInvalidPrevote         = sdkerrors.Register(ModuleName, 2, "invalid prevote hashes")
	ErrInvalidOracleData      = sdkerrors.Register(ModuleName, 3, "invalid oracle data")
	ErrNoPrecommit            = sdkerrors.Register(ModuleName, 4, "no precommit for validator")
	ErrHashMismatch           = sdkerrors.Register(ModuleName, 6, "precommit hash doesn't match commit hash")
	ErrUnsupportedDataType    = sdkerrors.Register(ModuleName, 8, "unsupported data type")
	ErrDuplicatedOracleData   = sdkerrors.Register(ModuleName, 10, "duplicated oracle data")
	ErrAlreadyCommitted       = sdkerrors.Register(ModuleName, 11, "allocation commit already provided")
	ErrInvalidOracleVote      = sdkerrors.Register(ModuleName, 12, "invalid oracle vote")
	ErrAggregatedDataNotFound = sdkerrors.Register(ModuleName, 13, "aggregated oracle data not found")
)
