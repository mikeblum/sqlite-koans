package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ForeignKeys struct {
	// foreign key constraints: 1 if foreign keys are enforced or 0 if it is not.
	Enabled int
}

func (k *KoansTest) ForeignKeysTest(t *testing.T) {
	rows, err := k.db.Query(PragmaForeignKeys)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		foreignKeys := ForeignKeys{}
		err := rows.Scan(&foreignKeys.Enabled)
		assert.Nil(t, err)
		defer rows.Close()
		assert.Equal(t, 1, foreignKeys.Enabled)
		n = n + 1
	}
	assert.True(t, n > 0)
}
