package models

type ResponseBook struct {
	Result Book `json:"result"`
}

type Book struct {
	Id           int64  `json:"id" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Author       string `json:"author" validate:"required"`
	ReleasedYear int64  `json:"releasedYear" validate:"required"`
}
