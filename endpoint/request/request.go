package request

type ComputationRequest struct {
	Expression string `json:"expression"`
}

func FromStringSlice(expressions []string) []ComputationRequest {
	requests := make([]ComputationRequest, 0)
	for _, expression := range expressions {
		requests = append(requests, ComputationRequest{
			Expression: expression,
		})
	}

	return requests
}
