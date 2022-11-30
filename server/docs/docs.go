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
        "/api/player": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new player.",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestNewPlayer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/player/{player_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get player by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "player id",
                        "name": "player_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Player"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/room": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new room.",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestNewRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/room/{room_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get room by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/room/{room_id}/join": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add player to a room.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestJoinRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/room/{room_id}/leave": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove player from a room.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestLeaveRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/room/{room_id}/start": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Start game in a room.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.RequestNewGame"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/rooms": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get all rooms.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Room"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "entity.Player": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "entity.RequestJoinRoom": {
            "type": "object",
            "properties": {
                "player_id": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                }
            }
        },
        "entity.RequestLeaveRoom": {
            "type": "object",
            "properties": {
                "player_id": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                }
            }
        },
        "entity.RequestNewGame": {
            "type": "object",
            "properties": {
                "player_time_ms": {
                    "type": "integer"
                }
            }
        },
        "entity.RequestNewPlayer": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                }
            }
        },
        "entity.RequestNewRoom": {
            "type": "object",
            "properties": {
                "room_name": {
                    "type": "string"
                }
            }
        },
        "entity.Room": {
            "type": "object",
            "properties": {
                "game_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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