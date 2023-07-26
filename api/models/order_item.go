package models

type Order_item struct {
	Order_id    int      `json:"order_id"`
	Item_id     int      `json:"item_id"`
	Product_id  int      `json:"product_id"`
	ProductData *Product `json:"product_data"`
	Quantity    float64  `json:"quantity"`
	List_price  float64  `json:"list_price"`
	Discount    float64  `json:"discount"`
}

type Order_itemPrimaryKey struct {
	Order_id int `json:"order_id"`
}

type CreateOrder_item struct {
	Order_id   int     `json:"order_id"`
	Item_id    int     `json:"item_id"`
	Product_id int     `json:"product_id"`
	Quantity   float64 `json:"quantity"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}

type UpdateOrder_item struct {
	Order_id   int     `json:"order_id"`
	Item_id    int     `json:"item_id"`
	Product_id int     `json:"product_id"`
	Quantity   float64 `json:"quantity"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}

type GetListOrder_itemRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrder_itemResponse struct {
	Count       int           `json:"count"`
	Order_items []*Order_item `json:"order_items"`
}
