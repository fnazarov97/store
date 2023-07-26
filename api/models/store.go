package models

type Store struct {
	Store_id   int     `json:"store_id"`
	Store_name string  `json:"store_name"`
	Phone      string `json:"phone"`
	Email      string  `json:"email"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_code   string  `json:"zip_code"`
}

type StorePrimaryKey struct {
	Store_id int `json:"store_id"`
}

type CreateStore struct {
	Store_name string  `json:"store_name"`
	Phone      string `json:"phone"`
	Email      string  `json:"email"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_code   string  `json:"zip_code"`
}

type UpdateStore struct {
	Store_id   int     `json:"store_id"`
	Store_name string  `json:"store_name"`
	Phone      string `json:"phone"`
	Email      string  `json:"email"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_code   string  `json:"zip_code"`
}

type GetListStoreRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStoreResponse struct {
	Count  int      `json:"count"`
	Stores []*Store `json:"stores"`
}
