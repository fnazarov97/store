package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO products(
			product_id, 
			product_name,
			brand_id,
			category_id,
			model_year,
			list_price
		)
		VALUES (
			(select max(product_id)+1 as id from products), 
			$1, $2, $3, $4, $5 ) returning product_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Product_name,
		req.Brand_id,
		req.Category_id,
		req.Model_year,
		req.List_price,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *ProductRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	var (
		query string
		resp  models.Product
	)

	query = `
		SELECT
			COALESCE(p.product_id, 0), 
			COALESCE(p.product_name, ''),
			COALESCE(p.brand_id, 0),
			COALESCE(b.brand_id, 0),

			COALESCE(b.brand_name, ''),
			COALESCE(p.category_id, 0),

			COALESCE(c.category_id, 0),
			COALESCE(c.category_name, ''),
			
			COALESCE(p.model_year, 0),
			COALESCE(p.list_price, 0)
		FROM brands as b natural join products as p
		natural join categories as c
		WHERE product_id = $1
	`
	resp.BrandData = &models.Brand{}
	resp.CategoryData = &models.Category{}
	err := r.db.QueryRow(ctx, query, req.Product_id).Scan(
		&resp.Product_id,
		&resp.Product_name,
		&resp.Brand_id,
		&resp.BrandData.Brand_id,
		&resp.BrandData.Brand_name,
		&resp.Category_id,
		&resp.CategoryData.Category_id,
		&resp.CategoryData.Category_name,
		&resp.Model_year,
		&resp.List_price,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *ProductRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(p.product_id, 0), 
			COALESCE(p.product_name, ''),
			COALESCE(p.brand_id, 0),
			COALESCE(b.brand_id, 0),

			COALESCE(b.brand_name, ''),
			COALESCE(p.category_id, 0),

			COALESCE(c.category_id, 0),
			COALESCE(c.category_name, ''),
			
			COALESCE(p.model_year, 0),
			COALESCE(p.list_price, 0)
		FROM brands as b natural join products as p
		natural join categories as c
	`

	if len(req.Search) > 0 {
		filter += " AND product_name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			product  models.Product
			brand    models.Brand
			category models.Category
		)

		var id int
		err = rows.Scan(
			&resp.Count,
			&product.Product_id,
			&product.Product_name,
			&product.Brand_id,
			&brand.Brand_id,
			&brand.Brand_name,
			&product.Category_id,
			&category.Category_id,
			&category.Category_name,
			&product.Model_year,
			&product.List_price,
		)
		product.Product_id = id
		product.BrandData = &brand
		product.CategoryData = &category
		if err != nil {
			return nil, err
		}
		resp.Products = append(resp.Products, &product)
	}

	return resp, nil
}

func (r *ProductRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			products
		SET
			product_name = :product_name,
			brand_id = :brand_id,
			category_id = :category_id,
			model_year = :model_year,
			list_price = :list_price
		WHERE product_id = :product_id
	`

	params = map[string]interface{}{
		"product_id":   req.Product_id,
		"product_name": req.Product_name,
		"brand_id":     req.Brand_id,
		"category_id":  req.Category_id,
		"model_year":   req.Model_year,
		"list_price":   req.List_price,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *ProductRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			products
		SET
		` + set + `
		WHERE product_id = :id
	`

	req.Fields["id"] = req.ID

	fmt.Println(req.Fields)

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *ProductRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM products WHERE product_id = $1", req.Product_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
