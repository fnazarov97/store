package unit_test

import (
	"app/config"
	"app/storage/postgresql"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	categoryTestRepo *postgresql.CategoryRepo
	brandTestRepo    *postgresql.BrandRepo
	productTestRepo  *postgresql.ProductRepo
	stockTestRepo    *postgresql.StockRepo
	customerTestRepo *postgresql.CustomerRepo
	storeTestRepo    *postgresql.StoreRepo
	staffTestRepo    *postgresql.StaffRepo
	orderTestRepo    *postgresql.OrderRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(pool)
	}

	categoryTestRepo = postgresql.NewCategoryRepo(pool)
	brandTestRepo = postgresql.NewBrandRepo(pool)
	productTestRepo = postgresql.NewProductRepo(pool)
	stockTestRepo = postgresql.NewStockRepo(pool)
	customerTestRepo = postgresql.NewCustomerRepo(pool)
	storeTestRepo = postgresql.NewStoreRepo(pool)
	staffTestRepo = postgresql.NewStaffRepo(pool)
	orderTestRepo = postgresql.NewOrderRepo(pool)

	os.Exit(m.Run())
}
