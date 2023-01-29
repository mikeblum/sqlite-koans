package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func (k *KoansTest) BackupWALTest(t *testing.T) {
	assert.Nil(t, k.db.Ping())
	err := k.Truncate()
	assert.Nil(t, err)
	k.InsertRecords(&testing.B{
		N: 1000,
	})
	assert.Nil(t, k.Koans.Backup())
}
