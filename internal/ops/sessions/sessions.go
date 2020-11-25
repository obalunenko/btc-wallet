package sessions

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/obalunenko/btc-wallet/internal/db/sessions"
	"github.com/obalunenko/btc-wallet/internal/db/users"
)

// New creates session for user.
func New(ctx context.Context, b Backends, userID int64) (sessions.Session, error) {
	expiration := time.Hour * 24

	dbc := b.DB()

	u, err := users.Lookup(ctx, dbc, userID)
	if err != nil {
		return sessions.Session{}, errors.Wrapf(err, "failed to find user [id=%d]", userID)
	}

	token, err := encode(u)
	if err != nil {
		return sessions.Session{}, errors.Wrap(err, "failed to generate token")
	}

	id, err := sessions.Create(ctx, dbc, u.ID, token, expiration)
	if err != nil {
		return sessions.Session{}, errors.Wrap(err, "failed to create session")
	}

	sess, err := sessions.Lookup(ctx, dbc, id)
	if err != nil {
		return sessions.Session{}, errors.Wrap(err, "session not created")
	}

	return sess, nil
}

// AuthRequired checks if user is authorized.
func AuthRequired(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		sess, err := GetSessionFromRequest(b, c.Request)
		if err != nil {
			log.Errorf("failed to find token: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")

			return
		}

		if !sess.Valid() {
			log.Debugf("expired token: [%s: %+v]", token, sess)

			c.AbortWithStatusJSON(http.StatusUnauthorized, "token expired")

			return
		}

		log.Debugf("authorized user: [id: %d]", sess.UserID)

		c.Next()
	}
}

// GetSessionFromRequest returns Session from http.Request header 'Authorization` token.
func GetSessionFromRequest(b Backends, r *http.Request) (sessions.Session, error) {
	dbc := b.DB()
	ctx := r.Context()

	token := r.Header.Get("Authorization")
	if token == "" {
		return sessions.Session{}, errors.New("token missed")
	}

	sess, err := sessions.LookupByToken(ctx, dbc, token)
	if err != nil {
		return sessions.Session{}, errors.Wrapf(err, "failed to find token")
	}

	return sess, nil
}
