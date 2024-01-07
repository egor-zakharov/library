package storage

import (
	"database/sql"
	"errors"

	"github.com/egor-zakharov/library/users/internal/models"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) Find(userId int64) (*models.User, error) {
	user := models.User{}
	err := s.db.QueryRow("select * from users where id = ?", userId).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *storage) FindAll() ([]models.User, error) {
	users := []models.User{}
	rows, err := s.db.Query("select * from users order by id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *storage) Add(in models.User) (*models.User, error) {
	result, err := s.db.Exec("insert into users (first_name, last_name) values (?,?)", in.FirstName, in.LastName)
	if err != nil {
		return nil, err
	}
	in.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (s *storage) Update(in models.User) (*models.User, error) {
	result, err := s.db.Exec("update users set first_name = ?, last_name = ? where id = ?", in.FirstName, in.LastName, in.Id)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, ErrorNothingToUpdate
	}
	return &in, nil
}

func (s *storage) Delete(userId int64) error {
	_, err := s.db.Exec("delete from users where id =?", userId)
	if err != nil {
		return err
	}
	return nil
}
