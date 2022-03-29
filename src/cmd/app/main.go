package main

import (
	"net/http"
	"os"
	"shopflowers/config"
	authUserHandler "shopflowers/src/internal/auth/delivery/http"
	authRepository "shopflowers/src/internal/auth/repository"
	authService "shopflowers/src/internal/auth/service"
	flowerHandler "shopflowers/src/internal/flower/delivery/http"
	flowerRepository "shopflowers/src/internal/flower/repository"
	flowerService "shopflowers/src/internal/flower/service"
	"shopflowers/src/pkg/db/psql"
	"shopflowers/src/pkg/logg"
)

func main() {
	newLogger := logg.NewCommonLogger()

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

	repositoryAuth := authRepository.NewAuthRepository(psqlDB, newLogger)
	serviceAuth := authService.NewAuthService(repositoryAuth, newLogger)
	handlerAuth := authUserHandler.NewAuthHandler(serviceAuth, newLogger)

	repositoryFlower := flowerRepository.NewFlowerRepository(psqlDB, newLogger)
	serviceFlower := flowerService.NewFlowerService(repositoryFlower, newLogger)
	handlerFlower := flowerHandler.NewFlowerHandler(serviceFlower, newLogger)

	defer psqlDB.Close()

	mux := http.NewServeMux()
	handlerAuth.RegisterRoutes(mux)
	handlerFlower.RegisterRoutes(mux)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/", http.StripPrefix("/web/",
		http.FileServer(http.Dir("src/web/"))))

	if err = http.ListenAndServe(":8000", mux); err != nil {
		newLogger.Error("http.ListenAndServe", err.Error())
	}
}
