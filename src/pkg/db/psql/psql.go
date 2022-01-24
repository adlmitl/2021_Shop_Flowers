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

var l = logg.NewLogg()

// NewPSQLDB - Конфигурация нового пул соединения с БД.
func NewPSQLDB(c *config.Config) (*pgxpool.Pool, error) {
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
		l.LogError("Error create pool", err.Error())
	}

	poolConfig.MaxConns = maxCountConn
	poolConfig.MinConns = minCountConn

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		l.LogError("Error connect to db", err.Error())
	}
	l.LogInfo("Connect OK!")

	if err = pool.Ping(ctx); err != nil {
		l.LogError("Error ping failed", err.Error())
	}
	l.LogInfo("Success!")

	return pool, nil
}
