package main

import (
	"fmt"
	"log"

	"wave_windy/openweather_api"
)

func main() {
	//ichinomiya, chiba
	lat := 35.377426
	lon := 140.390991
	//openweather API key
	//https://openweathermap.org/current
	apiKey := "ab439487caabe9c49c7d15b6fdf608ef"
	// 5days par 3hours
	result, err := openweather_api.GetWeather(lat, lon, apiKey)
	if err != nil {
		log.Fatalf("API取得失敗: %v", err)
	}
	fmt.Println(result)
}
