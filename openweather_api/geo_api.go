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

func GetGeoInfo(zipcode string) (string, error) {
	url := fmt.Sprintf("https://geoapi.heartrails.com/api/json?method=searchByPostal&postal=%s", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	lan, lon, err := GetGeolan_lon(string(body))
	if err != nil {
		return "", err
	}

	fmt.Println("lan:", lan, "lon:", lon)

	return string(body), nil
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

// zipcodeから緯度経度を取得するのは完了しているが、ここから風とうねりの向き関数呼び出す処理を追加すル〜
