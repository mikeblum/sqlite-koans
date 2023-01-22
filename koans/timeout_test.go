package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BusyTimeout struct {
	milliseconds int
}

func (k *KoansTest) BusyTimeoutTest(t *testing.T) {
	rows, err := k.db.Query(PragmaTimeout)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		timeout := BusyTimeout{}
		err := rows.Scan(&timeout.milliseconds)
		assert.Nil(t, err)
		defer rows.Close()
		assert.Equal(t, PragmaTimeoutMs, timeout.milliseconds)
		n = n + 1
	}
	assert.True(t, n > 0)
}
