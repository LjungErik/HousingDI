package model

import (
	"github.com/LjungErik/datainjestor/lib/encoding"

	log "github.com/sirupsen/logrus"
)

type Broker struct {
	PropertyLink string `json:"propertyLink"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Firm         string `json:"firm"`
}

type SpaceWithUnit struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type HousingForSale struct {
	PropId            uint64        `json:"propId"`
	Address           string        `json:"address"`
	Area              string        `json:"area"`
	City              string        `json:"city"`
	AskPrice          float64       `json:"askPrice"`
	AccommodationType string        `json:"accommodationType"`
	FormOfTenure      string        `json:"formOfTenure"`
	NumberOfRooms     float32       `json:"numberOfRooms"`
	LivingSpace       SpaceWithUnit `json:"livingSpace"`
	GrossFloorArea    SpaceWithUnit `json:"grossFloorArea"`
	PlotSize          SpaceWithUnit `json:"plotSize"`
	Balcony           bool          `json:"balcony"`
	Patio             bool          `json:"patio"`
	Floor             string        `json:"floor"`
	ConstructionYear  uint32        `json:"constructionYear"`
	HousingSociety    string        `json:"housingSociety"`
	LivingFee         float64       `json:"livingFee"`
	OperatingCost     float64       `json:"operatingCost"`
	PlotFee           float64       `json:"plotFee"`
	AreaLease         float64       `json:"areaLease"`
	PricePerSqm       float64       `json:"pricePerSqm"`
	NumberOfVisits    int32         `json:"numberOfVisits"`
	DaysOnHemnet      int32         `json:"daysOnHemnet"`
	Broker            Broker        `json:"broker"`
}

type HousingSold struct {
	PropId            uint64        `json:"propId"`
	Address           string        `json:"address"`
	Area              string        `json:"area"`
	City              string        `json:"city"`
	SaleDate          string        `json:"saleDate"`
	SalePrice         float64       `json:"salePrice"`
	AskPrice          float64       `json:"askPrice"`
	AccommodationType string        `json:"accommodationType"`
	FormOfTenure      string        `json:"formOfTenure"`
	NumberOfRooms     float32       `json:"numberOfRooms"`
	LivingSpace       SpaceWithUnit `json:"livingSpace"`
	GrossFloorArea    SpaceWithUnit `json:"grossFloorArea"`
	PlotSize          SpaceWithUnit `json:"plotSize"`
	Balcony           bool          `json:"balcony"`
	Patio             bool          `json:"patio"`
	Floor             string        `json:"floor"`
	ConstructionYear  uint32        `json:"constructionYear"`
	HousingSociety    string        `json:"housingSociety"`
	LivingFee         float64       `json:"livingFee"`
	OperatingCost     float64       `json:"operatingCost"`
	PlotFee           float64       `json:"plotFee"`
	AreaLease         float64       `json:"areaLease"`
	PricePerSqm       float64       `json:"pricePerSqm"`
	NumberOfVisits    int32         `json:"numberOfVisits"`
	DaysOnHemnet      int32         `json:"daysOnHemnet"`
	Broker            Broker        `json:"broker"`
}

func ParseHousingForSale(req *HousingForSaleRequest, decoder encoding.IDecoder) (*HousingForSale, error) {
	var m HousingForSale
	err := decoder.UnmarshalBytes(req.Data, &m)

	if err != nil {
		log.Warn("Failed to unmarshal for sale housing data")
		return nil, err
	}

	return &m, nil
}

func ParseHousingSold(req *HousingSoldRequest, decoder encoding.IDecoder) (*HousingSold, error) {
	var m HousingSold
	err := decoder.UnmarshalBytes(req.Data, &m)

	if err != nil {
		log.Warn("Failed to unmarshal for sale housing data")
		return nil, err
	}

	pe := newParsingErrors()
	longTime := m.SaleDate + " 00:00:00 " + req.TimeLocation
	m.SaleDate = parseInLocation(longTime, req.TimeLocation, "soldDate", pe)

	if pe.Error() != nil {
		log.WithFields(log.Fields{
			"SoldDateTime": longTime,
		}).Warn("Failed to convert sold date to utc timestamp")
		return nil, pe.Error()
	}

	return &m, nil
}
