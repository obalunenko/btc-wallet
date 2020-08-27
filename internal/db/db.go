package db

import (
	"database/sql"
	"flag"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var sockFile = getSocketFile()

var (
	dbURI = flag.String("db", "mysql://root@unix("+sockFile+")/test?",
		"Database URI")
	dbMaxOpenConns = flag.Int("db_max_open_conns", 100,
		"Maximum number of open database connections")
	dbMaxIdleConns = flag.Int("db_max_idle_conns", 50,
		"Maximum number of idle database connections")
	dbConnMaxLifetime = flag.Duration("db_conn_max_lifetime", time.Minute,
		"Maximum time a single database connection can be left open")
)

func getSocketFile() string {
	sock := "/tmp/mysql.sock"

	if _, err := os.Stat(sock); os.IsNotExist(err) {
		// try common linux/Ubuntu socket file location
		return "/var/run/mysqld/mysqld.sock"
	}

	return sock
}

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

	dbc.SetMaxOpenConns(*dbMaxOpenConns)
	dbc.SetMaxIdleConns(*dbMaxIdleConns)
	dbc.SetConnMaxLifetime(*dbConnMaxLifetime)

	return dbc, nil
}

// ConnectWithURI establishes a database connection to the given URI.  If
// readonly is provided, it will be used as a readonly replica.
func ConnectWithURI(uri string) (*sql.DB, error) {
	dbc, err := connect("mysql", uri)
	if err != nil {
		return nil, err
	}

	if err = dbc.Ping(); err != nil {
		return nil, err
	}

	return dbc, nil
}

// Connect establishes db connection.
func Connect() (*sql.DB, error) {
	return ConnectWithURI(*dbURI)
}
