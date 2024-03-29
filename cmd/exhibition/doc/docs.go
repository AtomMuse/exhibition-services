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
                "description": "Get a list of all exhibitions data is public only",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Get all exhibitions is public",
                "operationId": "GetExhibitionsIsPublic",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ResponseExhibition"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new exhibition data",
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
                        "name": "requestExhibition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestCreateExhibition"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseGetExhibitionId"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/all": {
            "get": {
                "description": "Get a list of all exhibitions data",
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ResponseExhibition"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/{id}": {
            "get": {
                "description": "Get exhibition data by exhibitionID",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibition"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update exhibition data by exhibitionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exhibitions"
                ],
                "summary": "Update exhibition by ID",
                "operationId": "UpdateExhibition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exhibition ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Exhibition data to update",
                        "name": "updateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestUpdateExhibition"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibition"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete exhibition data by exhibitionID",
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
                    "200": {
                        "description": "Delete Exhibition Success",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseGetExhibitionId"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/{id}/like": {
            "put": {
                "description": "Like exhibition by exhibitionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Like"
                ],
                "summary": "Like exhibition by ID",
                "operationId": "LikeExhibition",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibition"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/{id}/sections": {
            "get": {
                "description": "Get Sections By exhibitionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Get Sections By exhibitionID",
                "operationId": "GetSectionsByExhibitionID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ResponseExhibitionSection"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/exhibitions/{id}/unlike": {
            "put": {
                "description": "unlike exhibition by exhibitionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Unlike"
                ],
                "summary": "Unlike exhibition by ID",
                "operationId": "UnlikeExhibition",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibition"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/sections": {
            "post": {
                "description": "Create a new exhibitionSection data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Create a new exhibitionSection",
                "parameters": [
                    {
                        "description": "ExhibitionSection data to create",
                        "name": "requestExhibitionSection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestCreateExhibitionSection"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseGetExhibitionSectionId"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/sections/all": {
            "get": {
                "description": "Get a list of all exhibition sections data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Get all exhibitions sections",
                "operationId": "GetAllExhibitionSections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ResponseExhibitionSection"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        },
        "/api/sections/{id}": {
            "get": {
                "description": "Get exhibition data by sectionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Get exhibitionSection by ID",
                "operationId": "GetExhibitionSectionByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exhibition Section ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibitionSection"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update exhibitionSection data by sectionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Update exhibitionSection by sectionID",
                "operationId": "UpdateExhibitionSection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ExhibitionSection ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ExhibitionSection data to update",
                        "name": "updateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestUpdateExhibitionSection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseExhibition"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Section data by sectionID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sections"
                ],
                "summary": "Delete Section by ID",
                "operationId": "DeleteExhibitionSectionByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Section ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete Section Success",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseGetExhibitionId"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helper.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.APIError": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "model.CenterItem": {
            "type": "object",
            "properties": {
                "details": {
                    "$ref": "#/definitions/model.Details"
                },
                "previewType": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "model.Contents": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Details": {
            "type": "object",
            "properties": {
                "contents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Contents"
                    }
                },
                "img": {
                    "type": "string"
                }
            }
        },
        "model.ExhibitionSection": {
            "type": "object",
            "required": [
                "exhibitionId",
                "sectionType"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "background": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "exhibitionId": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "leftCol": {
                    "$ref": "#/definitions/model.LeftColumn"
                },
                "rightCol": {
                    "$ref": "#/definitions/model.RightColumn"
                },
                "sectionType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.LeftColumn": {
            "type": "object",
            "properties": {
                "contentType": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "imageDescription": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.LeftRightItem": {
            "type": "object",
            "properties": {
                "details": {
                    "$ref": "#/definitions/model.Details"
                },
                "previewType": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "model.RequestCreateExhibition": {
            "type": "object",
            "required": [
                "endDate",
                "exhibitionDescription",
                "exhibitionName",
                "isPublic",
                "layoutUsed",
                "startDate",
                "status",
                "userId"
            ],
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "exhibitionCategories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "exhibitionDescription": {
                    "type": "string"
                },
                "exhibitionName": {
                    "type": "string"
                },
                "exhibitionSectionsID": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "exhibitionTags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isPublic": {
                    "type": "boolean"
                },
                "layoutUsed": {
                    "type": "string"
                },
                "likeCount": {
                    "type": "integer"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Room"
                    }
                },
                "startDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "thumbnailImg": {
                    "type": "string"
                },
                "userId": {
                    "$ref": "#/definitions/model.UserID"
                },
                "visitedNumber": {
                    "type": "integer"
                }
            }
        },
        "model.RequestCreateExhibitionSection": {
            "type": "object",
            "required": [
                "exhibitionID",
                "sectionType"
            ],
            "properties": {
                "background": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "exhibitionID": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "leftCol": {
                    "$ref": "#/definitions/model.LeftColumn"
                },
                "rightCol": {
                    "$ref": "#/definitions/model.RightColumn"
                },
                "sectionType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.RequestUpdateExhibition": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "exhibitionCategories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "exhibitionDescription": {
                    "type": "string"
                },
                "exhibitionName": {
                    "type": "string"
                },
                "exhibitionSectionsID": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "exhibitionTags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isPublic": {
                    "type": "boolean"
                },
                "layoutUsed": {
                    "type": "string"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Room"
                    }
                },
                "startDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "thumbnailImg": {
                    "type": "string"
                },
                "userId": {
                    "$ref": "#/definitions/model.UserID"
                },
                "visitedNumber": {
                    "type": "integer"
                }
            }
        },
        "model.RequestUpdateExhibitionSection": {
            "type": "object",
            "required": [
                "exhibitionID",
                "sectionType"
            ],
            "properties": {
                "background": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "exhibitionID": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "leftCol": {
                    "$ref": "#/definitions/model.LeftColumn"
                },
                "rightCol": {
                    "$ref": "#/definitions/model.RightColumn"
                },
                "sectionType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ResponseExhibition": {
            "type": "object",
            "required": [
                "_id",
                "exhibitionName",
                "layoutUsed",
                "status",
                "userId"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "exhibitionCategories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "exhibitionDescription": {
                    "type": "string"
                },
                "exhibitionName": {
                    "type": "string"
                },
                "exhibitionSections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ExhibitionSection"
                    }
                },
                "exhibitionTags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isPublic": {
                    "type": "boolean"
                },
                "layoutUsed": {
                    "type": "string"
                },
                "likeCount": {
                    "type": "integer"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Room"
                    }
                },
                "startDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "thumbnailImg": {
                    "type": "string"
                },
                "userId": {
                    "$ref": "#/definitions/model.UserID"
                },
                "visitedNumber": {
                    "type": "integer"
                }
            }
        },
        "model.ResponseExhibitionSection": {
            "type": "object",
            "required": [
                "exhibitionID",
                "sectionType"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "background": {
                    "type": "string"
                },
                "contentType": {
                    "type": "string"
                },
                "exhibitionID": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "leftCol": {
                    "$ref": "#/definitions/model.LeftColumn"
                },
                "rightCol": {
                    "$ref": "#/definitions/model.RightColumn"
                },
                "sectionType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ResponseGetExhibitionId": {
            "type": "object",
            "required": [
                "_id"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                }
            }
        },
        "model.ResponseGetExhibitionSectionId": {
            "type": "object",
            "required": [
                "_id"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                }
            }
        },
        "model.RightColumn": {
            "type": "object",
            "properties": {
                "contentType": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "imageDescription": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Room": {
            "type": "object",
            "properties": {
                "center": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CenterItem"
                    }
                },
                "left": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LeftRightItem"
                    }
                },
                "mapThumbnail": {
                    "type": "string"
                },
                "right": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LeftRightItem"
                    }
                }
            }
        },
        "model.UserID": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
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
	Schemes:          []string{"http"},
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
