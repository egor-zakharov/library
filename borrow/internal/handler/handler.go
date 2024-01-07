package handler

import (
	"net/http"

	"github.com/egor-zakharov/library/borrow/internal/models"
	"github.com/egor-zakharov/library/borrow/internal/service"
	u "github.com/egor-zakharov/library/borrow/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type restHanlder struct {
	service service.Service
}

func NewHandler(service service.Service) RestHandler {
	return &restHanlder{
		service: service,
	}
}

// GetBookById godoc
// @Summary Get book
// @Produce json
// @Param id path int true "Book Id"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (h *restHanlder) GetBookById(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	result, err := h.service.FindBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// GetUserById godoc
// @Summary Get user
// @Produce json
// @Param id path int true "User Id"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func (h *restHanlder) GetUserById(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	result, err := h.service.FindUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// GetAllBorrows godoc
// @Summary Get all borrows
// @Produce json
// @Success 200 {array} models.Borrow
// @Router /borrow/ [get]
func (h *restHanlder) FindAll(c *gin.Context) {
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// AddBorrow godoc
// @Summary Add borrow
// @Produce json
// @Param book body models.Borrow true "Add borrow"
// @Success 200 {object} models.Borrow
// @Router /borrow/ [post]
func (h *restHanlder) Add(c *gin.Context) {
	param := &models.Borrow{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	result, err := h.service.Add(*param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// DeleteBorrow godoc
// @Summary Delete borrow
// @Produce json
// @Param book body models.Borrow true "Delete borrow"
// @Success 200
// @Router /borrow/ [delete]
func (h *restHanlder) Delete(c *gin.Context) {
	param := &models.Borrow{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.Delete(*param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, SuccessResponse{Result: "successfully deleted"})
}
