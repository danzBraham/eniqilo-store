package usererror

import "errors"

var (
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
)
