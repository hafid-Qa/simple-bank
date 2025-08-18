package db

import (
	"database/sql"
	"go_api/utils"
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

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.TestDbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	DropAllTables()
	// truncateAllData()                    // clean data, keep schema + version table
	runMigrations(config.TestDbSource) // ensures version table + schema
	testQueries = New(testDB)

	code := m.Run()
	os.Exit(code)
	DropAllTables()
}

func testDBSource() string {
	// If you use config, you can cache it or load once
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	return config.TestDbSource
}
func resetDB() {
	DropAllTables()
	runMigrations(testDBSource())
}

func runMigrations(dbSource string) {
	m, err := migrate.New(
		"file://../migrations", // adjust if your folder differs
		dbSource,
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

func DropAllTables() {
	const q = `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    (SELECT tablename FROM pg_tables
     WHERE schemaname='public'
      )
  LOOP
    EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
  END LOOP;
END $$;`
	if _, err := testDB.Exec(q); err != nil {
		log.Fatalf("cannot drop tables: %v", err)
	}
	log.Println("all tables dropped (schema + version table kept)")
}
