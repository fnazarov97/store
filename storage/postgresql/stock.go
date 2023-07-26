package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type StockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *StockRepo {
	return &StockRepo{
		db: db,
	}
}

func (r *StockRepo) Create(ctx context.Context, req *models.CreateStock) (string, error) {

	var (
		query string
		id    int
	)

	query = `
		INSERT INTO stocks(
			store_id, 
			product_id,
			quantity
		)
		VALUES ( 
			$1, $2, $3) returning store_id
	`

	err := r.db.QueryRow(ctx, query,
		req.Store_id,
		req.Product_id,
		req.Quantity,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}

func (r *StockRepo) GetByIdProductStock(ctx context.Context, storeId int, productId int) (resp *models.Stock, err error) {
	resp = &models.Stock{}
	resp.StoreData = &models.Store{}
	resp.ProductData = &models.Product{}

	query := `
		SELECT 
			s.store_id,

			st.store_id, 
			st.store_name,
			COALESCE(st.phone, ''),
			COALESCE(st.email, ''),
			COALESCE(st.street, ''),
			COALESCE(st.city, ''),
			COALESCE(st.state, ''),
			COALESCE(st.zip_code, ''),

			s.product_id,

			p.product_id, 
			p.product_name, 
			p.brand_id,
			p.category_id,
			p.model_year,
			p.list_price,
			
			s.quantity
		FROM stocks AS s
		JOIN stores AS st ON st.store_id = s.store_id
		JOIN products AS p ON p.product_id = s.product_id
		WHERE s.store_id = $1 AND s.product_id = $2
	`

	err = r.db.QueryRow(ctx, query, storeId, productId).Scan(
		&resp.Store_id,
		&resp.StoreData.Store_id,
		&resp.StoreData.Store_name,
		&resp.StoreData.Phone,
		&resp.StoreData.Email,
		&resp.StoreData.Street,
		&resp.StoreData.City,
		&resp.StoreData.State,
		&resp.StoreData.Zip_code,
		&resp.Product_id,
		&resp.ProductData.Product_id,
		&resp.ProductData.Product_name,
		&resp.ProductData.Brand_id,
		&resp.ProductData.Category_id,
		&resp.ProductData.Model_year,
		&resp.ProductData.List_price,
		&resp.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (r *StockRepo) GetByID(ctx context.Context, req *models.StockPrimaryKey) (*models.Stock, error) {
	var (
		query string
		resp  models.Stock
	)

	query = `
		SELECT
			COALESCE(s.store_id, 0), 

			COALESCE(st.store_id, 0),
			COALESCE(st.store_name, ''),
			COALESCE(st.phone, ''),
			COALESCE(st.email, ''),
			COALESCE(st.street, ''),
			COALESCE(st.city, ''),
			COALESCE(st.state, ''),
			COALESCE(st.zip_code, ''),

			COALESCE(s.product_id, 0), 

			COALESCE(p.product_id, 0),
			COALESCE(p.product_name, ''),
			COALESCE(p.brand_id, 0),UpdateStock
			COALESCE(p.category_id, 0),
			COALESCE(p.model_year, 0),
			COALESCE(p.list_price, 0)
		FROM stocks as s join stores as st 
		ON s.store_id = st.store_id join products as p
		ON s.product_id = p.product_id
		WHERE s.store_id = $1
	`
	resp.StoreData = &models.Store{}
	resp.ProductData = &models.Product{}
	err := r.db.QueryRow(ctx, query, req.Store_id).Scan(
		&resp.Store_id,
		&resp.StoreData.Store_id,
		&resp.StoreData.Store_name,
		&resp.StoreData.Phone,
		&resp.StoreData.Email,
		&resp.StoreData.Street,
		&resp.StoreData.City,
		&resp.StoreData.State,
		&resp.StoreData.Zip_code,
		&resp.Product_id,
		&resp.ProductData.Product_id,
		&resp.ProductData.Product_name,
		&resp.ProductData.Brand_id,
		&resp.ProductData.Category_id,
		&resp.ProductData.Model_year,
		&resp.ProductData.List_price,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *StockRepo) GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error) {

	resp = &models.GetListStockResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			COALESCE(s.store_id, 0), 

			COALESCE(st.store_id, 0),
			COALESCE(st.store_name, ''),
			COALESCE(st.phone, ''),
			COALESCE(st.email, ''),
			COALESCE(st.street, ''),
			COALESCE(st.city, ''),
			COALESCE(st.state, ''),
			COALESCE(st.zip_code, ''),

			COALESCE(s.product_id, 0), 

			COALESCE(p.product_id, 0),
			COALESCE(p.product_name, ''),
			COALESCE(p.brand_id, 0),
			COALESCE(p.category_id, 0),
			COALESCE(p.model_year, 0),
			COALESCE(p.list_price, 0),

			COALESCE(s.quantity, 0)			
		FROM stocks as s join stores as st 
		ON s.store_id = st.store_id join products as p
		ON s.product_id = p.product_id
	`

	if len(req.Search) > 0 {
		filter += " AND quantity ILIKE '%' || '" + req.Search + "' || '%' "
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

		var stock models.Stock
		stock.StoreData = &models.Store{}
		stock.ProductData = &models.Product{}
		err = rows.Scan(
			&resp.Count,
			&stock.Store_id,
			&stock.StoreData.Store_id,
			&stock.StoreData.Store_name,
			&stock.StoreData.Phone,
			&stock.StoreData.Email,
			&stock.StoreData.Street,
			&stock.StoreData.City,
			&stock.StoreData.State,
			&stock.StoreData.Zip_code,
			&stock.Product_id,
			&stock.ProductData.Product_id,
			&stock.ProductData.Product_name,
			&stock.ProductData.Brand_id,
			&stock.ProductData.Category_id,
			&stock.ProductData.Model_year,
			&stock.ProductData.List_price,
			&stock.Quantity,
		)
		if err != nil {
			return nil, err
		}
		resp.Stocks = append(resp.Stocks, &stock)
	}

	return resp, nil
}

func (r *StockRepo) Update(ctx context.Context, req *models.UpdateStock) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			stocks
		SET
			quantity = :quantity
		WHERE stock_id = :store_id AND product_id = :product_id
	`

	params = map[string]interface{}{
		"store_id":   req.Store_id,
		"product_id": req.Product_id,
		"quantity":   req.Quantity,
	}
	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StockRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			stocks
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

func (r *StockRepo) Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error) {

	rows, err := r.db.Exec(ctx,
		"DELETE FROM stocks WHERE store_id = $1", req.Store_id,
	)

	if err != nil {
		return rows.RowsAffected(), err
	}

	return rows.RowsAffected(), nil
}
