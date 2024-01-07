package bookwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/egor-zakharov/library/borrow/internal/models"
)

type wrapper struct {
}

func New() BookWrapper {
	return &wrapper{}
}

func (w *wrapper) FindBook(bookId int64) (*models.Book, error) {
	response, err := http.Get(fmt.Sprintf("http://localhost:8085/api/v1/book/%d", bookId))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, ErrorBookNotFound
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := models.ResponseBook{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
