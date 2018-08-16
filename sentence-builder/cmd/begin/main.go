package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/beeceej/sls-go-examples/sentence-builder/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type response struct {
	events.APIGatewayProxyResponse
}

func handler(context context.Context, event events.APIGatewayProxyRequest) (*response, error) {
	if event.HTTPMethod != "POST" {
		return &response{
			APIGatewayProxyResponse: events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Invalid HTTP Method",
			},
		}, nil
	}

	var req model.Data

	json.Unmarshal([]byte(event.Body), &req)

	session := session.Must(session.NewSession(&aws.Config{}))
	svc := sfn.New(session)

	out, err := svc.StartExecution(&sfn.StartExecutionInput{
		Name:            aws.String(uuid.NewV4().String()),
		Input:           aws.String(event.Body),
		StateMachineArn: aws.String(os.Getenv("STATE_MACHINE_ARN")),
	})

	return &response{
		APIGatewayProxyResponse: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       out.String(),
		},
	}, err
}

func main() {
	lambda.Start(handler)
}
