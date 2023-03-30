// Package transactions implements the transactions operations.
package transactions

import (
	"github.com/gin-gonic/gin"
)

// Create makes a transaction from one wallet to another
// ○ Transaction is free if transferred to own wallet.
// ○ Transaction costs 1.5% of the transferred amount (profit of the platform) if
// transferred to a wallet of another user.
func Create(_ Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// List returns user’s transactions.
func List(_ Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
