package handler

import (
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHander(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Sabili Maulana",
		"bio":  "Sedang belajar golang",
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	// title, price
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Discount:    b.Discount,
			Rating:      b.Rating,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}
