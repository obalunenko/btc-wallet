package state

import (
	"database/sql"

	"github.com/oleg-balunenko/btc-wallet/internal/db"
)

// State holds state specific information
// for a the service.
type State interface {
	DB() *sql.DB
}

type state struct {
	db *sql.DB
}

// DB returns DB connection.
func (s *state) DB() *sql.DB {
	return s.db
}

// New creates new state.
func New() (State, error) {
	dbc, err := db.Connect()
	if err != nil {
		return nil, err
	}

	return &state{
		db: dbc,
	}, nil
}
