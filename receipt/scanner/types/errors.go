package types

import (
	err "errors"
	"github.com/micro/micro/v3/service/errors"
)

var (
	UnfitParser        = err.New("UNFIT_PARSER")
	UnfitCode          = errors.BadRequest("UNFIT_CODE_FORMAT", "Unsupported code string passed")
	CodeNotResolved    = errors.BadRequest("CODE_NOT_RESOLVED", "Code data is not resolved")
	CodeResolveTimeout = errors.Timeout("CODE_RESOLVE_TIMEOUT", "Code resolve request timeout")
)
