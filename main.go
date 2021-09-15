package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Sabili Maulana",
			"bio":  "Sedang belajar golang",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "Belajar golang API with Gin",
		})
	})

	router.Run(":8888")
}
