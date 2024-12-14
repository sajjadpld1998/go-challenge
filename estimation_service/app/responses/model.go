package responses

type ResponseNoContent struct {
	// status of opration success
	Status bool `json:"status" example:"true"`
	// response code
	Code   int  `json:"code" example:"200"`
}

type ResponseContent struct {
	// status of opration success
	Status bool        `json:"status" example:"false"`
	// response code
	Code   int         `json:"code" example:"422"`
	// content
	Data   interface{} `json:"data"`
}
