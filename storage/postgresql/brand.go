package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type BrandRepo struct {
	db *pgxpool.Pool
}

func NewBrandRepo(db *pgxpool.Pool) *BrandRepo {
	return &BrandRepo{
		db: db,
	}
}

func (r *BrandRepo) Create(ctx context.Context, req *models.CreateBrand) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO brands(
			brand_id, 
			brand_name
		)
		VALUES ((select max(brand_id)+1 as id from brands), $1) returning brand_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Brand_name,
	).Scan(&id)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (r *BrandRepo) GetByID(ctx context.Context, req *models.BrandPrimaryKey) (*models.Brand, error) {
	var (
		query      string
		brand_id   sql.NullString
		brand_name sql.NullString
	)

	query = `
		SELECT
			brand_id,
			brand_name
		FROM brands
		WHERE brand_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Brand_id).Scan(
		&brand_id,
		&brand_name,
	)

	if err != nil {
		return nil, err
	}
	brand_i, _ := strconv.Atoi(brand_id.String)
	return &models.Brand{
		Brand_id:   brand_i,
		Brand_name: brand_name.String,
	}, nil
}

func (r *BrandRepo) GetList(ctx context.Context, req *models.GetListBrandRequest) (resp *models.GetListBrandResponse, err error) {

	resp = &models.GetListBrandResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			brand_id,
			brand_name
		FROM brands
	`

	if len(req.Search) > 0 {
		filter += " AND brand_name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var brand models.Brand
		var id int
		err = rows.Scan(
			&resp.Count,
			&id,
			&brand.Brand_name,
		)
		brand.Brand_id = id
		if err != nil {
			return nil, err
		}
		resp.Brands = append(resp.Brands, &brand)
	}

	return resp, nil
}

func (r *BrandRepo) Update(ctx context.Context, req *models.UpdateBrand) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			brands
		SET
			brand_name = :brand_name
		WHERE brand_id = :brand_id
	`

	params = map[string]interface{}{
		"brand_id":   req.Brand_id,
		"brand_name": req.Brand_name,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *BrandRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			brands
		SET
		` + set + `
		WHERE brand_id = :id
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

func (r *BrandRepo) Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM brands WHERE brand_id = $1", req.Brand_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
