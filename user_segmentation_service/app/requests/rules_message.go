package requests

var Rules = make(map[string]string)

func initRules() {
	Rules["required"] = "The :attribute field is required."
	Rules["required_with"] = "The :attribute field is required when :value is present."
	Rules["max"] = "The :attribute may not be greater than :value."
	Rules["min"] = "The :attribute must be at least :value."
	Rules["json"] = "The :attribute must be a valid JSON string."
	Rules["dive"] = "The :attribute must be an array."
	Rules["numeric"] = "The :attribute must be a number."
	Rules["boolean"] = "The :attribute field must be true or false."
	Rules["oneof"] = "The :attribute field does not exist in :value."

	Rules["unique"] = "The :attribute has already been taken."
	Rules["exists"] = "The selected :attribute is invalid."
	Rules["isjson"] = "The :attribute must be a valid JSON string."
}
