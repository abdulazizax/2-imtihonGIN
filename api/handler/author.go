package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	m "exam/models"

	"github.com/gin-gonic/gin"
)

type AuthorTest struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Biography string `json:"biography"`
}

func (h *handler) CreateAuthor(c *gin.Context) {
	newAuthor := AuthorTest{}

	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		log.Println("Error 1 handler")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !IsDateValid(newAuthor.BirthDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Date of birth is wrong!"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newAuthor.BirthDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Date of birth is wrong!"})
		return
	}

	result := m.Author{
		Name:      newAuthor.Name,
		BirthDate: parsedDate,
		Biography: newAuthor.Biography,
	}

	id, err := h.storage.Author().CreateAuthor(c.Request.Context(), result)
	if err != nil {
		log.Println("Error 2 handler")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func (h *handler) GetAuthors(c *gin.Context) {
	authors, err := h.storage.Author().GetAuthors(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if len(authors) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Result not found!"})
		return
	}
	c.IndentedJSON(http.StatusCreated, authors)
}

func (h *handler) GetAuthorByID(c *gin.Context) {
	id := c.Param("id")

	author, err := h.storage.Author().GetAuthorById(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, author)
}

func (h *handler) UpdateAuthoByID(c *gin.Context) {
	id := c.Param("id")
	newAuthor := AuthorTest{}

	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if !IsDateValid(newAuthor.BirthDate) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Date of birth is wrong!"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newAuthor.BirthDate)
	if err != nil {
		fmt.Println("Erro converting string to time.Time()", err.Error())
		return
	}

	result := m.Author{
		Name:      newAuthor.Name,
		BirthDate: parsedDate,
		Biography: newAuthor.Biography,
	}

	author, err := h.storage.Author().UpdateAuthorByID(c.Request.Context(), id, result)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, author)
}

func (h *handler) DeleteAuthorByID(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Author().DeleteAuthorByID(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "There is no author with this id"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Deleted author_id": id})
}

func (h *handler) GetAuthorByName(c *gin.Context) {
	name := c.Param("name")

	author, err := h.storage.Author().GetAuthorByName(c.Request.Context(), name)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, author)
}
