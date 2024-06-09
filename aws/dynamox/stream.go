package dynamox

import "github.com/aws/aws-lambda-go/events"

type StreamEvent struct {
	Records []StreamEventRecord `json:"Records"`
}

type StreamEventRecord struct {
	AWSRegion      string                       `json:"awsRegion"`
	Change         StreamRecord                 `json:"dynamodb"`
	EventID        string                       `json:"eventID"`
	EventName      string                       `json:"eventName"`
	EventSource    string                       `json:"eventSource"`
	EventVersion   string                       `json:"eventVersion"`
	EventSourceArn string                       `json:"eventSourceARN"`
	UserIdentity   *events.DynamoDBUserIdentity `json:"userIdentity,omitempty"`
}

type StreamRecord struct {
	ApproximateCreationDateTime events.SecondsEpochTime `json:"ApproximateCreationDateTime,omitempty"`
	Keys                        Record                  `json:"Keys,omitempty"`
	NewImage                    Record                  `json:"NewImage,omitempty"`
	OldImage                    Record                  `json:"OldImage,omitempty"`
	SequenceNumber              string                  `json:"SequenceNumber"`
	SizeBytes                   int64                   `json:"SizeBytes"`
	StreamViewType              string                  `json:"StreamViewType"`
}
