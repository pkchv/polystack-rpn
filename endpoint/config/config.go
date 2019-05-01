package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var BufferSize int
var LogErrors bool
var Timeout time.Duration
var NatsUrl string
var RequestSubject string
var ResponseSubject string

func init() {
	bufferSizeStr := optional("BUFFER_SIZE", "64")
	logErrorsStr := optional("LOG_ERRORS", "true")
	buffSize, buffSizeErr := strconv.Atoi(bufferSizeStr)
	abortIfError(buffSizeErr)
	BufferSize = buffSize
	LogErrors = toBool(logErrorsStr)
	timeout, timeoutErr := strconv.Atoi(bufferSizeStr)
	abortIfError(timeoutErr)
	Timeout = time.Duration(timeout) * time.Millisecond
	NatsUrl = required("NATS_URI")
	RequestSubject = required("REQ_SUBJECT")
	ResponseSubject = required("RES_SUBJECT")
}

func optional(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func required(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		fmt.Fprintf(os.Stderr, "Required environment variable is missing: %s\n", key)
		os.Exit(1)
	}

	return value
}

func toBool(value string) bool {
	if value == "true" {
		return true
	}

	return false
}

func abortIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
