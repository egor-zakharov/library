package service

import (
	"github.com/egor-zakharov/library/users/internal/models"
	"github.com/egor-zakharov/library/users/internal/storage"
	"github.com/egor-zakharov/library/users/internal/utils"
)

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) Find(userId int64) (*models.User, error) {
	return s.storage.Find(userId)
}

func (s *service) FindAll() ([]models.User, error) {
	return s.storage.FindAll()
}

func (s *service) Add(in models.User) (*models.User, error) {
	err := utils.ValidateAddableUser(in)
	if err != nil {
		return nil, err
	}
	return s.storage.Add(in)
}

func (s *service) Update(in models.User) (*models.User, error) {
	err := utils.ValidateUpdatableUser(in)
	if err != nil {
		return nil, err
	}
	return s.storage.Update(in)
}

func (s *service) Delete(userId int64) error {
	return s.storage.Delete(userId)
}
