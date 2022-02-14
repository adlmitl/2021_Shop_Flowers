package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

// Config - Configuration.
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

// GetConfigPath - Path to configuration file.
func GetConfigPath(cfgPath string) string {
	return "config/config"
}

// LoadConfigFile - Load configuration file.
func LoadConfigFile(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseFileConfig - Parsing configuration file.
func ParseFileConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
