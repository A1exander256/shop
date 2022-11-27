package service

import (
	"github.com/alexander256/shop/models"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/sirupsen/logrus"
)

type User interface {
	Create(user *models.User) (int, error)
	Update(user *models.User) error
	Delete(userId int) error
	GetById(userId int) (models.User, error)
	GetAll() ([]models.User, error)
}

type Order interface {
	Create(order *models.Order) (int, error)
}

type Product interface {
	Create(product *models.Product) (int, error)
	Update(product *models.Product) error
	Delete(productId int) error
}
type Service struct {
	log *logrus.Logger
	User
	Order
	Product
}

func NewService(storage *storage.Storage, log *logrus.Logger) *Service {
	return &Service{
		log:     log,
		User:    NewUserService(storage.User, log),
		Order:   NewOrderService(storage.Order, log),
		Product: NewProductService(storage.Product, log),
	}
}
