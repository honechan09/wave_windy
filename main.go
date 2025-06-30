package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	// "io"
	"log"
	"net/http"
	"os"
	"strings"
	"wave_windy/openweather_api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}
	//ichinomiya, chiba
	zipcode := "2960002" // 鴨川の郵便番号

	e := echo.New()
	InfoweatherApi(e, zipcode)

	e.Logger.Fatal(e.Start(":8080"))
	log.Println("サーバーが起動しました: http://localhost:8080")

}

func InfoweatherApi(e *echo.Echo, zipcode string) {
	apiKey := os.Getenv("API_KEY_WINDY")
	lat, lon, err := openweather_api.GetGeoInfo(zipcode)
	if err != nil {
		log.Fatal("緯度経度の取得に失敗しました: ", err)
	}
	e.GET("/", func(c echo.Context) error {
		result_windy, err_windy := openweather_api.GetWeather(lat, lon, apiKey)
		result_surge, err_surge := openweather_api.GetSurgeWether(lat, lon)
		if err_windy != nil {
			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_windy.Error())
		}
		if err_surge != nil {
			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_surge.Error())
		}
		surgeStr := strings.Join(result_surge, "\n")
		result := surgeStr + "\n" + result_windy
		if result == "" {
			return c.String(http.StatusInternalServerError, "結果が不足しています "+err_surge.Error())
		}
		return c.String(http.StatusOK, result)
	})
}
