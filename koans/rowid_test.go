package koans

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (k *KoansTest) WithoutRowIdStrictTablesTest(t *testing.T) {
	rows, err := k.db.Query(PragmaTableList)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		table := TableListRaw{}
		err := rows.Scan(&table.Schema,
			&table.Name,
			&table.Type,
			&table.NCol,
			&table.WR,
			&table.Strict)
		assert.Nil(t, err)
		expected := strings.ToLower(strings.Join([]string{
			table.Schema,
			TableTestWithoutRowIdStrict,
		}, "."))
		if table.Skip() || !strings.EqualFold(expected, table.TableName()) {
			continue
		}
		assert.Equal(t, 1, table.WR, fmt.Sprintf("table: %s.%s WITHOUT ROWID: 1 if the table is a WITHOUT ROWID table or 0 if it is not.", table.Schema, table.Name))
		assert.Equal(t, 1, table.Strict, fmt.Sprintf("table: %s.%s STRICT: 1 if the table is a STRICT table or 0 if it is not.", table.Schema, table.Name))
		n = n + 1
	}
	assert.True(t, n > 0)
}

func (k *KoansTest) EmptyPrimaryKeyTest(t *testing.T) {
	_, err := k.db.Exec(fmt.Sprintf(`INSERT INTO %s(id, name) VALUES(?, ?);`, TableTestWithoutRowIdStrict), "", "test")
	assert.NotNil(t, err)
	expected := "CHECK constraint failed: length(trim(id)) > 0"
	assert.EqualError(t, err, expected)
}

func (k *KoansTest) UpsertRecordsBench(b *testing.B) {
	// cleanup TableTestWithoutRowIdStrict
	_, err := k.db.Exec(fmt.Sprintf("DELETE FROM %s", TableTestWithoutRowIdStrict))
	assert.Nil(b, err)
	err = k.UpsertRecord(b)
	assert.Nil(b, err)
	result, err := k.db.Query(fmt.Sprintf(`SELECT COUNT(1) FROM %s;`, TableTestWithoutRowIdStrict))
	assert.Nil(b, err)
	count := 0
	for result.Next() {
		err := result.Scan(&count)
		assert.Nil(b, err)
	}
	assert.Equal(b, 1, count)
}
