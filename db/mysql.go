package db

import (
	"database/sql"
	"github.com/confetti-framework/errors"
	"github.com/go-sql-driver/mysql"
	"time"
)

// MySQL 4.1+ see https://github.com/go-sql-driver/mysql/
type MySQL struct {
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string
	Parameters map[string]string

	// When the open connection limit is reached, and all connections are in-use,
	// any new database tasks that your application needs to execute will be forced
	// to wait until a connection becomes free and marked as idle. To mitigate this
	// you can set a fixed, fast, timeout when making database calls.
	// Default 10 seconds
	QueryTimeout time.Duration

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0 means
	// that there is no maximum lifetime and the connection is reused forever.
	// Default 5 minutes
	ConnMaxLifetime time.Duration

	// Set the maximum number of concurrently open connections (in-use + idle)
	// Setting this to less than 0 will mean there is no maximum limit.
	// Default 25
	MaxOpenConnections int

	// Set the maximum number of concurrently idle connections. Setting this
	// to less than 0 will mean that no idle connections are retained.
	// Default MaxOpenConnections
	MaxIdleConnections int

	// pool is a database handle representing a pool of zero or more
	// underlying connections. It's safe for concurrent use by multiple
	// goroutines.
	pool *sql.DB
}

func (m *MySQL) Open() error {
	pool, err := sql.Open("mysql", m.NetworkAddress())
	if err != nil {
		return errors.Wrap(err, "can't open MySQL connection")
	}

	ConfigConnection(pool, m.ConnMaxLifetime, m.MaxOpenConnections, m.MaxIdleConnections)

	m.pool = pool

	return err
}

func (m MySQL) Pool() *sql.DB {
	return m.pool
}

func (m MySQL) Timeout() time.Duration {
	return m.QueryTimeout
}

func (m MySQL) NetworkAddress() string {
	config := mysql.NewConfig()
	config.User = m.Username
	config.Passwd = m.Password
	if m.Host != "" {
		config.Net = "tcp"
	}
	config.Addr = m.Host
	config.DBName = m.Database
	config.Params = m.Parameters

	return config.FormatDSN()
}
