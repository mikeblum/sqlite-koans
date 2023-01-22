package main

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	PragmaForeignKeys = true
	PragmaJournalMode = "WAL"
	PragmaSynchronous = "NORMAL"
	PragmaTimeout     = 5000
	SqliteCmd         = "sqlite3"
	DbFile            = "file:koans.db"
)

func main() {
	var url string
	var db *sql.DB
	var err error
	teardown()
	if url, err = dbUrl(); err != nil {
		log.Fatalf("failed to construct sqlite url: %v", err)
	}
	if db, err = setup(url); err != nil {
		log.Fatalf("failed to open sqlite3 conn: %v", err)
	}
	defer db.Close()
	createTableStmt := `
	CREATE TABLE IF NOT EXISTS test_strict (id integer not null primary key, name text) STRICT;
	`
	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, createTableStmt)
		return
	}
}

// dbUrl - construct a Sqlite DSN (Data Source Name) string
// https://github.com/mattn/go-sqlite3#connection-string
func dbUrl() (string, error) {
	var dataSourceName *url.URL
	var err error
	if dataSourceName, err = url.Parse(DbFile); err != nil {
		return "", err
	}
	pragmas := url.Values{}
	// mattn/go-sqlite3 DSN keys
	pragmas.Set("_busy_timeout", strconv.Itoa(PragmaTimeout))
	pragmas.Set("_foreign_keys", strconv.FormatBool(PragmaForeignKeys))
	pragmas.Set("_journal_mode", PragmaJournalMode)
	pragmas.Set("_synchronous", PragmaSynchronous)
	dataSourceName.RawQuery = pragmas.Encode()
	log.Println(dataSourceName.String())
	return dataSourceName.String(), err
}

func setup(dbUrl string) (*sql.DB, error) {
	return sql.Open(SqliteCmd, dbUrl)
}

func teardown() {
	os.Remove(DbFile)
}
