package main

import (
	"net/http"
	"os"
	"shopflowers/config"
	userHandler "shopflowers/src/internal/auth/delivery/http"
	"shopflowers/src/internal/auth/repository"
	"shopflowers/src/internal/auth/service"
	"shopflowers/src/pkg/db/psql"
	"shopflowers/src/pkg/logg"
)

func main() {
	newLogger := logg.NewCommonLogger()
	newLogger.InitLogger()

	cfgPath := config.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfigFile(cfgPath)
	if err != nil {
		newLogger.Error("config.LoadConfigFile", err.Error())
	}

	cfg, err := config.ParseFileConfig(cfgFile)
	if err != nil {
		newLogger.Error("config.ParseFileConfig", err.Error())
	}

	psqlDB, err := psql.NewPSQLDB(cfg)
	if err != nil {
		newLogger.Error("psql.NewPSQLDB", err.Error())
	}

	repositoryAuth := repository.NewAuthRepository(psqlDB, newLogger)
	serviceAuth := service.NewAuthService(repositoryAuth, newLogger)
	handlerAuth := userHandler.NewAuthHandler(serviceAuth, newLogger)

	defer psqlDB.Close()

	mux := http.NewServeMux()
	handlerAuth.RegisterRoutes(mux)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/", http.StripPrefix("/web/",
		http.FileServer(http.Dir("src/web/"))))

	if err = http.ListenAndServe(":8000", mux); err != nil {
		newLogger.Error("http.ListenAndServe", err.Error())
	}
}
