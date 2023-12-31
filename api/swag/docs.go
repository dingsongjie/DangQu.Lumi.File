// Code generated by swaggo/swag. DO NOT EDIT.

package swag

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
        "/Converter/GetFisrtImageByGavingKey": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "GetFisrtImageByGavingKey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetFisrtImageByGavingKey"
                ],
                "summary": "GetFisrtImageByGavingKey",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/converter.ConvertByGavingKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/converter.ConvertByGavingKeyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CommonErrorResponse"
                        }
                    }
                }
            }
        },
        "/Converter/GetPdfByGavingKey": {
            "post": {
                "description": "GetPdfByGavingKey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetPdfByGavingKey"
                ],
                "summary": "GetPdfByGavingKey",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/converter.ConvertByGavingKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/converter.ConvertByGavingKeyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CommonErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "converter.ConvertByGavingKeyRequest": {
            "type": "object",
            "required": [
                "items"
            ],
            "properties": {
                "items": {
                    "type": "array",
                    "maxItems": 10,
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/converter.ConvertByGavingKeyRequestItem"
                    }
                }
            }
        },
        "converter.ConvertByGavingKeyRequestItem": {
            "type": "object",
            "required": [
                "sourceKey",
                "targetKey"
            ],
            "properties": {
                "sourceKey": {
                    "type": "string"
                },
                "targetKey": {
                    "type": "string"
                }
            }
        },
        "converter.ConvertByGavingKeyResponse": {
            "type": "object",
            "required": [
                "isAllSucceed",
                "items"
            ],
            "properties": {
                "isAllSucceed": {
                    "type": "boolean"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/converter.ConvertByGavingKeyResponseItem"
                    }
                }
            }
        },
        "converter.ConvertByGavingKeyResponseItem": {
            "type": "object",
            "required": [
                "isSucceed",
                "sourceKey"
            ],
            "properties": {
                "isSucceed": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "sourceKey": {
                    "type": "string"
                }
            }
        },
        "models.AclError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "requestStack": {
                    "type": "string"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ValidationError"
                    }
                }
            }
        },
        "models.CommonErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.AclError"
                }
            }
        },
        "models.ValidationError": {
            "type": "object",
            "properties": {
                "members": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
