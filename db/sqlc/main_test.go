package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbsource = "postgresql://postgres:postgres@postgres:5432/go_api_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbsource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	runMigrations()   // ensures version table + schema
	truncateAllData() // clean data, keep schema + version table
	testQueries = New(testDB)

	code := m.Run()
	os.Exit(code)
}

func runMigrations() {
	m, err := migrate.New(
		"file://../migrations", // adjust if your folder differs
		dbsource,
	)
	if err != nil {
		log.Fatalf("cannot create migrate instance: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot apply migrations: %v", err)
	}
	log.Println("migrations applied (schema_migrations retained)")
}

func truncateAllData() {
	const q = `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    (SELECT tablename FROM pg_tables
     WHERE schemaname='public'
       AND tablename <> 'schema_migrations')
  LOOP
    EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE';
  END LOOP;
END $$;`
	if _, err := testDB.Exec(q); err != nil {
		log.Fatalf("cannot truncate tables: %v", err)
	}
	log.Println("data truncated (schema + version table kept)")
}
