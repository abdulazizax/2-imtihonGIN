package models

import "time"

type Book struct {
	AuthorName      string    `json:"author_name"`
	BookID          int       `json:"book_id"`
	Title           string    `json:"title"`            
	AuthorID        string    `json:"-"`        
	PublicationDate time.Time `json:"publication_date"` 
	ISBN            string    `json:"isbn"`             
	Description     string    `json:"description"`      
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
