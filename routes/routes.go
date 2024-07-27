package routes

import (
	"go-hipo/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.Hello)
}
