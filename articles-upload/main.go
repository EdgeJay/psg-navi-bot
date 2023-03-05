package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/EdgeJay/psg-navi-bot/articles-upload/sqs"
)

type MyEvent struct {
	Name string `json:"name"`
}

func isRunningInLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		var parsed sqs.MessageBody
		if err := json.NewDecoder(strings.NewReader(message.Body)).Decode(&parsed); err != nil {
			log.Fatalln(err)
		}

		log.Printf("object s3://%s/%s added\n", parsed.Records[0].S3.Bucket.Name, parsed.Records[0].S3.Object.Key)
	}

	return nil
}

func main() {
	log.Printf("Start article-upload")
	lambda.Start(handler)
}
