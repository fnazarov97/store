package models

type OrderItem struct {
	Order_id   int     `json:"order_id"`
	Item_id    int     `json:"item_id"`
	Product_id int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}

type OrderItemPrimaryKey struct {
	Order_id int `json:"order_id"`
	Item_id  int `json:"item_id"`
}
