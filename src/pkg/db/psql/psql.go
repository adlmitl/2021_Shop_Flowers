package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"shopflowers/config"
	"shopflowers/src/pkg/logg"
)

const (
	maxCountConn = 10
	minCountConn = 5
)

// NewPSQLDB - Configuring a new database connection pool.
func NewPSQLDB(c *config.Config) (*pgxpool.Pool, error) {
	newLogger := logg.NewCommonLogger()

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.Postgres.PgDriver,
		c.Postgres.PSQLUser,
		c.Postgres.PSQLPassword,
		c.Postgres.PSQLHost,
		c.Postgres.PSQLPort,
		c.Postgres.PSQLDBName,
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		newLogger.Error("pgxpool.ParseConfig", err.Error())
	}

	poolConfig.MaxConns = maxCountConn
	poolConfig.MinConns = minCountConn

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		newLogger.Error("pgxpool.ConnectConfig", err.Error())
	}
	newLogger.Info("Connect OK!")

	if err = pool.Ping(ctx); err != nil {
		newLogger.Error("pool.Ping", err.Error())
	}
	newLogger.Info("Success!")

	return pool, nil
}
