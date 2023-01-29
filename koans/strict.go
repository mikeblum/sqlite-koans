package koans

const (
	// STRICT must be appended to every CREATE TABLE
	// https://www.sqlite.org/stricttables.html
	CreateStrictTableStmt = `
	CREATE TABLE IF NOT EXISTS test_strict (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL,
		created INTEGER DEFAULT (datetime('unixepoch')),
		updated INTEGER
	) STRICT;
	`
)
