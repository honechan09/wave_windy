package openweather_api

import (
	"encoding/json"
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
	data_formatted, err := FormatWeatherData(body)
	return data_formatted, nil
}

type WeatherResponse struct {
	List []WeatherItem `json:"list"`
}
type WeatherItem struct {
	DtTxt string   `json:"dt_txt"`
	Wind  WindInfo `json:"wind"`
}
type WindInfo struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

func FormatWeatherData(body []byte) (string, error) {
	var resp WeatherResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	result := ""
	for _, item := range resp.List {
		result += fmt.Sprintf(
			"日時: %s, %s\n",
			item.DtTxt, FormatWindInfo(item.Wind),
		)
	}
	if result == "" {
		return "データなし", nil
	}
	return result, nil
	// return string(body), nil
}

// 風向への変換
func degToDirection(deg float64) string {
	directions := []string{
		"北", "北北東", "北東", "東北東",
		"東", "東南東", "南東", "南南東",
		"南", "南南西", "南西", "西南西",
		"西", "西北西", "北西", "北北西",
	}
	idx := int((deg+11.25)/22.5) % 16
	return directions[idx]
}

// 風速への変換
func FormatWindInfo(wind WindInfo) string {
	return fmt.Sprintf("風速: %.2f m/s, 風向: %s, 最大風速: %.2f m/s",
		wind.Speed, degToDirection(wind.Deg), wind.Gust)
}
