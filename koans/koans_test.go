package koans

import "testing"

// koans_test executes all koan tests

// TestKoans: test suite runner
func TestKoans(t *testing.T) {
	koans, teardown := SetupSuite(t)
	defer teardown(t)
	t.Run("koan=strict", koans.StrictTablesTest)
	t.Run("koan=rowid", koans.WithoutRowIdStrictTablesTest)
	t.Run("koan=timeout", koans.BusyTimeoutTest)
	t.Run("koan=encoding", koans.EncodingTest)
	t.Run("koan=foreign_keys", koans.ForeignKeysTest)
	t.Run("koan=synchronous", koans.SynchronousTest)
	t.Run("koan=journal_mode", koans.JournalModeTest)
}
