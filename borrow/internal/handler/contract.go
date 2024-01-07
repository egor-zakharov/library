package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetUserById(c *gin.Context)
	GetBookById(c *gin.Context)
	FindAll(c *gin.Context)
	Add(c *gin.Context)
	Delete(c *gin.Context)
}
