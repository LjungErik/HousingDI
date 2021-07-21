package encoding

import (
	"errors"

	"encoding/json"

	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

const EncodingCSV = 1
const EncodingJSON = 2

type IDecoder interface {
	UnmarshalBytes(in []byte, out interface{}) error
}

type CSVDecoder struct{}
type JSONDecoder struct{}

func NewDecoder(encodingType int) (IDecoder, error) {
	switch encodingType {
	case EncodingCSV:
		log.Debug("Creating CSV Decoder")
		return CSVDecoder{}, nil
	case EncodingJSON:
		log.Debug("Creating JSON Decoder")
		return JSONDecoder{}, nil
	}

	log.WithFields(log.Fields{
		"encodingType": encodingType,
	}).Warn("Failed to create decoder, unsupported encoding type")
	return nil, errors.New("Failed to create decoder, unsupported encoding type")
}

func (d CSVDecoder) UnmarshalBytes(in []byte, out interface{}) error {
	err := gocsv.UnmarshalBytes(in, out)
	if err != nil {
		log.Warn("Failed to unmarshal CSV data")
		return err
	}
	return nil
}

func (d JSONDecoder) UnmarshalBytes(in []byte, out interface{}) error {
	err := json.Unmarshal(in, out)
	if err != nil {
		log.Warn("Failed to unmarshal JSON data")
		return err
	}
	return nil
}
