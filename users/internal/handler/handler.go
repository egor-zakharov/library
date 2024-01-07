package handler

import (
	"fmt"
	"net/http"

	"github.com/egor-zakharov/library/users/internal/service"
	u "github.com/egor-zakharov/library/users/internal/utils"

	"github.com/egor-zakharov/library/users/internal/models"

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

// GetAllUsers godoc
// @Summary Get all users
// @Produce json
// @Success 200 {array} models.User
// @Router /user/ [get]
func (h *restHanlder) GetAllUsers(c *gin.Context) {
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
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
	result, err := h.service.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// AddUser godoc
// @Summary Add user
// @Produce json
// @Param user body models.User true "Add user"
// @Success 200 {object} models.User
// @Router /user/ [post]
func (h *restHanlder) AddUser(c *gin.Context) {
	param := &models.User{}
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

// UpdateUser godoc
// @Summary Update user
// @Produce json
// @Param id path int true "User Id"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func (h *restHanlder) UpdateUser(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	param := &models.User{}
	err = c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	param.Id = id
	result, err := h.service.Update(*param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// DeleteUserById godoc
// @Summary Delete user
// @Produce json
// @Param id path int true "User Id"
// @Success 200
// @Router /user/{id} [delete]
func (h *restHanlder) DeleteUserById(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("User: %d. successfully deleted", id)})
}
