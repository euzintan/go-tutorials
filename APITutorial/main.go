package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "Ikigai: The Secret To A Long and Healthy Life", Author: "Some Randoms Ikiguy", Quantity: 6},
	{ID: "2", Title: "Tender is the flesh", Author: "Some Hella Creepy", Quantity: 3},
	{ID: "3", Title: "The Martian", Author: "Andy Weir", Quantity: 1},
}

func getBooks(c *gin.Context) {
	// returning an indentedJSON created from the books array
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	queriedBook, err := getBookByID(c.Param("id"))

	if err != nil {
		// gin.H is a shortcut for writing custom JSON
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, queriedBook)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, fmt.Errorf("Book with ID: %v not found", id)
}

func checkoutById(c *gin.Context) {
	targetBook, err := getBookByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id of book not found"})
		return
	}

	if targetBook.Quantity <= 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "No more book available for loan"})
		return
	}

	targetBook.Quantity -= 1
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v has been checked out. There are %v left", targetBook.Title, targetBook.Quantity)})
}

func createBook(c *gin.Context) {
	var newBook book

	// bind the JSON we receive from the gin.Context to the newBook object
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// append in GO returns the appended slice
	books = append(books, newBook)

	// similar to returning to a GET but instead we return StatusCreated instead
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.POST("/checkout/:id", checkoutById)
	router.Run("localhost:8080")
}
