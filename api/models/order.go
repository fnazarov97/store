package models

type Order struct {
	Order_id      int          `json:"order_id"`
	Customer_id   int          `json:"customer_id"`
	CustomerData  *Customer    `json:"customer_data"`
	Order_status  int          `json:"order_status"`
	Order_date    interface{}  `json:"order_date"`
	Required_date interface{}  `json:"required_date"`
	Shipped_date  interface{}  `json:"shipped_date"`
	Store_id      int          `json:"store_id"`
	StoreData     *Store       `json:"store_data"`
	Staff_id      int          `json:"staff_id"`
	StaffData     *Staff       `json:"staff_data"`
	OrderItems    []*OrderItem `json:"order_items"`
}

type OrderPrimaryKey struct {
	Order_id int `json:"order_id"`
}

type CreateOrder struct {
	Customer_id   int         `json:"customer_id"`
	Order_status  int         `json:"order_status"`
	Order_date    interface{} `json:"order_date"`
	Required_date interface{} `json:"required_date"`
	Shipped_date  interface{} `json:"shipped_date"`
	Store_id      int         `json:"store_id"`
	Staff_id      int         `json:"staff_id"`
}

type UpdateOrder struct {
	Order_id      int         `json:"order_id"`
	Customer_id   int         `json:"customer_id"`
	Order_status  int         `json:"order_status"`
	Order_date    interface{} `json:"order_date"`
	Required_date interface{} `json:"required_date"`
	Shipped_date  interface{} `json:"shipped_date"`
	Store_id      int         `json:"store_id"`
	Staff_id      int         `json:"staff_id"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}
