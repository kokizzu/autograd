// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/submissions": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create a user submission",
                "operationId": "CreateSubmission",
                "parameters": [
                    {
                        "description": "submission request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpsvc.SubmissionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.SubmissionRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/submissions/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete a submission",
                "operationId": "DeleteSubmission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "submission id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.SubmissionRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.Error"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "post": {
                "description": "create a User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create (register) a User",
                "operationId": "CreateUser",
                "parameters": [
                    {
                        "description": "name",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpsvc.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.UserRes"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "login a User",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "login request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpsvc.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "your@email.com"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpsvc.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "httpsvc.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiredIn": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "tokenType": {
                    "type": "string"
                }
            }
        },
        "httpsvc.SubmissionReq": {
            "type": "object",
            "properties": {
                "assignmentID": {
                    "type": "string"
                },
                "fileURL": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "submittedBy": {
                    "type": "string"
                }
            }
        },
        "httpsvc.SubmissionRes": {
            "type": "object",
            "properties": {
                "assignmentID": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "feedback": {
                    "type": "string"
                },
                "fileURL": {
                    "type": "string"
                },
                "grade": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "submittedBy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "httpsvc.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "ADMIN",
                        "STUDENT"
                    ]
                }
            }
        },
        "httpsvc.UserRes": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Autograde API",
	Description: "API documentation for Autograde",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
