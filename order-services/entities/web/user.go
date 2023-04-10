package web

type UserSuccess struct {
	Code   int         `json:"code" form:"code"`
	Status string      `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
}

type UserFailure struct {
	Code   int    `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
	Error  error  `json:"error" form:"error"`
}
type UserResponse struct {
	Name string `json:"name" form:"name"`
}
