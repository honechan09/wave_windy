package main

import (
	// "fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"html/template"

	"log"
	"net/http"
	"os"
	"strings"
	"wave_windy/entity"
	"wave_windy/openweather_api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}
	//ichinomiya, chiba
	// zipcode := "2960002" // 鴨川の郵便番号

	e := echo.New()
	t := &entity.Template{
		Templates: template.Must(template.ParseGlob("templates/index.html")),
	}
	e.Renderer = t

	// トップページのみ
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/weather", func(c echo.Context) error {
		zipcode := c.FormValue("city")
		if zipcode == "" {
			return c.String(http.StatusBadRequest, "郵便番号が入力されていません")
		}
		status, result_sea, err := InfoweatherApi(zipcode)
		if err != nil {
			log.Println("API取得失敗: ", err)
			return c.String(http.StatusInternalServerError, "API取得失敗: "+err.Error())
		}
		return c.String(status, result_sea)
	})

	e.Logger.Fatal(e.Start(":8080"))
	log.Println("サーバーが起動しました: http://localhost:8080")

}

func InfoweatherApi(zipcode string) (int, string, error) {
	apiKey := os.Getenv("API_KEY_WINDY")
	lat, lon, err := openweather_api.GetGeoInfo(zipcode)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	result_windy, err_windy := openweather_api.GetWeather(lat, lon, apiKey)
	result_surge, err_surge := openweather_api.GetSurgeWether(lat, lon)
	if err_windy != nil {
		return http.StatusInternalServerError, "", err_windy
	}
	if err_surge != nil {
		return http.StatusInternalServerError, "", err_surge
	}
	surgeStr := strings.Join(result_surge, "\n")
	log.Println("surgeStr:", surgeStr)
	log.Println("result_windy:", result_windy)
	result := surgeStr + "\n" + result_windy
	if result == "" {
		return http.StatusInternalServerError, "", err_surge
	}
	return http.StatusOK, result, nil
}
