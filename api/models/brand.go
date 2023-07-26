package models

type Brand struct {
	Brand_id   int    `json:"brand_id"`
	Brand_name string `json:"brand_name"`
}

type BrandPrimaryKey struct {
	Brand_id int `json:"brand_id"`
}

type CreateBrand struct {
	Brand_name string `json:"brand_name"`
}

type UpdateBrand struct {
	Brand_id   int    `json:"brand_id"`
	Brand_name string `json:"brand_name"`
}

type GetListBrandRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListBrandResponse struct {
	Count  int      `json:"count"`
	Brands []*Brand `json:"brands"`
}
