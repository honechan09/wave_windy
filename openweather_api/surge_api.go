package openweather_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wave_windy/entity"
)

func GetSurgeWether(lat, lon float64) ([]string, error) {
	url := fmt.Sprintf("https://marine-api.open-meteo.com/v1/marine?latitude=%f&longitude=%f&current=swell_wave_height&hourly=swell_wave_height,swell_wave_direction&timezone=Asia%%2FTokyo", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var surgeResp entity.SurgeResponse
	err = json.Unmarshal(body, &surgeResp)
	if err != nil {
		return nil, err
	}

	n := len(surgeResp.Hourly.Time)
	var results []string
	for i := 0; i < n; i++ {
		time := surgeResp.Hourly.Time[i]
		height := surgeResp.Hourly.SwellWaveHeight[i]
		direction := surgeResp.Hourly.SwellWaveDirection[i]
		dirStr := surgeDirection(direction)
		results = append(results, fmt.Sprintf("時刻: %s, 波高: %.2f, 方向: %s", time, height, dirStr))
	}

	return results, nil
}

func surgeDirection(deg float64) string {
	directions := []string{
		"北", "北北東", "北東", "東北東",
		"東", "東南東", "南東", "南南東",
		"南", "南南西", "南西", "西南西",
		"西", "西北西", "北西", "北北西",
	}
	idx := int((deg+11.25)/22.5) % 16
	return directions[idx]
}
