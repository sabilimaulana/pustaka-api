package book

import (
	"errors"
	"strconv"
)

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	s.repository.Update(book)
	return book, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	if book.ID == 0 {
		return book, errors.New("Data with id: " + strconv.Itoa(ID) + " doesn't exist.")
	}

	if err != nil {
		return book, err
	}

	book, err = s.repository.Delete(book)

	return book, err
}
