package db

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_db_default_instance_from_app(t *testing.T) {
	app := setup()
	qb := app.Db()

	_, ok := qb.(inter.Database)
	require.True(t, ok)
}

func Test_db_default_instance_from_app_with_correct_connection(t *testing.T) {
	app := setup()
	qb := app.Db()

	_, ok := qb.(inter.Database).Connection().(*db.MySQL)
	require.True(t, ok)
}

func Test_db_instance_by_name(t *testing.T) {
	app := setup()
	qb := app.Db("mysql")

	_, ok := qb.(inter.Database).Connection().(*db.MySQL)
	require.True(t, ok)
}

func Test_db_other_instance_by_name(t *testing.T) {
	app := setup()
	qb := app.Db("postgresql")

	_, ok := qb.(inter.Database).Connection().(*db.PostgreSQL)
	require.True(t, ok)
}

func setup() inter.App {
	app := foundation.NewApp()
	app.Bind("config.Database.Default", "mysql")
	app.Bind("open_connections", map[string]inter.Connection{
		"mysql":      &db.MySQL{},
		"postgresql": &db.PostgreSQL{},
	})

	return app
}
