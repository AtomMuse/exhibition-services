// Package doc Code generated by swaggo/swag. DO NOT EDIT
package doc

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
        "/api/exhibitions": {
            "get": {
                "description": "Get a list of all exhibitions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Get all exhibitions",
                "operationId": "GetAllExhibitions",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Permission denied",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new exhibition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Create a new exhibition",
                "parameters": [
                    {
                        "description": "Exhibition data to create",
                        "name": "template",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Permission denied",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/{id}": {
            "get": {
                "description": "Get exhibition details by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Get exhibition by ID",
                "operationId": "GetExhibitionByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exhibition ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Permission denied",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete exhibition by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Delete exhibition by ID",
                "operationId": "DeleteExhibition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exhibition ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Delete Exhibition Success"
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Permission denied",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "Exhibition Service API",
	Description:      "Exhibition Service สำหรับขอจัดการเกี่ยวกับ Exhibition ทั้งการสร้าง แก้ไข ลบ exhibition",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}