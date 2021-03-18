package errors

import (
	"github.com/micro/micro/v3/service/errors"
)

var (
	CtxMetadataNotRetrieved = errors.Timeout("CTX_META_NOT_RETRIEVED", "Failed to retrieve metadata from context")
	AuthHeaderNotFound      = errors.Timeout("AUTH_HEADER_NOT_FOUND", "Authentication header not found in context")
	AuthHeaderInvalidFormat = errors.Timeout("AUTH_HEADER_INVALID_FMT", "Authentication header format is invalid, bearer schema expected")
)
