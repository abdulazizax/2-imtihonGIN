package api

import (
	h "exam/api/handler"
	"exam/storage"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Storage storage.IStorage
}

func New(option Option) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := h.New(&h.HandlerConfig{
		Storage: option.Storage,
	})

	crud := router.Group("")
	crud.POST("/author", handler.CreateAuthor)
	crud.POST("/book", handler.CreateBook)

	crud.GET("/author", handler.GetAuthors)
	crud.GET("/book", handler.GetBooks)

	crud.GET("/author/:id", handler.GetAuthorByID)
	crud.GET("/book/:id", handler.GetBookByID)

	crud.PUT("/author/:id", handler.UpdateAuthoByID)
	crud.PUT("/book/:id", handler.UpdateBookByID)

	crud.DELETE("/author/:id", handler.DeleteAuthorByID)
	crud.DELETE("/book/:id", handler.DeleteBookByID)

	return router
}
