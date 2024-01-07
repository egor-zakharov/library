package service

import (
	"github.com/egor-zakharov/library/borrow/internal/models"
	"github.com/egor-zakharov/library/borrow/internal/storage"
	bookwrapper "github.com/egor-zakharov/library/borrow/internal/wrapper/book_wrapper"
	userwrapper "github.com/egor-zakharov/library/borrow/internal/wrapper/user_wrapper"
)

type service struct {
	storage     storage.Storage
	bookWrapper bookwrapper.BookWrapper
	userWrapper userwrapper.UserWrapper
}

func New(storage storage.Storage, bookWrapper bookwrapper.BookWrapper, userWrapper userwrapper.UserWrapper) Service {
	return &service{
		storage:     storage,
		bookWrapper: bookWrapper,
		userWrapper: userWrapper,
	}
}

func (s *service) FindBook(bookId int64) (*models.Book, error) {
	return s.bookWrapper.FindBook(bookId)
}

func (s *service) FindUser(userId int64) (*models.User, error) {
	return s.userWrapper.FindUser(userId)
}

func (s *service) FindAll() ([]models.Borrow, error) {
	return s.storage.FindAll()
}

func (s *service) Add(in models.Borrow) (*models.Borrow, error) {
	_, err := s.bookWrapper.FindBook(in.BookId)
	if err != nil {
		return nil, err
	}
	_, err = s.userWrapper.FindUser(in.UserId)
	if err != nil {
		return nil, err
	}
	return s.storage.Add(in)
}

func (s *service) Delete(in models.Borrow) error {
	return s.storage.Delete(in)
}
