package response

import (
	"encoding/json"
	"fmt"
	"time"
)

type RawResponse struct {
	Result string `json:"result"`
}

type ComputationResponse struct {
	Duration time.Duration
	Result   string
}

func errorToResult(err error) string {
	if err.Error() == "nats: timeout" {
		return "Timeout"
	}

	return err.Error()
}

func durationToMilliseconds(duration time.Duration) float64 {
	return duration.Seconds() * float64(time.Second/time.Millisecond)
}

func PrettyPrint(cr *ComputationResponse) {
	fmt.Printf("%s,%.3f\n", cr.Result, durationToMilliseconds(cr.Duration))
}

func CreateFromData(duration time.Duration, data []byte) *ComputationResponse {
	cr := RawResponse{}

	err := json.Unmarshal(data, &cr)

	if err != nil {
		return &ComputationResponse{Duration: duration, Result: err.Error()}
	}

	return &ComputationResponse{Duration: duration, Result: cr.Result}
}

func CreateFromError(duration time.Duration, err error) *ComputationResponse {
	result := errorToResult(err)
	return &ComputationResponse{Duration: duration, Result: result}
}
