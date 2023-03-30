// Package config provide application configuration.
package config

import (
	"flag"
	"time"
)

const (
	// DefaultLogLevel is default log level.
	DefaultLogLevel = "INFO"
	// DefaultLogFormat is default log format.
	DefaultLogFormat = "text"
	// DefaultLogSentryTraceLevel is default log sentry trace level.
	DefaultLogSentryTraceLevel = "PANIC"
	// DefaultDBMaxOpenConns is default db max open connections.
	DefaultDBMaxOpenConns = 10
	// DefaultDBMaxIdleConns is default db max idle connections.
	DefaultDBMaxIdleConns = 10
	// DefaultDBConnMaxLifetime is default db connection max lifetime.
	DefaultDBConnMaxLifetime = time.Minute
)

var (
	// Log related configs.
	logFile        = flag.String("log_file", "", "log file name, if empty - only stdout output.")
	logLevel       = flag.String("log_level", DefaultLogLevel, "set log level of application")
	logFormat      = flag.String("log_format", DefaultLogFormat, "Format of logs (supported values: text, json")
	logSentryDSN   = flag.String("log_sentry_dsn", "", "Sentry DSN")
	logSentryTrace = flag.Bool("log_sentry_trace", false,
		"Enables sending stacktrace to sentry")
	logSentryTraceLevel = flag.String("log_sentry_trace_level", DefaultLogSentryTraceLevel,
		"The level at which to start capturing stacktraces")
	dbURI = flag.String("db", "mysql://root@tcp(mysql:3306)/btc_wallet?",
		"Database URI")
	dbMaxOpenConns = flag.Int("db_max_open_conns", DefaultDBMaxOpenConns,
		"Maximum number of open database connections")
	dbMaxIdleConns = flag.Int("db_max_idle_conns", DefaultDBMaxIdleConns,
		"Maximum number of idle database connections")
	dbConnMaxLifetime = flag.Duration("db_conn_max_lifetime", DefaultDBConnMaxLifetime,
		"Maximum time a single database connection can be left open")
)

// ensureFlags panics if env is checked before flags are parsed.
// Ok to panic since this should be caught in dev or staging.
func ensureFlags() {
	if !flag.Parsed() {
		panic("flags not parsed yet")
	}
}

// LogLevel config.
func LogLevel() string {
	ensureFlags()
	return *logLevel
}

// LogFile config.
func LogFile() string {
	ensureFlags()
	return *logFile
}

// LogSentryDSN config.
func LogSentryDSN() string {
	ensureFlags()
	return *logSentryDSN
}

// LogSentryEnabled config.
func LogSentryEnabled() bool {
	ensureFlags()
	return LogSentryDSN() != ""
}

// LogSentryTraceEnabled config.
func LogSentryTraceEnabled() bool {
	ensureFlags()
	return *logSentryTrace
}

// LogSentryTraceLevel config.
func LogSentryTraceLevel() string {
	ensureFlags()
	return *logSentryTraceLevel
}

// LogFormat config.
func LogFormat() string {
	ensureFlags()
	return *logFormat
}

// DBConnURI config.
func DBConnURI() string {
	ensureFlags()
	return *dbURI
}

// DBMaxOpenConns config.
func DBMaxOpenConns() int {
	ensureFlags()
	return *dbMaxOpenConns
}

// DBMaxIdleConns config.
func DBMaxIdleConns() int {
	ensureFlags()
	return *dbMaxIdleConns
}

// DBConnMaxLifetime config.
func DBConnMaxLifetime() time.Duration {
	ensureFlags()
	return *dbConnMaxLifetime
}

// DBConnMaxLifetimeSec config.
func DBConnMaxLifetimeSec() int {
	ensureFlags()
	return int(DBConnMaxLifetime().Seconds())
}

// Load loads application configuration.
func Load() {
	flag.Parse()
}
