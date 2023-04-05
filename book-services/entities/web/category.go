package web

type CategorySuccess struct {
	Code   int         `json:"code" form:"code"`
	Status string      `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
}

type CategoryFailure struct {
	Code   int    `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
}

type CategoryResponse struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"desc" form:"desc"`
}
