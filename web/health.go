package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (wc *Config) ready(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("[Ready Health]: OK")
	w.WriteHeader(http.StatusOK)
}

func (wc *Config) live(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("[Live Health]: OK")
	w.WriteHeader(http.StatusOK)
}
