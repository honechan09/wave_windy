package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//ichinomiya, chiba
	lat := 35.3727
	lon := 140.3685
	//openweather API key
	//https://openweathermap.org/current
	APIkey := "ab439487caabe9c49c7d15b6fdf608ef"

	url := fmt.Sprintf("https://pro.openweathermap.org/data/2.5/forecast/hourly?lat=%f&lon=%f&appid=%s", lat, lon, APIkey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("リクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("レスポンス読み込み失敗: %v", err)
	}

	fmt.Println(string(body))
}
