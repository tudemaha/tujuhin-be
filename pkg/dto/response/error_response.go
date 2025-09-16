package response

type ErrorResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ArrErrorResponse []ErrorResponse

func NewErrorResponseValue(key string, value string) ErrorResponse {
	return ErrorResponse{Key: key, Value: value}
}

func NewArrErrorResponse(errors ...ErrorResponse) ArrErrorResponse {
	arrError := ArrErrorResponse{}

	for _, v := range errors {
		arrError = append(arrError, v)
	}

	return arrError
}
