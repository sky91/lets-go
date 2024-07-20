module github.com/sky91/lets-go/aws/dynamox

go 1.21

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.32.8
	github.com/pkg/errors v0.9.1
	github.com/sky91/lets-go/gox v0.0.0-20240719115002-aab6624792e1
)

require github.com/aws/smithy-go v1.20.2 // indirect

replace github.com/sky91/lets-go/gox => ../../gox
