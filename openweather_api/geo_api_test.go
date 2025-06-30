package openweather_api

import (
	"testing"
)

const zipcode = "2994301"

func TestGetSurgeWether(t *testing.T) {
	lan, lon, err := GetGeoInfo(zipcode)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lan, lon)
}
