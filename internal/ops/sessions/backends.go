package sessions

import "database/sql"

// Backends used for dependency injection.
type Backends interface {
	DB() *sql.DB
}
