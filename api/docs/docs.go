// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-11-27 01:38:28.85774209 +0530 IST m=+0.199886916

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
        "/api/manage/all/": {
            "post": {
                "description": "Handles challenge management routes for all the challenges with actions which includes - DEPLOY, UNDEPLOY.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
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
                            "type": "JSON"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "type": "JSON"
                        }
                    }
                }
            }
        },
        "/api/manage/challenge/": {
            "post": {
                "description": "Handles challenge management routes with actions which includes - DEPLOY, UNDEPLOY.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
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
                            "type": "JSON"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "type": "JSON"
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
                            "type": "JSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "JSON"
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
                "summary": "Handles route related to beast static content serving container, takes action as route parameter and perform that action",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "JSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "JSON"
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
                            "type": "JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "JSON"
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
                            "type": "JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "JSON"
                        }
                    }
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
