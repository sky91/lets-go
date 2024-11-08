module github.com/sky91/lets-go/aws/dynamox

go 1.21

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.7.35
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.34.6
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.47.0
)

require (
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.15.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.22.5 // indirect
	github.com/aws/smithy-go v1.20.4 // indirect
	golang.org/x/text v0.16.0 // indirect
)

replace github.com/sky91/lets-go/gox => ../../gox
