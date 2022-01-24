package config

// Config - Конфигурация приложения.
type Config struct {
	Postgres PSQLConfig
}

// PSQLConfig - Конфигурация PSQL.
type PSQLConfig struct {
	PSQLHost     string
	PSQLPort     string
	PSQLUser     string
	PSQLPassword string
	PSQLDBName   string
	PgDriver     string
}
