package koans

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/google/uuid"
)

// koanstest
// common sql and functions across the koan tests

var koanstest *KoansTest

const (
	PragmaTableList  = "PRAGMA main.table_list;"
	UpsertRecordStmt = `
	INSERT INTO %s(id, name) VALUES(?, ?)
	ON CONFLICT DO
		UPDATE SET updated=datetime('unixepoch');`
)

type TableListRaw struct {
	// schema: the schema in which the table or view appears (for example "main" or "temp").
	Schema string
	// name: the name of the table or view.
	Name string
	// type: the type of object - one of "table", "view", "shadow" (for shadow tables), or "virtual" for virtual tables.
	Type string
	// ncol: the number of columns in the table, including generated columns and hidden columns.
	NCol int
	// wr: 1 if the table is a WITHOUT ROWID table or 0 if is not.
	WR int
	// strict: 1 if the table is a STRICT table or 0 if it is not.
	Strict int
}

func (t *TableListRaw) TableName() string {
	return strings.ToLower(strings.Join([]string{
		t.Schema,
		t.Name,
	}, "."))
}

// omit SQLite internal tables
// kv: schema.name -> skip == true
func (t *TableListRaw) Skip() bool {
	skipTables := map[string]bool{
		"main.sqlite_sequence": true,
		"main.sqlite_schema":   true,
	}
	skip, ok := skipTables[t.TableName()]
	return skip && ok
}

type KoansTest struct {
	Koans
}

func SetupSuite() (*KoansTest, func(t *testing.T) error, error) {
	if koanstest != nil {
		return koanstest, Teardown, nil
	}
	koans, err := New()
	koanstest = &KoansTest{
		*koans,
	}
	return koanstest, Teardown, err
}

func (k *KoansTest) UpsertRecord(b *testing.B) error {
	var stmt *sql.Stmt
	var tx *sql.Tx
	var err error
	if tx, err = k.db.Begin(); err != nil {
		return err
	}
	upsertStmt := fmt.Sprintf(UpsertRecordStmt, TableTestWithoutRowIdStrict)
	if stmt, err = tx.Prepare(upsertStmt); err != nil {
		log.Printf("failed to prepare stmt: %q: %s\n", err, upsertStmt)
		return err
	}
	defer stmt.Close()
	id := uuid.Must(uuid.NewRandom()).String()
	for i := 0; i < b.N; i++ {
		if _, err = stmt.Exec(id, i); err != nil {
			log.Printf("failed to insert record: %v", err)
			return err
		}
	}
	return tx.Commit()
}
