package storage

import (
	"errors"

	"github.com/egor-zakharov/library/books/internal/models"
)

var ErrorNotFound = errors.New("book not found")
var ErrorNothingToUpdate = errors.New("nothing book to update")

type Storage interface {
	Find(bookId int64) (*models.Book, error)
	FindAll() ([]models.Book, error)
	Add(in models.Book) (*models.Book, error)
	Update(in models.Book) (*models.Book, error)
	Delete(bookId int64) error
}
