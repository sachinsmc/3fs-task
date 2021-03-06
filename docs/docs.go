// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/v1/groups": {
            "get": {
                "description": "Get group list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Get all groups",
                "operationId": "get-groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "add group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Create group",
                "operationId": "create group",
                "parameters": [
                    {
                        "description": "Group",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/groups/{id}": {
            "put": {
                "description": "Update Group by group id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "update a group",
                "operationId": "update-group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Group",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete group by group id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "delete a group",
                "operationId": "delete-group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "Get users list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "operationId": "get-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "add user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "operationId": "get-users",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "put": {
                "description": "Update user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update a user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete a user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.CreateGroupRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserRequest": {
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
                }
            }
        },
        "models.Group": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "routes.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
