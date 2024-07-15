package usererror

import "errors"

var (
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	ErrUserNotFound             = errors.New("user not found")
	ErrInvalidPassword          = errors.New("invalid password")
)
