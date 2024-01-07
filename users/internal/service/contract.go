package service

import "github.com/egor-zakharov/library/users/internal/models"

type Service interface {
	Find(userId int64) (*models.User, error)
	FindAll() ([]models.User, error)
	Add(in models.User) (*models.User, error)
	Update(in models.User) (*models.User, error)
	Delete(userId int64) error
}
