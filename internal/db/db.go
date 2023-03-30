// Package db provides a database connection pool.
package db

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func build(connectStr string) (string, error) {
	const prefix = "mysql://"

	if !strings.HasPrefix(connectStr, prefix) {
		return "", errors.New("db: URI is missing mysql:// prefix")
	}

	connectStr = connectStr[len(prefix):]

	dsn, err := mysql.ParseDSN(connectStr)
	if err != nil {
		return "", err
	}

	if dsn.Params == nil {
		dsn.Params = make(map[string]string)
	}

	dsn.Params["parseTime"] = "true"               // time.Time for datetime
	dsn.Params["collation"] = "utf8mb4_general_ci" // non-BMP unicode chars

	return dsn.FormatDSN(), nil
}

func connect(driver, connectStr string) (*sql.DB, error) {
	connectStr, err := build(connectStr)
	if err != nil {
		return nil, err
	}

	dbc, err := sql.Open(driver, connectStr)
	if err != nil {
		return nil, err
	}

	return dbc, nil
}

// ConnectParams holds the parameters for connecting to a database.
type ConnectParams struct {
	ConnURI         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// ConnectWithURI establishes a database connection to the given URI.  If
// readonly is provided, it will be used as a readonly replica.
func ConnectWithURI(p ConnectParams) (*sql.DB, error) {
	dbc, err := connect("mysql", p.ConnURI)
	if err != nil {
		return nil, err
	}

	dbc.SetMaxOpenConns(p.MaxOpenConns)
	dbc.SetMaxIdleConns(p.MaxIdleConns)
	dbc.SetConnMaxLifetime(p.ConnMaxLifetime)

	if err := dbc.Ping(); err != nil {
		return nil, err
	}

	return dbc, nil
}
