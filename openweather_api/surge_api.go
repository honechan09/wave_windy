// うねりのAPI
// https://open-meteo.co
package openweather_api

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetSurgeWether(lat, lon float64) (string, error) {
	url := fmt.Sprintf("https://marine-api.open-meteo.com/v1/marine?latitude=%f&longitude=%f&current=swell_wave_height&hourly=swell_wave_height,swell_wave_direction&timezone=Asia%%2FTokyo", lat, lon)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
