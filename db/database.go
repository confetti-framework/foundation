package db

import (
	"context"
	"database/sql"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support"
)

type Database struct {
	connection inter.Connection
	app        inter.App
}

func NewDatabase(app inter.App, connection inter.Connection) *Database {
	return &Database{app: app, connection: connection}
}

func (d Database) Connection() inter.Connection {
	return d.connection
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d Database) Exec(sql string, args ...interface{}) sql.Result {
	result, err := d.ExecE(sql, args...)
	if err != nil {
		panic(err)
	}
	return result
}

// ExecE executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d Database) ExecE(sql string, args ...interface{}) (sql.Result, error) {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	execContext, err := connection.Pool().ExecContext(ctx, sql, args...)
	if err != nil {
		err = errors.WithMessage(errors.WithStack(err), "can't execute database query")
	}
	return execContext, err
}

// Query executes a query that returns rows, typically a SELECT. The args are
// for any placeholder parameters in the query.
func (d Database) Query(sql string, args ...interface{}) support.Collection {
	result, err := d.QueryE(sql, args...)
	if err != nil {
		panic(err)
	}
	return result
}

// QueryE executes a query that returns rows or an error, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d Database) QueryE(sql string, args ...interface{}) (support.Collection, error) {
	result := support.NewCollection()

	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	rows, err := connection.Pool().QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		err := rows.Scan(columnPointers...)
		if err != nil {
			return result, errors.WithStack(err)
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		result = result.Push(m)
	}

	return result, nil
}
