package web

type BookSuccess struct {
	Code   int         `json:"code" form:"code"`
	Status string      `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
}

type BookFailure struct {
	Code   int    `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
}

type BookResponse struct {
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`
	Author string `json:"author" form:"author"`
}

type SellerResponse struct {
	Name string `json:"name" form:"name"`
}
