package wallets

import (
	"github.com/gin-gonic/gin"
)

// Create create BTC wallet for the authenticated user. 1 BTC (or 100000000
// satoshi ) is automatically granted to the new wallet upon creation. User may register up to
// 10 wallets.
// Returns wallet address and current balance in BTC and USD.
func Create(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// Lookup returns wallet address and current balance in BTC and USD.
func Lookup(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// ListTransactions returns transactions related to the wallet.
func ListTransactions(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
