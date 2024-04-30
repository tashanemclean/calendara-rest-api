package db

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/tashanemclean/calendara-rest-api-api/util/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

var PGPool *pgxpool.Pool
var PGBunDB *bun.DB

func initPostgres() {
	migrate()
	setupDBPool()

	// Init Bun DB
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.Config.DatabaseConnectionURL)))
	PGBunDB = bun.NewDB(sqldb, pgdialect.New())
}

func setupDBPool() {
	if PGPool != nil {
		return
	}

	var err error
	PGPool, err = pgxpool.New(context.Background(), config.Config.DatabaseConnectionURL)

	if err != nil {
		log.Fatal(err)
	}
}

func migrate() {
	u, err := url.Parse(config.Config.DatabaseConnectionURL)

	if err != nil {
		log.Fatal(err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	configPath := filepath.Join(basepath, "migrations")

	migrator := dbmate.New(u)
	migrator.MigrationsDir = []string{configPath}
	migrator.SchemaFile = filepath.Join(basepath, "schema.sql")

	err = migrator.CreateAndMigrate()

	if err != nil {
		log.Fatal(err)
	}
}