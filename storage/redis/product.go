package redis

import (
	"app/api/models"
	"encoding/json"

	"github.com/go-redis/redis"
)

type ProductRepo struct {
	db *redis.Client
}

func NewProductRepo(db *redis.Client) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(req *models.GetListProductResponse) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = r.db.Set("product_list", body, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo) GetList() (*models.GetListProductResponse, error) {

	resp := &models.GetListProductResponse{}

	body, err := r.db.Get("product_list").Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ProductRepo) Delete() error {

	err := r.db.Del("product_list").Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo) Exists() (bool, error) {

	exist, err := r.db.Exists("product_list").Result()
	if err != nil {
		return false, err
	}

	if exist <= 0 {
		return false, nil
	}

	return true, nil
}
