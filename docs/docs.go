// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/advertisements/": {
            "get": {
                "description": "Get all advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements"
                ],
                "summary": "Get advertisement",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.advertisementRoutes"
                        }
                    }
                }
            }
        },
        "/api/v1/advertisements/:id": {
            "get": {
                "description": "Get advertisement by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements"
                ],
                "summary": "Get advertisement",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.advertisementRoutes"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.getAdvertisementResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.getAdvertisementResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.getAdvertisementResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/advertisements/create": {
            "post": {
                "description": "Create advertisement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements"
                ],
                "summary": "Create advertisement",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.advertisementRoutes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.createAdvertisementResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.createAdvertisementResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.createAdvertisementResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.advertisementRoutes": {
            "type": "object"
        },
        "v1.createAdvertisementResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.getAdvertisementResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pictures": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "price": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Advertisement Management Service",
	Description:      "Test task from avito.tech for a Backend developer trainee",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
