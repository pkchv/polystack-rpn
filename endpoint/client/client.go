package client

import (
	"encoding/json"
	"errors"
	nats "github.com/nats-io/go-nats"
)

var client *nats.Conn
var subscription *nats.Subscription
var options *ClientOptions

func Open(opts *ClientOptions) error {
	options = opts
	conn, err := nats.Connect(options.NatsUrl)

	if err != nil {
		return err
	}

	client = conn

	return nil
}

func Publish(request interface{}) error {
	if client == nil {
		return errors.New("No open connection")
	}

	payload, err := json.Marshal(request)

	if err != nil {
		return errors.New("Cannot serialize struct")
	}

	errRequest := client.Publish(options.RequestSubject, payload)
	if errRequest != nil {
		return errRequest
	}

	return client.Flush()
}

func CreateSyncSubscription() error {
	sub, err := client.SubscribeSync(options.ResponseSubject)

	if err != nil {
		return err
	}

	subscription = sub

	return err
}

func GetMessage() (*nats.Msg, error) {

	if subscription == nil {
		return nil, errors.New("No subscription")
	}

	return subscription.NextMsg(options.Timeout)
}

func Close() {
	client.Close()
	client = nil
}
