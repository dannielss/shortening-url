// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shorten": {
            "post": {
                "description": "Takes a long URL and returns a shortened version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Shorten a URL",
                "parameters": [
                    {
                        "description": "URL to shorten",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ShortUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UrlMapping"
                        }
                    }
                }
            }
        },
        "/{shorted_url}": {
            "get": {
                "description": "Takes a short URL and redirects to the original long URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Redirect to Long URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shortened URL",
                        "name": "shorted_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Redirects to the original URL"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ShortUrlRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "domain.UrlMapping": {
            "type": "object",
            "required": [
                "long_url",
                "short_url"
            ],
            "properties": {
                "long_url": {
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "URL Shortener api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
