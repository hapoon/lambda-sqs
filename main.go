package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	SqsURL string
)

func main() {
	lambda.Start(HandleRequest)
}

type LambdaEvent struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}

func init() {
	SqsURL = os.Getenv("SQS_URL")
	if len(SqsURL) == 0 {
		log.Fatal("SQS_URL is empty")
	}
}

func HandleRequest(ctx context.Context, event LambdaEvent) (string, error) {
	sess := session.Must(session.NewSession())
	sqsSvc := sqs.New(sess)
	sqsSvc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Key1": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Sample value1"),
			},
			"Key2": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("111"),
			},
		},
		MessageBody: aws.String("Sample message body"),
		QueueUrl:    &SqsURL,
	})

	return "ok", nil
}
