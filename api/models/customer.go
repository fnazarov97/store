package models

type Customer struct {
	Customer_id int     `json:"customer_id"`
	First_name  string  `json:"first_name"`
	Last_name   string  `json:"last_name"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zip_code    float64 `json:"zip_code"`
}

type CustomerPrimaryKey struct {
	Customer_id int `json:"customer_id"`
}

type CreateCustomer struct {
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_code   float64 `json:"zip_code"`
}

type UpdateCustomer struct {
	Customer_id int     `json:"customer_id"`
	First_name  string  `json:"first_name"`
	Last_name   string  `json:"last_name"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zip_code    float64 `json:"zip_code"`
}

type GetListCustomerRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCustomerResponse struct {
	Count     int         `json:"count"`
	Customers []*Customer `json:"customers"`
}
