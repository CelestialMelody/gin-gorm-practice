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
        "contact": {
            "name": "API Support"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get multiple articles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Keyword",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Edit an article",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Desc",
                        "name": "desc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "ModifiedBy",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "State",
                        "name": "state",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"ok\",\"data\":{},\"article_msg\":\"edit an article\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add an article",
                "parameters": [
                    {
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Desc",
                        "name": "desc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "State",
                        "name": "state",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get a single article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code:200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an article",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "description": "获取所有标签",
                "produces": [
                    "application/json"
                ],
                "summary": "获取所有标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "添加多个文章标签",
                "produces": [
                    "application/json"
                ],
                "summary": "添加多个文章标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "description": "编辑多个文章标签",
                "produces": [
                    "application/json"
                ],
                "summary": "编辑多个文章标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ModifiedBy",
                        "name": "modified_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除多个文章标签",
                "summary": "DeleteTags",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
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
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gin-gorm-practice",
	Description:      "gin-gorm-practice",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
