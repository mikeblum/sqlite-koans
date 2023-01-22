package koans

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	UpsertRecordStmt = `
	INSERT INTO test_strict(name,created) VALUES(?, ?)
	ON CONFLICT(name) DO
		UPDATE SET updated=now();`
)

func UpsertRecord(db *sql.DB) error {
	var stmt *sql.Stmt
	var tx *sql.Tx
	var err error
	if tx, err = db.Begin(); err != nil {
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
