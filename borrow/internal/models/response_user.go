package models

type ResponseUser struct {
	Result User `json:"result"`
}

type User struct {
	Id        int64  `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
