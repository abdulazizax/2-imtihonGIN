package postgres

import (
	"context"
	"errors"
	m "exam/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	DB *sqlx.DB
}

func NewBookrepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{
		DB: db,
	}
}

func (b *BookRepo) CreateBook(ctx context.Context, bok m.Book) (m.Book, string, error) {
	var authorID string

	tx, err := b.DB.BeginTx(ctx, nil)
	if err != nil {
		return m.Book{}, "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
		SELECT author_id FROM authors WHERE name = $1
	`
	row := tx.QueryRowContext(ctx, query, bok.AuthorName)
	err = row.Scan(&authorID)
	if err != nil {
		return m.Book{}, "", err
	}

	query = `
		INSERT INTO books (
			title, 
			author_id, 
			publication_date,
			isbn,
			description)
		VALUES($1, $2, $3, $4, $5)
		RETURNING book_id, title, author_id, publication_date, isbn, description, created_at, updated_at;
	`

	var newBook m.Book
	err = tx.QueryRowContext(ctx, query, strings.ToLower(bok.Title), authorID, bok.PublicationDate, bok.ISBN, bok.Description).Scan(
		&newBook.BookID,
		&newBook.Title,
		&newBook.AuthorID,
		&newBook.PublicationDate,
		&newBook.ISBN,
		&newBook.Description,
		&newBook.CreatedAt,
		&newBook.UpdatedAt,
	)
	if err != nil {
		return m.Book{}, "", err
	}

	return newBook, bok.AuthorName, nil
}

func (b *BookRepo) GetBookById(ctx context.Context, id string) (m.Book, error) {
	query := `
        SELECT books.book_id, books.title, authors.name AS author_name, books.publication_date, books.isbn, books.description, books.created_at, books.updated_at
        FROM books
        INNER JOIN authors ON books.author_id = authors.author_id
        WHERE books.book_id = $1;
    `

	intValue, err := strconv.Atoi(id)
	if err != nil {
		return m.Book{}, err
	}

	row := b.DB.QueryRowContext(ctx, query, intValue)

	book := m.Book{}

	err = row.Scan(&book.BookID, &book.Title, &book.AuthorName, &book.PublicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookRepo) GetBooks(ctx context.Context) ([]m.Book, error) {
	query := `
		SELECT books.book_id, books.title, authors.name AS author_name, books.publication_date, books.isbn, books.description, books.created_at, books.updated_at
		FROM books
		INNER JOIN authors ON books.author_id = authors.author_id;
	`
	rows, err := b.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []m.Book

	for rows.Next() {
		var book m.Book
		err := rows.Scan(&book.BookID, &book.Title, &book.AuthorName, &book.PublicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookRepo) UpdateBookById(ctx context.Context, id string, updatedBook m.Book) (m.Book, string, error) {
	query := `
		UPDATE books
		SET title = $1, 
		    author_id = (SELECT author_id FROM authors WHERE name = $2),
		    publication_date = $3, 
		    isbn = $4, 
		    description = $5, 
		    updated_at = CURRENT_TIMESTAMP
		WHERE book_id = $6
		RETURNING book_id, title, author_id, publication_date, isbn, description, created_at, updated_at;
	`

	intValue, err := strconv.Atoi(id)
	if err != nil {
		return m.Book{}, "", err
	}

	row := b.DB.QueryRowContext(ctx, query, updatedBook.Title, updatedBook.AuthorName, updatedBook.PublicationDate, updatedBook.ISBN, updatedBook.Description, intValue)

	var newBook m.Book
	err = row.Scan(
		&newBook.BookID,
		&newBook.Title,
		&newBook.AuthorID,
		&newBook.PublicationDate,
		&newBook.ISBN,
		&newBook.Description,
		&newBook.CreatedAt,
		&newBook.UpdatedAt,
	)
	if err != nil {
		return m.Book{}, "", err
	}

	return newBook, updatedBook.AuthorName, nil
}

func (b *BookRepo) DeleteBookByID(ctx context.Context, id string) error {
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid book ID: " + err.Error())
	}

	query := `
		DELETE FROM books
		WHERE book_id = $1;
	`

	result, err := b.DB.ExecContext(ctx, query, bookID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no book found with the given ID")
	}

	return nil
}

func (b *BookRepo) GetBookByTitle(ctx context.Context, title string) (m.Book, error) {
	query := `
        SELECT books.book_id, books.title, authors.name AS author_name, books.publication_date, books.isbn, books.description, books.created_at, books.updated_at
        FROM books
        INNER JOIN authors ON books.author_id = authors.author_id
        WHERE books.title = $1;
    `

	row := b.DB.QueryRowContext(ctx, query, title)

	book := m.Book{}

	err := row.Scan(&book.BookID, &book.Title, &book.AuthorName, &book.PublicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}
