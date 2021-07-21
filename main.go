package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/LjungErik/datainjestor/lib/config"
	"github.com/LjungErik/datainjestor/web"

	"github.com/LjungErik/datainjestor/sql"
	log "github.com/sirupsen/logrus"
)

func getenv(name, defaultValue string) string {
	value := os.Getenv(fmt.Sprintf("DATAINJESTOR_%s", name))
	if value == "" {
		return defaultValue
	}
	return value
}

func initializeLogging() {
	formatter := log.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(&formatter)
	if getenv("VERBOSE_LOGGING", "0") == "1" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetOutput(os.Stdout)
}

func initializeSqlClient(cfg *config.Config) *sql.Client {

	if len(cfg.SQL.Connection) == 0 {
		log.Error("Missing connection string cannot create sql client")
		panic(errors.New("Missing connection string for sql client"))
	}

	log.Info("SQL database: ENABLED")
	sqlconf := sql.Config{
		ConnString: cfg.SQL.Connection,
	}

	sqlclient := sql.NewClient(&sqlconf)

	return sqlclient
}

func initializeWebConfig(cfg *config.Config) *web.Config {
	sqlclient := initializeSqlClient(cfg)
	wcfg := web.NewConfig(sqlclient)
	return wcfg
}

func main() {
	initializeLogging()
	cfgFile := getenv("CONFIG_FILE", "config.yml")
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Errorf("Failed to parse config at: %v", cfgFile)
		panic(err)
	}

	wcfg := initializeWebConfig(cfg)
	router := wcfg.InitRouter()

	log.Infof("Starting up server at :9090")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":9090"), router))
}
