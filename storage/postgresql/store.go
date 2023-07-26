package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type StoreRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) *StoreRepo {
	return &StoreRepo{
		db: db,
	}
}

func (r *StoreRepo) Create(ctx context.Context, req *models.CreateStore) (string, error) {

	query := `
		INSERT INTO stores(
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8) returning store_name
	`

	name, err := r.db.Exec(ctx, query,
		req.Store_name,
		req.Phone,
		req.Email,
		req.Street,
		req.City,
		req.State,
		req.Zip_code,
	)

	if err != nil {
		return "", err
	}

	return string(name), nil
}

func (r *StoreRepo) GetByID(ctx context.Context, req *models.StorePrimaryKey) (resp *models.Store, err error) {
	resp = &models.Store{}
	query := `
		SELECT
			COALESCE(store_id, 0),
			COALESCE(store_name,''),
			COALESCE(phone,''),
			COALESCE(email,''),
			COALESCE(street,''),
			COALESCE(city,''),
			COALESCE(state,''),
			COALESCE(zip_code,'')
		FROM stores
		WHERE store_id = $1
	`

	err = r.db.QueryRow(ctx, query, req.Store_id).Scan(
		&resp.Store_id,
		&resp.Store_name,
		&resp.Phone,
		&resp.Email,
		&resp.Street,
		&resp.City,
		&resp.State,
		&resp.Zip_code,
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *StoreRepo) GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error) {

	resp = &models.GetListStoreResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(store_id, 0),
			COALESCE(store_name,''),
			COALESCE(phone,''),
			COALESCE(email,''),
			COALESCE(street,''),
			COALESCE(city,''),
			COALESCE(state,''),
			COALESCE(zip_code,'')
		FROM stores
	`

	if len(req.Search) > 0 {
		filter += " AND store_name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var store models.Store
		err = rows.Scan(
			&resp.Count,
			&store.Store_id,
			&store.Store_name,
			&store.Phone,
			&store.Email,
			&store.Street,
			&store.City,
			&store.State,
			&store.Zip_code,
		)
		if err != nil {
			return nil, err
		}
		resp.Stores = append(resp.Stores, &store)
	}

	return resp, nil
}

func (r *StoreRepo) Update(ctx context.Context, req *models.UpdateStore) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			stores
		SET
			store_name = :store_name,
			phone = :phone,
			email = :email,
			street = :street,
			city = :city,
			state = :state,
			zip_code = :zip_code
		WHERE store_id = :store_id
	`

	params = map[string]interface{}{
		"store_id":   req.Store_id,
		"first_name": req.Store_name,
		"phone":      req.Phone,
		"email":      req.Email,
		"street":     req.Street,
		"city":       req.City,
		"state":      req.State,
		"zip_code":   req.Zip_code,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StoreRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			stores
		SET
		` + set + `
		WHERE store_id = :id
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

func (r *StoreRepo) Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM stores WHERE store_id = $1", req.Store_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
