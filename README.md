{
  "openapi": "3.0.3",
  "info": {
    "title": "Category RESTful API",
    "description": "API Spec for Category RESTful API",
    "version": "1.0.0"
  },
  "servers": [
    {
      "url": "http://localhost:3000/api"
    }
  ],
  "paths": {
    "/categories": {
      "get": {
        "security": [{
          "CategoryAuth": []
        }],
        "tags": ["Category API"],
        "description": "List All Categories",
        "summary": "List All Categories",
        "responses": {
          "200": {
            "description": "Success get all categories",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "number"
                    },
                    "status": {
                      "type": "string"
                    },
                    "data": {
                      "type": "array",
                      "items": {
                        "$ref": "#/components/schemas/Category"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      },
      "post": {
        "security": [{
          "CategoryAuth": []
        }],
        "tags": ["Category API"],
        "description": "Create new Category",
        "summary": "Create new Category",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CreateOrUpdateCategory"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Success Create Category",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "number"
                    },
                    "status": {
                      "type": "string"
                    },
                    "data": {
                      "$ref": "#/components/schemas/Category"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/categories/{categoryId}": {
      "get": {
        "security": [{
          "CategoryAuth": []
        }],
        "tags": ["Category API"],
        "description": "Get Category by Id",
        "summary": "Get category by Id",
        "parameters": [
          {
            "name": "categoryId",
            "in": "path",
            "description": "Category Id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success get Category",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "number"
                    },
                    "status": {
                      "type": "string"
                    },
                    "data": {
                      "$ref": "#/components/schemas/Category"
                    }
                  }
                }
              }
            }
          }
        }
      },
      "put": {
        "security": [{
          "CategoryAuth": []
        }],
        "tags": ["Category API"],
        "summary": "Update Category by Id",
        "description": "Update Category by Id",
        "parameters": [
          {
            "name": "categoryId",
            "in": "path",
            "description": "Category Id"
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CreateOrUpdateCategory"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Success get Category",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "number"
                    },
                    "status": {
                      "type": "string"
                    },
                    "data": {
                      "$ref": "#/components/schemas/Category"
                    }
                  }
                }
              }
            }
          }
        }
      },
      "delete": {
        "security": [{
          "CategoryAuth": []
        }],
        "tags": ["Category API"],
        "description": "Delete Category by Id",
        "summary": "Delete category by Id",
        "parameters": [
          {
            "name": "categoryId",
            "in": "path",
            "description": "Category Id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success delete Category",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "number"
                    },
                    "status": {
                      "type": "string"
                    }
                  }
                }
              }
            }}
        }
      }
    }
  },
  "components": {
    "securitySchemes": {
      "CategoryAuth": {
        "type": "apiKey",
        "in": "header",
        "name": "X-API-Key",
        "description": "Authentication for Category API"
      }
    },
    "schemas": {
      "CreateOrUpdateCategory": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          }
        }
      },
      "Category": {
        "type": "object",
        "properties": {
          "id": {
            "type": "number"
          },
          "name": {
            "type": "string"
          }
        }
      }
    }
  }
}
