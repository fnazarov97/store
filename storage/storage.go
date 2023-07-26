package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Brand() BrandRepoI
	Product() ProductRepoI
	Customer() CustomerRepoI
	Store() StoreRepoI
	Staff() StaffRepoI
	Order() OrderRepoI
	Stock() StockRepoI
	User() UserRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) (int64, error)
}

type BrandRepoI interface {
	Create(context.Context, *models.CreateBrand) (string, error)
	GetByID(context.Context, *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(context.Context, *models.GetListBrandRequest) (*models.GetListBrandResponse, error)
	Update(context.Context, *models.UpdateBrand) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.BrandPrimaryKey) (int64, error)
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) (int64, error)
}

type CustomerRepoI interface {
	Create(context.Context, *models.CreateCustomer) (string, error)
	GetByID(context.Context, *models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(context.Context, *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	Update(context.Context, *models.UpdateCustomer) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.CustomerPrimaryKey) (int64, error)
}

type StoreRepoI interface {
	Create(context.Context, *models.CreateStore) (string, error)
	GetByID(context.Context, *models.StorePrimaryKey) (*models.Store, error)
	GetList(context.Context, *models.GetListStoreRequest) (*models.GetListStoreResponse, error)
	Update(context.Context, *models.UpdateStore) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.StorePrimaryKey) (int64, error)
}

type StaffRepoI interface {
	Create(context.Context, *models.CreateStaff) (string, error)
	GetByID(context.Context, *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(context.Context, *models.GetListStaffRequest) (*models.GetListStaffResponse, error)
	Update(context.Context, *models.UpdateStaff) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.StaffPrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimaryKey) (*models.Order, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(context.Context, *models.UpdateOrder) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.OrderPrimaryKey) (int64, error)
	AddOrderItem(ctx context.Context, req *models.OrderItem) (string, error)
	RemoveOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error)
}

type StockRepoI interface {
	Create(context.Context, *models.CreateStock) (string, error)
	GetByID(context.Context, *models.StockPrimaryKey) (*models.Stock, error)
	GetList(context.Context, *models.GetListStockRequest) (*models.GetListStockResponse, error)
	GetByIdProductStock(ctx context.Context, storeId int, productId int) (resp *models.Stock, err error)
	Update(context.Context, *models.UpdateStock) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.StockPrimaryKey) (int64, error)
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) (int64, error)
}
