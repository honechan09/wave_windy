package openweather_api

import (
	"testing"
)

const lat = 35.377426
const lon = 140.390991

func TestGetSurgeWether(t *testing.T) {
	results, err := GetSurgeWether(lat, lon)
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range results {
		t.Log(line)
	}
}
