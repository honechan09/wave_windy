package openweather_api

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// const zipcode = "2994301"

func GetGeoInfo(zipcode string) (float64, float64, error) {
	url := fmt.Sprintf("https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=%s", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body := mustReadAll(resp.Body)
	lat, lon, err := GetGeolan_lon(string(body))
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}

func GetGeolan_lon(geo_resp string) (float64, float64, error) {
	// レスポンス用の構造体
	type geoResponse struct {
		Response struct {
			Location []struct {
				X string `json:"x"`
				Y string `json:"y"`
			} `json:"location"`
		} `json:"response"`
	}

	var resp geoResponse
	err := json.Unmarshal([]byte(geo_resp), &resp)
	if err != nil {
		return 0, 0, err
	}
	if len(resp.Response.Location) == 0 {
		return 0, 0, nil // データなし
	}
	lon, err := strconv.ParseFloat(resp.Response.Location[0].X, 64)
	if err != nil {
		return 0, 0, err
	}
	lan, err := strconv.ParseFloat(resp.Response.Location[0].Y, 64)
	if err != nil {
		return 0, 0, err
	}
	return lan, lon, nil
}

func mustReadAll(r io.Reader) []byte {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err) // テストや開発用。運用では適切なエラーハンドリングを
	}
	return b
}
