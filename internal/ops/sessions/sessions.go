// Package sessions provides a simple session management system.
package sessions

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/obalunenko/logger"

	"github.com/obalunenko/btc-wallet/internal/db/sessions"
	"github.com/obalunenko/btc-wallet/internal/db/users"
)

// New creates session for user.
func New(ctx context.Context, b Backends, userID int64) (sessions.Session, error) {
	expiration := time.Hour * 24

	dbc := b.DB()

	u, err := users.Lookup(ctx, dbc, userID)
	if err != nil {
		return sessions.Session{}, fmt.Errorf("failed to find user [id=%d]: %w", userID, err)
	}

	token, err := encode(u)
	if err != nil {
		return sessions.Session{}, fmt.Errorf("failed to generate token: %w", err)
	}

	id, err := sessions.Create(ctx, dbc, u.ID, token, expiration)
	if err != nil {
		return sessions.Session{}, fmt.Errorf("failed to create session: %w", err)
	}

	sess, err := sessions.Lookup(ctx, dbc, id)
	if err != nil {
		return sessions.Session{}, fmt.Errorf("session not created: %w", err)
	}

	return sess, nil
}

// AuthRequired checks if user is authorized.
func AuthRequired(b Backends) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		token := c.GetHeader("Authorization")

		sess, err := GetSessionFromRequest(b, c.Request)
		if err != nil {
			log.WithError(ctx, err).Error("Failed to find token")

			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")

			return
		}

		if !sess.Valid() {
			log.WithFields(ctx, log.Fields{
				"token": token,
				"sess":  sess,
			}).Debug("Token expired")

			c.AbortWithStatusJSON(http.StatusUnauthorized, "token expired")

			return
		}

		log.WithField(ctx, "id", sess.UserID).
			Debug("Authorized user")

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
		return sessions.Session{}, fmt.Errorf("failed to find token: %w", err)
	}

	return sess, nil
}
