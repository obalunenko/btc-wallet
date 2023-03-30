// Package routes provides a set of routes for the application.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/obalunenko/btc-wallet/internal/ops/sessions"
	"github.com/obalunenko/btc-wallet/internal/ops/transactions"
	"github.com/obalunenko/btc-wallet/internal/ops/users"
	"github.com/obalunenko/btc-wallet/internal/ops/wallets"
	"github.com/obalunenko/btc-wallet/internal/state"
)

// Register creates a new router with registered routes.
func Register(st state.State) *gin.Engine {
	r := gin.Default()

	r.RedirectTrailingSlash = true

	// creates user.
	// Returns a token that will authenticate all other requests for this user.
	r.POST("/users", users.Create(st))

	// create BTC wallet for the authenticated user. 1 BTC (or 100000000
	// satoshi ) is automatically granted to the new wallet upon creation. User may register up to
	// 10 wallets.
	// ○ Returns wallet address and current balance in BTC and USD.
	r.POST("/wallets", sessions.AuthRequired(st), wallets.Create(st))

	// returns wallet address and current balance in BTC and USD.
	r.GET("/wallets/:address", sessions.AuthRequired(st), wallets.LookupAddress(st))

	// returns list of wallets for user.
	r.GET("/wallets", sessions.AuthRequired(st), wallets.List(st))

	// makes a transaction from one wallet to another
	// ○ Transaction is free if transferred to own wallet.
	// ○ Transaction costs 1.5% of the transferred amount (profit of the platform) if
	// transferred to a wallet of another user.
	r.POST("/transactions", sessions.AuthRequired(st), transactions.Create(st))

	// returns user’s transactions
	r.GET("/transactions", sessions.AuthRequired(st), transactions.List(st))

	// returns transactions related to the wallet
	r.GET("/wallets/:address/transactions", sessions.AuthRequired(st), wallets.ListTransactions(st))

	return r
}
