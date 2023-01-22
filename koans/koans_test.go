package koans

import "testing"

// koans_test executes all koan tests

// TestKoans: test suite runner
func TestKoans(t *testing.T) {
	koans, teardown := SetupSuite(t)
	defer teardown(t)
	t.Run("koan=strict", koans.StrictTablesTest)
}
