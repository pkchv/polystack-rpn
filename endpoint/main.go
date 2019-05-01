package main

import (
	"endpoint/config"
	"endpoint/cli"
	"endpoint/client"
	"endpoint/request"
	"endpoint/response"
	sw "github.com/fatih/stopwatch"
)

func main() {
	parseOptions := cli.ParseOptions{
		LogErrors: config.LogErrors,
		BufferSize: config.BufferSize,
	}

	clientOptions := client.ClientOptions{
		NatsUrl: config.NatsUrl,
		RequestSubject: config.RequestSubject,
		ResponseSubject: config.ResponseSubject,
		Timeout: config.Timeout,
		LogErrors: config.LogErrors,
	}

	client.Open(&clientOptions)
	client.CreateSyncSubscription()

	input := cli.ParseInput(&parseOptions)
	expressions := request.FromStringSlice(input)
	stopwatches := make([]*sw.Stopwatch, 0)

	for _, expression := range expressions {
		stopwatches = append(stopwatches, sw.Start(0))
		client.Publish(expression)	
	}

	for _, stopwatch := range stopwatches {
		msg, err := client.GetMessage()
		stopwatch.Stop()
		duration := stopwatch.ElapsedTime()

		if err == nil {
			res := response.CreateFromData(duration, msg.Data)
			response.PrettyPrint(res)
		} else {
			res := response.CreateFromError(duration, err)
			response.PrettyPrint(res)
		}
	}

	client.Close()
}
