// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/login": {
            "post": {
                "description": "login User with \"email\" and \"password\" to get token for authentication user to use other endpoints",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "LOGIN USER",
                "parameters": [
                    {
                        "description": " ",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.PairToken"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "register User for use api",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "REGISTER USER",
                "parameters": [
                    {
                        "description": " ",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.PairToken"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/category/create": {
            "post": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "create category based on params. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "CREATE CATEGORY",
                "parameters": [
                    {
                        "description": " ",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateCategoryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/category/getall": {
            "get": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "get all category. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "GETALL CATEGORY",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "delete": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "delete category based on category Id. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "DELETE CATEGORY",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/todo/create": {
            "post": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "create todo based on params. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "CREATE TODO",
                "parameters": [
                    {
                        "description": " ",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateTodoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/todo/getall": {
            "get": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "get all todo based on offset for pagination. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "GETALL TODO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "minimum number for offset is '0', defualt limit is '5'",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        },
        "/todo/{id}": {
            "put": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "update todo based on todo Id. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "UPDATE TODO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateTodoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "delete todo based on todo Id. (auth required)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "DELETE TODO",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/delivery.jsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.jsonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "types.CreateCategoryReq": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Work"
                }
            }
        },
        "types.CreateTodoReq": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer",
                    "example": 2
                },
                "description": {
                    "type": "string",
                    "example": "solve all problems in chapter 2"
                },
                "title": {
                    "type": "string",
                    "example": "do homeWork"
                }
            }
        },
        "types.LoginReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@gmai.com"
                },
                "password": {
                    "type": "string",
                    "example": "111222333444"
                }
            }
        },
        "types.PairToken": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTY5Nzc2MjEsImZuYW1lIjoiU2hhbiIsImxuYW1lIjoiVml2IiwidXNlciI6ImFzZEB0ZXN0LmNvbSJ9.tdhUL-KpDmzSNtV9z6XhUgoTKcVabuOPS3fHAySjSXQ"
                },
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTY5Nzc2MjEsImZuYW1lIjoiU2hhbiIsImxuYW1lIjoiVml2IiwidXNlciI6ImFzZEB0ZXN0LmNvbSJ9.tdhUL-KpDmzSNtV9z6XhUgoTKcVabuOPS3fHAySjSXQ"
                }
            }
        },
        "types.RegisterReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@gmai.com"
                },
                "password": {
                    "type": "string",
                    "example": "111222333444"
                },
                "username": {
                    "type": "string",
                    "example": "example5040"
                }
            }
        },
        "types.UpdateTodoReq": {
            "type": "object",
            "properties": {
                "categoryid": {
                    "type": "integer",
                    "example": 2
                },
                "description": {
                    "type": "string",
                    "example": "solve all problems in chapter 2"
                },
                "due_date": {
                    "type": "string",
                    "example": "2022-09-25 20:35:01"
                },
                "priority": {
                    "type": "integer",
                    "example": 2
                },
                "status": {
                    "type": "boolean",
                    "example": true
                },
                "title": {
                    "type": "string",
                    "example": "do homework"
                }
            }
        }
    },
    "securityDefinitions": {
        "apiKey": {
            "type": "apiKey",
            "name": "Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Todo API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
