// Package transactions provides a set of functions to manage transactions
package transactions

import (
	"database/sql"
)

// TX ...
type TX struct {
	ID                 int64
	UserID             int64
	SourceAddress      string
	DestinationAddress string
	Amount             int64
	Fee                int64
	CreatedAt          sql.NullTime
	Type               TxType
}

//go:generate stringer -type=TxType

// TxType represents type of money transfer.
type TxType uint

const (
	// txTypeUnknown - default value.
	txTypeUnknown TxType = iota
	// TxTypeInternalTransfer - transfer between user's wallet - no fee.
	TxTypeInternalTransfer
	// TxTypeExternalTransfer - transfer to other user wallet - fee should be added.
	TxTypeExternalTransfer
	// txTypeSentinel - Sentinel bound value. Should be always  last.
	txTypeSentinel
)
