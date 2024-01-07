package storage

import (
	"errors"

	"github.com/egor-zakharov/library/users/internal/models"
)

var ErrorNotFound = errors.New("user not found")
var ErrorNothingToUpdate = errors.New("nothing users to update")

type Storage interface {
	Find(userId int64) (*models.User, error)
	FindAll() ([]models.User, error)
	Add(in models.User) (*models.User, error)
	Update(in models.User) (*models.User, error)
	Delete(userId int64) error
}
