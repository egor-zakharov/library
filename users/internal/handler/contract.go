package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetUserById(c *gin.Context)
	GetAllUsers(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUserById(c *gin.Context)
}
