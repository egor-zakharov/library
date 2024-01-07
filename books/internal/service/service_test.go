package service

import (
	"errors"
	"testing"

	"github.com/egor-zakharov/library/books/internal/models"
	"github.com/egor-zakharov/library/books/internal/storage"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var id int64 = 1
	book := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	err := storage.ErrorNotFound

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Find(id).Return(&book, nil)
		got, _ := service.Find(id)
		assert.Equal(t, &book, got)
	})

	t.Run("error not found", func(t *testing.T) {
		mock.EXPECT().Find(id).Return(nil, err)
		_, errs := service.Find(id)
		assert.Equal(t, err, errs)
	})
}

func TestFindAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var err = errors.New("some error")
	var books = []models.Book{
		{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020},
		{Id: 2, Title: "title_2", Author: "author_2", ReleasedYear: 2021},
	}

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().FindAll().Return(books, nil)
		got, _ := service.FindAll()
		assert.Equal(t, books, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().FindAll().Return(nil, err)
		_, got := service.FindAll()
		assert.Equal(t, err, got)
	})
}

func TestAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	err := errors.New("some error")
	validateErr := models.ErrorIncorrectYear
	book := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	inBook := models.Book{Title: "title", Author: "author", ReleasedYear: 2020}
	inBookErr := models.Book{Title: "title", Author: "author", ReleasedYear: 1}

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Add(inBook).Return(&book, nil)
		got, _ := service.Add(inBook)
		assert.Equal(t, &book, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().Add(inBook).Return(nil, err)
		_, got := service.Add(inBook)
		assert.Equal(t, err, got)
	})

	t.Run("validate error", func(t *testing.T) {
		mock.EXPECT().Add(inBookErr).AnyTimes().Return(nil, validateErr)
		_, got := service.Add(inBookErr)
		assert.Equal(t, validateErr, got)
	})
}

func TestUpdate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var err = errors.New("some error")
	book := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	inBookErr := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 1}
	validateErr := models.ErrorIncorrectYear

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Update(book).Return(&book, nil)
		got, _ := service.Update(book)
		assert.Equal(t, &book, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().Update(book).Return(nil, err)
		_, got := service.Update(book)
		assert.Equal(t, err, got)
	})

	t.Run("validate error", func(t *testing.T) {
		mock.EXPECT().Update(inBookErr).AnyTimes().Return(nil, validateErr)
		_, got := service.Update(inBookErr)
		assert.Equal(t, validateErr, got)
	})
}

func TestDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var id int64 = 1
	var err = errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Delete(id).Return(nil)
		err := service.Delete(id)
		assert.Nil(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mock.EXPECT().Delete(id).Return(err)
		errs := service.Delete(id)
		assert.Equal(t, err, errs)
	})

}
