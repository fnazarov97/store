package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type Store struct {
	db       *pgxpool.Pool
	category storage.CategoryRepoI
	brand    storage.BrandRepoI
	product  storage.ProductRepoI
	customer storage.CustomerRepoI
	store    storage.StoreRepoI
	staff    storage.StaffRepoI
	order    storage.OrderRepoI
	stock    storage.StockRepoI
	user     storage.UserRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pgpool,
		category: NewCategoryRepo(pgpool),
		brand:    NewBrandRepo(pgpool),
		product:  NewProductRepo(pgpool),
		customer: NewCustomerRepo(pgpool),
		store:    NewStoreRepo(pgpool),
		staff:    NewStaffRepo(pgpool),
		order:    NewOrderRepo(pgpool),
		stock:    NewStockRepo(pgpool),
		user:     NewUserRepo(pgpool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Brand() storage.BrandRepoI {

	if s.brand == nil {
		s.brand = NewBrandRepo(s.db)
	}

	return s.brand
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Customer() storage.CustomerRepoI {

	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) Store() storage.StoreRepoI {

	if s.store == nil {
		s.store = NewStoreRepo(s.db)
	}

	return s.store
}

func (s *Store) Staff() storage.StaffRepoI {

	if s.staff == nil {
		s.staff = NewStaffRepo(s.db)
	}

	return s.staff
}

func (s *Store) Order() storage.OrderRepoI {

	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Stock() storage.StockRepoI {

	if s.stock == nil {
		s.stock = NewStockRepo(s.db)
	}

	return s.stock
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

// GORM
// ROW
// SQLBUILDER
// SQLX
// PGXPOOL
