package storage

import (
	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type User interface {
	Create(user *models.User) (int, error)
	Update(user *models.User) error
	Delete(userId int) error
}

type Storage struct {
	log *logrus.Logger
	User
}

func NewStorage(db *sqlx.DB, log *logrus.Logger) *Storage {
	return &Storage{
		log:  log,
		User: NewUserStorage(db, log),
	}
}
