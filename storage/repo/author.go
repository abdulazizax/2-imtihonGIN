package repo

import (
	"context"
	m "exam/models"
)

type AuthorsStorageI interface {
	CreateAuthor(ctx context.Context, aut m.Author) (string, error)
	GetAuthors(ctx context.Context) ([]m.Author, error)
	GetAuthorById(ctx context.Context, id string) (m.Author, error)
	UpdateAuthorByID(ctx context.Context, id string, author m.Author) (m.Author, error)
	DeleteAuthorByID(ctx context.Context, id string) error
	GetAuthorByName(ctx context.Context, name string) (m.Author, error)
}
