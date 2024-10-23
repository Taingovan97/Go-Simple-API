package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Tay Du Ky", Author: "Ngo Thua An", Quantity: 2},
	{ID: "2", Title: "Hong Lau Mong", Author: "Tao Tuyet Can", Quantity: 5},
	{ID: "3", Title: "Tam Quoc", Author: "La Quan Trung", Quantity: 6},
	{ID: "4", Title: "Thuy Hu", Author: "Thi Nai Am", Quantity: 8},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, book)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func bookByID(id string) (*book, error) {
	for _, v := range books {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}

	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookByID)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}
