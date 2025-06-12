package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"wave_windy/openweather_api"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	//ichinomiya, chiba
	lat := 35.377426
	lon := 140.390991
	//openweather API key
	//https://openweathermap.org/current
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}
	apiKey := os.Getenv("API_KEY")
	// 5days par 3hours
	result, err := openweather_api.GetWeather(lat, lon, apiKey)
	if err != nil {
		log.Fatalf("API取得失敗: %v", err)
	}
	fmt.Println(result)
}
