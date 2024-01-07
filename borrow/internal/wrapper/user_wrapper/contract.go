package userwrapper

import (
	"errors"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

var ErrorUserNotFound = errors.New("user not found")

type UserWrapper interface {
	FindUser(userId int64) (*models.User, error)
}
