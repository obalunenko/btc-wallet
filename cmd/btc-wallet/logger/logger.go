// Package logger provides logger initialize.
package logger

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	log "github.com/obalunenko/logger"
	"github.com/obalunenko/version"

	"github.com/obalunenko/btc-wallet/cmd/btc-wallet/internal/config"
)

type writercloser struct {
	writers io.Writer
	closers []func() error
}

func (w writercloser) Write(p []byte) (n int, err error) {
	return w.writers.Write(p)
}

func (w writercloser) Close() error {
	var errs error

	for i := range w.closers {
		if err := w.closers[i](); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

// Init initiates logger and add format options.
func Init(ctx context.Context) {
	const (
		wnum             = 2
		perm os.FileMode = 0600
	)

	var (
		out    = make([]io.Writer, 0, wnum)
		closer = make([]func() error, 0, wnum)
	)

	out = append(out, os.Stdout)
	closer = append(closer, os.Stdout.Close)

	filename := config.LogFile()

	f, err := os.OpenFile(filepath.Clean(filename), os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		log.WithField(ctx, "filename", filename).
			WithField("error", err).
			Warn("Failed to open file - will write logs to stdout")
	} else {
		out = append(out, f)

		closer = append(closer, f.Close)
	}

	w := writercloser{
		writers: io.MultiWriter(out...),
		closers: closer,
	}

	log.Init(ctx, log.Params{
		Writer: w,
		Level:  config.LogLevel(),
		Format: config.LogFormat(),
		SentryParams: log.SentryParams{
			Enabled:      config.LogSentryEnabled(),
			DSN:          config.LogSentryDSN(),
			TraceEnabled: config.LogSentryTraceEnabled(),
			TraceLevel:   config.LogSentryTraceLevel(),
			Tags: map[string]string{
				"app_name":     version.GetAppName(),
				"go_version":   version.GetGoVersion(),
				"version":      version.GetVersion(),
				"build_date":   version.GetBuildDate(),
				"short_commit": version.GetShortCommit(),
			},
		},
	})
}
