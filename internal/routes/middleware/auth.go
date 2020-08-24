package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	dbUsers "github.com/oleg-balunenko/btc-wallet/internal/db/users"
	opsUsers "github.com/oleg-balunenko/btc-wallet/internal/ops/users"
	"github.com/oleg-balunenko/btc-wallet/internal/state"
)

// AuthRequired checks if user is authorized.
func AuthRequired(st state.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		var u dbUsers.User

		if err := opsUsers.Decode(token, &u); err != nil {
			log.Errorf("failed to decode token: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")

			return
		}

		if _, err := dbUsers.Lookup(c.Request.Context(), st.DB(), u.ID); err != nil {
			log.Errorf("failed to get user [id: %d]: %v", u.ID, err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "user not exist")

			return
		}

		log.Infof("decoded token: %+v", u)

		c.Next()
	}
}
