// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/purchase/get/{id}": {
            "get": {
                "description": "get purchase with id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "purchases"
                ],
                "summary": "Get purchase with id",
                "operationId": "api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Uid of purchase",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Delivery": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "chrt_id": {
                    "description": "PurchaseUid string ` + "`" + `json:\"-\"` + "`" + `",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nm_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rid": {
                    "type": "string"
                },
                "sale": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "bank": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "custom_fee": {
                    "type": "integer"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "goods_total": {
                    "type": "integer"
                },
                "payment_dt": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                }
            }
        },
        "models.PurchaseDTO": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "delivery": {
                    "$ref": "#/definitions/models.Delivery"
                },
                "delivery_service": {
                    "type": "string"
                },
                "entry": {
                    "type": "string"
                },
                "internal_signature": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "locale": {
                    "type": "string"
                },
                "oof_shard": {
                    "type": "string"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment": {
                    "$ref": "#/definitions/models.Payment"
                },
                "shardkey": {
                    "type": "string"
                },
                "sm_id": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Wildberries L0",
	Description:      "API for Getting purchases",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
