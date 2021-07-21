package model

import "testing"

func TestParseInLocation(t *testing.T) {
	expected := "2020-01-01 05:01:00"
	p := newParsingErrors()
	timestamp := parseInLocation("2020-01-01 00:01:00 EST", "EST", "timestamp", p)

	if len(p.failedFields) != 0 {
		t.Errorf("Failed to parse date, %v", p.Error())
	}

	if timestamp != expected {
		t.Errorf("Parsing timestamp was incorrect, got %v, want: %v", timestamp, expected)
	}
}

func TestParseInLocationCET(t *testing.T) {
	expected := "2020-01-01 00:01:00"
	p := newParsingErrors()
	timestamp := parseInLocation("2020-01-01 01:01:00 CET", "CET", "timestamp", p)

	if len(p.failedFields) != 0 {
		t.Errorf("Failed to parse date, %v", p.Error())
	}

	if timestamp != expected {
		t.Errorf("Parsing timestamp was incorrect, got %v, want: %v", timestamp, expected)
	}
}

func TestParseInt(t *testing.T) {
	var expected int64 = 1999
	p := newParsingErrors()
	str := "1999"
	result := parseInt(str, "year", p)

	if len(p.failedFields) != 0 {
		t.Errorf("Failed to parse date, %v", p.Error())
	}

	if result != expected {
		t.Errorf("Parsing int was incorrect, got %v,  want %v", result, expected)
	}
}

func TestParseFloat(t *testing.T) {
	var expected float64 = 23.56
	p := newParsingErrors()
	str := "23.56"
	result := parseFloat(str, "year", p)

	if len(p.failedFields) != 0 {
		t.Errorf("Failed to parse date, %v", p.Error())
	}

	if result != expected {
		t.Errorf("Parsing int was incorrect, got %v,  want %v", result, expected)
	}
}
