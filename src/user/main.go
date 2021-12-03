package main

import (
	pb "github.com/iqubb/src/user/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = "5001"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
