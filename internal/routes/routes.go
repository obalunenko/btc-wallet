package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/oleg-balunenko/btc-wallet/internal/ops/transactions"
	"github.com/oleg-balunenko/btc-wallet/internal/ops/users"
	"github.com/oleg-balunenko/btc-wallet/internal/ops/wallets"
	"github.com/oleg-balunenko/btc-wallet/internal/routes/middleware"
	"github.com/oleg-balunenko/btc-wallet/internal/state"
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
	r.POST("/wallets", middleware.AuthRequired(st), wallets.Create(st))

	// returns wallet address and current balance in BTC and USD.
	r.GET("/wallets/:address", middleware.AuthRequired(st), wallets.Lookup(st))

	// makes a transaction from one wallet to another
	// ○ Transaction is free if transferred to own wallet.
	// ○ Transaction costs 1.5% of the transferred amount (profit of the platform) if
	// transferred to a wallet of another user.
	r.POST("/transactions", middleware.AuthRequired(st), transactions.Create(st))

	// returns user’s transactions
	r.GET("/transactions", middleware.AuthRequired(st), transactions.List(st))

	// returns transactions related to the wallet
	r.GET("/wallets/:address/transactions", middleware.AuthRequired(st), wallets.ListTransactions(st))

	return r
}
