package koans

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JournalMode struct {
	// PRAGMA schema.journal_mode = DELETE | TRUNCATE | PERSIST | MEMORY | WAL | OFF
	Mode string
}

func (k *KoansTest) JournalModeTest(t *testing.T) {
	rows, err := k.db.Query(PragmaJournalMode)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		journalMode := JournalMode{}
		err := rows.Scan(&journalMode.Mode)
		assert.Nil(t, err)
		defer rows.Close()
		assert.Equal(t, strings.ToLower(JournalModeWAL), journalMode.Mode)
		n = n + 1
	}
	assert.True(t, n > 0)
}
