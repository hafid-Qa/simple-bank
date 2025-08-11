package main

import (
	"database/sql"
	"go_api/api"
	db "go_api/db/sqlc"
	"go_api/utils"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbsource      = "postgresql://postgres:postgres@localhost:5432/go_api_db?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Can not start the server", err)
	}
}
