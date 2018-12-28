// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-27 22:18:22.492990836 +0530 IST m=+0.027250624

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Beast the automatic deployment tool for backdoor",
        "title": "Beast API",
        "contact": {
            "name": "SDSLabs",
            "url": "https://chat.sdslabs.co",
            "email": "contact.sdslabs.co.in"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "beast.sdslabs.co",
    "basePath": "/",
    "paths": {
        "/api/config/reaload/": {
            "patch": {
                "description": "Populates beast gobal config map by reparsing the config file $HOME/.beast/config.toml.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Reloads any changes in beast global configuration, located at ~/.beast/config.toml.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/manage/all/": {
            "post": {
                "description": "Handles challenge management routes for all the challenges with actions which includes - DEPLOY, UNDEPLOY.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manage"
                ],
                "summary": "Handles challenge management actions for multiple(all) challenges.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Action for the challenge",
                        "name": "action",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/manage/challenge/": {
            "post": {
                "description": "Handles challenge management routes with actions which includes - DEPLOY, UNDEPLOY, PURGE.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manage"
                ],
                "summary": "Handles challenge management actions.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the challenge to be managed, here name is the unique identifier for challenge",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Action for the challenge",
                        "name": "action",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/manage/deploy/local": {
            "post": {
                "description": "Handles deployment of a challenge using the absolute directory path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manage"
                ],
                "summary": "Deploy a local challenge using the path provided in the post parameter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Challenge Directory",
                        "name": "challenge_dir",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/manage/static/": {
            "post": {
                "description": "Handles beast static content serving container routes.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manage"
                ],
                "summary": "Handles route related to beast static content serving container, takes action as route parameter and perform that action",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/remote/reset/": {
            "post": {
                "description": "Resets local copy of remote git directory, it first deletes the existing directory and then clone from the remote again.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Resets beast local copy of remote git repository.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/api/remote/sync/": {
            "post": {
                "description": "Syncs beasts local challenges database with the remote git repository(hack) the local copy of the challenge database is located at $HOME/.beast/remote/$REMOTE_NAME.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Syncs beast's local copy of remote git repository for challenges.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        },
        "/auth/": {
            "get": {
                "description": "Sends challenge for authorization of user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Handles challenge generation",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.AuthorizationChallengeResp"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            },
            "post": {
                "description": "JWT can be received by sending back correct answer to challenge",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Handles solution check and token production",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPAuthorizeResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.HTTPPlainResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AuthorizationChallengeResp": {
            "type": "object",
            "properties": {
                "challenge": {
                    "type": "string",
                    "example": "Challenge String"
                },
                "message": {
                    "type": "string",
                    "example": "Response message"
                }
            }
        },
        "api.HTTPAuthorizeResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Response message"
                },
                "token": {
                    "type": "string",
                    "example": "YOUR_AUTHENTICATION_TOKEN"
                }
            }
        },
        "api.HTTPPlainResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Messsage in response to your request"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}