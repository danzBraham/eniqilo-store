package validator

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	validate.RegisterValidation("image_url", validateImageURL)
	validate.RegisterValidation("phone_number", validatePhoneNumber)
}

func formatValidationError(err error) string {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		var sb strings.Builder
		for _, fieldError := range validationError {
			sb.WriteString(fmt.Sprintf("field %s failed on the '%s' tag; ", fieldError.Field(), fieldError.Tag()))
		}
		return sb.String()
	}
	return err.Error()
}

func ValidatePayload(payload any) error {
	if err := validate.Struct(payload); err != nil {
		return fmt.Errorf(formatValidationError(err))
	}
	return nil
}

func validateImageURL(fl validator.FieldLevel) bool {
	u, err := url.ParseRequestURI(fl.Field().String())
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	if u.Host == "" {
		return false
	}
	ext := path.Ext(u.Path)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return false
	}
	return true
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	re := regexp.MustCompile(`^\+\d{1,3}-?\d{4,13}$`)
	return re.MatchString(phoneNumber)
}
