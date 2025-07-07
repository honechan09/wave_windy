package openweather_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
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

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	group := make(map[string]*entity.Agg)

	n := len(surgeResp.Hourly.Time)
	for i := 0; i < n; i++ {
		t, err := time.ParseInLocation("2006-01-02T15:04", surgeResp.Hourly.Time[i], jst)
		if err != nil {
			continue
		}
		if t.Before(now) {
			continue // 過去データは除外
		}
		// 3時間ごとのキーを作成（例: 0:00, 3:00, ...）
		hour := t.Hour() / 3 * 3
		groupKey := fmt.Sprintf("%04d-%02d-%02d %02d:00", t.Year(), t.Month(), t.Day(), hour)
		if group[groupKey] == nil {
			group[groupKey] = &entity.Agg{}
		}
		group[groupKey].Count++
		group[groupKey].HeightSum += surgeResp.Hourly.SwellWaveHeight[i]
		group[groupKey].DirSum += surgeResp.Hourly.SwellWaveDirection[i]
	}
	// キーを昇順に並べる
	var keys []string
	for k := range group {
		keys = append(keys, k)
	}
	// 並び替え
	sort.Strings(keys)

	var results []string
	for _, k := range keys {
		a := group[k]
		avgHeight := a.HeightSum / float64(a.Count)
		avgDir := a.DirSum / float64(a.Count)
		dirStr := surgeDirection(avgDir)
		results = append(results, fmt.Sprintf("時刻: %s, 波高: %.2f, 方向: %s", k, avgHeight, dirStr))
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
