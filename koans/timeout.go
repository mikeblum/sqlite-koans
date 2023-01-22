package koans

// PRAGMA busy_timeout;
// PRAGMA busy_timeout = milliseconds;

// Query or change the setting of the busy timeout.

// Each database connection can only have a single busy handler.
// This PRAGMA sets the busy handler for the process, possibly overwriting any previously set busy handler.

const (
	PragmaTimeoutMs   = 5000
	PragmaTimeoutStmt = "PRAGMA busy_timeout = %d;"
	PragmaTimeout     = "PRAGMA busy_timeout;"
)
