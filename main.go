package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&book.Book{})
	fmt.Println("Database connection succeed")

	// define repository
	bookRepository := book.NewRepository(db)

	// Create
	// book := book.Book{
	// 	Title:       "Kertas Basah",
	// 	Description: "Kumpulan puisi.",
	// 	Price:       50000,
	// 	Rating:      5,
	// 	Discount:    5,
	// }
	// bookRepository.Create(book)

	// FindById
	// var book book.Book

	// book, _ = bookRepository.FindById(3)
	// fmt.Println("Title :", book.Title)
	// fmt.Println("Description :", book.Description)

	// FindAll
	books, _ := bookRepository.FindAll()
	for _, book := range books {
		fmt.Println("Title :", book.Title, ", Description :", book.Description)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/book", handler.PostBookHandler)

	router.Run(":8888")
}
