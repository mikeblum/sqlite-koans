package main

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mikeblum/sqlite-koans/koans"
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
	Teardown()
	if url, err = DbUrl(); err != nil {
		log.Fatalf("failed to construct sqlite url: %v", err)
	}
	if db, err = Setup(url); err != nil {
		log.Fatalf("failed to open sqlite3 conn: %v", err)
	}
	defer db.Close()
}

// DbUrl - construct a Sqlite DSN (Data Source Name) string
// https://github.com/mattn/go-sqlite3#connection-string
func DbUrl() (string, error) {
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

func Setup(dbUrl string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if db, err = sql.Open(SqliteCmd, dbUrl); err != nil {
		return nil, err
	}
	defer db.Close()
	if _, err = db.Exec(koans.CreateStrictTableStmt); err != nil {
		log.Printf("%q: %s\n", err, koans.CreateStrictTableStmt)
		return db, err
	}
	return db, err
}

func Teardown() {
	os.Remove(DbFile)
}
