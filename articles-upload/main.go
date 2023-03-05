package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/EdgeJay/psg-navi-bot/articles-upload/articles"
	awsUtils "github.com/EdgeJay/psg-navi-bot/articles-upload/aws"
	"github.com/EdgeJay/psg-navi-bot/articles-upload/sqs"
)

func isRunningInLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func parseFile(ch chan *articles.Article, path awsUtils.S3Path) {
	reader := articles.NewReader(path)
	article, err := reader.LoadAndParseFile()
	if err != nil {
		log.Fatalln("unable to parse file", err)
	}
	log.Printf("s3://%s/%s loaded and parsed\n", *path.Bucket, *path.Key)
	ch <- article
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	parsedArticles := make([]*articles.Article, 0)
	parser := sqs.NewRecordsParser(sqsEvent.Records)
	paths := parser.GetS3Paths()
	ch := make(chan *articles.Article, len(paths))

	for _, path := range paths {
		go parseFile(ch, path)
	}

	for i := 0; i < len(paths); i += 1 {
		article := <-ch
		parsedArticles = append(parsedArticles, article)
	}

	log.Println("Handler execution done")

	return nil
}

func main() {
	log.Printf("Start article-upload")
	lambda.Start(handler)
}
