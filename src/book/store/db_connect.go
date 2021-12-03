package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "postgres"
	password      = "gogol"
	database_name = "books"
)

func CreateConnection() (*sqlx.DB, error) {
	config := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)
	db, err := sqlx.Connect("postgres", config)
	if err != nil {
		return nil, err
	}
	return db, nil
}
