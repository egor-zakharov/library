package userwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

type wrapper struct {
}

func New() UserWrapper {
	return &wrapper{}
}

func (w *wrapper) FindUser(userId int64) (*models.User, error) {
	response, err := http.Get(fmt.Sprintf("http://localhost:8086/api/v1/user/%d", userId))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, ErrorUserNotFound
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res models.ResponseUser
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res.Result, nil
}
