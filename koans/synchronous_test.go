package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Synchronous struct {
	// PRAGMA schema.synchronous = 0 | OFF | 1 | NORMAL | 2 | FULL | 3 | EXTRA;
	Mode int
}

func (k *KoansTest) SynchronousTest(t *testing.T) {
	rows, err := k.db.Query(PragmaSynchronous)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		synchronous := Synchronous{}
		err := rows.Scan(&synchronous.Mode)
		assert.Nil(t, err)
		defer rows.Close()
		assert.Equal(t, 1, synchronous.Mode)
		n = n + 1
	}
	assert.True(t, n > 0)
}
