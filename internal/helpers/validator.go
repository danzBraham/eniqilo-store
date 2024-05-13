package helpers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidatePhoneNumber(phoneNumber string) (bool, error) {
	regexPattern := `^\+\d{1,3}-?\d{9,14}$`
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return false, err
	}
	return regex.MatchString(phoneNumber), nil
}

func ValidatePayload(payload interface{}) error {
	if err := validate.Struct(payload); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
