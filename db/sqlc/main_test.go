package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // Import the PostgreSQL driver

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Replace with your database driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbsource = "postgresql://postgres:postgres@localhost:5433/go_api_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbsource)
	if err != nil {
		{
			log.Fatalf("cannot connect to db: %v", err)
		}
	}
	// Run migrations
	runMigrations()
	testQueries = New(conn)
	code := m.Run()
	os.Exit(code)
}

func runMigrations() {
	migration, err := migrate.New(
		"file://../migrations", // adjust if needed
		"postgres://postgres:postgres@localhost:5433/go_api_db?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("cannot create migrate instance: %v", err)
	}
	defer migration.Close()

	// Drop the database if it exists
	if err := migration.Down(); err != nil && err != migrate.ErrNoChange {
		log.Printf("cannot drop database: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot run migrate up: %v", err)
	}

	log.Println("migrations completed successfully")
}
