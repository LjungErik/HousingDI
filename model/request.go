package model

type HousingForSaleRequest struct {
	TimeLocation string
	Data         []byte
}

type HousingSoldRequest struct {
	TimeLocation string
	Data         []byte
}
