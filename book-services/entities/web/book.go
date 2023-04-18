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
	Id    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
type BookDeleteRespons struct {
	Id uint `json:"id" form:"id"`
}

type SellerResponse struct {
	Name string `json:"name" form:"name"`
}
