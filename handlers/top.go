package handlers

import (
	"net/http"
	"os"
	"wave_windy/openweather_api"

	"github.com/labstack/echo/v4"
)

func TopPageHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title": "トップページ",
	}
	return c.Render(http.StatusOK, "top.html", data)
}

func OpenWindyApi(c echo.Context) error {
	lat := 35.377426
	lon := 140.390991
	apiKey := os.Getenv("API_KEY")
	result, err := openweather_api.GetWeather(lat, lon, apiKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, "API取得失敗")
	}
	return c.String(http.StatusOK, result)

}
