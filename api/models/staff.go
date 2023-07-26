package models

type Staff struct {
	Staff_id   int    `json:"staff_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Active     int    `json:"active"`
	Store_id   int    `json:"store_id"`
	StoreData  *Store `json:"store_data"`
	Manager_id int    `json:"manager_id"`
}

type StaffPrimaryKey struct {
	Staff_id int `json:"staff_id"`
}

type CreateStaff struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Active     string `json:"active"`
	Store_id   int    `json:"store_id"`
	Manager_id int    `json:"manager_id"`
}

type UpdateStaff struct {
	Staff_id   int    `json:"staff_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Active     string `json:"active"`
	Store_id   int    `json:"store_id"`
	Manager_id int    `json:"manager_id"`
}

type GetListStaffRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStaffResponse struct {
	Count  int      `json:"count"`
	Staffs []*Staff `json:"staffs"`
}
