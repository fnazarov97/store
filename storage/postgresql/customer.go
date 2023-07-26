package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type CustomerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) Create(ctx context.Context, req *models.CreateCustomer) (string, error) {

	query := `
		INSERT INTO customers(
			first_name,
			last_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8) returning first_name
	`

	name, err := r.db.Exec(ctx, query,
		req.First_name,
		req.Last_name,
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

func (r *CustomerRepo) GetByID(ctx context.Context, req *models.CustomerPrimaryKey) (resp *models.Customer, err error) {
	resp = &models.Customer{}
	query := `
		SELECT
			COALESCE(customer_id, 0),
			COALESCE(first_name,''),
			COALESCE(last_name,''),
			COALESCE(phone,''),
			COALESCE(email,''),
			COALESCE(street,''),
			COALESCE(city,''),
			COALESCE(state,''),
			COALESCE(zip_code,0)
		FROM customers
		WHERE customer_id = $1
	`

	err = r.db.QueryRow(ctx, query, req.Customer_id).Scan(
		&resp.Customer_id,
		&resp.First_name,
		&resp.Last_name,
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

func (r *CustomerRepo) GetList(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) {

	resp = &models.GetListCustomerResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(customer_id, 0),
			COALESCE(first_name,''),
			COALESCE(last_name,''),
			COALESCE(phone,''),
			COALESCE(email,''),
			COALESCE(street,''),
			COALESCE(city,''),
			COALESCE(state,''),
			COALESCE(zip_code,0)
		FROM customers
	`

	if len(req.Search) > 0 {
		filter += " AND customer_name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var customer models.Customer
		err = rows.Scan(
			&resp.Count,
			&customer.Customer_id,
			&customer.First_name,
			&customer.Last_name,
			&customer.Phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.Zip_code,
		)
		if err != nil {
			return nil, err
		}
		resp.Customers = append(resp.Customers, &customer)
	}

	return resp, nil
}

func (r *CustomerRepo) Update(ctx context.Context, req *models.UpdateCustomer) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			customers
		SET
			first_name = :first_name,
			last_name = :last_name,
			phone = :phone,
			email = :email,
			street = :street,
			city = :city,
			state = :state,
			zip_code = :zip_code
		WHERE customer_id = :customer_id
	`

	params = map[string]interface{}{
		"customer_id": req.Customer_id,
		"first_name":  req.First_name,
		"last_name":   req.Last_name,
		"phone":       req.Phone,
		"email":       req.Email,
		"street":      req.Street,
		"city":        req.City,
		"state":       req.State,
		"zip_code":    req.Zip_code,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *CustomerRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			customers
		SET
		` + set + `
		WHERE customer_id = :id
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

func (r *CustomerRepo) Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM customers WHERE customer_id = $1", req.Customer_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), err
}
