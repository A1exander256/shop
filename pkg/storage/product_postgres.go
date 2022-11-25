package storage

import (
	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductStorage struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewProductStorage(db *sqlx.DB, log *logrus.Logger) *ProductStorage {
	return &ProductStorage{
		log: log,
		db:  db,
	}
}

func (s *ProductStorage) Create(product *models.Product) (int, error) {
	return 0, nil
}
