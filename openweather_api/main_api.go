package openweather_api

import (
	"fmt"
	"io"
	"net/http"
)

func GetWeather(lat, lon float64, apiKey string) (string, error) {
	// 5days par 3hours
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&lang=ja", lat, lon, apiKey)
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
