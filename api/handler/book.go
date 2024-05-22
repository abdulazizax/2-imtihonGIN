package handler

import (
	m "exam/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookTest struct {
	AuthorName      string `json:"author_name"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	ISBN            string `json:"isbn"`
	Description     string `json:"description"`
}

func (h *handler) CreateBook(c *gin.Context) {
	newBook := BookTest{}

	if err := c.ShouldBind(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !IsDateValid(newBook.PublicationDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Publication date is wrong!"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newBook.PublicationDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := m.Book{
		AuthorName:      newBook.AuthorName,
		Title:           newBook.Title,
		PublicationDate: parsedDate,
		ISBN:            newBook.ISBN,
		Description:     newBook.Description,
	}

	book, authorName, err := h.storage.Book().CreateBook(c.Request.Context(), result)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.AuthorName = authorName

	c.IndentedJSON(http.StatusCreated, book)
}

func (h *handler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.storage.Book().GetBookById(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func (h *handler) GetBooks(c *gin.Context) {
	books, err := h.storage.Book().GetBooks(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(books) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Result not found!!!"})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func (h *handler) UpdateBookByID(c *gin.Context) {
	id := c.Param("id")
	newBook := BookTest{}

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !IsDateValid(newBook.PublicationDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Publication date is wrong!"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newBook.PublicationDate)
	if err != nil {
		fmt.Println("Erro converting string to time.Time()", err.Error())
		return
	}

	result := m.Book{
		AuthorName:      newBook.AuthorName,
		Title:           newBook.Title,
		PublicationDate: parsedDate,
		ISBN:            newBook.ISBN,
		Description:     newBook.Description,
	}

	book, authorName, err := h.storage.Book().UpdateBookById(c.Request.Context(), id, result)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.AuthorName = authorName
	c.IndentedJSON(http.StatusOK, book)
}

func (h *handler) DeleteBookByID(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Book().DeleteBookByID(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "There is no book with this id"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Deleted book id": id})
}
