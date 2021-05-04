package db

import (
	"database/sql"
	"time"
)

func ConfigConnection(
	db *sql.DB,
	connMaxLifetime time.Duration,
	maxOpenConnections int,
	maxIdleConnections int,
) {
	// Set the maximum amount of time a connection may be reused.
	if connMaxLifetime == 0 {
		connMaxLifetime = 5 * time.Minute
	}
	db.SetConnMaxLifetime(connMaxLifetime)

	// Set the maximum number of open connections to the database.
	if maxOpenConnections == 0 {
		maxOpenConnections = 25
	}
	db.SetMaxOpenConns(maxOpenConnections)

	// Sets the maximum number of connections in the idle connection pool.
	if maxIdleConnections == 0 {
		maxIdleConnections = maxOpenConnections
	}
	db.SetMaxIdleConns(maxOpenConnections)
}
