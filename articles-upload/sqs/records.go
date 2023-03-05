package sqs

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"

	awsUtils "github.com/EdgeJay/psg-navi-bot/articles-upload/aws"
)

type RecordsParser struct {
	EventRecords []events.SQSMessage
}

func NewRecordsParser(records []events.SQSMessage) *RecordsParser {
	return &RecordsParser{
		EventRecords: records,
	}
}

func (r *RecordsParser) GetS3Paths() []awsUtils.S3Path {
	paths := make([]awsUtils.S3Path, 0)

	for _, message := range r.EventRecords {
		// fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		var parsed MessageBody
		if err := json.NewDecoder(strings.NewReader(message.Body)).Decode(&parsed); err == nil {
			paths = append(paths, awsUtils.S3Path{
				Bucket: aws.String(parsed.Records[0].S3.Bucket.Name),
				Key:    aws.String(parsed.Records[0].S3.Object.Key),
			})
		}
	}

	return paths
}
