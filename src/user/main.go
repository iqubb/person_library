package main

import (
	"github.com/iqubb/src/user/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ""
)

func main() {

	db, err := store.CreateConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	//repository := store.NewPostgresRepository(db)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
