package models

import "errors"

var ErrorNegativeId = errors.New("bookId and userId must be more than 1")
var ErrorIncorrectId = errors.New("bookId and userId must be absent")

type Borrow struct {
	BookId int64 `json:"bookId"`
	UserId int64 `json:"userId"`
}
