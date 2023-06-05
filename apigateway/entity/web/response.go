package web

type ErrorResponse struct {
	Code   int
	Status string
	Error  error
}

type SuccessResponse struct {
	Code   int
	Status string
	Data   interface{}
}
