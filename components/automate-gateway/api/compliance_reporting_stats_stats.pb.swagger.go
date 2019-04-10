package api

func init() {
	Swagger.Add("compliance_reporting_stats_stats", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/compliance/reporting/stats/stats.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/compliance/reporting/stats/failures": {
      "post": {
        "operationId": "ReadFailures",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1Failures"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "StatsService"
        ]
      }
    },
    "/compliance/reporting/stats/profiles": {
      "post": {
        "summary": "should cover /profiles, profiles/:profile-id/summary, profiles/:profile-id/controls",
        "operationId": "ReadProfiles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1Profile"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "StatsService"
        ]
      }
    },
    "/compliance/reporting/stats/summary": {
      "post": {
        "summary": "should cover /summary, /summary/nodes, /summary/controls",
        "operationId": "ReadSummary",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1Summary"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "StatsService"
        ]
      }
    },
    "/compliance/reporting/stats/trend": {
      "post": {
        "summary": "should cover /trend/nodes, /trend/controls",
        "operationId": "ReadTrend",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1Trends"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "StatsService"
        ]
      }
    }
  },
  "definitions": {
    "QueryOrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "v1ControlStats": {
      "type": "object",
      "properties": {
        "control": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "passed": {
          "type": "integer",
          "format": "int32"
        },
        "failed": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        },
        "impact": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "v1ControlsSummary": {
      "type": "object",
      "properties": {
        "failures": {
          "type": "integer",
          "format": "int32"
        },
        "majors": {
          "type": "integer",
          "format": "int32"
        },
        "minors": {
          "type": "integer",
          "format": "int32"
        },
        "criticals": {
          "type": "integer",
          "format": "int32"
        },
        "passed": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1FailureSummary": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "failures": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string"
        },
        "profile": {
          "type": "string"
        }
      }
    },
    "v1Failures": {
      "type": "object",
      "properties": {
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1FailureSummary"
          }
        },
        "platforms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1FailureSummary"
          }
        },
        "controls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1FailureSummary"
          }
        },
        "environments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1FailureSummary"
          }
        }
      }
    },
    "v1ListFilter": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "type": {
          "type": "string"
        }
      }
    },
    "v1NodeSummary": {
      "type": "object",
      "properties": {
        "compliant": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        },
        "noncompliant": {
          "type": "integer",
          "format": "int32"
        },
        "high_risk": {
          "type": "integer",
          "format": "int32"
        },
        "medium_risk": {
          "type": "integer",
          "format": "int32"
        },
        "low_risk": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1Profile": {
      "type": "object",
      "properties": {
        "profile_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ProfileList"
          }
        },
        "profile_summary": {
          "$ref": "#/definitions/v1ProfileSummary"
        },
        "control_stats": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ControlStats"
          }
        }
      }
    },
    "v1ProfileList": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "failures": {
          "type": "integer",
          "format": "int32"
        },
        "majors": {
          "type": "integer",
          "format": "int32"
        },
        "minors": {
          "type": "integer",
          "format": "int32"
        },
        "criticals": {
          "type": "integer",
          "format": "int32"
        },
        "passed": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1ProfileSummary": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "license": {
          "type": "string"
        },
        "maintainer": {
          "type": "string"
        },
        "copyright": {
          "type": "string"
        },
        "copyright_email": {
          "type": "string"
        },
        "summary": {
          "type": "string"
        },
        "supports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Support"
          }
        },
        "stats": {
          "$ref": "#/definitions/v1ProfileSummaryStats"
        }
      }
    },
    "v1ProfileSummaryStats": {
      "type": "object",
      "properties": {
        "failed": {
          "type": "integer",
          "format": "int32"
        },
        "passed": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        },
        "failed_nodes": {
          "type": "integer",
          "format": "int32"
        },
        "total_nodes": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1Query": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        },
        "interval": {
          "type": "integer",
          "format": "int32"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ListFilter"
          }
        },
        "order": {
          "$ref": "#/definitions/QueryOrderType"
        },
        "sort": {
          "type": "string"
        },
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "per_page": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1ReportSummary": {
      "type": "object",
      "properties": {
        "stats": {
          "$ref": "#/definitions/v1Stats"
        },
        "status": {
          "type": "string"
        },
        "duration": {
          "type": "number",
          "format": "double"
        },
        "start_date": {
          "type": "string"
        }
      }
    },
    "v1Stats": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "string",
          "format": "int64"
        },
        "platforms": {
          "type": "integer",
          "format": "int32"
        },
        "environments": {
          "type": "integer",
          "format": "int32"
        },
        "profiles": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1Summary": {
      "type": "object",
      "properties": {
        "controls_summary": {
          "$ref": "#/definitions/v1ControlsSummary"
        },
        "node_summary": {
          "$ref": "#/definitions/v1NodeSummary"
        },
        "report_summary": {
          "$ref": "#/definitions/v1ReportSummary"
        }
      }
    },
    "v1Support": {
      "type": "object",
      "properties": {
        "os_name": {
          "type": "string"
        },
        "os_family": {
          "type": "string"
        },
        "release": {
          "type": "string"
        },
        "inspec_version": {
          "type": "string"
        },
        "platform_name": {
          "type": "string"
        },
        "platform_family": {
          "type": "string"
        },
        "platform": {
          "type": "string"
        }
      }
    },
    "v1Trend": {
      "type": "object",
      "properties": {
        "report_time": {
          "type": "string"
        },
        "passed": {
          "type": "integer",
          "format": "int32"
        },
        "failed": {
          "type": "integer",
          "format": "int32"
        },
        "skipped": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "v1Trends": {
      "type": "object",
      "properties": {
        "trends": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Trend"
          }
        }
      }
    }
  }
}
`)
}
