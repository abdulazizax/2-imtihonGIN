package models

import (
	"time"
)

type Author struct {
	AuthorID   string    `json:"author_id"`
	Name       string    `json:"name"`       
	BirthDate  time.Time `json:"birth_date"` 
	Biography  string    `json:"biography"`  
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
