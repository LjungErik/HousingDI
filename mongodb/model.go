package mongodb

type WithUnits struct {
	Value float64 `bson:"value"`
	Unit  string  `bson:"unit"`
}

type Broker struct {
	PropertyLink string `bson:"propertyLink,omitempty"`
	Name         string `bson:"name,omitempty"`
	Link         string `bson:"link,omitempty"`
	Firm         string `bson:"firm,omitempty"`
}

type HousingForSaleDoc struct {
	PropId            uint64    `bson:"propId,omitempty"`
	Address           string    `bson:"address,omitempty"`
	Area              string    `bson:"area,omitempty"`
	City              string    `bson:"city,omitempty"`
	AskPrice          float64   `bson:"askPrice,omitempty"`
	AccommodationType string    `bson:"accommodationType,omitempty"`
	FormOfTenure      string    `bson:"formOfTenure,omitempty"`
	NumberOfRooms     float32   `bson:"numberOfRooms,omitempty"`
	LivingSpace       WithUnits `bson:"livingSpace,omitempty"`
	GrossFloorArea    WithUnits `bson:"grossFloorArea,omitempty"`
	PlotSize          WithUnits `bson:"plotSize,omitempty"`
	Balcony           bool      `bson:"balcony,omitempty"`
	Patio             bool      `bson:"patio,omitempty"`
	Floor             string    `bson:"floor,omitempty"`
	ConstructionYear  uint32    `bson:"constructionYear,omitempty"`
	HousingSociety    string    `bson:"housingSociety,omitempty"`
	LivingFee         float64   `bson:"livingFee,omitempty"`
	OperatingCost     float64   `bson:"operatingCost,omitempty"`
	PlotFee           float64   `bson:"plotFee,omitempty"`
	AreaLease         float64   `bson:"areaLease,omitempty"`
	PricePerSqm       float64   `bson:"pricePerSqm,omitempty"`
	NumberOfVisits    int32     `bson:"numberOfVisits,omitempty"`
	DaysOnHemnet      int32     `bson:"daysOnHemnet,omitempty"`
	Broker            Broker    `bson:"broker,omitempty"`
}

type HousingSoldDoc struct {
	PropId            uint64    `bson:"propId,omitempty"`
	Address           string    `bson:"address,omitempty"`
	Area              string    `bson:"area,omitempty"`
	City              string    `bson:"city,omitempty"`
	SaleDate          string    `bson:"saleDate,omitempty"`
	SalePrice         float64   `bson:"salePrice,omitempty"`
	AskPrice          float64   `bson:"askPrice,omitempty"`
	AccommodationType string    `bson:"accommodationType,omitempty"`
	FormOfTenure      string    `bson:"formOfTenure,omitempty"`
	NumberOfRooms     float32   `bson:"numberOfRooms,omitempty"`
	LivingSpace       WithUnits `bson:"livingSpace,omitempty"`
	GrossFloorArea    WithUnits `bson:"grossFloorArea,omitempty"`
	PlotSize          WithUnits `bson:"plotSize,omitempty"`
	Balcony           bool      `bson:"balcony,omitempty"`
	Patio             bool      `bson:"patio,omitempty"`
	Floor             string    `bson:"floor,omitempty"`
	ConstructionYear  uint32    `bson:"constructionYear,omitempty"`
	HousingSociety    string    `bson:"housingSociety,omitempty"`
	LivingFee         float64   `bson:"livingFee,omitempty"`
	OperatingCost     float64   `bson:"operatingCost,omitempty"`
	PlotFee           float64   `bson:"plotFee,omitempty"`
	AreaLease         float64   `bson:"areaLease,omitempty"`
	PricePerSqm       float64   `bson:"pricePerSqm,omitempty"`
	NumberOfVisits    int32     `bson:"numberOfVisits,omitempty"`
	DaysOnHemnet      int32     `bson:"daysOnHemnet,omitempty"`
	Broker            Broker    `bson:"broker,omitempty"`
}
