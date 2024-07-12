package autherror

import "errors"

var (
	ErrMissingAuthHeader          = errors.New("missing Authorization header")
	ErrInvalidAuthHeader          = errors.New("invalid Authorization header")
	ErrInvalidToken               = errors.New("invalid token")
	ErrUnknownClaims              = errors.New("unknown claims type")
	ErrUserIDNotFoundInTheContext = errors.New("user ID not found in the context")
)
