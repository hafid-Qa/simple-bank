package main

import (
	"database/sql"
	"go_api/api"
	db "go_api/db/sqlc"
	"go_api/utils"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("Can not create server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Can not start the server", err)
	}
}
