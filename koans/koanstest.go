package koans

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// koanstest
// common sql and functions across the koan tests

var koanstest *KoansTest

const (
	UpsertRecordStmt = `
	INSERT INTO test_strict(name,created) VALUES(?, ?)
	ON CONFLICT(name) DO
		UPDATE SET updated=now();`
)

type KoansTest struct {
	Koans
}

func SetupSuite(t *testing.T) (*KoansTest, func(t *testing.T) error) {
	if koanstest != nil {
		return koanstest, Teardown
	}
	koans, err := New()
	assert.Nil(t, err)
	koanstest = &KoansTest{
		*koans,
	}
	return koanstest, Teardown
}

func (k *KoansTest) UpsertRecord() error {
	var stmt *sql.Stmt
	var tx *sql.Tx
	var err error
	if tx, err = k.db.Begin(); err != nil {
		return err
	}
	if stmt, err = tx.Prepare(UpsertRecordStmt); err != nil {
		log.Printf("failed to prepare stmt: %q: %s\n", err, UpsertRecordStmt)
		return err
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		if _, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i)); err != nil {
			log.Printf("failed to insert record: %v", err)
			return err
		}

	}
	return tx.Commit()
}
