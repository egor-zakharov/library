package service

import "github.com/egor-zakharov/library/borrow/internal/models"

type Service interface {
	FindBook(bookId int64) (*models.Book, error)
	FindUser(userId int64) (*models.User, error)
	FindAll() ([]models.Borrow, error)
	Add(in models.Borrow) (*models.Borrow, error)
	Delete(in models.Borrow) error
}
