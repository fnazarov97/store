package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type StaffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *StaffRepo {
	return &StaffRepo{
		db: db,
	}
}

func (r *StaffRepo) Create(ctx context.Context, req *models.CreateStaff) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO staffs(
			staff_id, 
			first_name,
			last_name,
			email,
			phone,
			active,
			store_id,
			manager_id
		)
		VALUES (
			(select max(staff_id)+1 as id from staffs), 
			$1, $2, $3, $4, $5, $6, $7) returning staff_id
	`

	err := r.db.QueryRow(ctx, query,
		req.First_name,
		req.Last_name,
		req.Email,
		req.Phone,
		req.Active,
		req.Store_id,
		req.Manager_id,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *StaffRepo) GetByID(ctx context.Context, req *models.StaffPrimaryKey) (resp *models.Staff, err error) {
	resp = &models.Staff{}
	query := `
		SELECT
			COALESCE(sta.staff_id, 0), 
			COALESCE(sta.first_name, ''),
			COALESCE(sta.last_name, ''),
			COALESCE(sta.email, ''),
			COALESCE(sta.phone, ''),
			COALESCE(sta.active, 0),
			COALESCE(sta.store_id, 0),

			COALESCE(sto.store_id, 0), 
			COALESCE(sto.store_name, ''),
			COALESCE(sto.phone, ''),
			COALESCE(sto.email, ''),
			COALESCE(sto.street, ''),
			COALESCE(sto.city, ''),
			COALESCE(sto.state, ''),
			COALESCE(sto.zip_code, ''),

			COALESCE(sta.manager_id, 0)
		FROM staffs as sta join stores as sto 
		ON sta.store_id = sto.store_id
		WHERE sta.staff_id = $1
	`
	resp.StoreData = &models.Store{}
	err = r.db.QueryRow(ctx, query, req.Staff_id).Scan(
		&resp.Staff_id,
		&resp.First_name,
		&resp.Last_name,
		&resp.Email,
		&resp.Phone,
		&resp.Active,
		&resp.Store_id,
		&resp.StoreData.Store_id,
		&resp.StoreData.Store_name,
		&resp.StoreData.Phone,
		&resp.StoreData.Email,
		&resp.StoreData.Street,
		&resp.StoreData.City,
		&resp.StoreData.State,
		&resp.StoreData.Zip_code,
		&resp.Manager_id,
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *StaffRepo) GetList(ctx context.Context, req *models.GetListStaffRequest) (resp *models.GetListStaffResponse, err error) {

	resp = &models.GetListStaffResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(staff_id, 0), 
			COALESCE(first_name, ''),
			COALESCE(last_name, ''),
			COALESCE(email, ''),
			COALESCE(phone, ''),
			COALESCE(active, 0),
			COALESCE(store_id, 0),
			COALESCE(manager_id, 0)
		FROM staffs
	`

	if len(req.Search) > 0 {
		filter += " AND staff_name ILIKE '%' || '" + req.Search + "' || '%' "
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

		var staff models.Staff
		err = rows.Scan(
			&resp.Count,
			&staff.Staff_id,
			&staff.First_name,
			&staff.Last_name,
			&staff.Email,
			&staff.Phone,
			&staff.Active,
			&staff.Store_id,
			&staff.Manager_id,
		)
		if err != nil {
			return nil, err
		}
		resp.Staffs = append(resp.Staffs, &staff)
	}

	return resp, nil
}

func (r *StaffRepo) Update(ctx context.Context, req *models.UpdateStaff) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staffs
		SET
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			phone = :phone,
			active = :active,
			store_id = :store_id,
			manager_id = :manager_id
		WHERE staff_id = :staff_id
	`

	params = map[string]interface{}{
		"staff_id":   req.Staff_id,
		"first_name": req.First_name,
		"last_name":  req.Last_name,
		"email":      req.Email,
		"phone":      req.Phone,
		"active":     req.Active,
		"store_id":   req.Store_id,
		"manager_id": req.Manager_id,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StaffRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			staffs
		SET
		` + set + `
		WHERE staff_id = :id
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

func (r *StaffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM staffs WHERE staff_id = $1", req.Staff_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
