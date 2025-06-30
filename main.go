package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"html/template"

	"log"
	"net/http"
	"wave_windy/entity"
	// "os" api用
	// "strings" api用
	// "wave_windy/openweather_api" api用
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

	e.Logger.Fatal(e.Start(":8080"))
	log.Println("サーバーが起動しました: http://localhost:8080")

}

// InfoweatherApi(e, zipcode)
// func InfoweatherApi(e *echo.Echo, zipcode string) {
// 	apiKey := os.Getenv("API_KEY_WINDY")
// 	lat, lon, err := openweather_api.GetGeoInfo(zipcode)
// 	if err != nil {
// 		log.Fatal("緯度経度の取得に失敗しました: ", err)
// 	}
// 	e.GET("/", func(c echo.Context) error {
// 		result_windy, err_windy := openweather_api.GetWeather(lat, lon, apiKey)
// 		result_surge, err_surge := openweather_api.GetSurgeWether(lat, lon)
// 		if err_windy != nil {
// 			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_windy.Error())
// 		}
// 		if err_surge != nil {
// 			return c.String(http.StatusInternalServerError, "API取得失敗: "+err_surge.Error())
// 		}
// 		surgeStr := strings.Join(result_surge, "\n")
// 		result := surgeStr + "\n" + result_windy
// 		if result == "" {
// 			return c.String(http.StatusInternalServerError, "結果が不足しています "+err_surge.Error())
// 		}
// 		return c.String(http.StatusOK, result)
// 	})
// }
