package web

type BookSuccess struct {
	Code   int         `json:"code" form:"code"`
	Status string      `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
}

type BookFailuer struct {
	Code   int    `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
}

type BookResponse struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
