// Code generated by go-swagger; DO NOT EDIT.

package rest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Kubernikus",
    "version": "1.0.0"
  },
  "paths": {
    "/api/": {
      "get": {
        "summary": "List available api versions",
        "operationId": "ListAPIVersions",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ApiVersions"
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/api/v1/clusters/": {
      "get": {
        "summary": "List available clusters",
        "operationId": "ListClusters",
        "security": [
          {
            "keystone": []
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Cluster"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "post": {
        "summary": "Create a cluster",
        "operationId": "CreateCluster",
        "security": [
          {
            "keystone": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      }
    },
    "/api/v1/clusters/{name}": {
      "get": {
        "summary": "Show the specified cluser",
        "operationId": "ShowCluster",
        "security": [
          {
            "keystone": []
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "put": {
        "summary": "Update the specified cluser",
        "operationId": "UpdateCluster",
        "security": [
          {
            "keystone": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "delete": {
        "summary": "Terminate the specified cluster",
        "operationId": "TerminateCluster",
        "security": [
          {
            "keystone": []
          }
        ],
        "responses": {
          "202": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/clusters/{name}/credentials": {
      "get": {
        "summary": "Get user specific credentials to access the cluster",
        "operationId": "GetClusterCredentials",
        "security": [
          {
            "keystone": []
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/clusters/{name}/info": {
      "get": {
        "summary": "Get user specific info about the cluster",
        "operationId": "GetClusterInfo",
        "security": [
          {
            "keystone": []
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ClusterInfo"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/info": {
      "get": {
        "summary": "Get info about Kubernikus",
        "operationId": "Info",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Info"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiVersions": {
      "required": [
        "versions"
      ],
      "properties": {
        "versions": {
          "description": "versions are the api versions that are available.",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Cluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "description": "name of the cluster",
          "type": "string",
          "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
        },
        "spec": {
          "type": "object",
          "properties": {
            "nodePools": {
              "type": "array",
              "items": {
                "type": "object",
                "required": [
                  "name",
                  "size",
                  "flavor"
                ],
                "properties": {
                  "flavor": {
                    "type": "string"
                  },
                  "image": {
                    "type": "string"
                  },
                  "name": {
                    "type": "string",
                    "pattern": "^[a-z]([a-z0-9]*)?$"
                  },
                  "size": {
                    "type": "integer",
                    "maximum": 127,
                    "minimum": 0
                  }
                }
              }
            }
          }
        },
        "status": {
          "type": "object",
          "properties": {
            "kluster": {
              "properties": {
                "message": {
                  "type": "string"
                },
                "state": {
                  "description": "status of the cluster",
                  "type": "string"
                }
              }
            },
            "nodePools": {
              "type": "array",
              "items": {
                "required": [
                  "name",
                  "size",
                  "running",
                  "healthy",
                  "schedulable"
                ],
                "properties": {
                  "healthy": {
                    "type": "integer"
                  },
                  "name": {
                    "type": "string"
                  },
                  "running": {
                    "type": "integer"
                  },
                  "schedulable": {
                    "type": "integer"
                  },
                  "size": {
                    "type": "integer"
                  }
                }
              }
            }
          },
          "readOnly": true
        }
      }
    },
    "ClusterInfo": {
      "properties": {
        "binaries": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "links": {
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "link": {
                      "type": "string"
                    },
                    "platform": {
                      "type": "string"
                    }
                  }
                }
              },
              "name": {
                "type": "string"
              }
            }
          }
        },
        "setupCommand": {
          "type": "string"
        }
      }
    },
    "Credentials": {
      "type": "object",
      "properties": {
        "kubeconfig": {
          "type": "string"
        }
      }
    },
    "Info": {
      "properties": {
        "version": {
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "account": {
          "description": "account id",
          "type": "string"
        },
        "authUrl": {
          "description": "Identity Endpoint",
          "type": "string"
        },
        "domain": {
          "description": "user's domain name",
          "type": "string"
        },
        "id": {
          "description": "userid",
          "type": "string"
        },
        "name": {
          "description": "username",
          "type": "string"
        },
        "roles": {
          "description": "list of roles the user has in the given scope",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "error": {
      "description": "the error model is a model for all the error responses coming from Kubernikus \n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "code": {
          "description": "The error code",
          "type": "integer"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "errorResponse": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "keystone": {
      "description": "OpenStack Keystone Authentication",
      "type": "apiKey",
      "name": "x-auth-token",
      "in": "header"
    }
  }
}`))
}
