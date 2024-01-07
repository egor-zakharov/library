package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/egor-zakharov/library/borrow/config"
	"github.com/egor-zakharov/library/borrow/internal/handler"
	"github.com/egor-zakharov/library/borrow/internal/service"
	"github.com/egor-zakharov/library/borrow/internal/storage"
	bookwrapper "github.com/egor-zakharov/library/borrow/internal/wrapper/book_wrapper"
	userwrapper "github.com/egor-zakharov/library/borrow/internal/wrapper/user_wrapper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/egor-zakharov/library/borrow/docs"
	_ "github.com/go-sql-driver/mysql"

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
	h := handler.NewHandler(service.New(storage.New(db), bookwrapper.New(), userwrapper.New()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("api/v1")
	{
		v1.GET("/book/:id", h.GetBookById)
		v1.GET("/user/:id", h.GetUserById)
		v1.GET("/borrow/", h.FindAll)
		v1.POST("/borrow/", h.Add)
		v1.DELETE("/borrow/", h.Delete)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
