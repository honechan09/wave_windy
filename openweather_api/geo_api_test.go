package openweather_api

import (
	"testing"
)

const test_zipcode = "2994301"

func TestGeoAPI(t *testing.T) {
	result, err := GetGeoInfo(test_zipcode)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
