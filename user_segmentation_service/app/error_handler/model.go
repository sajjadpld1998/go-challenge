package error_handler

type messageGenerator func(rule, attribute, rule_value string) string

type ResponseError struct {
	// status of opration success
	Status bool `json:"status" example:"false"`
	// response code
	Code int `json:"code" example:"422"`
	// content
	Error interface{} `json:"error"`
}

type ResponseErrorDetail struct {
	// error message
	Message string `json:"message" example:"Server Error"`
	// detail of error
	Detail interface{} `json:"detail"`
}

type bodyValidationError struct {
	// field name
	StructField string `json:"field" example:"name"`
	// rule title that denied
	ActualTag string `json:"rule" example:"min"`
	// value of the field
	Value interface{} `json:"field_value"`
	// rule param
	Param string `json:"rule_param" example:"10"`
	// message param
	Message string `json:"message" example:"The name field is required."`
}

type routeValidationError struct {
	// field name
	StructField string `json:"field" example:"name"`
	// rule title that denied
	ActualTag string `json:"rule" example:"required"`
	// message param
	Message string `json:"message" example:"The selected item does not exist!"`
}

type serverError struct {
	// message param
	Message string `json:"message" example:"The selected item does not exist!"`
	err     interface{}
}

type forbiddenError struct {
	// message param
	Message string `json:"message" example:"The selected item does not exist!"`
}

type unauthorizedError struct {
	// message param
	Message string `json:"message" example:"The selected item does not exist!"`
}
