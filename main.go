package main

import (
	"log"
	"net/http"
	"os"

	"wave_windy/openweather_api"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}
	//ichinomiya, chiba
	apiKey := os.Getenv("API_KEY_WINDY")
	lat := 35.377426
	lon := 140.390991

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result_windy, err_windy := openweather_api.GetWeather(lat, lon, apiKey)
		result_surge, err_surge := openweather_api.GetSurgeWether(lat, lon)
		if err_windy != nil {
			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_windy.Error())
		}

		if err_surge != nil {
			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_surge.Error())

		}
		result := result_surge + "\n" + result_windy
		if result == "" {
			return c.String(http.StatusInternalServerError, "結果が不足しています "+err_surge.Error())
		}

		return c.String(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":8080"))
	log.Println("サーバーが起動しました: http://localhost:8080")

}
