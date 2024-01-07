package utils

import "github.com/egor-zakharov/library/users/internal/models"

func ValidateAddableUser(in models.User) error {
	if !in.IsDraft() {
		return models.ErrorIncorrectId
	}
	if !in.IsValidFirstName() {
		return models.ErrorIncorrectFirstName
	}
	if !in.IsValidLastName() {
		return models.ErrorIncorrectLastName
	}
	return nil
}

func ValidateUpdatableUser(in models.User) error {
	if !in.IsValidId() {
		return models.ErrorNegativeId
	}
	if !in.IsValidFirstName() {
		return models.ErrorIncorrectFirstName
	}
	if !in.IsValidLastName() {
		return models.ErrorIncorrectLastName
	}
	return nil
}
