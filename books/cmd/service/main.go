package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/egor-zakharov/library/books/config"
	"github.com/egor-zakharov/library/books/internal/handler"
	"github.com/joho/godotenv"

	"github.com/egor-zakharov/library/books/internal/service"
	"github.com/egor-zakharov/library/books/internal/storage"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/egor-zakharov/library/books/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Books API
// @version 1.0
// @description Swagger API for Golang Project Books.
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
		v1.GET("/book", h.GetAllBooks)
		v1.GET("/book/:id", h.GetBookById)
		v1.POST("/book/", h.AddBook)
		v1.PUT("/book/:id", h.UpdateBook)
		v1.DELETE("/book/:id", h.DeleteBookById)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
