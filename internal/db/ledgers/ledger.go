package ledgers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	log "github.com/obalunenko/logger"
)

const (
	table = " ledgers "
	cols  = " (wallet_id, balance, available_balance, created_at, updated_at) "
)

// Ledger represents users.User ledger info, where balance is reflected.
type Ledger struct {
	ID               int64        `sql:"id"`
	WalletID         int64        `sql:"wallet_id"`
	Balance          uint64       `sql:"balance"`
	AvailableBalance uint64       `sql:"available_balance"`
	CreatedAt        sql.NullTime `sql:"created_at"`
	UpdatedAt        sql.NullTime `sql:"updated_at"`
}

// NULL represents empty Ledger default value.
var NULL = Ledger{}

// Create inserts new Ledger for wallets.Wallet into database.
func Create(ctx context.Context, dbc *sql.DB, walletID int64) (int64, error) {
	tx, err := dbc.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err = tx.Rollback(); err != nil {
			log.WithError(ctx, err).Error("failed to rollback transaction")
		}
	}()

	now := time.Now()

	res, err := tx.ExecContext(ctx, "INSERT INTO"+table+cols+"VALUES (?, ?, ?, ?, ?)",
		walletID, 0, 0, now, now)
	if err != nil {
		return 0, fmt.Errorf("failed to insert data: %w", err)
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

// Update sets new values of balance and available balance for Ledger.
func Update(ctx context.Context, dbc *sql.DB, id int64, newBalance, newAvailable string) error {
	tx, err := dbc.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err = tx.Rollback(); err != nil {
			log.WithError(ctx, err).Error("failed to rollback transaction")
		}
	}()

	now := time.Now()
	res, err := tx.ExecContext(ctx,
		"UPDATE"+table+"SET balance = ?, available_balance = ?, updated_at = ? WHERE id = ?",
		newBalance, newAvailable, now, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	} else if count == 0 {
		return errors.New("no rows updated")
	}

	if _, err = res.LastInsertId(); err != nil {
		return err
	}

	return tx.Commit()
}

// Lookup returns Ledger by its ID.
func Lookup(ctx context.Context, dbc *sql.DB, id int64) (Ledger, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM "+table+" WHERE id=?", id)

	return scan(row)
}

// LookupWalletID returns Ledger by walletID.
func LookupWalletID(ctx context.Context, dbc *sql.DB, walletID int64) (Ledger, error) {
	row := dbc.QueryRowContext(ctx, "SELECT * FROM "+table+" WHERE wallet_id=?", walletID)

	return scan(row)
}

func scan(row *sql.Row) (Ledger, error) {
	var l Ledger

	if err := row.Err(); err != nil {
		return NULL, err
	}

	if err := row.Scan(&l.ID, &l.WalletID, &l.Balance, &l.AvailableBalance, &l.CreatedAt, &l.UpdatedAt); err != nil {
		return NULL, err
	}

	return l, nil
}
