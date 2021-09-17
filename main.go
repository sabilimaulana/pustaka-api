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

	// book := book.Book{}
	// book.Title = "Hidup begitu indah dan hanya itu yang kita punya"
	// book.Description = "Buku ini adalah bunga rampai nonfiksi pertama Dea Anugrah."
	// book.Price = 50000
	// book.Discount = 5
	// book.Rating = 5

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// var book book.Book
	// var books []book.Book

	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// for _, b := range books {
	// 	fmt.Println(b.Title)
	// }

	// fmt.Println(books)

	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	book.Title = "Kertas Basah"
	book.Description = "Kumpulan puisi."

	err = db.Debug().Save(&book).Error

	if err != nil {
		fmt.Println(err.Error())
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
