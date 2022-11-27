package service

import (
	"github.com/alexander256/shop/models"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	log     *logrus.Logger
	storage storage.Product
}

func NewProductService(storage storage.Product, log *logrus.Logger) *ProductService {
	return &ProductService{
		log:     log,
		storage: storage,
	}
}

func (s *ProductService) Create(product *models.Product) (int, error) {
	return s.storage.Create(product)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.storage.Update(product)
}

func (s *ProductService) Delete(productId int) error {
	return s.storage.Delete(productId)
}
