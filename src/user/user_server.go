package main

import (
	"context"
	pb "github.com/iqubb/src/user/proto"
)

type UserServer struct {
	//add DB
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, usReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	//add logic
	return nil, nil
}

func (s *UserServer) Create(ctx context.Context, newReq *pb.User) (*pb.CreateUserResponse, error) {
	//add logic
	return nil, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, allReq *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	//add logic
	return nil, nil
}