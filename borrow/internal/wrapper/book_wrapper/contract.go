package bookwrapper

import (
	"errors"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

var ErrorBookNotFound = errors.New("book not found")

type BookWrapper interface {
	FindBook(bookId int64) (*models.Book, error)
}
