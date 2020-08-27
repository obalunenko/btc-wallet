package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/btc-wallet/internal/db/users"
	"github.com/oleg-balunenko/btc-wallet/internal/ops/sessions"
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

		sess, err := sessions.New(ctx, b, u.ID)
		if err != nil {
			log.Errorf("failed to generate session: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create session")

			return
		}

		token := sess.Token

		c.JSON(http.StatusCreated, response{
			ID:    u.ID,
			Token: token,
		})

		log.Infof("user created: %+v [%s]", u, token)
	}
}
