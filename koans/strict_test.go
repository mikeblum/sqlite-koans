package koans

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (k *KoansTest) StrictTablesTest(t *testing.T) {
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
		if table.Skip() {
			continue
		}
		assert.Equal(t, 1, table.Strict, fmt.Sprintf("table: %s.%s STRICT: 1 if the table is a STRICT table or 0 if it is not.", table.Schema, table.Name))
		n = n + 1
	}
	assert.True(t, n > 0)
}
