package requests

import (
	"strings"
	"user_segmentation_service/error_handler"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func InitValidation() {
	validate = validator.New()
	customValidations(validate)
	initAttributes()
	initRules()
}

func ValidateRequestBody(validationStruct interface{}) {
	err := validate.Struct(validationStruct)

	if err != nil {
		error_handler.ThrowBodyValidationError(err.(validator.ValidationErrors), GenerateValidationMessage)
	}
}

func ValidateRouteParam(validationStruct interface{}) {
	err := validate.Struct(validationStruct)

	if err != nil {
		error_handler.ThrowRouteValidationError(err.(validator.ValidationErrors))
	}
}

func GenerateValidationMessage(rule, attribute, rule_value string) (message string) {
	message, rule_exists := Rules[rule]

	if rule_exists {
		attribute_message, attribute_exists := Attributes[attribute]

		if !attribute_exists {
			attribute_message = attribute
		}

		message = strings.Replace(message, ":attribute", attribute_message, 1)
		message = strings.Replace(message, ":value", rule_value, 1)
	}

	return
}
