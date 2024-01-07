package service

import (
	"errors"
	"testing"

	"github.com/egor-zakharov/library/users/internal/models"
	"github.com/egor-zakharov/library/users/internal/storage"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var id int64 = 1
	user := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	err := storage.ErrorNotFound

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Find(id).Return(&user, nil)
		got, _ := service.Find(id)
		assert.Equal(t, &user, got)
	})

	t.Run("error not found", func(t *testing.T) {
		mock.EXPECT().Find(id).Return(nil, err)
		_, errs := service.Find(id)
		assert.Equal(t, err, errs)
	})
}

func TestFindAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var err = errors.New("some error")
	var users = []models.User{
		{Id: 1, FirstName: "fname", LastName: "lname"},
		{Id: 2, FirstName: "fname_2", LastName: "lname_2"},
	}

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().FindAll().Return(users, nil)
		got, _ := service.FindAll()
		assert.Equal(t, users, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().FindAll().Return(nil, err)
		_, got := service.FindAll()
		assert.Equal(t, err, got)
	})
}

func TestAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	err := errors.New("some error")
	validateErr := models.ErrorIncorrectFirstName
	user := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	inUser := models.User{FirstName: "fname", LastName: "lname"}
	inUserErr := models.User{FirstName: "", LastName: "lname"}

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Add(inUser).Return(&user, nil)
		got, _ := service.Add(inUser)
		assert.Equal(t, &user, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().Add(inUser).Return(nil, err)
		_, got := service.Add(inUser)
		assert.Equal(t, err, got)
	})

	t.Run("validate error", func(t *testing.T) {
		mock.EXPECT().Add(inUserErr).AnyTimes().Return(nil, validateErr)
		_, got := service.Add(inUserErr)
		assert.Equal(t, validateErr, got)
	})
}

func TestUpdate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var err = errors.New("some error")
	user := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	inUserErr := models.User{Id: 1, FirstName: "fname", LastName: "lname"}
	validateErr := models.ErrorIncorrectFirstName

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Update(user).Return(&user, nil)
		got, _ := service.Update(user)
		assert.Equal(t, &user, got)
	})

	t.Run("error", func(t *testing.T) {
		mock.EXPECT().Update(user).Return(nil, err)
		_, got := service.Update(user)
		assert.Equal(t, err, got)
	})

	t.Run("validate error", func(t *testing.T) {
		mock.EXPECT().Update(inUserErr).AnyTimes().Return(nil, validateErr)
		_, got := service.Update(inUserErr)
		assert.Equal(t, validateErr, got)
	})
}

func TestDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mock := NewMockService(ctl)
	service := New(mock)
	var id int64 = 1
	var err = errors.New("some error")

	t.Run("success", func(t *testing.T) {
		mock.EXPECT().Delete(id).Return(nil)
		err := service.Delete(id)
		assert.Nil(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		mock.EXPECT().Delete(id).Return(err)
		errs := service.Delete(id)
		assert.Equal(t, err, errs)
	})

}
