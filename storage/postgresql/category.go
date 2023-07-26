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

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO categories(
			category_id, 
			category_name
		)
		VALUES ((select max(category_id)+1 as id from categories), $1) returning category_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Category_name,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *CategoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query         string
		category_id   sql.NullString
		category_name sql.NullString
	)

	query = `
		SELECT
			category_id,
			category_name
		FROM categories
		WHERE category_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Category_id).Scan(
		&category_id,
		&category_name,
	)

	if err != nil {
		return nil, err
	}
	category_i, _ := strconv.Atoi(category_id.String)
	return &models.Category{
		Category_id:   category_i,
		Category_name: category_name.String,
	}, nil
}

func (r *CategoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {

	resp = &models.GetListCategoryResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			category_id,
			category_name
		FROM categories
	`

	if len(req.Search) > 0 {
		filter += " AND category_name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var category models.Category
		var id int
		err = rows.Scan(
			&resp.Count,
			&id,
			&category.Category_name,
		)
		category.Category_id = id
		if err != nil {
			return nil, err
		}
		resp.Categories = append(resp.Categories, &category)
	}

	return resp, nil
}

func (r *CategoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			categories
		SET
			category_name = :category_name
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"category_id":   req.Category_id,
		"category_name": req.Category_name,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *CategoryRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			categories
		SET
		` + set + `
		WHERE category_id = :id
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

func (r *CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM categories WHERE category_id = $1", req.Category_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
