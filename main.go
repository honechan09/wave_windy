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
	apiKey := os.Getenv("API_KEY")
    lat := 35.377426
    lon := 140.390991
	//openweather API key

	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        result, err := openweather_api.GetWeather(lat, lon, apiKey)
        if err != nil {
            return c.String(http.StatusInternalServerError, "API取得失敗: "+err.Error())
        }
        return c.String(http.StatusOK, result)
    })
    e.Logger.Fatal(e.Start(":8080"))
	log.Println("サーバーが起動しました: http://localhost:8080")

}
