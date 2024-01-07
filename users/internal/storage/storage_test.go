package storage

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/egor-zakharov/library/users/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFindSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	var id int64 = 1
	wantUser := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	row := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).AddRow(1, "fname", "lname")
	wantNotFoundErr := ErrorNotFound
	someErr := errors.New("some err")
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from users").
			WithArgs(id).
			WillReturnRows(row)
		got, _ := s.Find(id)
		assert.Equal(t, &wantUser, got)
	})

	t.Run("error book not found", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from users").
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)
		_, got := s.Find(id)
		assert.Equal(t, wantNotFoundErr, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("select (.+) from users").
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
		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
			AddRow(1, "fname", "lname").
			AddRow(2, "fname_2", "lname_2")
		mock.ExpectQuery("select (.+) from users order by id").
			WillReturnRows(rows)
		want := []models.User{
			{Id: 1, FirstName: "fname", LastName: "lname"},
			{Id: 2, FirstName: "fname_2", LastName: "lname_2"},
		}
		got, _ := s.FindAll()
		assert.Equal(t, want, got)
	})

	t.Run("error", func(t *testing.T) {
		wantErr := errors.New("some error")
		mock.ExpectQuery("select (.+) from users order by id").
			WillReturnError(wantErr)
		_, got := s.FindAll()
		assert.Equal(t, wantErr, got)
	})

	t.Run("scan error", func(t *testing.T) {
		wantErr := errors.New("sql: Scan error on column index 0, name \"id\": converting NULL to int64 is unsupported")
		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
			AddRow(nil, "fname", "lname")
		mock.ExpectQuery("select (.+) from users order by id").
			WillReturnRows(rows)
		_, got := s.FindAll()
		assert.Equal(t, wantErr.Error(), got.Error())
	})
}

func TestAdd(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantUser := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("insert into users").
			WithArgs(wantUser.FirstName, wantUser.LastName).WillReturnResult(sqlmock.NewResult(1, 1))
		got, _ := s.Add(wantUser)
		assert.Equal(t, wantUser, *got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("insert into users").
			WithArgs(wantUser.FirstName, wantUser.LastName).WillReturnError(wantErr)
		_, got := s.Add(wantUser)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error with last insertedId", func(t *testing.T) {
		mock.ExpectExec("insert into users").
			WithArgs(wantUser.FirstName, wantUser.LastName).WillReturnResult(sqlmock.NewErrorResult(wantErr))
		_, got := s.Add(wantUser)
		assert.Equal(t, wantErr, got)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	s := New(db)
	wantUser := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	wantErr := errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("update users").
			WithArgs(wantUser.FirstName, wantUser.LastName, wantUser.Id).WillReturnResult(sqlmock.NewResult(0, 1))
		got, _ := s.Update(wantUser)
		assert.Equal(t, wantUser, *got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("update users").
			WithArgs(wantUser.FirstName, wantUser.LastName, wantUser.Id).WillReturnError(wantErr)
		_, got := s.Update(wantUser)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error affected", func(t *testing.T) {
		mock.ExpectExec("update users").
			WithArgs(wantUser.FirstName, wantUser.LastName, wantUser.Id).WillReturnResult(sqlmock.NewErrorResult(wantErr))
		_, got := s.Update(wantUser)
		assert.Equal(t, wantErr, got)
	})

	t.Run("error nothing to update", func(t *testing.T) {
		wantErr = ErrorNothingToUpdate
		mock.ExpectExec("update users").
			WithArgs(wantUser.FirstName, wantUser.LastName, wantUser.Id).WillReturnResult(sqlmock.NewResult(0, 0))
		_, got := s.Update(wantUser)
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
		mock.ExpectExec("delete from users").
			WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
		got := s.Delete(id)
		assert.Nil(t, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec("delete from users").
			WithArgs(id).WillReturnError(wantErr)
		got := s.Delete(id)
		assert.Equal(t, wantErr, got)
	})

}
