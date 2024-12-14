package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"skeleton/error_handler"
	"skeleton/repositories"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	paramOpeningChar        = '('
	paramClosingChar        = ')'
	paramSepratorValuesChar = '-'
)

func getTagParamValues(param string) (params []string, err error) {
	if param[0] != paramOpeningChar && param[len(param)-1] != paramClosingChar {
		err = errors.New("custom validatio has syntax error: param is not in '(' and ')' charecters.")

		return
	}

	param = param[1 : len(param)-1]

	params = strings.Split(param, "-")

	return
}

func getTagNumberParams(number int, param string) (params []string, err error) {
	params, err = getTagParamValues(param)

	if err != nil {
		return
	}

	if len(params) >= number {
		params = params[0:number]
	} else {
		err = errors.New(fmt.Sprintf("custom validatio has syntax error: params length are less than %v", number))
	}

	return
}

func getExistsParams(param string) (table, column string, err error) {
	params, err := getTagNumberParams(2, param)

	if err != nil {
		return
	}

	table = params[0]
	column = params[1]

	return
}

func exists(fl validator.FieldLevel) bool {

	table, column, err := getExistsParams(fl.Param())

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	value := fl.Field().String()

	return repositories.CheckExistsCustom(table, column, value)
}

func unique(fl validator.FieldLevel) bool {
	return !exists(fl)
}

func isJson(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	var jsonObject map[string]interface{}
	var jsonArray []interface{}

	errObject := json.Unmarshal([]byte(value), &jsonObject)
	errArray := json.Unmarshal([]byte(value), &jsonArray)

	if errObject != nil && errArray != nil {
		return false
	} else {
		return true
	}
}

func customValidations(Validate *validator.Validate) {
	validate.RegisterValidation("exists", exists)
	validate.RegisterValidation("unique", unique)
	validate.RegisterValidation("isjson", isJson)
}
