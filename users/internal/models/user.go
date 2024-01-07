package models

import "errors"

var ErrorNegativeId = errors.New("id must be more than 1")
var ErrorIncorrectId = errors.New("id must be absent")
var ErrorIncorrectFirstName = errors.New("firstName must be absent")
var ErrorIncorrectLastName = errors.New("lastName must be absent")

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) IsDraft() bool {
	return u.Id == 0
}

func (u *User) IsValidId() bool {
	return u.Id > 0
}

func (u *User) IsValidFirstName() bool {
	return u.FirstName != ""
}

func (u *User) IsValidLastName() bool {
	return u.LastName != ""
}
