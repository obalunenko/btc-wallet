package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/btc-wallet/internal/db/users"
)

// Create creates user.
// Returns a token that will authenticate all other requests for this user.
func Create(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbc := b.DB()

		ctx := c.Request.Context()

		id, err := users.Create(ctx, dbc)
		if err != nil {
			log.Errorf("failed to create user: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create user")

			return
		}

		u, err := users.Lookup(ctx, dbc, id)
		if err != nil {
			log.Errorf("failed to get user: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to get user")

			return
		}

		token, err := Encode(u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create token")

			return
		}

		c.JSON(http.StatusOK, response{
			ID:    u.ID,
			Token: token,
		})

		log.Infof("user created: %+v [%s]", u, token)
	}
}
