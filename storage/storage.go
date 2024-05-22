package storage

import (
	"exam/storage/postgres"
	"exam/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Author() repo.AuthorsStorageI
	Book() repo.BooksStorageI
}

type storagePg struct {
	db         *sqlx.DB
	authorRepo repo.AuthorsStorageI
	bookRepo   repo.BooksStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		authorRepo: postgres.NewAuthorrepo(db),
		bookRepo:   postgres.NewBookrepo(db),
	}
}

func (s *storagePg) Author() repo.AuthorsStorageI {
	return s.authorRepo
}

func (s *storagePg) Book() repo.BooksStorageI {
	return s.bookRepo
}
