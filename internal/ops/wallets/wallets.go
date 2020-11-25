package wallets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/obalunenko/btc-wallet/internal/db/wallets"
	"github.com/obalunenko/btc-wallet/internal/ops/sessions"
)

const (
	maxWalletsForUser = 10
)

// Create create BTC wallet for the authenticated user. 1 BTC (or 100000000
// satoshi ) is automatically granted to the new wallet upon creation. User may register up to
// 10 wallets.
// Returns wallet address and current balance in BTC and USD.
func Create(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		dbc := b.DB()

		sess, err := sessions.GetSessionFromRequest(b, c.Request)
		if err != nil {
			log.Errorf("failed to get session: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid session")

			return
		}

		uID := sess.UserID

		count, err := wallets.CountForUser(ctx, dbc, uID)
		if err != nil {
			log.Errorf("failed to get count of wallets [user_id: %d]", uID)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create wallet")

			return
		}

		if count >= maxWalletsForUser {
			log.Errorf("wallets limit reached [user_id: %d]", uID)

			c.AbortWithStatusJSON(http.StatusBadRequest, "wallets limit reached")

			return
		}

		addr := uuid.New().String()

		id, err := wallets.Create(ctx, dbc, uID, addr)
		if err != nil {
			log.Errorf("failed to create wallet [user_id: %d]: %v", uID, err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create wallet")

			return
		}

		w, err := wallets.Lookup(ctx, dbc, id)
		if err != nil {
			log.Errorf("wallet not created [user_id: %d], [address: %s], [id: %d]: %v", uID, addr, id, err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create wallet")

			return
		}

		// TODO(oleg.balunenko): add real rates and update balance for new wallet.
		c.JSON(http.StatusCreated, responseWallet{
			Address: w.Address,
			Balance: struct {
				USD string `json:"usd"`
				BTC string `json:"btc"`
			}{
				USD: "usd mock",
				BTC: "btc mock",
			},
		})
	}
}

// LookupAddress returns wallet address and current balance in BTC and USD.
func LookupAddress(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		dbc := b.DB()

		addr := c.Param("address")
		if addr == "" {
			log.Error("empty address in request")

			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid address value")

			return
		}

		w, err := wallets.LookupAddress(ctx, dbc, addr)
		if err != nil {
			log.WithFields(map[string]interface{}{
				"address": addr,
				"error":   err,
			}).Error("failed to find wallet")

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to find wallet")

			return
		}

		// TODO(oleg.balunenko): add real rates and update balance for new wallet.
		c.JSON(http.StatusOK, responseWallet{
			Address: w.Address,
			Balance: struct {
				USD string `json:"usd"`
				BTC string `json:"btc"`
			}{
				USD: "usd mock",
				BTC: "btc mock",
			},
		})
	}
}

// List returns list of wallets for user.
func List(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		dbc := b.DB()

		sess, err := sessions.GetSessionFromRequest(b, c.Request)
		if err != nil {
			log.Errorf("failed to get session: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid session")

			return
		}

		uID := sess.UserID

		list, err := wallets.ListForUser(ctx, dbc, uID)
		if err != nil {
			log.WithFields(map[string]interface{}{
				"user_id": uID,
				"error":   err,
			}).Error("failed to list wallets")

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to list wallets")

			return
		}

		addresses := make([]string, 0, len(list))
		for _, w := range list {
			addresses = append(addresses, w.Address)
		}

		// TODO(oleg.balunenko): add real rates and update balance for new wallet.
		c.JSON(http.StatusOK, responseWallets{
			Wallets: addresses,
			Count:   len(addresses),
		})
	}
}

// ListTransactions returns transactions related to the wallet.
func ListTransactions(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
