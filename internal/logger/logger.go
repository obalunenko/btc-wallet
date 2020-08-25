package logger

import (
	"flag"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var (
	logLevel = flag.String("log_level", "INFO", "set log level of application")
	logFile  = flag.String("log_file", "", "log file name, if empty - only stdout output.")
)

// Init initiates logger and add format options.
func Init() {
	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)

	filename := *logFile
	if filename == "" {
		return
	}

	formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	formatter.FullTimestamp = true

	formatter.QuoteEmptyFields = true
	formatter.ForceColors = true

	var out []io.Writer

	out = append(out, os.Stdout)

	f, err := os.OpenFile(filepath.Clean(filename), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.WithField("filename", *logFile).
			WithField("error", err).
			Warnf("failed to open file - will write logs to stdout")
	} else {
		out = append(out, f)
	}

	log.SetOutput(io.MultiWriter(out...))
}
