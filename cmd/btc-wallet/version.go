package main

import (
	"context"

	log "github.com/obalunenko/logger"
)

const unset = "unset"

var ( // build info
	version = unset
	date    = unset
	commit  = unset
)

func printVersion(ctx context.Context) {
	log.WithFields(ctx, log.Fields{
		"version": version,
		"date":    date,
		"commit":  commit,
	}).Info("Build info")
}
