package koans

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

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

type Koans struct {
	db *sql.DB
}

func New() (*Koans, error) {
	var url string
	var db *sql.DB
	var err error
	if url, err = DbUrl(); err != nil {
		log.Printf("failed to construct sqlite url: %v\n", err)
		return nil, err
	}
	if db, err = Setup(url); err != nil {
		log.Printf("failed to open sqlite3 conn: %v\n", err)
		return nil, err
	}
	return &Koans{
		db: db,
	}, nil
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
	// not every PRAGMA has a DSN equivalent
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
	if _, err = db.Exec(CreateStrictTableStmt); err != nil {
		log.Printf("%q: %s\n", err, CreateStrictTableStmt)
		return db, err
	}
	return db, err
}

func Teardown(t *testing.T) error {
	err := os.Remove(DbFile)
	// if t != nil {
	// 	assert.Nil(t, err)
	// }
	return err
}
