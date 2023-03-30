// Package sessions provides database access to sessions.
package sessions

import (
	"context"
	"database/sql"
	"time"

	log "github.com/obalunenko/logger"
)

const (
	table      = ` sessions `
	colsInsert = ` (user_id, token, created_at, expires_at) `
)

// Session represents users.User session.
type Session struct {
	ID        int64        `sql:"id"`
	UserID    int64        `sql:"user_id"`
	Token     string       `sql:"token"`
	CreatedAt sql.NullTime `sql:"created_at"`
	ExpiresAt sql.NullTime `sql:"expires_at"`
}

// Valid checks if session valid by expiration date.
func (s Session) Valid() bool {
	if s.ExpiresAt.Time.IsZero() {
		return true
	}

	return s.ExpiresAt.Time.After(time.Now())
}

// Create inserts new Session into database.
func Create(ctx context.Context, dbc *sql.DB, userID int64, token string, expiration time.Duration) (int64, error) {
	tx, err := dbc.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err = tx.Rollback(); err != nil {
			log.WithError(ctx, err).Error("Failed to rollback transaction")
		}
	}()

	now := time.Now()
	exp := now.Add(expiration)

	res, err := tx.ExecContext(ctx, `
	INSERT INTO `+table+colsInsert+
		`VALUES(?, ?, ?, ?)`, userID, token, now, exp)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else if count == 0 {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

// Lookup returns Session by it's ID.
func Lookup(ctx context.Context, dbc *sql.DB, id int64) (Session, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM"+table+"WHERE id=?", id)

	return scan(row)
}

// LookupByToken returns Session by it's token.
func LookupByToken(ctx context.Context, dbc *sql.DB, token string) (Session, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM"+table+"WHERE token=?", token)

	return scan(row)
}

func scan(row *sql.Row) (Session, error) {
	var s Session

	if err := row.Scan(&s.ID, &s.UserID, &s.Token, &s.CreatedAt, &s.ExpiresAt); err != nil {
		return Session{}, err
	}

	return s, nil
}
