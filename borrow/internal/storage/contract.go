package storage

import (
	"errors"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

var ErrorNotFound = errors.New("book or user not found")

type Storage interface {
	FindBook(bookId int64) (*models.Borrow, error)
	FindUser(userId int64) (*models.Borrow, error)
	FindAll() ([]models.Borrow, error)
	Add(in models.Borrow) (*models.Borrow, error)
	Delete(in models.Borrow) error
}
