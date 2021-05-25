package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

const (
	table = " users "
	cols  = " (created_at) "
)

// User represents user's model.
type User struct {
	ID        int64        `sql:"id"`
	CreatedAt sql.NullTime `sql:"created_at"`
}

// NULL represents empty User default value.
var NULL = User{}

// Create inserts new User into database and returns its ID.
func Create(ctx context.Context, dbc *sql.DB) (int64, error) {
	tx, err := dbc.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	now := time.Now()

	res, err := tx.ExecContext(ctx, "INSERT INTO"+table+cols+"VALUES (?)", now)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert data")
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else if count == 0 {
		return 0, errors.New("no rows updated")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

// Lookup returns User by its ID.
func Lookup(ctx context.Context, dbc *sql.DB, id int64) (User, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM "+table+" WHERE id=?", id)

	return scan(row)
}

func scan(row *sql.Row) (User, error) {
	var u User

	if err := row.Err(); err != nil {
		return NULL, err
	}

	if err := row.Scan(&u.ID, &u.CreatedAt); err != nil {
		return NULL, err
	}
	return u, nil
}
