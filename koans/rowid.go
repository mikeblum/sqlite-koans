package koans

const (
	// STRICT must be appended to every CREATE TABLE
	// https://www.sqlite.org/stricttables.html
	// WITHOUT ROWID avoids null primary keys (CHECK needed for empty strings)
	// https://www.sqlite.org/withoutrowid.html
	TableTestWithoutRowIdStrict       = "test_without_rowid_strict"
	CreateStrictWithoutRowIdTableStmt = `
	CREATE TABLE IF NOT EXISTS test_without_rowid_strict (
		id text NOT NULL PRIMARY KEY CHECK (length(trim(id)) > 0), 
		name TEXT,
		created INTEGER,
		updated INTEGER
	) WITHOUT ROWID, STRICT;
	`
)
