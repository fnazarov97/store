package storage

import (
	"app/api/models"
)

type CacheStorageI interface {
	CloseDB()
	ProductCache() ProductCachaRepoI
}

type ProductCachaRepoI interface {
	Create(req *models.GetListProductResponse) error
	GetList() (*models.GetListProductResponse, error)
	Delete() error
	Exists() (bool, error)
}
