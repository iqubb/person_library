package book_store

import (
	"context"
	_ "github.com/iqubb/src/book/proto"
	pb "github.com/iqubb/src/book/proto"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id      int64  `sql:"id"`
	Title   string `sql:"title"`
	Author  string `sql:"author"`
	isOrder bool   `sql:"is_order"`
}

type MongoRepository struct {
	db *mongo.Collection
}

type Repository interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBook(ctx context.Context, id int64) (*Book, error)
	GetAllBooks(ctx context.Context) ([]*Book, error)
	GetBooksByAuthor(ctx context.Context, author string) ([]*Book, error)
}

func (repos *MongoRepository) CreateBook(ctx context.Context, book *Book) {
	repos.db.InsertOne(ctx, book)
}

func (repos *MongoRepository) GetBook(ctx context.Context, id int64) *Book {
	var book *Book
	cursor, _ := repos.db.Find(ctx, bson.M{"id": id})
	cursor.All(ctx, &book)
	return book
}

func (repos *MongoRepository) GetAllBooks(ctx context.Context) []*Book {
	books := make([]*Book, 0)
	cursor, _ := repos.db.Find(ctx, bson.M{})
	cursor.All(ctx, &books)
	return books
}

func (repos *MongoRepository) GetBooksByAuthor(ctx context.Context, author string) []*Book {
	books := make([]*Book, 0)
	cursor, _ := repos.db.Find(ctx, bson.M{"author": author})
	cursor.All(ctx, &books)
	return books
}

func ParsingBook(book *pb.Book) *Book {
	return &Book{
		Id:      book.Id,
		Title:   book.Title,
		Author:  book.Author,
		isOrder: book.IsOrder,
	}
}

func UnparsingBook(book *Book) *pb.Book {
	return &pb.Book{
		Id:      book.Id,
		Title:   book.Title,
		Author:  book.Author,
		IsOrder: book.isOrder,
	}
}

func ParsingAllBooks(books []*pb.Book) []*Book {
	result := make([]*Book, len(books))
	for _, val := range books {
		result = append(result, ParsingBook(val))
	}
	return result
}

func UnparsingAllBooks(books []*Book) []*pb.Book {
	result := make([]*pb.Book, len(books))
	for _, val := range books {
		result = append(result, UnparsingBook(val))
	}
	return result
}
