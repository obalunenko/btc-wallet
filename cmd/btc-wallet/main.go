package main

import (
	"flag"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/btc-wallet/internal/logger"
	"github.com/oleg-balunenko/btc-wallet/internal/routes"
	"github.com/oleg-balunenko/btc-wallet/internal/state"
)

func main() {
	flag.Parse()

	logger.Init()

	printVersion()

	st, err := state.New()
	if err != nil {
		log.Fatalf("failed to init state: %v", err)
	}

	r := routes.Register(st)

	if err := r.Run(net.JoinHostPort("", "8080")); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
