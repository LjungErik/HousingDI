package web

import (
	"io/ioutil"
	"net/http"

	"github.com/LjungErik/datainjestor/model"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (wc *Config) writeHousingForSale(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query()
	tl := q.Get("timeLocation")
	if tl == "" {
		log.Error("Missing valid time location")
		http.Error(w, "Invalid time location", http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read body")
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	hr := model.HousingForSaleRequest{
		TimeLocation: tl,
		Data:         b,
	}

	_, err = model.ParseHousingForSale(&hr, wc.decoder)
	if err != nil {
		log.Error("Failed to parse housing for sale data ")
		http.Error(w, "Cannot parse Housing For Sale data", http.StatusBadRequest)
		return
	}

	//Do something with the for sale housing data

	w.WriteHeader(http.StatusNoContent)
}

func (wc *Config) writeHousingSold(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query()
	tl := q.Get("timeLocation")
	if tl == "" {
		log.Error("Missing valid time location")
		http.Error(w, "Invalid time location", http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read body")
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	hr := model.HousingSoldRequest{
		TimeLocation: tl,
		Data:         b,
	}

	_, err = model.ParseHousingSold(&hr, wc.decoder)
	if err != nil {
		log.Error("Failed to parse housing sold data ")
		http.Error(w, "Cannot parse Housing Sold data", http.StatusBadRequest)
		return
	}

	//Do something with the sold housing data

	w.WriteHeader(http.StatusNoContent)
}
