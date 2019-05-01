package client 

import (
	"time"
)

type ClientOptions struct {
	NatsUrl string
	RequestSubject string
	ResponseSubject string
	LogErrors bool
	Timeout time.Duration
}
