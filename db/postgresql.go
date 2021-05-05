package db

import (
	"crypto/tls"
	"database/sql"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"time"
)

type PostgreSQL struct {
	// Host (e.g. localhost) or absolute path to unix domain socket directory (e.g. /private/tmp)
	Host     string
	Port     int
	Database string
	Username string
	Password string

	// nil disables TLS
	TLSConfig      *tls.Config
	ConnectTimeout time.Duration

	// e.g. net.Dialer.DialContext
	DialFunc pgconn.DialFunc

	// e.g. net.Resolver.LookupHost
	LookupFunc    pgconn.LookupFunc
	BuildFrontend pgconn.BuildFrontendFunc

	// Run-time parameters to set on connection as session default values (e.g. search_path or application_name)
	RuntimeParams map[string]string

	Fallbacks []*pgconn.FallbackConfig

	// ValidateConnect is called during a connection attempt after a successful authentication with the PostgreSQL server.
	// It can be used to validate that the server is acceptable. If this returns an error the connection is closed and the next
	// fallback config is tried. This allows implementing high availability behavior such as libpq does with target_session_attrs.
	ValidateConnect pgconn.ValidateConnectFunc

	// AfterConnect is called after ValidateConnect. It can be used to set up the connection (e.g. Set session variables
	// or prepare statements). If this returns an error the connection attempt fails.
	AfterConnect pgconn.AfterConnectFunc

	// OnNotice is a callback function called when a notice response is received.
	OnNotice pgconn.NoticeHandler

	// OnNotification is a callback function called when a notification from the LISTEN/NOTIFY system is received.
	OnNotification pgconn.NotificationHandler

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

func (m *PostgreSQL) Open() error {

	config, err := pgx.ParseConfig("")
	if err != nil {
		return err
	}

	m.setConfig(config)

	pool := stdlib.OpenDB(*config)

	ConfigConnection(pool, m.ConnMaxLifetime, m.MaxOpenConnections, m.MaxIdleConnections)
	m.pool = pool

	return nil
}

func (m PostgreSQL) Pool() *sql.DB {
	return m.pool
}

func (m PostgreSQL) Timeout() time.Duration {
	return m.QueryTimeout
}

func (m *PostgreSQL) setConfig(config *pgx.ConnConfig) {
	if m.Host != "" {
		config.Config.Host = m.Host
	}
	if m.Port != 0 {
		config.Config.Port = uint16(m.Port)
	}
	if m.Database != "" {
		config.Config.Database = m.Database
	}
	if m.Username != "" {
		config.Config.User = m.Username
	}
	if m.Password != "" {
		config.Config.Password = m.Password
	}
	if m.TLSConfig != nil {
		config.Config.TLSConfig = m.TLSConfig
	}
	if m.ConnectTimeout != 0 {
		config.Config.ConnectTimeout = m.ConnectTimeout
	}
	if m.DialFunc != nil {
		config.Config.DialFunc = m.DialFunc
	}
	if m.LookupFunc != nil {
		config.Config.LookupFunc = m.LookupFunc
	}
	if m.BuildFrontend != nil {
		config.Config.BuildFrontend = m.BuildFrontend
	}
	if len(m.RuntimeParams) != 0 {
		config.Config.RuntimeParams = m.RuntimeParams
	}
	if len(m.Fallbacks) != 0 {
		config.Config.Fallbacks = m.Fallbacks
	}
	if m.ValidateConnect != nil {
		config.Config.ValidateConnect = m.ValidateConnect
	}
	if m.AfterConnect != nil {
		config.Config.AfterConnect = m.AfterConnect
	}
	if m.OnNotice != nil {
		config.Config.OnNotice = m.OnNotice
	}
	if m.OnNotification != nil {
		config.Config.OnNotification = m.OnNotification
	}
}
