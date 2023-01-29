package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// koans_test executes all koan tests

// TestKoans: test suite runner
func TestKoans(t *testing.T) {
	koans, teardown, err := SetupSuite()
	defer teardown(t)
	assert.Nil(t, err)
	t.Run("koan=strict", koans.StrictTablesTest)
	t.Run("koan=rowid", koans.WithoutRowIdStrictTablesTest)
	t.Run("koan=primary_keys", func(t *testing.T) {
		koans.EmptyPrimaryKeyTest(t)
	})
	t.Run("koan=timeout", koans.BusyTimeoutTest)
	t.Run("koan=encoding", koans.EncodingTest)
	t.Run("koan=foreign_keys", koans.ForeignKeysTest)
	t.Run("koan=synchronous", koans.SynchronousTest)
	t.Run("koan=journal_mode", koans.JournalModeTest)
	t.Run("koan=backup", koans.BackupWALTest)
}

func BenchmarkKoans(b *testing.B) {
	koans, teardown, err := SetupSuite()
	defer teardown(nil)
	assert.Nil(b, err)
	b.Run("koan=upsert", koans.UpsertRecordsBench)
}
