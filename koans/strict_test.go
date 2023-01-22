package koans

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	PragmaTableList = "PRAGMA main.table_list;"
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

func TestStrictTables(t *testing.T) {
	koans, teardown := SetupSuite(t)
	defer teardown(t)
	rows, err := koans.db.Query(PragmaTableList)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	// omit SQLite internal tables
	// kv: schema.name -> skip == true
	skipTables := map[string]bool{
		"main.sqlite_sequence": true,
		"main.sqlite_schema":   true,
	}
	for rows.Next() {
		table := TableListRaw{}
		err := rows.Scan(&table.Schema,
			&table.Name,
			&table.Type,
			&table.NCol,
			&table.WR,
			&table.Strict)
		assert.Nil(t, err)
		if skip, ok := skipTables[table.TableName()]; skip && ok {
			continue
		}
		assert.Equal(t, 1, table.Strict, fmt.Sprintf("table: %s.%s STRICT: 1 if the table is a STRICT table or 0 if it is not.", table.Schema, table.Name))
		n = n + 1
	}
	assert.True(t, n > 0)
}
