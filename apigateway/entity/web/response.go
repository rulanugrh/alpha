package web

type ErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  error  `json:"error"`
}

type SuccessResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
