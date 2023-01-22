package main

import (
	"database/sql"
	"log"

	"github.com/mikeblum/sqlite-koans/koans"
)

func main() {
	var url string
	var db *sql.DB
	var err error
	koans.Teardown(nil)
	if url, err = koans.DbUrl(); err != nil {
		log.Fatalf("failed to construct sqlite url: %v", err)
	}
	if db, err = koans.Setup(url); err != nil {
		log.Fatalf("failed to open sqlite3 conn: %v", err)
	}
	defer db.Close()
}
