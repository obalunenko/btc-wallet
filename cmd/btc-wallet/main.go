// btc-wallet is a simple service that allows to create and manage wallets.
package main

import (
	"context"
	"net"

	log "github.com/obalunenko/logger"

	"github.com/obalunenko/btc-wallet/cmd/btc-wallet/internal/config"
	"github.com/obalunenko/btc-wallet/cmd/btc-wallet/logger"
	"github.com/obalunenko/btc-wallet/internal/db"
	"github.com/obalunenko/btc-wallet/internal/routes"
	"github.com/obalunenko/btc-wallet/internal/state"
)

func main() {
	ctx := context.Background()

	config.Load()

	logger.Init(ctx)

	printVersion(ctx)

	st, err := state.New(db.ConnectParams{
		ConnURI:         config.DBConnURI(),
		MaxOpenConns:    config.DBMaxOpenConns(),
		MaxIdleConns:    config.DBMaxIdleConns(),
		ConnMaxLifetime: config.DBConnMaxLifetime(),
	})
	if err != nil {
		log.WithError(ctx, err).Fatal("Failed to init state")
	}

	r := routes.Register(st)

	if err := r.Run(net.JoinHostPort("", "8080")); err != nil {
		log.WithError(ctx, err).Fatal("Failed to init state")
	}
}
