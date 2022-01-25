package main

import (
	"net/http"
	"os"
	http2 "shopflowers/src/internal/auth/delivery/http"
	"shopflowers/src/internal/auth/repository"
	"shopflowers/src/internal/auth/service"
	"shopflowers/src/pkg/db/psql"
	"shopflowers/src/pkg/logg"
	"shopflowers/src/pkg/util"
)

func main() {
	cfgPath := util.GetConfigPath(os.Getenv("config"))
	l := logg.NewLogg()
	cfgFile, err := util.LoadConfigFile(cfgPath)
	if err != nil {
		l.LogError("LoadConfigFile", err)
	}

	cfg, err := util.ParseFileConfig(cfgFile)
	if err != nil {
		l.LogError("ParseFileConfig", err)
	}

	psqlDB, err := psql.NewPSQLDB(cfg)
	if err != nil {
		l.LogError("Postgresql init: %s", err.Error())
	}
	repositoryAuth := repository.NewAuthRepository(psqlDB, l)
	authService := service.NewAuthService(repositoryAuth, l)
	handlerAuth := http2.NewAuthHandler(authService, l)

	defer psqlDB.Close()

	mux := http.NewServeMux()
	handlerAuth.Register(mux)

	if err = http.ListenAndServe(":8000", mux); err != nil {
		l.LogError("Error start server:", err.Error())
	}
}
