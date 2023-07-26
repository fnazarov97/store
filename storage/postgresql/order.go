package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO orders(
			order_id, 
			customer_id,
			order_status,
			order_date,
			required_date,
			shipped_date,
			store_id,
			staff_id
		)
		VALUES (
			(select max(order_id)+1 as id from orders), 
			$1, $2, $3, $4, $5, $6, $7 ) returning order_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Customer_id,
		req.Order_status,
		req.Order_date,
		req.Required_date,
		req.Shipped_date,
		req.Store_id,
		req.Staff_id,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *OrderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (resp *models.Order, err error) {
	resp = &models.Order{}
	q1 := ` 
		SELECT	
			order_id,
			item_id,
			product_id,
			quantity,
			list_price,
			discount
		FROM order_items
	`
	id := strconv.Itoa(req.Order_id)
	wh := " WHERE order_id = " + id
	q1 = q1 + wh
	rows, err := r.db.Query(ctx, q1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			orderItem models.OrderItem
		)
		err = rows.Scan(
			&orderItem.Order_id,
			&orderItem.Item_id,
			&orderItem.Product_id,
			&orderItem.Quantity,
			&orderItem.List_price,
			&orderItem.Discount,
		)

		if err != nil {
			return nil, err
		}
		resp.OrderItems = append(resp.OrderItems, &orderItem)
	}
	//-----------------------
	query := `
		SELECT
			COALESCE(o.order_id, 0), 
			COALESCE(o.customer_id, 0),

			COALESCE(c.customer_id, 0),
			COALESCE(c.first_name,''),
			COALESCE(c.last_name,''),
			COALESCE(c.phone,''),
			COALESCE(c.email,''),
			COALESCE(c.street,''),
			COALESCE(c.city,''),
			COALESCE(c.state,''),
			COALESCE(c.zip_code,0),

			COALESCE(o.order_status, 0), 
			o.order_date,
			o.required_date,
			o.shipped_date,
			COALESCE(o.store_id, 0),

			COALESCE(sto.store_id, 0),
			COALESCE(sto.store_name,''),
			COALESCE(sto.phone,''),
			COALESCE(sto.email,''),
			COALESCE(sto.street,''),
			COALESCE(sto.city,''),
			COALESCE(sto.state,''),
			COALESCE(sto.zip_code,''),

			COALESCE(o.staff_id, 0),

			COALESCE(sta.staff_id, 0), 
			COALESCE(sta.first_name, ''),
			COALESCE(sta.last_name, ''),
			COALESCE(sta.email, ''),
			COALESCE(sta.phone, ''),
			COALESCE(sta.active, 0),
			COALESCE(sta.store_id, 0),
			COALESCE(sta.manager_id, 0)
		FROM orders as o join customers as c 
		ON o.customer_id = c.customer_id join stores as sto 
		ON o.store_id = sto.store_id join staffs as sta
		ON o.staff_id = sta.staff_id
		WHERE o.order_id = $1
	`
	resp.CustomerData = &models.Customer{}
	resp.StoreData = &models.Store{}
	resp.StaffData = &models.Staff{}
	err = r.db.QueryRow(ctx, query, req.Order_id).Scan(
		&resp.Order_id,
		&resp.Customer_id,
		&resp.CustomerData.Customer_id,
		&resp.CustomerData.First_name,
		&resp.CustomerData.Last_name,
		&resp.CustomerData.Phone,
		&resp.CustomerData.Email,
		&resp.CustomerData.Street,
		&resp.CustomerData.City,
		&resp.CustomerData.State,
		&resp.CustomerData.Zip_code,
		&resp.Order_status,
		&resp.Order_date,
		&resp.Required_date,
		&resp.Shipped_date,
		&resp.Store_id,
		&resp.StoreData.Store_id,
		&resp.StoreData.Store_name,
		&resp.StoreData.Phone,
		&resp.StoreData.Email,
		&resp.StoreData.Street,
		&resp.StoreData.City,
		&resp.StoreData.State,
		&resp.StoreData.Zip_code,
		&resp.Staff_id,
		&resp.StaffData.Staff_id,
		&resp.StaffData.First_name,
		&resp.StaffData.Last_name,
		&resp.StaffData.Email,
		&resp.StaffData.Phone,
		&resp.StaffData.Active,
		&resp.StaffData.Store_id,
		&resp.StaffData.Manager_id,
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

	resp = &models.GetListOrderResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(o.order_id, 0), 
			COALESCE(o.customer_id, 0),

			COALESCE(c.customer_id, 0),
			COALESCE(c.first_name,''),
			COALESCE(c.last_name,''),
			COALESCE(c.phone,''),
			COALESCE(c.email,''),
			COALESCE(c.street,''),
			COALESCE(c.city,''),
			COALESCE(c.state,''),
			COALESCE(c.zip_code,0),

			COALESCE(o.order_status, 0), 
			o.order_date,
			o.required_date,
			o.shipped_date,
			COALESCE(o.store_id, 0),

			COALESCE(sto.store_id, 0),
			COALESCE(sto.store_name,''),
			COALESCE(sto.phone,''),
			COALESCE(sto.email,''),
			COALESCE(sto.street,''),
			COALESCE(sto.city,''),
			COALESCE(sto.state,''),
			COALESCE(sto.zip_code,''),

			COALESCE(o.staff_id, 0),

			COALESCE(sta.staff_id, 0), 
			COALESCE(sta.first_name, ''),
			COALESCE(sta.last_name, ''),
			COALESCE(sta.email, ''),
			COALESCE(sta.phone, ''),
			COALESCE(sta.active, 0),
			COALESCE(sta.store_id, 0),
			COALESCE(sta.manager_id, 0)
		FROM orders as o join customers as c 
		ON o.customer_id = c.customer_id join stores as sto 
		ON o.store_id = sto.store_id join staffs as sta
		ON o.staff_id = sta.staff_id 
	`

	if len(req.Search) > 0 {
		filter += " AND order_name ILIKE '%' || '" + req.Search + "' || '%' "
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
			order    models.Order
			customer models.Customer
			store    models.Store
			staff    models.Staff
		)
		err = rows.Scan(
			&resp.Count,
			&order.Order_id,
			&order.Customer_id,
			&customer.Customer_id,
			&customer.First_name,
			&customer.Last_name,
			&customer.Phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.Zip_code,
			&order.Order_status,
			&order.Order_date,
			&order.Required_date,
			&order.Shipped_date,
			&order.Store_id,
			&store.Store_id,
			&store.Store_name,
			&store.Phone,
			&store.Email,
			&store.Street,
			&store.City,
			&store.State,
			&store.Zip_code,
			&order.Staff_id,
			&staff.Staff_id,
			&staff.First_name,
			&staff.Last_name,
			&staff.Email,
			&staff.Phone,
			&staff.Active,
			&staff.Store_id,
			&staff.Manager_id,
		)
		order.CustomerData = &customer
		order.StoreData = &store
		order.StaffData = &staff

		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, &order)
	}

	return resp, nil
}

func (r *OrderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			customer_id = :customer_id,
			order_status = :order_status,
			order_date = :order_date,
			required_date = :required_date,
			shipped_date = :shipped_date,
			store_id = :store_id,
			staff_id = :staff_id
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id":      req.Order_id,
		"customer_id":   req.Customer_id,
		"order_status":  req.Order_status,
		"order_date":    req.Order_date,
		"required_date": req.Required_date,
		"shipped_date":  req.Shipped_date,
		"store_id":      req.Store_id,
		"staff_id":      req.Staff_id,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *OrderRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fieldManager_ids")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			orders
		SET
		` + set + `
		WHERE order_id = :id
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

func (r *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM orders WHERE order_id = $1", req.Order_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}

func (r *OrderRepo) AddOrderItem(ctx context.Context, req *models.OrderItem) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO order_items(
			order_id, 
			item_id,
			product_id,
			quantity,
			list_price,
			discount
		)
		VALUES (
			$1,
			(
				SELECT COALESCE(MAX(item_id), 0) + 1 FROM order_items WHERE order_id = $1
			), 
			$2, $3, $4, $5) returning item_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Order_id,
		req.Product_id,
		req.Quantity,
		req.List_price,
		req.Discount,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *OrderRepo) RemoveOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM order_items WHERE order_id = $1 AND item_id = $2", req.Order_id, req.Item_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
