package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response leverages AWS Lambda Proxy Request
type Response events.APIGatewayProxyResponse

// Handler is lambda handler invoked by `lambda.Start`
func Handler(ctx context.Context) (Response, error) {
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "response body",
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
