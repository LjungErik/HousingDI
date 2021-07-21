package model

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	longTimeFormat  = "2006-01-02 15:04:05 MST"
	timestampFormat = "2006-01-02 15:04:05"
)

// Timestamp in string format

type parsingErrors struct {
	failedFields []string
}

func newParsingErrors() *parsingErrors {
	return &parsingErrors{
		failedFields: make([]string, 0),
	}
}

func (p *parsingErrors) Error() error {
	if len(p.failedFields) > 0 {
		return fmt.Errorf("Failed parsing fields: %+q", p.failedFields)
	}
	return nil
}

func parseInt(value, fieldName string, out *parsingErrors) int64 {
	if len(value) == 0 {
		return 0
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"field": fieldName,
		}).Warnf(err.Error())
		out.failedFields = append(out.failedFields, fieldName)
	}

	return parsed
}

func parseFloat(value, fieldName string, out *parsingErrors) float64 {
	if len(value) == 0 {
		return 0.0
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"field": fieldName,
		}).Warnf(err.Error())
		out.failedFields = append(out.failedFields, fieldName)
	}

	return parsed
}

func asTimestamp(t time.Time) string {
	return t.UTC().Format(timestampFormat)
}

// ParseInLocation Parses long time format (:longTime) in location (:timeLocation)
// as Timestamp
func parseInLocation(longTime, timeLocation, fieldName string, out *parsingErrors) string {
	location, err := time.LoadLocation(timeLocation)
	if err != nil {
		log.WithFields(log.Fields{
			"longTime":     longTime,
			"timeLocation": timeLocation,
			"fieldName":    fieldName,
		}).Warnf("Failed to parse location: %v", err)

		out.failedFields = append(out.failedFields, fieldName)

		return ""
	}

	parsedTime, err := time.ParseInLocation(longTimeFormat, longTime, location)
	if err != nil {
		log.WithFields(log.Fields{
			"longTime":     longTime,
			"timeLocation": timeLocation,
		}).Warnf("Failed to parse time: %v", err)

		out.failedFields = append(out.failedFields, fieldName)

		return ""
	}

	return asTimestamp(parsedTime)
}
