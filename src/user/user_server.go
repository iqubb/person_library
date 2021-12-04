package main

import (
	"context"
	pb "github.com/iqubb/src/user/proto"
	"github.com/iqubb/src/user/store"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	repository store.Repository
}

func (s *UserServer) GetUser(ctx context.Context, usReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	result, _ := s.repository.GetUser(ctx, usReq.Id)
	user := store.UnparsingUser(result)
	return &pb.GetUserResponse{
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		Password:     user.Password,
		OrderedBooks: user.OrderedBooks,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, newReq *pb.User) (*pb.CreateUserResponse, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(newReq.Password), bcrypt.DefaultCost)
	newReq.Password = string(pass)
	_ = s.repository.CreateUser(ctx, store.ParsingUser(newReq))
	return &pb.CreateUserResponse{
		User: newReq,
	}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, allReq *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	results, _ := s.repository.GetAllUsers(ctx)
	users := store.UnparsingAllUsers(results)
	return &pb.GetAllUsersResponse{
		AllUsers: users,
	}, nil
}
