package main

import (
	"net/http"
	"os"
	"shopflowers/src/pkg/db/psql"
	"shopflowers/src/pkg/logg"
	"shopflowers/src/pkg/util"
)

var l = logg.NewLogg()

func main() {
	cfgPath := util.GetConfigPath(os.Getenv("config"))

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
	defer psqlDB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	if err = http.ListenAndServe(":8000", mux); err != nil {
		l.LogError("Error start server:", err.Error())
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Flowers"))
	if err != nil {
		l.LogError("Error start server:", err.Error())
	}
}
