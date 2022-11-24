package service

import (
	"github.com/alexander256/shop/models"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	log     *logrus.Logger
	storage storage.Order
}

func NewOrderService(storage storage.Order, log *logrus.Logger) *OrderService {
	return &OrderService{
		log:     log,
		storage: storage,
	}
}

func (s *OrderService) Create(order *models.Order) (int, error) {
	return s.storage.Create(order)
}
