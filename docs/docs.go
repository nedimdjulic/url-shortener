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
        "/count/{short_url}": {
            "get": {
                "description": "Returns the count of redirections from shortened to original URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get redirections count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.countRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.message"
                        }
                    }
                }
            }
        },
        "/create": {
            "post": {
                "description": "Returns a shortened URL mapped to the original URL sent in request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates a short URL",
                "parameters": [
                    {
                        "description": "Create new short URL",
                        "name": "url-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.createRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.message"
                        }
                    }
                }
            }
        },
        "/delete/{short_url}": {
            "delete": {
                "description": "Removes the URL mapping of short to original",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete short link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.message"
                        }
                    }
                }
            }
        },
        "/{short_url}": {
            "get": {
                "description": "Redirection the to original URL, provided the shortened link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Redirects to original URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.countRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "handlers.createReq": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "handlers.createRes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "shortened": {
                    "type": "string"
                }
            }
        },
        "handlers.message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
