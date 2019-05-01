package main

import (
	"endpoint/cli"
	"endpoint/client"
	"endpoint/request"
	"endpoint/response"
	"time"
)

func main() {
	logErrors := true
	bufferSize := 10
	timeout := 2 * time.Second
	natsUrl := "nats://nats:4222"
	requestSubject := "worker-computation-requests"
	responseSubject := "worker-computation-responses"

	parseOptions := cli.ParseOptions{
		LogErrors: logErrors,
		BufferSize: bufferSize,
	}

	clientOptions := client.ClientOptions{
		NatsUrl: natsUrl,
		RequestSubject: requestSubject,
		ResponseSubject: responseSubject,
		Timeout: timeout,
		LogErrors: true,
	}

	client.Open(&clientOptions)
	client.CreateSyncSubscription()

	input := cli.ParseInput(&parseOptions)
	expressions := request.FromStringSlice(input)

	for _, expression := range expressions {
		start := time.Now()
		client.Publish(expression)
		msg, err := client.GetMessage()
		end := time.Now()
		dur := end.Sub(start)

		if err == nil {
			res := response.CreateFromData(dur, msg.Data)
			response.PrettyPrint(res)
		} else {
			errRes := response.CreateFromError(dur, err)
			response.PrettyPrint(errRes)
		}
	}

	client.Close()
}
