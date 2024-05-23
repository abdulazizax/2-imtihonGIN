package postgres

import (
	"context"
	"errors"
	m "exam/models"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthorRepo struct {
	DB *sqlx.DB
}

func NewAuthorrepo(db *sqlx.DB) *AuthorRepo {
	return &AuthorRepo{
		DB: db,
	}
}

func (a *AuthorRepo) CreateAuthor(ctx context.Context, aut m.Author) (string, error) {
	query := `
        INSERT INTO authors (
            name, 
            birth_date, 
            biography)
        VALUES($1, $2, $3)
        RETURNING author_id;
    `

	var id uuid.UUID

	row := a.DB.QueryRowContext(ctx, query, strings.ToLower(aut.Name), aut.BirthDate, aut.Biography)
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (a *AuthorRepo) GetAuthors(ctx context.Context) ([]m.Author, error) {
	query := `
        SELECT author_id, name, birth_date, biography, created_at, updated_at
        FROM authors;
    `
	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []m.Author

	for rows.Next() {
		var author m.Author
		err := rows.Scan(&author.AuthorID, &author.Name, &author.BirthDate, &author.Biography, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *AuthorRepo) GetAuthorById(ctx context.Context, id string) (m.Author, error) {
	query := `
		SELECT author_id, name, birth_date, biography, created_at, updated_at
		FROM authors
		WHERE author_id = $1;
    `
	row := a.DB.QueryRowContext(ctx, query, id)

	author := m.Author{}

	err := row.Scan(&author.AuthorID, &author.Name, &author.BirthDate, &author.Biography, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return author, err
	}

	return author, nil
}

func (a *AuthorRepo) UpdateAuthorByID(ctx context.Context, id string, author m.Author) (m.Author, error) {
	query := `
        UPDATE authors
        SET name = $2, 
		birth_date = $3, 
		biography = $4, 
		updated_at = CURRENT_TIMESTAMP
        WHERE author_id = $1
		RETURNING author_id, name, birth_date, biography, created_at, updated_at;
    `

	var updatedAuthor m.Author

	err := a.DB.QueryRowContext(ctx, query, id, author.Name, author.BirthDate, author.Biography).Scan(
		&updatedAuthor.AuthorID,
		&updatedAuthor.Name,
		&updatedAuthor.BirthDate,
		&updatedAuthor.Biography,
		&updatedAuthor.CreatedAt,
		&updatedAuthor.UpdatedAt,
	)
	if err != nil {
		return m.Author{}, err
	}

	return updatedAuthor, nil
}

func (a *AuthorRepo) DeleteAuthorByID(ctx context.Context, id string) error {
	query := `
        DELETE FROM authors
        WHERE author_id = $1;
    `

	result, err := a.DB.ExecContext(ctx, query, id)
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

func (a *AuthorRepo) GetAuthorByName(ctx context.Context, name string) (m.Author, error) {
	query := `
		SELECT author_id, name, birth_date, biography, created_at, updated_at
		FROM authors
		WHERE name = $1;
	`
	row := a.DB.QueryRowContext(ctx, query, name)

	author := m.Author{}

	err := row.Scan(&author.AuthorID, &author.Name, &author.BirthDate, &author.Biography, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return author, err
	}

	return author, nil
}