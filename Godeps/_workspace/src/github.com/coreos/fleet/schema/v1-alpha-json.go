package schema
//
// This file is automatically generated by contrib/schema-generator
//
// **** DO NOT EDIT ****
//
const DiscoveryJSON = `{
  "kind": "discovery#restDescription",
  "discoveryVersion": "v1-alpha",
  "id": "fleet:v1-alpha",
  "name": "schema",
  "version": "v1-alpha",
  "title": "Fleet API",
  "description": "",
  "documentationLink": "http://github.com/coreos/fleet",
  "protocol": "rest",
  "icons": {
    "x16": "",
    "x32": ""
  },
  "labels": [],
  "baseUrl": "$ENDPOINT/v1-alpha/",
  "basePath": "/v1-alpha/",
  "rootUrl": "$ENDPOINT/",
  "servicePath": "v1-alpha/",
  "batchPath": "batch",
  "parameters": {},
  "auth": {},
  "schemas": {
    "Machine": {
      "id": "Machine",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "primaryIP": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "properties": {},
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "MachinePage": {
      "id": "MachinePage",
      "type": "object",
      "properties": {
        "machines": {
          "type": "array",
          "items": {
            "$ref": "Machine"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "UnitOption": {
      "id": "UnitOption",
      "type": "object",
      "properties": {
        "section": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "Unit": {
      "id": "Unit",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "options": {
          "type": "array",
          "items": {
            "$ref": "UnitOption"
          }
        },
        "desiredState": {
          "type": "string",
          "enum": [
            "inactive",
            "loaded",
            "launched"
          ]
        },
        "currentState": {
          "type": "string",
          "enum": [
            "inactive",
            "loaded",
            "launched"
          ]
        },
        "machineID": {
          "type": "string",
          "required": true
        }
      }
    },
    "UnitPage": {
      "id": "UnitPage",
      "type": "object",
      "properties": {
        "units": {
          "type": "array",
          "items": {
            "$ref": "Unit"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "UnitState": {
      "id": "UnitState",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "hash": {
          "type": "string"
        },
        "machineID": {
          "type": "string"
        },
        "systemdLoadState": {
          "type": "string"
        },
        "systemdActiveState": {
          "type": "string"
        },
        "systemdSubState": {
          "type": "string"
        }
      }
    },
    "UnitStatePage": {
      "id": "UnitStatePage",
      "type": "object",
      "properties": {
        "states": {
          "type": "array",
          "items": {
            "$ref": "UnitState"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    }
  },
  "resources": {
    "Machines": {
      "methods": {
        "List": {
          "id": "fleet.Machine.List",
          "description": "Retrieve a page of Machine objects.",
          "httpMethod": "GET",
          "path": "machines",
          "parameters": {
            "nextPageToken": {
              "type": "string",
              "location": "query"
            }
          },
          "response": {
            "$ref": "MachinePage"
          }
        }
      }
    },
    "Units": {
      "methods": {
        "List": {
          "id": "fleet.Unit.List",
          "description": "Retrieve a page of Unit objects.",
          "httpMethod": "GET",
          "path": "units",
          "parameters": {
            "nextPageToken": {
              "type": "string",
              "location": "query"
            }
          },
          "response": {
            "$ref": "UnitPage"
          }
        },
        "Get": {
          "id": "fleet.Unit.Get",
          "description": "Retrieve a single Unit object.",
          "httpMethod": "GET",
          "path": "units/{unitName}",
          "parameters": {
            "unitName": {
              "type": "string",
              "location": "path",
              "required": true
            }
          },
          "parameterOrder": [
            "unitName"
          ],
          "response": {
            "$ref": "Unit"
          }
        },
        "Delete": {
          "id": "fleet.Unit.Delete",
          "description": "Delete the referenced Unit object.",
          "httpMethod": "DELETE",
          "path": "units/{unitName}",
          "parameters": {
            "unitName": {
              "type": "string",
              "location": "path",
              "required": true
            }
          },
          "parameterOrder": [
            "unitName"
          ]
        },
        "Set": {
          "id": "fleet.Unit.Set",
          "description": "Create or update a Unit.",
          "httpMethod": "PUT",
          "path": "units/{unitName}",
          "parameters": {
            "unitName": {
              "type": "string",
              "location": "path",
              "required": true
            }
          },
          "parameterOrder": [
            "unitName"
          ],
          "request": {
            "$ref": "Unit"
          }
        }
      }
    },
    "UnitState": {
      "methods": {
        "List": {
          "id": "fleet.UnitState.List",
          "description": "Retrieve a page of UnitState objects.",
          "httpMethod": "GET",
          "path": "state",
          "parameters": {
            "nextPageToken": {
              "type": "string",
              "location": "query"
            },
            "unitName": {
              "type": "string",
              "location": "query"
            },
            "machineID": {
              "type": "string",
              "location": "query"
            }
          },
          "response": {
            "$ref": "UnitStatePage"
          }
        }
      }
    }
  }
}
`