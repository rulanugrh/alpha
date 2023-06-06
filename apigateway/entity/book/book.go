package book

type BookModel struct {
	Id         uint   `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Price      int    `json:"price" form:"price"`
	Stock      int    `json:"stock_quantity" form:"stock_quantity"`
	Author     string `json:"author_name" form:"author_name" validate:"required"`
	Examplar   int    `json:"examplar" form:"examplar" validate:"required"`
	SellerID   uint   `json:"seller_id" form:"seller_id"`
	CategoryID uint   `json:"category_id" form:"category_id"`
}

type Category struct {
	Id          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" name:"description"`
}
