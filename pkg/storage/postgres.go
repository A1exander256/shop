package storage

import (
	"fmt"

	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
)

var (
	tableUsers    = "users"
	tableOrders   = "orders"
	tableProducts = "products"
	tablePrices   = "prices"
)

func NewPostgresDB(cfg *models.Postgres) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.UserName,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode)
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
