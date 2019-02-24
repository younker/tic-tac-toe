package main

import (
    "encoding/json"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/younker/tic-tac-toe/internal/game"
)

type RequestBody struct {
    Board [9]int `json:"board"`
}

func main() {
    lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // Request Event
    // Exposing this via Amazon API Gateway, I enabled the Lambda Proxy
    // Integration setting for convenience. This means a couple of things but
    // the primary thing to remember is that it converts each HTTP request to
    // JSON and passes it to Lambda as the event object. It is from this event
    // that we get request context (body), stage vars, headers, params, etc. It
    // assumes nothing about the body (based on any aspect of the request, ie
    // content-type header) thus the request body is an escaped string, not
    // JSON.
    // https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html

    // [noob@] Unmarshal takes a slice of bytes so we need to convert that
    // before we can parse the json body.
    bs := []byte(request.Body)
    var body RequestBody
    err := json.Unmarshal(bs, &body)
    if err != nil {
        return events.APIGatewayProxyResponse{}, err
    }

    m := game.GetNextMove(body.Board, game.Bot)
    body.Board[m.Index] = m.Player

    // Again, the events api will assume nothing for us so we need to convert
    // our RequestBody to JSON
    resp, err := json.Marshal(body)
    if err != nil {
        return events.APIGatewayProxyResponse{}, err
    }

    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       string(resp),
    }, nil
}
