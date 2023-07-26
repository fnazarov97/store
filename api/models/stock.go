package models

type Stock struct {
	Store_id    int         `json:"store_id"`
	StoreData   *Store      `json:"store_data"`
	Product_id  int         `json:"product_id"`
	ProductData *Product    `json:"product_data"`
	Quantity    interface{} `json:"quantity"`
}

type StockPrimaryKey struct {
	Store_id int `json:"store_id"`
}

type CreateStock struct {
	Store_id   int         `json:"store_id"`
	Product_id int         `json:"product_id"`
	Quantity   interface{} `json:"quantity"`
}

type UpdateStock struct {
	Store_id   int         `json:"store_id"`
	Product_id int         `json:"product_id"`
	Quantity   interface{} `json:"quantity"`
}

type GetListStockRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStockResponse struct {
	Count  int      `json:"count"`
	Stocks []*Stock `json:"stocks"`
}
