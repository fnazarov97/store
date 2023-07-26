package models

type Category struct {
	Category_id   int    `json:"category_id"`
	Category_name string `json:"category_name"`
}

type CategoryPrimaryKey struct {
	Category_id int `json:"category_id"`
}

type CreateCategory struct {
	Category_name string `json:"category_name"`
}

type UpdateCategory struct {
	Category_id   int    `json:"category_id"`
	Category_name string `json:"category_name"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}
