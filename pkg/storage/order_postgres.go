package storage

import (
	"fmt"

	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrderStorage struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewOrderStorage(db *sqlx.DB, log *logrus.Logger) *OrderStorage {
	return &OrderStorage{
		log: log,
		db:  db,
	}
}

func (s *OrderStorage) Create(order *models.Order) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING uuid", tableOrders)

	row := s.db.QueryRow(query, order.UserId)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
