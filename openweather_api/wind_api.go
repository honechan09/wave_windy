package openweather_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"wave_windy/entity"
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

func FormatWeatherData(body []byte) (string, error) {
	var resp entity.WeatherResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	result := ""

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	for _, item := range resp.List {
		// item.DtTxtは "2006-01-02 15:04:05" 形式
		t, err := time.ParseInLocation("2006-01-02 15:04:05", item.DtTxt, jst)
		if err != nil {
			continue // パースできなければスキップ
		}
		if t.Before(now) {
			continue // 過去データはスキップ
		}
		result += fmt.Sprintf(
			"日時: %s, %s\n",
			item.DtTxt, FormatWindInfo(item.Wind),
		)
	}
	if result == "" {
		return "データなし", nil
	}
	return result, nil
}

// 風向への変換
func windDirection(deg float64) string {
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
func FormatWindInfo(wind entity.WindInfo) string {
	return fmt.Sprintf("風速: %.2f m/s, 風向: %s, 最大風速: %.2f m/s",
		wind.Speed, windDirection(wind.Deg), wind.Gust)
}
