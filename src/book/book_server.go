package main

import (
	"context"
	"github.com/iqubb/src/book/book_store"
	pb "github.com/iqubb/src/book/proto"
)

type BookServer struct {
	repository book_store.Repository
}

func (s *BookServer) GetBook(ctx context.Context, bReq *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	result, _ := s.repository.GetBook(ctx, bReq.Id)
	book := book_store.UnparsingBook(result)
	return &pb.GetBookResponse{
		Title:   book.GetTitle(),
		Author:  book.GetAuthor(),
		IsOrder: book.IsOrder,
	}, nil
}

func (s *BookServer) CreateBook(ctx context.Context, newReq *pb.Book) (*pb.CreateBookResponse, error) {
	_ = s.repository.CreateBook(ctx, book_store.ParsingBook(newReq))
	return &pb.CreateBookResponse{
		Book: newReq,
	}, nil
}

func (s *BookServer) GetAllBooks(ctx context.Context, allReq *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	books_, _ := s.repository.GetAllBooks(ctx)
	books := book_store.UnparsingAllBooks(books_)
	return &pb.GetAllBooksResponse{
		AllBooks: books,
	}, nil
}

func (s *BookServer) GetBooksByAuthor(ctx context.Context, auReq *pb.GetBooksByAuthorRequest) (*pb.GetBooksByAuthorResponse, error) {
	books_, _ := s.repository.GetBooksByAuthor(ctx, auReq.Author)
	books := book_store.UnparsingAllBooks(books_)
	return &pb.GetBooksByAuthorResponse{
		Books: books,
	}, nil
}
