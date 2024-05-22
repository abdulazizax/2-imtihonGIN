package repo

import (
	"context"
	m "exam/models"
)

type BooksStorageI interface {
	CreateBook(ctx context.Context, bok m.Book) (m.Book, string, error)
	GetBookById(ctx context.Context, id string) (m.Book, error)
	GetBooks(ctx context.Context) ([]m.Book, error)
	UpdateBookById(ctx context.Context, id string, updatedBook m.Book) (m.Book, string, error)
	DeleteBookByID(ctx context.Context, id string) error
}
