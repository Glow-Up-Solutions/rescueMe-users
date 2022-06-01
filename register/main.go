package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Greeting string `json:"greeting"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var nick Message
	_ = json.Unmarshal([]byte(request.Body), &nick)

	newMes := Message{
		Nickname: nick.Nickname,
		Name:     "Jonathan Espinosa",
		Age:      25,
		Greeting: "Handshake Update",
	}
	a, _ := json.Marshal(newMes)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(a),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
