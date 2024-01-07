package storage

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/egor-zakharov/library/books/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFindSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	var id int64 = 1
	wantBook := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	row := sqlmock.NewRows([]string{"id", "title", "author", "released_year"}).AddRow(1, "title", "author", 2020)
	wantNotFoundErr := ErrorNotFound
	someErr := errors.New("some err")
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from books").
			WithArgs(id).
			WillReturnRows(row)
		got, _ := s.Find(id)
		assert.Equal(t, &wantBook, got)
	})

	t.Run("error book not found", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from books").
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)
		_, got := s.Find(id)
		assert.Equal(t, wantNotFoundErr, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from books").
			WithArgs(id).
			WillReturnError(someErr)
		_, got := s.Find(id)
		assert.Equal(t, someErr, got)
	})
}

func TestFindAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "author", "released_year"}).
			AddRow(1, "title", "author", 2020).
			AddRow(2, "title_2", "author_2", 2021)
		mock.ExpectQuery("select (.+) from books order by id").
			WillReturnRows(rows)
		want := []models.Book{
			{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020},
			{Id: 2, Title: "title_2", Author: "author_2", ReleasedYear: 2021},
		}
		got, _ := s.FindAll()
		assert.Equal(t, want, got)
	})

	t.Run("error", func(t *testing.T) {
		wantErr := errors.New("some error")
		mock.ExpectQuery("select (.+) from books order by id").
			WillReturnError(wantErr)
		_, got := s.FindAll()
		assert.Equal(t, wantErr, got)
	})

	t.Run("scan error", func(t *testing.T) {
		wantErr := errors.New("sql: Scan error on column index 0, name \"id\": converting NULL to int64 is unsupported")
		rows := sqlmock.NewRows([]string{"id", "title", "author", "released_year"}).
			AddRow(nil, "title", "author", 2020)
		mock.ExpectQuery("select (.+) from books order by id").
			WillReturnRows(rows)
		_, got := s.FindAll()
		assert.Equal(t, wantErr.Error(), got.Error())
	})
}

func TestAdd(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantBook := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("insert into books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear).WillReturnResult(sqlmock.NewResult(1, 1))
		got, _ := s.Add(wantBook)
		assert.Equal(t, wantBook, *got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("insert into books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear).WillReturnError(wantErr)
		_, got := s.Add(wantBook)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error with last insertedId", func(t *testing.T) {
		mock.ExpectExec("insert into books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear).WillReturnResult(sqlmock.NewErrorResult(wantErr))
		_, got := s.Add(wantBook)
		assert.Equal(t, wantErr, got)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantBook := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("update books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear, wantBook.Id).WillReturnResult(sqlmock.NewResult(0, 1))
		got, _ := s.Update(wantBook)
		assert.Equal(t, wantBook, *got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("update books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear, wantBook.Id).WillReturnError(wantErr)
		_, got := s.Update(wantBook)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error affected", func(t *testing.T) {
		mock.ExpectExec("update books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear, wantBook.Id).WillReturnResult(sqlmock.NewErrorResult(wantErr))
		_, got := s.Update(wantBook)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error nothing to update", func(t *testing.T) {
		wantErr = ErrorNothingToUpdate
		mock.ExpectExec("update books").
			WithArgs(wantBook.Title, wantBook.Author, wantBook.ReleasedYear, wantBook.Id).WillReturnResult(sqlmock.NewResult(0, 0))
		_, got := s.Update(wantBook)
		assert.Equal(t, wantErr, got)
	})
}

func TestDelete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	var id int64 = 1
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("delete from books").
			WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
		got := s.Delete(id)
		assert.Nil(t, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("delete from books").
			WithArgs(id).WillReturnError(wantErr)
		got := s.Delete(id)
		assert.Equal(t, wantErr, got)
	})

}
