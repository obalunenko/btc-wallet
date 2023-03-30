// Package wallets provides access to wallets table.
package wallets

import (
	"context"
	"database/sql"
	"errors"
)

const (
	table      = " wallets "
	colsInsert = " (user_id, address) VALUES (?, ?) "
)

// Wallet represents user's wallet.
type Wallet struct {
	ID      int64  `sql:"id"`
	Address string `sql:"address"`
	UserID  int64  `sql:"user_id"`
}

// Create creates Wallet for users.User with passed defined address.
func Create(ctx context.Context, dbc *sql.DB, userID int64, address string) (int64, error) {
	tx, err := dbc.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	res, err := tx.ExecContext(ctx, "INSERT INTO "+table+colsInsert, userID, address)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else if count == 0 {
		return 0, errors.New("no rows affected")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

// Lookup returns Wallet by it's ID.
func Lookup(ctx context.Context, dbc *sql.DB, id int64) (Wallet, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM"+table+"WHERE id=?", id)

	return scan(row)
}

// LookupAddress returns Wallet by it's address.
func LookupAddress(ctx context.Context, dbc *sql.DB, address string) (Wallet, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM"+table+"WHERE address=?", address)

	return scan(row)
}

// ListForUser returns list of Wallet for users.User but his ID.
func ListForUser(ctx context.Context, dbc *sql.DB, userID int64) ([]Wallet, error) {
	rows, err := dbc.QueryContext(ctx, "SELECT * FROM"+table+"WHERE user_id=?", userID)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	return list(rows)
}

// CountForUser returns number of Wallet for user by his ID.
func CountForUser(ctx context.Context, dbc *sql.DB, userID int64) (int, error) {
	var res int

	row := dbc.QueryRowContext(ctx, "SELECT COUNT(*) FROM"+table+"WHERE user_id=?", userID)

	if err := row.Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return -1, err
	}

	return res, nil
}

func list(rows *sql.Rows) ([]Wallet, error) {
	var res []Wallet

	for rows.Next() {
		var w Wallet

		if err := rows.Scan(&w.ID, &w.UserID, &w.Address); err != nil {
			return nil, err
		}

		res = append(res, w)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func scan(row *sql.Row) (Wallet, error) {
	var w Wallet

	if err := row.Scan(&w.ID, &w.UserID, &w.Address); err != nil {
		return Wallet{}, err
	}

	return w, nil
}
