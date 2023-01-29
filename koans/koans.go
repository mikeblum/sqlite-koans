package koans

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

const (
	DbPrefix      = "file:"
	DbName        = "koans.db"
	BackupDbName  = "koans-backup.db"
	DbFile        = DbPrefix + DbName
	DbBackupFile  = DbPrefix + BackupDbName
	DbFileShm     = DbFile + "-shm"
	DbFileWal     = DbFile + "-wal"
	SqliteCmd     = "sqlite3"
	SqliteSchema  = "main"
	SqliteConnErr = "failed to open db connection to %s"
)

var DbFiles []string = []string{
	DbFile, DbFileShm, DbFileWal, DbBackupFile,
}

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
	pragmas.Set("_busy_timeout", strconv.Itoa(PragmaTimeoutMs))
	pragmas.Set("_foreign_keys", strconv.FormatBool(true))
	pragmas.Set("_journal_mode", JournalModeWAL)
	pragmas.Set("_synchronous", PragmaSynchronousNormal)
	dataSourceName.RawQuery = pragmas.Encode()
	log.Println(dataSourceName.String())
	return dataSourceName.String(), err
}

// BackupDBUrl - construct a Sqlite DSN (Data Source Name) string as a backup target
// https://www.sqlite.org/backup.html
func BackupDbUrl() (string, error) {
	var dataSourceName *url.URL
	var err error
	if dataSourceName, err = url.Parse(DbBackupFile); err != nil {
		return "", err
	}
	pragmas := url.Values{}
	// mattn/go-sqlite3 DSN keys
	// not every PRAGMA has a DSN equivalent
	pragmas.Set("_busy_timeout", strconv.Itoa(PragmaTimeoutMs))
	pragmas.Set("_foreign_keys", strconv.FormatBool(true))
	pragmas.Set("_journal_mode", JournalModeDelete)
	pragmas.Set("_synchronous", PragmaSynchronousFull)
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
	if _, err = db.Exec(CreateStrictWithoutRowIdTableStmt); err != nil {
		log.Printf("%q: %s\n", err, CreateStrictWithoutRowIdTableStmt)
		return db, err
	}
	return db, err
}

func Teardown(t *testing.T) error {
	var err error
	for i := range DbFiles {
		file := DbFiles[i]
		err = os.Remove(strings.Replace(file, DbPrefix, "", 1))
		if t != nil {
			assert.Nil(t, err)
		}
	}
	return err
}

// Backup - backup WAL mode SQLite DB
// https://github.com/mattn/go-sqlite3/blob/master/_example/hook/hook.go
func (k *Koans) Backup() error {
	src, err := k.db.Conn(context.Background())
	if err != nil {
		return err
	}
	backupDbUrl, err := BackupDbUrl()
	if err != nil {
		return err
	}
	var backupDb *sql.DB
	backupDb, err = sql.Open(SqliteCmd, backupDbUrl)
	if err != nil {
		return err
	}
	target, err := backupDb.Conn(context.Background())
	if err != nil {
		return err
	}
	return target.Raw(func(destConn any) error {
		return src.Raw(func(srcConn any) error {
			var srcDriver *sqlite3.SQLiteConn
			var destDriver *sqlite3.SQLiteConn
			var backup *sqlite3.SQLiteBackup
			var ok bool
			var err error
			if destDriver, ok = destConn.(*sqlite3.SQLiteConn); !ok {
				return fmt.Errorf(SqliteConnErr, DbBackupFile)
			}
			if srcDriver, ok = srcConn.(*sqlite3.SQLiteConn); !ok {
				return fmt.Errorf(SqliteConnErr, DbFile)
			}
			if backup, err = destDriver.Backup(SqliteSchema, srcDriver, SqliteSchema); err != nil {
				return err
			}
			if _, err = backup.Step(-1); err != nil {
				return err
			}
			return backup.Finish()
		})
	})
}
