package models

type Product struct {
	Product_id   int       `json:"product_id"`
	Product_name string    `json:"product_name"`
	Brand_id     int       `json:"brand_id"`
	BrandData    *Brand    `json:"brand_data"`
	Category_id  int       `json:"category_id"`
	CategoryData *Category `json:"category_data"`
	Model_year   int       `json:"model_year"`
	List_price   float64   `json:"list_price"`
}

type ProductPrimaryKey struct {
	Product_id int `json:"product_id"`
}

type CreateProduct struct {
	Product_name string  `json:"product_name"`
	Brand_id     int     `json:"brand_id"`
	Category_id  int     `json:"category_id"`
	Model_year   int     `json:"model_year"`
	List_price   float64 `json:"list_price"`
}

type UpdateProduct struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Brand_id     int     `json:"brand_id"`
	Category_id  int     `json:"category_id"`
	Model_year   int     `json:"model_year"`
	List_price   float64 `json:"list_price"`
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
