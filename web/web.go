package web

import (
	"github.com/LjungErik/datainjestor/lib/encoding"
	"github.com/LjungErik/datainjestor/sql"
	"github.com/julienschmidt/httprouter"
)

// Config config for handling clients
type Config struct {
	sql     *sql.Client
	decoder encoding.IDecoder
}

// NewConfig Generates a new webconfig with the provided metric clients
func NewConfig(c *sql.Client) *Config {
	d, err := encoding.NewDecoder(encoding.EncodingCSV)
	if err != nil {
		// handle error when decoder can not be created
	}

	return &Config{
		sql:     c,
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
