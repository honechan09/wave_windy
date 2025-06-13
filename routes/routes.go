package routes

import (
	"wave_windy/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/", handlers.TopPageHandler)
}
