package web

import (
	"github.com/LjungErik/datainjestor/lib/encoding"
	"github.com/LjungErik/datainjestor/mongodb"
	"github.com/LjungErik/datainjestor/sql"
	"github.com/julienschmidt/httprouter"
)

// Config config for handling clients
type Config struct {
	sql     *sql.Client
	mongo   *mongodb.Client
	decoder encoding.IDecoder
}

// NewConfig Generates a new webconfig with the provided metric clients
func NewConfig(sqlclient *sql.Client, mongoclient *mongodb.Client) *Config {
	d, err := encoding.NewDecoder(encoding.EncodingJSON)
	if err != nil {
		// handle error when decoder can not be created
	}

	return &Config{
		sql:     sqlclient,
		mongo:   mongoclient,
		decoder: d,
	}
}

// InitRouter Initializes router for the web config
func (wc *Config) InitRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/data/housing/forsale", wc.writeHousingForSale)
	router.POST("/data/housing/sold", wc.writeHousingSold)
	router.GET("/health/live", wc.live)
	router.GET("/health/ready", wc.ready)
	return router
}
