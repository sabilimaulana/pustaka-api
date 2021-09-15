package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/book/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/book", postBookHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Sabili Maulana",
		"bio":  "Sedang belajar golang",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar golang API with Gin",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {
	// title, price
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
