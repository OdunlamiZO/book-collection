package database

import (
	"database/sql"
	"os"
	"strings"

	"odunlamizo/book-collection/internal/model"
	"odunlamizo/book-collection/internal/util"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func Initialize() error {
	var err error
	if db, err = sql.Open("pgx", os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	return nil
}

func GetBooksByTitle(title string) ([]model.Book, error) {
	var query string = "SELECT * FROM books WHERE lower(title) = $1"
	rows, err := db.Query(query, strings.ToLower(title))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractBooks(rows)
}

func GetBooksByAuthor(author string) ([]model.Book, error) {
	var query string = "SELECT * FROM books WHERE lower(author) = $1 OR $2::varchar[] <@ lower(co_authors::varchar)::varchar[]"
	rows, err := db.Query(query, strings.ToLower(author), util.ArrayToString([]string{strings.ToLower(author)}))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return extractBooks(rows)
}

func AddBook(book *model.Book) error {
	var query string = "INSERT INTO books (title, url, author, co_authors) VALUES ($1, $2, $3, $4::varchar[]) ON CONFLICT (title, author, co_authors) DO NOTHING RETURNING id"
	if err := db.QueryRow(query, book.Title, book.Url, book.Author, util.ArrayToString(book.CoAuthors)).Scan(&book.Id); err != nil {
		return err
	}
	return nil
}

func DeleteBookById(id int) (*model.Book, error) {
	var book model.Book
	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}
	var query string = "DELETE FROM books WHERE id = $1 RETURNING *"
	if err := txn.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Url, &book.Author, &book.CoAuthors); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return &book, nil
}

func extractBooks(rows *sql.Rows) ([]model.Book, error) {
	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Url, &book.Author, &book.CoAuthors); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
