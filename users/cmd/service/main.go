package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/egor-zakharov/library/users/config"
	"github.com/egor-zakharov/library/users/internal/handler"
	"github.com/joho/godotenv"

	"github.com/egor-zakharov/library/users/internal/service"
	"github.com/egor-zakharov/library/users/internal/storage"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/egor-zakharov/library/users/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Users API
// @version 1.0
// @description Swagger API for Golang Project users.
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v1
func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("Error loading .env file")
	}
	cfg := config.New()

	//DSN full form username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(:%s)/%s", cfg.DBUserName, cfg.DBPassword, cfg.DBPort, cfg.DBDatabaseName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	h := handler.NewHandler(service.New(storage.New(db)))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("api/v1")
	{
		v1.GET("/user", h.GetAllUsers)
		v1.GET("/user/:id", h.GetUserById)
		v1.POST("/user/", h.AddUser)
		v1.PUT("/user/:id", h.UpdateUser)
		v1.DELETE("/user/:id", h.DeleteUserById)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
