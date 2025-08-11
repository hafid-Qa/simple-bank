package main

import (
	"database/sql"
	"go_api/api"
	db "go_api/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbsource      = "postgresql://postgres:postgres@localhost:5432/go_api_db?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbsource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Can not start the server", err)
	}
}
