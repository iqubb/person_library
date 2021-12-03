package main

import (
	"context"
	_ "github.com/iqubb/src/book/proto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Book struct {
	ID      int64  `sql:"id"`
	Title   string `sql:"title"`
	Author  string `sql:"author"`
	isOrder bool   `sql:"is_order"`
}

type PostgresRepository struct {
	db *sqlx.DB
}

func (repos *PostgresRepository) GetBook(ctx context.Context, id int64) (*Book, error) {
	var book *Book
	if err := repos.db.GetContext(ctx, &book, "select * from books where id = $1", id); err != nil {
		return nil, err
	}
	return book, nil
}

func (repos *PostgresRepository) CreateBook(ctx context.Context, book *Book) error {
	log.Println(book)
	query := "insert into books (id, title, author, isOrder) values ($1, $2, $3, $4, $5)"
	_, err := repos.db.ExecContext(ctx, query, book.ID, book.Title, book.Author, book.isOrder)
	return err
}

func (repos *PostgresRepository) GetAllBooks(ctx context.Context) ([]*Book, error) {
	books := make([]*Book, 0)
	if err := repos.db.GetContext(ctx, books, "select * from books"); err != nil {
		return books, err
	}
	return books, nil
}

func (repos *PostgresRepository) GetBooksByAuthor(ctx context.Context, author string) ([]*Book, error) {
	books := make([]*Book, 0)
	if err := repos.db.GetContext(ctx, books, "select * from books where author = $1", author); err != nil {
		return books, err
	}
	return books, nil
}
