package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validator = validator.New()

func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func init() {
	Validator.RegisterValidation("notBlank", notBlank)
}

// func ValidateRequest(req any) map[string]string {
// 	if err := Validator.Struct(req); err != nil {
// 		errors := make(map[string]string)
// 		for _, err := range err.(validator.ValidationErrors) {
// 			errors[err.Field()] = fmt.Sprintf("failed on '%s' tag", err.Tag())
// 		}
// 		return errors
// 	}
// 	return nil
// }
