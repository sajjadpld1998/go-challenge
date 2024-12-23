// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/segments": {
            "post": {
                "description": "Create new segment pair endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segment"
                ],
                "summary": "Add a new segment pair",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserSegmentPair"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseNoContent"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/error_handler.ResponseError"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/error_handler.ResponseErrorDetail"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "detail": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/error_handler.bodyValidationError"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/error_handler.ResponseError"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/error_handler.ResponseErrorDetail"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "detail": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/error_handler.serverError"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "error_handler.ResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code",
                    "type": "integer",
                    "example": 422
                },
                "error": {
                    "description": "content"
                },
                "status": {
                    "description": "status of opration success",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "error_handler.ResponseErrorDetail": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "detail of error"
                },
                "message": {
                    "description": "error message",
                    "type": "string",
                    "example": "Server Error"
                }
            }
        },
        "error_handler.bodyValidationError": {
            "type": "object",
            "properties": {
                "field": {
                    "description": "field name",
                    "type": "string",
                    "example": "name"
                },
                "field_value": {
                    "description": "value of the field"
                },
                "message": {
                    "description": "message param",
                    "type": "string",
                    "example": "The name field is required."
                },
                "rule": {
                    "description": "rule title that denied",
                    "type": "string",
                    "example": "min"
                },
                "rule_param": {
                    "description": "rule param",
                    "type": "string",
                    "example": "10"
                }
            }
        },
        "error_handler.serverError": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "message param",
                    "type": "string",
                    "example": "The selected item does not exist!"
                }
            }
        },
        "requests.UserSegmentPair": {
            "type": "object",
            "required": [
                "segment",
                "user_id"
            ],
            "properties": {
                "segment": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "user_id": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "responses.ResponseNoContent": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code",
                    "type": "integer",
                    "example": 200
                },
                "status": {
                    "description": "status of opration success",
                    "type": "boolean",
                    "example": true
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
