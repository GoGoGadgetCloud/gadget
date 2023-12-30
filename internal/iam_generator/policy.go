package iam_generator

type (
	Effect  string
	Action  string
	Service string
)

var Allow Effect = "Allow"
var Deny Effect = "Deny"

var AllActions Action = "*"
var AllLambdaActions Action = "lambda:*"
var AssumeRole Action = "sts:AssumeRole"

var Lambda Service = "lambda.amazonaws.com"
var ApiGateway Service = "apigateway.amazonaws.com"

func GenerateAssumeRoleForService(service Service) map[string]interface{} {
	return map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			generatePrincipalStatement(Allow, AssumeRole, service),
		},
	}
}

func GeneratePermissionForReource(effect Effect, action Action, resource string) map[string]interface{} {
	return map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			generateResourceStatemenet(effect, action, resource),
		},
	}
}

func generateResourceStatemenet(effect Effect, action Action, resource string) map[string]interface{} {
	return map[string]interface{}{
		"Effect":   string(effect),
		"Action":   string(action),
		"Resource": resource,
	}
}

func generatePrincipalStatement(effect Effect, action Action, service Service) map[string]interface{} {
	return map[string]interface{}{
		"Effect": string(effect),
		"Action": string(action),
		"Principal": map[string]interface{}{
			"Service": string(service),
		},
	}
}
