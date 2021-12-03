package main

import (
	"context"
	pb "github.com/iqubb/src/book/proto"
)

type BookServer struct {
	//add DB
	pb.UnimplementedBookServiceServer
}

func (s *BookServer) GetBook(ctx context.Context, bReq *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	//add logic
	return nil, nil
}

func (s *BookServer) CreateBook(ctx context.Context, newReq *pb.Book) (*pb.CreateBookResponse, error) {
	//add logic
	return nil, nil
}

func (s *BookServer) GetAllBooks(ctx context.Context, allReq *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	//add logic
	return nil, nil
}

func (s *BookServer) GetAllBooksByAuthor(ctx context.Context, auReq *pb.GetBooksByAuthorRequest) (*pb.GetBooksByAuthorResponse, error) {
	//add logic
	return nil, nil
}
