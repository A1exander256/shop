package service

import (
	"github.com/alexander256/shop/models"
	"github.com/alexander256/shop/pkg/storage"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	log     *logrus.Logger
	storage storage.User
}

func NewUserService(storage storage.User, log *logrus.Logger) *UserService {
	return &UserService{
		log:     log,
		storage: storage,
	}
}

func (s *UserService) Create(user *models.User) (int, error) {
	return s.storage.Create(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.storage.Update(user)
}

func (s *UserService) Delete(userId int) error {
	return s.storage.Delete(userId)
}

func (s *UserService) GetById(userId int) (models.User, error) {
	return s.storage.GetById(userId)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.storage.GetAll()
}
