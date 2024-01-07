package storage

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/egor-zakharov/library/borrow/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFindBook(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	var id int64 = 1
	wantBorrow := models.Borrow{BookId: 1, UserId: 2}
	row := sqlmock.NewRows([]string{"book_id", "user_id"}).AddRow(1, 2)
	wantNotFoundErr := ErrorNotFound
	someErr := errors.New("some err")
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnRows(row)
		got, _ := s.FindBook(id)
		assert.Equal(t, &wantBorrow, got)
	})

	t.Run("error book not found", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)
		_, got := s.FindBook(id)
		assert.Equal(t, wantNotFoundErr, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnError(someErr)
		_, got := s.FindBook(id)
		assert.Equal(t, someErr, got)
	})
}

func TestFindUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	var id int64 = 1
	wantBorrow := models.Borrow{BookId: 1, UserId: 2}
	row := sqlmock.NewRows([]string{"book_id", "user_id"}).AddRow(1, 2)
	wantNotFoundErr := ErrorNotFound
	someErr := errors.New("some err")
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnRows(row)
		got, _ := s.FindUser(id)
		assert.Equal(t, &wantBorrow, got)
	})

	t.Run("error book not found", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)
		_, got := s.FindUser(id)
		assert.Equal(t, wantNotFoundErr, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from borrows").
			WithArgs(id).
			WillReturnError(someErr)
		_, got := s.FindUser(id)
		assert.Equal(t, someErr, got)
	})
}

func TestFindAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"book_id", "user_id"}).AddRow(1, 2).AddRow(3, 4)
		mock.ExpectQuery("select (.+) from borrows order by user_id, book_id").
			WillReturnRows(rows)
		want := []models.Borrow{
			{BookId: 1, UserId: 2},
			{BookId: 3, UserId: 4},
		}
		got, _ := s.FindAll()
		assert.Equal(t, want, got)
	})

	t.Run("error", func(t *testing.T) {
		wantErr := errors.New("some error")
		mock.ExpectQuery("select (.+) from borrows order by user_id, book_id").
			WillReturnError(wantErr)
		_, got := s.FindAll()
		assert.Equal(t, wantErr, got)
	})

	t.Run("scan error", func(t *testing.T) {
		wantErr := errors.New("sql: Scan error on column index 0, name \"book_id\": converting NULL to int64 is unsupported")
		rows := sqlmock.NewRows([]string{"book_id", "user_id"}).
			AddRow(nil, 2)
		mock.ExpectQuery("select (.+) from borrows order by user_id, book_id").
			WillReturnRows(rows)
		_, got := s.FindAll()
		assert.Equal(t, wantErr.Error(), got.Error())
	})
}

func TestAdd(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantBorrow := models.Borrow{BookId: 1, UserId: 2}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("insert into borrows").
			WithArgs(wantBorrow.BookId, wantBorrow.UserId).WillReturnResult(sqlmock.NewResult(1, 1))
		got, _ := s.Add(wantBorrow)
		assert.Equal(t, wantBorrow, *got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("insert into borrows").
			WithArgs(wantBorrow.BookId, wantBorrow.UserId).WillReturnError(wantErr)
		_, got := s.Add(wantBorrow)
		assert.Equal(t, wantErr, got)
	})
}

func TestDelete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantBorrow := models.Borrow{BookId: 1, UserId: 2}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("delete from borrows").
			WithArgs(wantBorrow.BookId, wantBorrow.UserId).WillReturnResult(sqlmock.NewResult(0, 1))
		got := s.Delete(wantBorrow)
		assert.Nil(t, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("delete from borrows").
			WithArgs(wantBorrow.BookId, wantBorrow.UserId).WillReturnError(wantErr)
		got := s.Delete(wantBorrow)
		assert.Equal(t, wantErr, got)
	})

}
