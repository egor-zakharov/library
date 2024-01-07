package storage

import (
	"database/sql"
	"errors"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) FindBook(bookId int64) (*models.Borrow, error) {
	borrow := models.Borrow{}
	err := s.db.QueryRow("select * from borrows where book_id = ?", bookId).Scan(&borrow.BookId, &borrow.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}
	return &borrow, nil
}

func (s *storage) FindUser(userId int64) (*models.Borrow, error) {
	borrow := models.Borrow{}
	err := s.db.QueryRow("select * from borrows where user_id = ?", userId).Scan(&borrow.BookId, &borrow.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}
	return &borrow, nil
}

func (s *storage) FindAll() ([]models.Borrow, error) {
	borrows := []models.Borrow{}
	rows, err := s.db.Query("select * from borrows order by user_id, book_id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		borrow := models.Borrow{}
		err := rows.Scan(&borrow.BookId, &borrow.UserId)
		if err != nil {
			return nil, err
		}
		borrows = append(borrows, borrow)
	}
	return borrows, nil
}

func (s *storage) Add(in models.Borrow) (*models.Borrow, error) {
	_, err := s.db.Exec("insert into borrows (book_id, user_id) values (?,?)", in.BookId, in.UserId)
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (s *storage) Delete(in models.Borrow) error {
	_, err := s.db.Exec("delete from borrows where book_id =? and user_id=?", in.BookId, in.UserId)
	if err != nil {
		return err
	}
	return nil
}
