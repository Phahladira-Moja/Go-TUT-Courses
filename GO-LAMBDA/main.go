package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	log.Println("Hello World")

// 	response := events.APIGatewayProxyResponse{
// 		StatusCode: 200,
// 	}

// 	return response, nil
// }

func main() {
	lambda.Start(HandleLambdaEvent)
}
